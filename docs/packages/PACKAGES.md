# Package Layout

## Problem with cyclic dependencies

In every language that has something like a package, and one of the packages depends on another, there's always a possibility of having a cyclic dependency. Cyclic dependency occurs when `Package A` want's to use some features of `Package B`, but `Package B` also wants some features from `Package A`. Some languages (Python, NodeJS) allow cyclic imports and let the programmer live with the consequences :) And there are a few[^1]:

- Circular dependencies can cause many unwanted effects in software programs. Most problematic from a software design point of view is the tight coupling of the mutually dependent modules, which reduces or makes impossible the separate re-use of a single module.
- Circular dependencies can cause a domino effect when a small local change in one module spreads into other modules and has unwanted global effects (program errors, compile errors). Circular dependencies can also result in infinite recursions or other unexpected failures.
- Circular dependencies may also cause memory leaks by preventing certain very primitive automatic garbage collectors (those that use reference counting) from deallocating unused objects.

Go language is strict about that: the compiler will forbid you to add cyclic imports. So we're forced to have a "proper" package design. But how to do that?

## A package isn't a directory

A common misconception about organizing packages in Go results from treating the package like a directory. Sometimes these "directories" are used to put together a bunch of things with a common name, like "models" or "validators". But this approach will bite you eventually. Sooner or later this will lead to a cyclic dependency problem. And the usual solution is to create another package, like "utils". This "fix" will work for some time, but later another conflict will force you to create another "extracted" package. And finally, you'll end up with a project with accidental structure.

The solution is actually simple. A package is not a group. It's a layer!

I recommend you [this](https://www.gobeyond.dev/standard-package-layout/) and [this](https://www.gobeyond.dev/packages-as-layers/) articles from Ben Johnson's blog, explaining the concept. But for now, let's move on.


## How go standard library does it?

Let's cover some packages from go's standard library. How do Go core developers organize packages that are related?

We can identify key patterns for package organization common in Go source code:

1. Some packages "build" on their parent. Parent package with lower abstraction code + child package with higher abstraction code (`net/http` => `net`).
2. Some packages closely related are grouped together (`text/template`, `text/scanner` => `text`).
3. Some packages depend on other packages within the same "group" (`image/jpeg` imports `image/color`).
4. Some packages depend on other packages outside of their "part of a tree" (`image/jpeg` imports `io`).

Go's standard library also has one crucial feature. Packages never import their children [^2]. This property is the key to have cyclic free package imports!

I find the first point the most interesting. Take a look at the `net/http` package. The core problem solved here is to "provide a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets" (quote from the Go's docs). So `net/http` package defines multiple types and functions that allow you to communicate using HTTP protocol. But this package has to abstract all the connection details away. It has to operate at a higher level of abstraction, using types like `Listener`, `Conn`, `Dialer`. So it "sits" on top of a package that defines these types: `net`. Take a note that all other packages in the standard library that build on those network types also "sit" on top of the `net` package (`net/mail`, `net/smtp`, e.t.c.). So to summarize: `net` provides a layer that abstracts base network communication, and `net/*` sub-packages use that layer to build their specialized layers for communication using higher-level protocols. This approach beautifully reflects the network layering (`net` is a transport layer, `net/http` is an application layer).

Other similar examples are:

- `crypto/*` packages build on `crypto`.
- `hash/*` packages build on `hash`.
- `log/syslog` packages build on `log`.
- `image/jpeg` packages build on `image`. Here we have also "horizontal" dependency: it imports `image/color`!
- `encoding/json` package builds on `encoding`. This example is a bit more subtle, though. The `encoding/json` doesn't import `encoding` directly but rather implements its interfaces.


## Key takeaways

1. Organize packages by their function. Following SRP, a package has to have one functionality to provide.
2. If you have 2 or more packages that rely on a common base, organize packages layer-like: `base`, `base/package1`, `base/package2`. Remove dependencies beetween them by abstracting common types to `base` layer.
3. The "deeper" the package, the more specialized it should be.



[^1]: [wikipedia.org/wiki/Circular_dependency](https://en.wikipedia.org/wiki/Circular_dependency).

[^2]: There is at least one exception of course :) Package `database/sql` imports its child - `database/sql/driver`.
