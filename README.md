# Go developers incomplete guide to writing typical backend service ![Build](https://github.com/nglogic/go-example-project/workflows/Build/badge.svg)


Table of contents
- [Go developers incomplete guide to writing typical backend service !Build](#go-developers-incomplete-guide-to-writing-typical-backend-service-)
  - [Intro](#intro)
    - [What is this repository?](#what-is-this-repository)
    - [Why this guide might be useful to you](#why-this-guide-might-be-useful-to-you)
    - [Guide goal](#guide-goal)
    - [Will this example application always work for me?](#will-this-example-application-always-work-for-me)
    - [How to read this repository](#how-to-read-this-repository)
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
    - [Other high-level concepts of go programming](#other-high-level-concepts-of-go-programming)
      - [Style and linters. Optimize for reading, not for writing](#style-and-linters-optimize-for-reading-not-for-writing)
      - [Error handling](#error-handling)
      - [Context](#context)
      - [Overusing language features](#overusing-language-features)
      - [Always optimize code for better performance!](#always-optimize-code-for-better-performance)
  - [Links to other guides](#links-to-other-guides)
    - [High abstraction level](#high-abstraction-level)
    - [Medium abstraction level](#medium-abstraction-level)
    - [Low abstraction level](#low-abstraction-level)

## Intro

### What is this repository?

This repository works as a guide explaining how to write the most common type of backend service in go. The guide consists of 2 complementary parts:

- A set of documents explaining various aspects of typical go backend service,
- Fully working codebase, implementing these documents in practice.

The topic of the example project is a **Bike rental service backend**.

The purpose of this project is to:

- Show how to structure medium to big go projects.
- Explain some high-level concepts of go programming, such as organizing packages, error handling, passing context, etc.
- Explain how to embrace good design principles in a project, such as clean architecture and SOLID principles.

### Why this guide might be useful to you

Go is a great language. It's simple, easy to learn and the code is straightforward. You can write a simple application in just `main.go`. But when you want to write a bigger project, there isn't any single guide or framework that can tell you exactly how to organize it. All of the projects are different. Some of them are great but, usually, programmers struggle with this freedom. There are many examples of "transplanting" code pieces from other languages/frameworks into go projects (`models` package!).

There are many great articles on how to write good go code, but there aren't that many sources explaining how to put all the good stuff together. One of the sources we recommend is Ben Johnson's blog: https://www.gobeyond.dev. He also uses a repository with an example code and has multiple posts that are worth reading. But this guide is going to be a little different. The other good source with series of blog posts is https://threedots.tech - check this as well.

### Guide goal

The goal of this guide is to explain all the important parts of a typical Go project. And to show in that context how to design and write readable and maintainable code, also explaining some topics specific to Go.
Later we'll also explain some problems common to Go projects (logging, caching, metrics, terminating goroutines, etc.).
This guide will hopefully be useful for experienced programmers switching from other languages.

The structure of code presented in this repository is designed to be flexible to use in various projects. The idea is that you can copy it, replace some application logic, customize adapters (explained later), and then you have a new project with a familiar structure and (hopefully) good design.

But there's one caveat: remember that this example application is overengineered. It's done on purpose, to show some concepts. In real life, in a similar application, you can merge some packages for simplicity.

TODO: Point here to the chapter about simplification.

### Will this example application always work for me?

This project structure is designed for medium to large-size applications. It's not a good idea to apply all the concepts and packages for:

1. Libraries
   Libraries are just different; We're not going to cover library design in this guide.
2. Very small applications
   If you just want to write `hello word` service, or you don't care about testing that much, or you want to write simple POC for some quick demo - don't copy this project. Later, in this guide, you'll find some tips for how to collapse some packages from this example to make things simpler.
3. Very big projects
   The author just lacks the experience to tell how does this guide relates to complex codebases.

### How to read this repository

Start with this README file. Read it up to the chapter explaining example project design. After that point you can:

- browse and run the code,
- continue reading chapter by chapter,
- or pick any chapter you want - order is not relevant

You'll find multiple `README.md` files in this repository. They contain explanations for some concepts in code. We recommend you to check them as well!

## Example project

### [Business requirements](/docs/businessrequirements/REQUIREMENTS.md)

### Design

The design for our example project is described in a separate document. Please read it first before browsing any code.

[System design documentation](/docs/systemdesign/SYSTEMDESIGN.md)

### Project structure

TODO, mention https://github.com/golang-standards/project-layout

## Typical backend service design

TODO

1. Break down of typical backend service
2. Intro to clean architecture
   1. https://herbertograca.com/2017/07/03/the-software-architecture-chronicles/ - we build on this!
3. Typical backend service organized in layers
4. Architecture for the example project

## Writing application in go

### Package layout

[Package layout documentation](/docs/packages/PACKAGES.md)

TODO

1. What is a package? (not a directory!)
2. Problem with circular dependencies
3. Package hierarchy explained using go standard library (example: `net` to `net/http`, `crypto` to `crypto/md5`, `encoding` to `encoding/json` etc.)
   1. Explain packages as layers (`http` builds on top of `net`!)
      1. Important `net` can never import its child packages! Explain why.
   2. How it solves the problem with circular dependencies?
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
      1. Unless there is a special infrastructure to create structured logs, each log is just one line in the app's output stream
      2. These lines of text are usually collected by some aggregator from multiple running instances
      3. If one instance logs 3 lines, those lines will often be spread across other lines from other instances
   2. Conclusion: one log should contain all the information about an event
      1. Don't log messages like "function started" or "function ended". The result aggregated from all running instances will be rather useless.
2. Standard error logging
   1. https://blog.golang.org/go1.13-errors
3. Other logs
   1. What to log? (Actually, more importantly, what not to log)
      1. Incoming requests
      2. Outgoing requests
      3. System state changes
   2. How it relates to app layers
   3. Put log together into stories using trace id
      1. Later in microservice architecture - distributed transaction ids

#### Caching

TODO

1. App or adapters? App, of course! Explain why.
2. How adding cache affects application logic (hint: it doesn't!)

#### Instrumentation

TODO

1. How it relates to app layers (similar to logging)


### Other high-level concepts of go programming

TODO

#### Style and linters. Optimize for reading, not for writing

TODO

#### Error handling

TODO

https://blog.golang.org/go1.13-errors

Nice talk: https://www.youtube.com/watch?v=IKoSsJFdRtI

#### Context

TODO: Primarily for signaling end of execution to goroutines

#### Overusing language features

TODO

1. Channels: use mutex whenever it makes things simple
2. Named returns: exception, not a rule

#### Always optimize code for better performance!

Just kidding, don't do that. Optimize for reading; care more about your coworkers than CPU cycles.


## Links to other guides

TODO: **need more links**

### High abstraction level

1.  https://www.gobeyond.dev/ - example repository and a series of blog posts.
2.  https://threedots.tech/ - example repository and a series of blog posts.
### Medium abstraction level

1.  https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html

### Low abstraction level

1.  https://github.com/golang/go/wiki/CodeReviewComments
