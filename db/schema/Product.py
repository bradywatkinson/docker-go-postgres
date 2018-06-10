from .DB import db, Base, CommittedTimestampMixin


class Product(Base, CommittedTimestampMixin):
    id = db.Column(db.Integer(), primary_key=True)
    name = db.Column(db.UnicodeText(), nullable=False)
    price = db.Column(db.Numeric(30, 16), nullable=False)
