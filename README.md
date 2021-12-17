# Go developers incomplete guide to writing a typical backend service ![Build](https://github.com/nglogic/go-application-guide/workflows/Build/badge.svg)


Table of contents
- [Go developers incomplete guide to writing a typical backend service !Build](#go-developers-incomplete-guide-to-writing-a-typical-backend-service-)
  - [Intro](#intro)
    - [What is this repository?](#what-is-this-repository)
    - [Why this guide might be helpful to you](#why-this-guide-might-be-helpful-to-you)
    - [Guide goal](#guide-goal)
    - [Will this example application always work for me?](#will-this-example-application-always-work-for-me)
    - [How to read this repository](#how-to-read-this-repository)
    - [Repository structure](#repository-structure)
  - [Business requirements and initial design](#business-requirements-and-initial-design)
  - [Guide to Go application design](#guide-to-go-application-design)
  - [Guide to writing Go packages hierarchy](#guide-to-writing-go-packages-hierarchy)
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

### Why this guide might be helpful to you

Go is a great language. It's simple, easy to learn, and the code is straightforward. You can write a simple application in just `main.go`. But when you want to write a bigger project, there isn't any single guide or framework that can tell you exactly how to organize it. All of the projects are different. Some of them are great but, usually, programmers struggle with this freedom. There are many examples of "transplanting" code pieces from other languages/frameworks into go projects (`models` package!).

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

### Repository structure

This repository's structure closely follows [github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout) guide for organizing project in top-level directories. I strongly suggest using it in your project. It has few advanages:

- It's well known and broadly accepted standard in Go community.
- When you join a project following this guide, you can instantly feel familiar with the repository.
- On the other hand, when someone's joining your team, there's a high chance he knows this guide and will be more confident and productive faster.

## Business requirements and initial design

Let's start with explanation of example project, that will be used to talk about other important stuff here. Have a look at the [business requirements and initial design](/docs/businessrequirements/REQUIREMENTS.md) for our demo app.

## Guide to Go application design

[Guide to Go application design](/docs/appdesign/DESIGN.md)

## Guide to writing Go packages hierarchy

[Guide to Go packaging](/docs/packages/PACKAGES.md)

[Packages in example app](/docs/packages/APP.PACKAGES.md)

## Testing

TODO: **need help here, open for any discussion**

### Unit tests

TODO

1. When
2. How

### Integration tests

TODO

1. How our architecture helps with tests
2. When
3. How

## Common functionalities in backend services

### Logging

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

### Caching

TODO

1. App or adapters? App, of course! Explain why.
2. How adding cache affects application logic (hint: it doesn't!)

### Instrumentation

TODO

1. How it relates to app layers (similar to logging)


## Other high-level concepts of go programming

TODO

### Style and linters. Optimize for reading, not for writing

TODO

### Error handling

Go is handling errors in probably the most reasonable way: **it doesn't**. It gives the developer the ability to handle
them based on the developer's requirements. For years the most popular languages like JavaScript or Python tried to move
away the burden of error handling from the developer. This resulted in no errors being handled at all and overusing the
try..catch blocks.

Go proposed a different, yet old-fashioned, way to handle errors: **errors has been promoted to fist class citizens**.
This, together with multiple return values, allows to treat the error as part of the logic.

#### Error type

The error is an interface defined in the standard library like this:

```go
type error interface {
Error() string
}
```

The `errors` package provides the most basic implementation of `error` interface with its unexported `errorString` type.
Every time the `errors.New` method is used it simply returns a new `errorString`.

Thanks to error being defined as a simple interface it is possible to use other types in place of an error. We will come
back to this later.

#### Returning errors

The most common way to return an error is to return it as the **last return value**. Most IDE and linters will complain
if an error is returned between other return values:

```go
func SomeFunction(input1, input2 int) (output1, output2, error) { // ... }
```

Thanks to this standardisation it is easier to check error only if we don't care about the result of the method, by
using blank identifier:

```go
if _, _, err := SomeFunction(0, 1); err != nil { // ... }
```

The way Go methods returns the error attracted a lot of attention and criticism, because one of its drawbacks (probably
the only one) is the amount of boilerplate code it produces:

```go
something, err := DoSomething()
if err != nil {
// handle error
return err
}

somethingElse, err := DoSomethingElse()
if err != nil {
// handle error, again
return err
}
```

But let's be honest: what's wrong with handling errors explicitly? Did any of us expected the code to always work
flawlessly and never return an error? \
This approach encourages developer to actually **think** what could go wrong and **prepare** for it. It is not forcing
anyone to handle the error, as it is allowed to use the blank identifier for errors, too. This has some drawbacks, some
of them being the linters constantly complaining about ignored error and teammate's glancing hatefully.

#### Sentinel errors

Now ask ourselves a question: do we like to figure out what happened based on the returned string? No?

That's where the **sentinel errors** will help. A sentinel error is an exported, predefined error that can be used to
compare it against the returned error. Some well-known examples of sentinel errors are `os.ErrExists` or `sql.ErrNoRows`
. \
They're simply defined using standard `errors.New` function,
e.g. `var ErrNoRows = errors.New("sql: no rows in result set")` and thanks to being predefined they can be compared with
whatever will be returned from the underlying library.

It is expected for libraries to define and document sentinel errors. Should application define such errors as well? The
answer is: **yes**, especially when the application exposes an API.

Using sentinel errors inside the application's or library's code might help, but is not mandatory. Also predefined error
lacks one important feature: it carries no details on what actually happened.

#### Wrapping errors

One issue of sentinel errors is lack of accurate details of what went wrong. That's why Go gives the developer an
ability to wrap an error with some additional information. Logging or printing a wrapper error usually exposes a path
that can be used to pinpoint the source of the issue:

```go
var (
sentinelError = errors.New("this is some error")
)

func One() error {
return fmt.Errorf("One: %w", sentinelError)
}

func Two() error {
return fmt.Errorf("Two: %w", One())
}

func main() {
fmt.Println(Two()) // Two: One: this is some error
}
```

As of Go 1.13 the preferred way for wrapping an error is simply using the `fmt.Errorf` function with `%w` format
placeholder. `%w` works in a similar way as `%v`, but allows the original error to be recovered with `errors.Unwrap()`.
Unwrapping an error returns the previous error.

Because wrapping an error attaches the original error to the new one it is important to know when to stop. The general
rule is to only wrap errors if there is a way or plan to handle them in the caller and using `%v` when it is not
possible to handle it.

#### SOLID compliant errors

Wrapping an error is definitely helpful when it comes to debugging a failed request, but comes with a risk of exposing
the implementation details. This breaks the SOLID principles and causes the library or application to behave in
unexpected ways. In worst case scenario it can lead to serious issues on the caller's end.

Let's imagine our example app returns a database error. The client will receive an error (let's say it will
be `fmt.Errorf("bike not found %w", pq.ErrInFailedTransaction`). The client will accommodate for this error and will
display a nice message to the customer. At some point we will decide to change the internal database, but won't change
the code, so we will return `fmt.Errorf("bike not found", mysql.ErrInvalidConn)`. This will probably cause the client to
display an unexpected exception error.

That's why errors should be written with the SOLID principle in mind:

- should be defined and exported in the package that is supposed to be imported/used by the client
- should hide the implementation details

#### Custom error types

// TODO: When to use, and when to avoid custom error types. How to define them. Include grpc.Status example maybe?

#### Checking error type

// TODO: How to use the errors.Is and errors.As methods with our custom types.

#### Handling unexpected panics

// TODO: Last but not least: panics and deferring recovery as last resort; useful for dependencies we do not control.

#### Anti-patterns

// TODO: what should be avoided at all cost?

- returning _empty_ or _default_ values in case of an error
- returning errors that leaks the implementation details
- return errors without any trace of where they happened
- returning errors from underlying libraries (similar to point 2)

https://blog.golang.org/go1.13-errors

Nice talk: https://www.youtube.com/watch?v=IKoSsJFdRtI

### Context

TODO: Primarily for signaling end of execution to goroutines

### Overusing language features

TODO

1. Channels: use mutex whenever it makes things simple
2. Named returns: exception, not a rule

### Always optimize code for better performance!

Just kidding, don't do that. Optimize for reading; care more about your coworkers than CPU cycles.


## Links to other guides

TODO: **need more links**
TODO: How to make this section short and to the point? We don't want 100+ links here.

### High abstraction level

1.  https://www.gobeyond.dev/ - example repository and a series of blog posts.
2.  https://threedots.tech/ - example repository and a series of blog posts.
### Medium abstraction level

1.  https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html

### Low abstraction level

1.  https://github.com/golang/go/wiki/CodeReviewComments
