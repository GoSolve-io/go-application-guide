# Go developers incomplete guide to writing typical backend service ![Build](https://github.com/nglogic/go-example-project/workflows/Build/badge.svg)


Table of contents
- [Go developers incomplete guide to writing typical backend service !Build](#go-developers-incomplete-guide-to-writing-typical-backend-service-)
  - [What is this repository?](#what-is-this-repository)
  - [Example project](#example-project)
    - [Business requirements](#business-requirements)
    - [Design](#design)
    - [Project structure](#project-structure)
  - [Typical backend service design](#typical-backend-service-design)
  - [Writing application in go](#writing-application-in-go)
    - [Package layout](#package-layout)
    - [Testing](#testing)
      - [Unit tests](#unit-tests)
      - [Integration tests](#integration-tests)
    - [Common functionalities in backend services](#common-functionalities-in-backend-services)
      - [Logging](#logging)
      - [Caching](#caching)
      - [Instrumentation](#instrumentation)
    - [Other high level concepts of go programming](#other-high-level-concepts-of-go-programming)
      - [Style, linters. Optimize for reading, not for writing.](#style-linters-optimize-for-reading-not-for-writing)
      - [Error handling](#error-handling)
      - [Context](#context)
      - [Overusing language features](#overusing-language-features)
      - [Always optimize code for better performance!](#always-optimize-code-for-better-performance)
  - [Links to other guides](#links-to-other-guides)
    - [High abstraction level](#high-abstraction-level)
    - [Medium abstraction level](#medium-abstraction-level)
    - [Low abstraction level](#low-abstraction-level)

## What is this repository?

This repository contains example of fully working application. App exposes GRPC and REST APIs, that implements some imaginary business requirements.

The topic of the project is a **Bike rental service backend**.

The purpose of this project is to:

- Show how to structure medium to big go projects.
- Explain some high level concepts of go programming, such as organizing packages, error handling, passing context, etc.
- Explain how to embrace good design principles in a project, such as clean architecture and SOLID principles. 

TODO

 1. How to read this document and the code.
    1. Readme files in important packages to explain some concepts.
 2. Problem statement
    1. Go is great, but there is no single framework forcing the project structure
    2. Many great articles on software architecture, not many good full guides for typical backend services
       1. Ben Johnson has similar example project and a list of great blog posts: https://www.gobeyond.dev. Explain what is different here. Anyway this work will be heavily inspired by his work.
       2. Three Dots Labs has similar project with series of blog posts: https://threedots.tech/. Blog posts are great, and there will be similarities with this work.
    3. Goals
       1. Guide for developers coming into go from other languages
       2. Reusable code structure, that will be familiar across multiple services developed within a company (this is the main difference from Ben Johnson's approach)
       3. Hints and tips about common problems (logging, caching, metrics, terminating goroutines, etc)
 3. Benefits of presented code structure
 4. When presented code structure doesn't apply?
    1. Libraries (Explanation)
    2. Very small services (in general, how to collapse packages from this example into "bigger" chunks, up to single `main.go` file)
    3. Very big services (not enough experience to judge)
    4. Other kind of services? Any ideas?

## Example project

### Business requirements

Requirements are described in separate document.

[Business requirements](/docs/businessrequirements/requirements.md)

### Design

Design for this example project is described in separate document. Please read it first before browsing any code.

[System design documentation](/docs/systemdesign/systemdesign.md)

### Project structure 

TODO, mention https://github.com/golang-standards/project-layout

## Typical backend service design

TODO

1. Break down of typical backend service
2. Intro to clean architecture
   1. https://herbertograca.com/2017/07/03/the-software-architecture-chronicles/ - we build on this!
3. Typical backend service organized in layers
4. Architecture for example project

## Writing application in go

### Package layout

TODO

1. What is a package? (not a directory!)
2. Problem with circular dependencies
3. Package hierarchy explained using go standard library (example: `net` to `net/http`, `crypto` to `crypto/md5`, `encoding` to `encoding/json` etc.)
   1. Explain packages as layers (`http` builds on top of `net`!)
      1. Important `net` can never import it's child packages! Explain why.
   2. How it solves problem with circular dependencies?
4. Example project packages breakdown (package layers diagram)
   1. Diagram with all packages stacked on top of each other
      1. Show dependency direction
      2. Show control flow direction
5. Sources
   1. https://www.gobeyond.dev/standard-package-layout/
   2. ...and continuation: https://www.gobeyond.dev/packages-as-layers/


### Testing

TODO: **need help here!**

#### Unit tests

TODO
1. When
2. How


#### Integration tests

TODO

1. How our architecture helps with tests
2. When
3. How

### Common functionalities in backend services

#### Logging

TODO

1. What does "log" mean?
   1. Common misconception: this is not the same as output in your terminal
      1. Unless there are prepared infrastructure to create structured logs, each log is just one line in text
      2. These lines of text are usually collected by some aggregator from multiple running instances
      3. If one instance logs 3 lines, those lines will often be spread across other lines from other instances
   2. Conclusion: one log should contain all the information about an event
      1. Don't log like: "function started", "function ended". Result aggregated from all running instances will be useless. 
2. Standard error logging
   1. https://blog.golang.org/go1.13-errors
3. Other logs
   1. What to log? (Actually more importantly what not to log)
      1. Incoming requests
      2. Outgoing requests
      3. System state changes
   2. How it relates to app layers
   3. Put log together into stories using trace id
      1. Later in microservice architecture - distributed transaction id

#### Caching

TODO

1. App or adapters? App of course and explain why.
2. How adding cache affects application logic (hint: it doesn't!)
#### Instrumentation

TODO

1. How it relates to app layers (similar to logging)


### Other high level concepts of go programming

TODO

#### Style, linters. Optimize for reading, not for writing.

TODO

#### Error handling

TODO
https://blog.golang.org/go1.13-errors

#### Context

TODO: Primarily for signalling end of execution to goroutines

#### Overusing language features

TODO

1. Channels: use mutex whenever it makes things simple
2. Named returns: exception, not a rule

#### Always optimize code for better performance!

Just kidding, don't do that. Optimize for reading, care about fellow coworkers, not CPU cycles.


## Links to other guides 

TODO: **need more links**

### High abstraction level

1.  https://www.gobeyond.dev/ - example repository and a series of blog posts.
2.  https://threedots.tech/ - example repository and a series of blog posts.
### Medium abstraction level

1.  https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html

### Low abstraction level

1.  https://github.com/golang/go/wiki/CodeReviewComments
