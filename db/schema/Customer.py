from .DB import db, Base, GormModelMixin


class Customer(Base, GormModelMixin):
    id = db.Column(db.Integer(), primary_key=True)
    first_name = db.Column(db.UnicodeText(), nullable=False)
    last_name = db.Column(db.UnicodeText(), nullable=False)
