# Database Migrations Service (DMS)

A simple database migration utility using `Flask-Migrate`.

## Overview

DMS is used to manage database migrations. It utilises the wonderful [`Flask-Migrate`](https://github.com/miguelgrinberg/Flask-Migrate) to generate an [`Alembic`](https://github.com/zzzeek/alembic) _migration script_ which in turn can be used to generate the SQL for a _migration_.

A _migration_ is a set of database transformations which is described in a _migration script_. `Flask-Migrate` can generate a _migration script_ from changes to the set of `SQLAlchemy` models that describes a database. A migration is reference by it's revision id.

## Usage

Create a migration:

    docker-compose run alembic migrate -m 'Migration comment'

Perform a migration:

    docker-compose run alembic upgrade

Generate SQL for a migration:

    docker-compose run alembic upgrade --sql <revision id>

Generate SQL and export it to `/migrations/transforms`

    docker-compose run alembic export

## Commands

    upgrade             Upgrade to a later version
    migrate             Alias for 'revision --autogenerate'
    current             Display the current revision for each database.
    stamp               'stamp' the revision table with the given revision;
                        dont run any migrations
    init                Generates a new migration
    downgrade           Revert to a previous version
    history             List changeset scripts in chronological order.
    revision            Create a new revision file.


## Initial Schema

    CREATE TABLE alembic_version (
        version_num VARCHAR(32) NOT NULL,
        CONSTRAINT alembic_version_pkc PRIMARY KEY (version_num)
    );

    CREATE TABLE product (
        committed_timestamp TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
        id INT GENERATED BY DEFAULT AS IDENTITY NOT NULL,
        name TEXT NOT NULL,
        price NUMERIC(30, 16) NOT NULL,
        PRIMARY KEY (id)
    );

    INSERT INTO alembic_version (version_num) VALUES ('6173b7ed3174');

## Connect to the instance

    $ docker-compose exec db psql -U dev -h db
