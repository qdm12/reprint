# Reprint

*Reprint is a Go library to deep copy any object THE RIGHT WAY :tm:*

[![reprint](https://github.com/qdm12/reprint/raw/master/title.png)](https://github.com/qdm12/reprint)

[![Join Slack channel](https://img.shields.io/badge/slack-@qdm12-yellow.svg?logo=slack)](https://join.slack.com/t/qdm12/shared_invite/enQtODMwMDQyMTAxMjY1LTU1YjE1MTVhNTBmNTViNzJiZmQwZWRmMDhhZjEyNjVhZGM4YmIxOTMxOTYzN2U0N2U2YjQ2MDk3YmYxN2NiNTc)
[![Build Status](https://travis-ci.org/qdm12/reprint.svg?branch=master)](https://travis-ci.org/qdm12/reprint)

[![GitHub last commit](https://img.shields.io/github/last-commit/qdm12/reprint.svg)](https://github.com/qdm12/reprint/issues)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/qdm12/reprint.svg)](https://github.com/qdm12/reprint/issues)
[![GitHub issues](https://img.shields.io/github/issues/qdm12/reprint.svg)](https://github.com/qdm12/reprint/issues)

## Features

Unlike most libraries out there, this one deep copies by assigning new pointers to all data structures
nested in a given object, hence doing it **THE RIGHT WAY :tm:**

It works with:

- slices
- maps
- pointers
- nested pointers
- structs (even with unexported fields)
- arrays
- functions (cannot change the pointer though)
- channels (does not deep copy elements IN the channel for now)

## Setup

```sh
go get -u github.com/qdm12/reprint
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/qdm12/reprint"
)

func main() {
    one := 1
    two := 2
    type myType struct{ A *int }

    // reprint.FromTo usage:
    var x, y myType
    x.A = &one
    reprint.FromTo(&x, &y)
    y.A = &two
    fmt.Println(x.A, *x.A) // 0xc0000a0010 1
    fmt.Println(y.A, *y.A) // 0xc0000a0018 2

    // reprint.This usage:
    x2 := myType{&one}
    out := reprint.This(x2)
    y2 := out.(myType)
    y2.A = &two
    fmt.Println(x2.A, *x2.A) // 0xc0000a0010 1
    fmt.Println(y2.A, *y2.A) // 0xc0000a0018 2
}
```

## Limits

- Does not support `UintPtr` and `UnsafePointer` types (untested)

## Development

1. Install [Docker](https://docs.docker.com/install/)
    - On Windows, share a drive with Docker Desktop and have the project on that partition
    - On OSX, share your project directory with Docker Desktop
1. With [Visual Studio Code](https://code.visualstudio.com/download), install the [remote containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
1. In Visual Studio Code, press on `F1` and select `Remote-Containers: Open Folder in Container...`
1. Your dev environment is ready to go!... and it's running in a container :+1:

## TODOs

- Verify it works for types:
    - [x] Func
    - [ ] Chan
    - [ ] Array
- Finish `FromTo` corner cases (nil pointers etc.)
- `forceCopyValue` might not be needed
