# Go developers incomplete guide to writing typical backend service

This repository contains example of fully working application. App exposes GRPC and REST APIs, that implements some imaginary business requirements.

The topic of the project is a **Bike rental service backend**.

## Plan

1. What is this repository?
   1. How to read this document and the code.
      1. Readme files in important packages to explain some concepts.
   2. Problem statement
      1. Go is great, but there is no single framework forcing the project structure
      2. Many great articles on software architecture, no full guides for typical backend services (?)
      3. Want to have a guide for developers coming into go from other languages
   3. When presented code structure doesn't apply?
      1. Libraries (Explanation)
      2. Very small services (in general, how to collapse packages from this example into "bigger" chunks, up to single `main.go` file)
      3. Very big services (not enough experience to judge)
      4. Other kind of services? Any ideas?
2. Example project
   1. Business requirements made up for this repository
   2. High level repository layout
      1. https://github.com/golang-standards/project-layout/tree/master/internal
3. Application architecture
   1. Break down of typical backend service
   2. Intro to clean architecture
      1. https://herbertograca.com/2017/07/03/the-software-architecture-chronicles/ - we build on this!
   3. Typical backend service organized in layers
   4. Architecture for example project
4. Package layout
   1. What is a package? (not a directory!)
   2. Problem with circular dependencies
   3. Package hierarchy explained using go standard library (example: `net` to `net/http`, `crypto` to `crypto/md5`)
      1. Explain packages as layers (`http` builds on top of `net`!)
         1. Important `net` can never import it's child packages! Explain why.
      2. How it solves problem with circular dependencies?
   4. Example project packages breakdown (package layers diagram)
      1. Diagram with all packages stacked on top of each other
         1. Show dependency direction
         2. Show control flow direction
   5. Sources
      1. https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
      2. ...and continuation: https://www.gobeyond.dev/packages-as-layers/
5. Testing (**need help here!**)  
   1. Integration tests
      1. How our architecure helps with tests
6. Caching
   1. App or adapters? App of course and explain why.
   2. How adding cache affects application logic (hint: it doesn't!)
7. Logging
   1. What does "log" mean?
      1. Common misconception: this is not the same as output in your terminal
         1. Unless there are prepared infrastructure to create structured logs, each log is just one line in text
         2. These lines of text are usually collected by some aggregator from multiple running instances
         3. If one instance logs 3 lines, those lines will often be spread across other lines from other instances
      2. Conclusion: one log should contain all the information about an event
         1. Don't log like: "function started", "function ended". Result aggregated from all running instances will be useless. 
   2. Standard error logging
   3. Other logs
      1. What to log? (Actually more importantly what not to log)
         1. Incoming requests
         2. Outgoing requests
         3. System state changes
      2. How it relates to app layers
      3. Put log together into stories using request id
         1. Later in microservice architecure - distributed transaction id
8. Instrumentation
   1. How it relates to app layers (similar to logging)
9.  Other high level concepts of go programming
   2. Style, linters. Optimize for reading, not for writing.
   3. Error handling
   4. Context
      1. Primarily for signalling end of execution to goroutines
   5. Overusing language features
      1. Channels: use mutex whenever it makes things simple
      2. Named returns: exception, not a rule
10. Links to other guides (**need more links**)
    1.  "Higher" level
        1.  https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html
    2.  "Lower" level
        1.  https://github.com/golang/go/wiki/CodeReviewComments

## What this is repository?

The purpose of this project is to:

- Show how to structure medium to big go projects.
- Explain some high level concepts of go programming, such as organizing packages, error handling, passing context, etc.
- Explain how to embrace good design principles in a project, such as clean architecture and SOLID principles. 

## Business requirements

Requirements are described in separate document.

[Business requirements](/docs/businessrequirements/requirements.md)
## System design

System design is described in separate document. Please read it before browsing any code.

[System design documentation](/docs/systemdesign/systemdesign.md)
## Project layout

TODO, mention https://github.com/golang-standards/project-layout
