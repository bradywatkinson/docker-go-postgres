from flask_sqlalchemy import SQLAlchemy
from sqlalchemy import MetaData
from sqlalchemy.schema import CreateColumn
from sqlalchemy.ext.compiler import compiles
from sqlalchemy.ext.declarative import declared_attr, declarative_base


db = SQLAlchemy()


metadata = MetaData(
    naming_convention={
        "ix": "ix_%(table_name)s_%(column_0_label)s",
        "uq": "uq_%(table_name)s_%(column_0_name)s",
        "ck": "ck_%(table_name)s_%(constraint_name)s",
        "fk": "fk_%(table_name)s_%(column_0_name)s_%(referred_table_name)s",
        "pk": "pk_%(table_name)s"
    }
)


Base = declarative_base(metadata=metadata, cls=db.Model)


@compiles(CreateColumn, 'postgresql')
def use_identity(element, compiler, **kw):
    text = compiler.visit_create_column(element, **kw)
    text = text.replace("SERIAL", "INT GENERATED BY DEFAULT AS IDENTITY")
    return text


class GormModelMixin(object):
    created_at = db.Column(
        db.DateTime(),
        server_default=db.func.now(),
        nullable=False,
    )
    updated_at = db.Column(
        db.DateTime(),
        server_default=db.func.now(),
        nullable=False,
    )
    deleted_at = db.Column(
        db.DateTime(),
        nullable=True,
    )

    @declared_attr
    def __table_args__(cls):
        return (db.Index('ix_{}_deleted_at'.format(cls.__tablename__), 'deleted_at'), )
