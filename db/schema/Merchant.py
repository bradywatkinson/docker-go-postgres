from .DB import db, Base, CommittedTimestampMixin


class Merchant(Base, CommittedTimestampMixin):
    id = db.Column(db.Integer(), primary_key=True)
    name = db.Column(db.UnicodeText(), nullable=False)
