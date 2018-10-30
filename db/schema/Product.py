from .DB import db, Base, GormModelMixin


class Product(Base, GormModelMixin):
    id = db.Column(db.Integer(), primary_key=True)
    name = db.Column(db.UnicodeText(), nullable=False)
    price = db.Column(db.Numeric(30, 16), nullable=False)

    merchant_id = db.Column(
        db.Integer(),
        db.ForeignKey('merchant.id'),
        nullable=False,
    )
