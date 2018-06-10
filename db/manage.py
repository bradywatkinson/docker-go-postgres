from time import sleep
from flask import Flask
from flask_script import Manager
from flask_migrate import Migrate, MigrateCommand
from sqlalchemy.exc import OperationalError

from schema import db

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://dev:dev@postgres/dev'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db.init_app(app)

migrate = Migrate(app, db)

manager = Manager(app)
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
