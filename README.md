# Gopher Trade (WIP)

Gopher Trade is a currency conversion API built in Go. References are scattered as links through this README. Hopefully when I'm finished I will create a References section at the end of this document.

## Contents

* [Features](#features)
  * [Development](#development)
* [Testing](#testing)
* [Run locally](#run-locally)
  * [Requirements (running without Docker)](#requirements-running-without-docker)
  * [Execution](#execution)
* [Application Dependencies](#aplication-dependencies)
* [Tech Specs](#tech-specs)
  * [Architecture](#architecture)
  * [Number Types](#number-types)
* [Workflow](#workflow)

## Features

Project (development) and product (API) features:

### Development

* CI (GitHub Actions) - Runs tests and linters on Pull Requests to main (which is a protected branch);
* Lint (golangci-lint) - Enforces style and best practices;
* Makefile - Simplifies running tests, dependencies and application.

### API

coming soon...

## Testing

This project does not aim to have 100% test coverage. Also, its development did not follow strict TDD doctrine, instead it [aimed to test behaviour](https://dave.cheney.net/paste/absolute-unit-test-london-gophers.pdf).

## Run locally

### Requirements (running without Docker)

* [Golang](https://go.dev/dl/) v1.18+

### Execution

The app is not runnable yet. You can execute all unit tests by running:

```bash
make test
```

## Aplication Dependencies

This project imports external packages:

* [Testify](https://github.com/stretchr/testify) - to simplify test assertions;
* [Decimal](https://github.com/shopspring/decimal) - to handle decimal values. See [number types section](#number-types) below;
* [UUID]("https://github.com/google/uuid") - for generating unique IDs.

## Tech Specs

### Architecture

Aiming to make this project more scalable and maintainable, it is being developed based on [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

### Number types

The first idea in this project was to use integer values for monetary values (representing in cents), for simplicity. However, specificities of currency exchange made this a bad alternative, i.e., non-fiat currencies (such as criptocurrencies) and exchange rates being commonly measured and represented with more than 2 decimal points.

Considering the [floating-point arithmetic problem](https://floating-point-gui.de/), the [decimal](https://pkg.go.dev/github.com/shopspring/decimal) package was used (initially using the standard lib `math/big` package was considered, but given the scope of the project and the time constraint, using an external lib was considered the best trade-off).

## Workflow

A very basic [Kanban board](https://www.atlassian.com/agile/kanban/boards) was used to keep track of priorities and deadline for the project. It was kept simples since the project is being developed by one person. For this purpose, [GitHub Projects](https://docs.github.com/en/issues/planning-and-tracking-with-projects) showed to be enough.

Given the time constraint (deadline and hours available during the day), an [MVP](https://www.productplan.com/glossary/minimum-viable-product/#:~:text=A%20minimum%20viable%20product%2C%20or,iterate%20and%20improve%20the%20product.) was planned, an some improvements indicated in the Backlog.

The [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) standard was followed during the development of the project. Pull Requests were opened for each feature (although without code review, but for the sake of having a stable main branch - and prettier commit log), mocking a [trunk-based development](https://www.atlassian.com/continuous-delivery/continuous-integration/trunk-based-development).
