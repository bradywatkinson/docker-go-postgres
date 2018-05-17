from .DB import db, Base, GormModelMixin


class Merchant(Base, GormModelMixin):
    id = db.Column(db.Integer(), primary_key=True)
    name = db.Column(db.UnicodeText(), nullable=False)
