# Golang Database Service Boilerplate

A simple, powerful foundation to get started with a database service, built for developer experience.

## Overview

The repo is comprised of two parts; Golang Database Service (GDB) which is located under `/app` and Database Migration Service (DMS), located under `/db`. See each individual README for usage.

## Motivation

> Persistent data is perhaps the most valuable part of an application, and therefore accessing persistent data is perhaps the most critical of functionalities.

_- Brady Watkinson_

In a previous project I worked on, we happily used micro(ish)services as it was cool and trendy and solved a few problems like deployments, collaboration in a large team and extensibility (we were slowly deprecating an older code base so carving out functionality into microservices worked very well). It solved a number of problems relatively well and there was much rejoicing.

All our microservices used SQLAlchemy to access the DB, defined in a set of SQLAlchemy models that was essentially copied and pasted between each service. As development continued, it quickly became apparent that we were running into the problem of having to keep the models in sync; deploying changes to a model for one service meant updating all other services that relied on that model individually (a task that increasingly became difficult to track). At some point in time, a fellow developer suggested a shared or "common" library to maintain the set of models that represented the DB, hosted on an internal python package repository, and used as a dependency for all of our microservices. The common library made updating the models as easy as bumping a version number and ensured changes propagated to where they were needed quickly and easily, and there was much rejoicing.

For a time, we were happy, however, a specter loomed; testing and deploying changes to the shared library itself was not a simple task. Making improvements or adding functionality to the common library meant testing the delta in each service and sometimes had unexpected side effects based on different usage patterns. Optimisations to the SQLAlchemy models became impossible as we needed to support so many use-cases. Over time, the common library version diverged across services as developers only selectively updated the services that needed the change, resulting in erratic upgrades when a seemingly thin change added a huge, unexpected delta. The benefits were slowly drowned by the cost...

Thankfully, I no longer work on that project.

Having thought on the subject extensively, I came to the realisation that the entire problem could be solved not as a static dependency that was nondiscriminatorily injected into each service, but instead as a service unto itself; one that could cover a wide variety of use cases including extensibility and backwards compatibility. Such a service would also provide the side benefit of natural place for logging and performance metrics, access control and connection pooling. For such a service to be net benefit, it would have to have a small footprint and a simple interface.

This lead me to develop Golang Database Service Boilerplate; a simple service to broker database access. As I wanted to focus on developer experience, I decided to also add database migration tools, including an opinionated migration workflow.

## Issues

If you have any questions, comments, or suggestions, join my slack (www.slack.bradywatkinson.com) and join the `go-db-service-boilerplate` channel.

## License

This project is licensed under the MIT license, Copyright (c) 2018 Brady Watkinson. For more information see LICENSE.md.
