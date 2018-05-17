# Go Postgres JSON Server

This project implements the example as outlined by [this tutorial](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql) with a couple of important differences:
- it is dockerized so you dont have to deal with GOPATH
- it uses [dep](https://github.com/golang/dep) for dependency management
- it uses a single postgres instance with two databases; dev for actual operation, and test to be used in the test suite
- it adds an opinionated `.editorconfig` (it seems tabs are the "go", but I hate tabs...)

## Issues

If you have any questions, comments, or suggestions, join my slack (www.slack.bradywatkinson.com) and join the `docker-go-postgres` channel.
