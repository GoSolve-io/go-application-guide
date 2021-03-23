# HTTP package

## Why this package exists, what problems does it solve?

TODO

Explain common utility layer for making http requests.

What are the alternatives? Maybe `internal/utilities/http`? I don't think so. We need http utilities **only** in context of adapters!
Adapters is the right place, why `weather` and `incidents` are inside `http`? Because they build on top of this package + `net/http`!

Things to note:
- closing response body is important while doing http request. Look at this code, there's only one place we do actual requests! So there's only one place to check if we do them properly!

## Why `Doer` interface?

TODO

Explain in context of testing.
