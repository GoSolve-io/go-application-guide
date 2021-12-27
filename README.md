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

Go is a great language. It's simple, easy to learn, and the code is straightforward. You can write a simple application in just `main.go`. But when you want to write a bigger project, there isn't any single guide or framework that can tell you exactly how to organize it. All the projects are different. Some of them are great but, usually, programmers struggle with this freedom. There are many examples of "transplanting" code pieces from other languages/frameworks into go projects (`models` package!).

There are many great articles on how to write good go code, but there aren't that many sources explaining how to put all the good stuff together. One of the sources we recommend is Ben Johnson's blog: https://www.gobeyond.dev. He also uses a repository with an example code and has multiple posts that are worth reading. But this guide is going to be a little different. The other good source with series of blog posts is https://threedots.tech - check this as well.

### Guide goal

The goal of this guide is to explain all the important parts of a typical Go project. And to show in that context how to design and write readable and maintainable code, also explaining some topics specific to Go.
Later we'll also explain some problems common to Go projects (logging, caching, metrics, terminating goroutines, etc.).
This guide will hopefully be useful for experienced programmers switching from other languages.

The structure of code presented in this repository is designed to be flexible to use in various projects. The idea is that you can copy it, replace some application logic, customize adapters (explained later), and then you have a new project with a familiar structure and (hopefully) good design.

But there's one caveat: remember that this example application is over-engineered. It's done on purpose, to show some concepts. In real life, in a similar application, you can merge some packages for simplicity.

TODO: Point here to the chapter about simplification.

### Will this example application always work for me?

This project structure is designed for medium to large-size applications. It's not a good idea to apply all the concepts and packages for:

1. Libraries
   These are just different; We're not going to cover library design in this guide.
2. Very small applications
   If you just want to write `hello word` service, or you don't care about testing that much, or you want to write simple POC for some quick demo - don't copy this project. Later, in this guide, you'll find some tips for how to collapse some packages from this example to make things simpler.
3. Very big projects
   The author just lacks the experience to tell how does this guide relates to complex code bases.

### How to read this repository

Start with this README file. Read it up to the chapter explaining example project design. After that point you can:

- browse and run the code,
- continue reading chapter by chapter,
- or pick any chapter you want - order is not relevant

You'll find multiple `README.md` files in this repository. They contain explanations for some concepts in code. We recommend you to check them as well!

### Repository structure

This repository's structure closely follows [github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout) guide for organizing project in top-level directories. I strongly suggest using it in your project. It has few advantages:

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

### Caching

TODO

