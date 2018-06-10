from .DB import db, Base, CommittedTimestampMixin


class Vendor(Base, CommittedTimestampMixin):
    id = db.Column(db.Integer(), primary_key=True)
    name = db.Column(db.UnicodeText(), nullable=False)
