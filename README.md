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

TODO

https://blog.golang.org/go1.13-errors

Nice talk: https://www.youtube.com/watch?v=IKoSsJFdRtI

### Context

TODO: Primarily for signaling end of execution to goroutines

Context is a well known concept from other programming languages. To provide some background for someone who never used
contexts before: it is a way to tie all the requests together and provide a way to stop the execution if needed. It is
sometimes used to carry request scope variables, but this should not be overused and will be explained in more details
later.

In Go, it is mostly used to signal end of execution to goroutines and when passing through domain boundaries.

#### Signaling end of execution

Thanks to Go being designed with the concurrency in mind it is common to use goroutines. Goroutines are small and fast
and benefits gained by using them definitely surpass drawbacks and risks.

Similar to how the errors should be always returned as the last parameter, the methods that are using context should
always accept it as the first parameter. Also, it makes sense to pass the context even if we don't plan to use it at
this moment.

There is one thing that might lead to issues: it is not clear when - and if - all goroutines have finished. Long-running
goroutines can lead to unexpected issues, like unexpectedly altering the object's state or trying to write to an already
closed channel. Leaving goroutines unattended can lead to resource leak and cause the system's health to degrade over
time.

That's where the **context** comes to the rescue. Thanks to its ability to signal end of execution we can manage the
goroutines to a great extent.

The first step is to create a context and decide how do we want to cancel it. There are at least two ways for doing
this:

- using **cancel** method
- defining **deadline**

The first use case is very simple and allows the developer to stop the execution at any time by calling the **cancel**
method returned from **WithCancel** method (other methods return the cancel as well, but let's focus on this one for
now):

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

It is important to always cancel the context, that's why deferring call to this method is the usual way of writing the
code.

Other solution is to define a deadline on the context, so it will get automatically canceled when the time comes. This
is often used within servers to make sure the caller will not wait forever for the request to be processed:

```go
ctx, cancel := context.WithDeadline(context.Background(), time.Date(2022, 12, 31, 0, 0, 0))
defer cancel()
```

There are two ways of setting the deadline: with the method mentioned above, or with `WithTimeout` that accepts the
parent context and a `time.Duration`, so it is possible to set the deadline e.g. 30 seconds in the future:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
defer cancel()
```

Now, once the context is ready to use, the goroutines (or any other receiver of the context) needs to simply check
the `Err` method or listen to the `Done()` channel. By default, as soon as the context gets canceled the `Err()` will
return `DeadlineExceeded` or `Canceled` error. It is a simple way to check if the execution should proceed:

```go
func DoSomething(ctx context.Context) error {
if err := ctx.Err(); err != nil {
// simply stop
return fmt.Errorf("do something: %w", err)
}
// do stuff
}
```

This should work fine in most cases, but what if this method is supposed to take some time? Calling `Err` on every loop
step will simply litter the code.

The solution for more complex tasks is to listen to the `Done()` channel of the context object. Once context gets
canceled or times out, the `Done()` channel is closed so all blocked reads will be unlocked. This is the best solution
when using channels:

// TODO: provide better example here

```go
func Parent() {
ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
defer cancel()

ch := make(chan int)

go Child(ctx, ch)

for v := range ch {
fmt.Println(v)
}
}

func Child(ctx context.Context, ch chan<- int) {
for {
select {
case <-time.After(time.Second):
ch<-rand.Int()
case <-ctx.Done():
close(ch)
return
}
}
}
```

#### Passing request-scoped values

Another use case for context is to pass request-scoped data across multiple layers and domains. This is often used
together with tracing, logging and authorization middlewares.

To add a request-scoped value to the context simply use `context.WithValue` method, retrieve it using the
context's `Value`:

```go
func Parent(ctx context.Context) {
valuedContext := context.WithValue(ctx, contextKey, contextValue)

Child(valuedContext)
}

func Child(ctx context.Context) {
fmt.Println(ctx.Value(contextKey))
}
```

This feature can be used to add trace ID to the context, as shown in the
example [code](https://github.com/GoSolve-io/go-application-guide/blob/master/internal/transport/grpc/httpgateway/middleware.go#L14)
and used later with [logger](https://github.com/GoSolve-io/go-application-guide/blob/master/internal/app/log.go#L26).

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