1. App or adapters? App, of course! Explain why.
2. How adding cache affects application logic (hint: it doesn't!)

### Instrumentation

TODO

1. How it relates to app layers (similar to logging)

Instrumentation allows to enhance the developer's visibility on actual application performance and behavior. It consists of several elements:
- logging
- tracing
- metrics

It is important to understand one thing: instrumentation is not only supposed to support developers. It is also supposed to provide vital information to the operators of the application. Even for smaller apps that are supposed to serve content to small number of clients it's still worth to implement it.

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

#### Tracing

Over past several years the concept of distributed tracing became one of the most important buzzwords. Go offers a lot of ready to use libraries that will allow to use the tracing to its extents, but the very basic tracing is simpler than it might sound. 

For basic distributed tracing a unique key, shared between all application's modules, has to be logged. This approach can be seen in [this](https://github.com/GoSolve-io/go-application-guide/blob/master/internal/transport/grpc/httpgateway/middleware.go#L14) example. This will not produce fancy diagrams or maps, but it's a good starting point.

It is also possible to use paid service providers, like [DataDog](https://www.datadoghq.com/), that allows to not only trace each request inside the application, but also provides ready to use libraries for other languages, like Flutter, Java or iOS, so it is possible to monitor every step of the process. This kind of services generally provides more details and are capable of creating flow charts based on the tracing informations. 

#### Metrics

Metrics can be used to monitor the performance and error rate of the application. By simply counting the number and duration of the requests it is possible to predict issues that will come. This is a perfect tool to help with planning the future of the application when the number of clients will start to grow. 

## Other high-level concepts of go programming

TODO

### Style and linters. Optimize for reading, not for writing

TODO

### Error handling

Go is handling errors in probably the most reasonable way: **it doesn't**. It gives the developer the ability to handle
them based on the developer's requirements. For years the most popular languages like JavaScript or Python tried to move
away the burden of error handling from the developer. This resulted in no errors being handled at all and overusing the
try...catch blocks.

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

In some cases a simple error string is not enough. To provide more context on the error it is possible to use custom
error types. As mentioned before an error is each type that implements the error interface.

To use a custom type as an error simply add `Error() string` method to it:

```go
type CustomError struct {
    Details string
}

func (c CustomError) Error() string {
    return fmt.Sprintf("details: %s", c.Details)
}
```

Popular use case is to define custom type for validation errors and use them to give more details on what validation has
failed.

Some libraries expose their custom error types as well. Good examples
are [pq.Error](https://github.com/lib/pq/blob/master/error.go#L25)
and [MySQL error](https://github.com/go-sql-driver/mysql/blob/master/errors.go#L58). Both provides similar
functionality: expose the internal database error code, so it can be handled better in the application code. Remember to
only use this error types in the database adapters to not break the SOLID.

#### Checking error type

Thanks to improvements introduced in Go 1.13 it is now much easier to work with custom errors. Two new methods has been
added: **Is** and **As**.

**Is** checks if the error, or any error it is wrapping, is of the specified type, while **As** not only checks if the
error is of given type, but also fills provided structure with the actual content if the types match.

A sample use case of the **As** method is checking the details of the above error type:

```go
var customError CustomError

err := DoSomething() // 
if errors.As(err, &customError) {
    // do something with the customError.Details field
}
```

Both **errors.Is** and **errors.As** methods are better than comparing the errors using `==` because they unwrap the
error if possible.

#### Handling unexpected panics

The last thing worth to mention in the terms of error handling are panics. If the program faces an issue that was not
expected at all - e.g. accessing a nil pointer or using an index outside the array boundaries - it will panic. A panic
is a special case that will immediately stop the current operation and will propagate to the caller. If the caller is
not prepared to handle the panic it will panic as well and so on, up to the main method, causing the whole program to
stop. This is as bad as it sounds: if the application is an HTTP server it will not inform the client of the panic and
the client will simply wait for the server to respond.

Fortunately panic is not affecting the deferred methods. This allows the developer to use **recover** method. **
Recovery** catches the panic, returns the initial error and allows the method to return in an expected way. This further
allows the application to handle error as usual.

Handling a panic is very easy, but should not be overused. It resembles the try...catch method of other languages, but
just because something work in other languages should not make it being the preferred way of doing this in Go.

```go
func SomeMethod() (err error) {
	defer func () {
		if something := recover(); something != nil {
			// do something with something: it might be an error or something else
			switch e := something.(type) {
			case error:
				err = e
			case string:
				err = fmt.Errorf("my error: %s", e)
			default:
				err = fmt.Errorf("something strange happened: %v", something)
			}
		}
	}()

	panic("this will be a panic")
}
```

This code snippet can be used within the server's middleware as well. This way each failed request will at least let the
client know something bad happened.

#### Anti-patterns

Because handling errors in Go produces some boilerplate code ignoring errors became rather popular solution. This is
definitely not a good approach. If you don't know how to handle error simply wrap it and return to the caller. If the
error has been returned it means the application encountered unexpected issue, and writing software **is** handling
unexpected cases. Otherwise, there would be no `if..else`.

Use the Go feature to return multiple values instead of defaulting to some default or empty value in case of an error.
This is often used in all kinds of creators, but doesn't play well with the logic of an application. If the creation
failed because of an invalid parameter it should return a validation error instead of a blank object.

Something that has been mentioned earlier: handle the errors with SOLID in mind. There's no need for the caller to know
how the called method is implemented, so instead of returning the raw error map it to something the caller can
understand better.

Another bad practice when handling errors is to return the errors as-is, without wrapping them first. Imagine a case
where an aggregate - that's supposed to pull data from different sources - returns *Not Found* error. What exactly is
missing? Is it missing all pieces, or only one of them? Wrapping errors definitely helps to trace down the issue, and
using **Is** and **As** methods still allows to handle them accordingly.

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

It is worth to mention that using context is concurrency-safe, so calling `Err` or using the `Done()` channel is safe
within multiple goroutines.

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

#### Context inheritance

Some context operations returns a new context: a child of the parent context. This is particularly useful in cases where
some part of the task should be guaranteed to finish before others, but with one caveat: if the parent context will be
canceled, all child contexts will be canceled as well.

It is possible to avoid cascade failures using this feature, though. By using context's `Deadline` method it is possible
to obtain the deadline of the parent context, and set child's context deadline slightly smaller than the parent one:

```go
parentDeadline, hasDeadline := parentCtx.Deadline()
if !hasDeadline {
    parentDeadline = time.Now().Add(5 * time.Second)
}

childContext, childCancel := context.WithDeadline(parentCtx, parentDeadline.Sub(time.Second))
```

This way it is possible to ensure the called methods will always finish before the parent context will time out.

### Overusing language features

Even if a language is great it can quickly become a headache if not used properly. That's why this document is supposed
to promote a healthy use of Go's features. Remember to stick to simple code where possible and keep the advanced
technologies for later, when they are really useful.

#### Channels: use mutex whenever it makes things simple

Channels are great. Working with streams of data never has been easier. But this doesn't mean to use them all the time,
in every single place where a goroutine is in use. Equally good, and in much simpler way, results can be achieved by
using `sync.Mutex`. When working with a single object or small set it's simply safer to manage the access to the data
using locking mechanism.

#### Named returns: exception, not a rule

Using named results can be helpful when we plan to defer the recovery after panic. Giving names to return values might
reduce the number of lines of code by 1, but might raise its complexity a lot in comparison to simply defining a
variable in the method's body. Also, it is worth to mention named value is not magically receiving a valid value: it is
still required to initialize the variable with whatever is needed, otherwise it will be `nil`. And this can lead to
unexpected failures if the initialization will be skipped, e.g. setting attribute of a `nil` struct will cause panic.

#### Adding method parameters to context values

Context is great at passing request-scope values, but it is important not to clutter it with some random data. Put the
values where they belong: if the variable is supposed to be used by some middleware or instrumentalization, then context
might be the place where it should live. If a variable is only needed in one or two methods that belong to the
application domain, then it's definitely better to keep the dependency graph clean and just add the variable to the
method's definition or some parameter struct.

#### Using panics as a substitute for try...catch

Go's lack of try...catch and the existence of panics might cause some developer to think it's a good opportunity to use
recover in place of catch. This is wrong on many levels:

- Go developer is supposed to write easy to read code
- panic can be handled in completely different place and figuring it will be really hard

Treating errors the way they should be treated - as values - is much easier to read and follow.

### Always optimize code for better performance!

Just kidding, don't do that. Optimize for reading; care more about your coworkers than CPU cycles.

## Links to other guides

TODO: **need more links**
TODO: How to make this section short and to the point? We don't want 100+ links here.

Nice talk about error handling: https://www.youtube.com/watch?v=IKoSsJFdRtI

### High abstraction level

1. https://www.gobeyond.dev/ - example repository and a series of blog posts.
2. https://threedots.tech/ - example repository and a series of blog posts.

### Medium abstraction level

1. https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html

### Low abstraction level

1. https://github.com/golang/go/wiki/CodeReviewComments
