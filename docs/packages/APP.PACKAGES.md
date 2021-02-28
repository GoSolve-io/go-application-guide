# Package Layout in the example application

This document builds on [guide to Go packaging](/docs/packages/PACKAGES.md) and [guide to Go app design](/docs/appdesign/DESIGN.md).

## Breakdown of example app's packages

We want to reflect "explicit architecture" somehow in our package structure, and at the same time we want to build on good package practices from previous guide. So first thing we'll start with is creating a base for our application "core", primary and secondary adapters. We'll create them as `internal` packages, because we don't want the code to be imported by other repositories (we won't be able to provide stable interface!).

- `internal/app`
- `internal/transport`
- `internal/adapter`

### internal/app

The package `internal/app` and its sub-packages has to provide a few things.

#### A common language for primary and secondary adapters

The **outer core** of the "common language" are things like:

- a base for error handling,
- a base for logging,
- a base for handling trace ids,

The **inner core** of the "common language" is the domain - all things related to managing and renting bikes (like `Bike` type).
If there are multiple, independent "chunks of the domain", we can build them on top of the domain core. That's why the package `internal/app/bikerental/discount` exists and is separated from its parent.

Please bear in mind that in the example existence of sub-packages of `internal/app/bikerental` is over-engineering. They are created only to explain the concept of packaging in Go. In reality, this domain is so small that there's no point in splitting it into smaller chunks.

#### Application logic

For example:

- How to interpret errors?
- How to log with some additional context
- How to pass trace id through all the calls between app and adapters

#### Domain logic

For example:

- What is a Bike/Customer/Reservation?
- How to make a reservation? What is a valid request for creating a reservation? When can a reservation be created, and when it can't?
- How to calculate a discount for a reservation?

### internal/transport

The package `internal/transport` and its sub-packages provide a way for the outside world to interact with the application. It is responsible for creating a transport that exposes some kind of API.

### internal/adapter

The package `internal/adapter` and its sub-packages provide a way for the application core to communicate with the outside world. It is responsible for:

- Calling external APIs, and adapting their response format to "core" language.
- Managing database connections, managing internal representation of system state in those databases.

## Diagram of the packages working together

![Packages breakdown](apppackages.svg)


TODO: Finish this doc.

- Show dependency direction
- Show control flow direction
