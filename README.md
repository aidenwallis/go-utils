# go-utils

[![codecov](https://codecov.io/gh/aidenwallis/go-utils/branch/main/graph/badge.svg?token=AT2T41NQ7K)](https://codecov.io/gh/aidenwallis/go-utils) [![Go Reference](https://pkg.go.dev/badge/github.com/aidenwallis/go-utils.svg)](https://pkg.go.dev/github.com/aidenwallis/go-ratelimiting)

A set of common Go utils I use throughout my projects, such as [Fossabot](https://fossabot.com).

I originally created this library to help me break down my large project monorepos into smaller repos, and share more code between them, but decided to open source it, in case some of these utils are useful to others.

A lot of these utils use [generics](https://go.dev/doc/tutorial/generics) - a new feature introduced in Go 1.18, therefore, this package is only supported on Go 1.18+.