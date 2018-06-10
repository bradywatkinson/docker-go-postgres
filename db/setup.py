from setuptools import setup


NAME = 'db'
VERSION = '0.0.1'
AUTHOR = 'Brady Watkinson'
INSTALL_DEPENDENCIES = [
    'flake8==3.3.0',
    'Flask==0.12.2',
    'Flask-Migrate==2.0.3',
    'Flask-Script==2.0.5',
    'Flask-SQLAlchemy==2.2',
    'psycopg2==2.7.4',
]


setup(
    name=NAME,
    version=VERSION,
    author=AUTHOR,
    install_requires=INSTALL_DEPENDENCIES,
)
