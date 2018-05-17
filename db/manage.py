import sys
import os
import io
from time import sleep
from flask import Flask, current_app
from flask_script import Manager
from flask_migrate import Migrate, MigrateCommand
from sqlalchemy.exc import OperationalError
from alembic import command

from schema import db

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://dev:dev@db/dev'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db.init_app(app)

migrate = Migrate(app, db)

manager = Manager(app)


class Capturing(list):
    def __enter__(self):
        self._stdout = sys.stdout
        sys.stdout = self._stringio = io.StringIO()
        return self

    def __exit__(self, *args):
        self.extend(self._stringio.getvalue().splitlines())
        del self._stringio    # free up some memory
        sys.stdout = self._stdout


@MigrateCommand.command
def export():
    """Export generated sql"""
    config = current_app.extensions['migrate'].migrate.get_config(None)
    with Capturing() as output:
        command.upgrade(config, 'head', sql=True)

    revision_id = db.engine.execute('select * from alembic_version')
    file_name = '{}.sql'.format(revision_id.fetchone()[0])
    dir = os.path.dirname(__file__)
    file = os.path.join(dir, 'migrations', 'transforms', file_name)
    with open(file, "w") as fh:
        fh.write("\n".join(output))


manager.add_command('db', MigrateCommand)


if __name__ == '__main__':
    connected = False
    retries = 0
    while not connected and retries < 10:
        retries += 1
        try:
            manager.run()
            connected = True
        except OperationalError:
            print('* Could not connect to database. Retrying...')
            sleep(5)

    print('* Database did not come up in time')
