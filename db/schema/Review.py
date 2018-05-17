from .DB import db, Base, GormModelMixin


class Review(Base, GormModelMixin):
    id = db.Column(db.Integer(), primary_key=True)
    rating = db.Column(db.Integer(), nullable=False)
    review = db.Column(db.UnicodeText(), nullable=True)

    customer_id = db.Column(
        db.Integer(),
        db.ForeignKey('customer.id'),
        nullable=False,
    )
    product_id = db.Column(
        db.Integer(),
        db.ForeignKey('product.id'),
        nullable=False,
    )
