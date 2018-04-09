# Bring Your Own Log (BYOLOG)

This package aims to provide a useful abstraction over various logging libraries.

## Godoc

https://godoc.org/github.com/gordonmleigh/byolog

## Why?

1. Logging should be simple
2. A package dependency shouldn't force you to use a certain logging library

The external dependencies of a given library should be replaceable.  A package depending on `byolog.Logger` can use any logging library that an implementation has been written for (see [output](./output/)).


## What loggers can I use?

See [output/ folder](output/).

## Contributions

PRs welcome for broader output support.
