# Reprint

*Reprint is a Go library to deep copy any object THE RIGHT WAY :tm:*

[![reprint](https://github.com/qdm12/reprint/raw/master/title.png)](https://github.com/qdm12/reprint)

[![Join Slack channel](https://img.shields.io/badge/slack-@qdm12-yellow.svg?logo=slack)](https://join.slack.com/t/qdm12/shared_invite/enQtOTE0NjcxNTM1ODc5LTYyZmVlOTM3MGI4ZWU0YmJkMjUxNmQ4ODQ2OTAwYzMxMTlhY2Q1MWQyOWUyNjc2ODliNjFjMDUxNWNmNzk5MDk)
[![Build Status](https://travis-ci.org/qdm12/reprint.svg?branch=master)](https://travis-ci.org/qdm12/reprint)

[![GitHub last commit](https://img.shields.io/github/last-commit/qdm12/reprint.svg)](https://github.com/qdm12/reprint/issues)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/qdm12/reprint.svg)](https://github.com/qdm12/reprint/issues)
[![GitHub issues](https://img.shields.io/github/issues/qdm12/reprint.svg)](https://github.com/qdm12/reprint/issues)

## Features

Unlike most libraries out there, this one deep copies by assigning new pointers to all data structures
nested in a given object, hence doing it **THE RIGHT WAY :tm:**

It works with slices, arrays, maps, pointers, nested pointers, nils, structs (with unexported fields too), functions  and channels.

*Limits*:

- Functions pointers are not changed but that's by design
- Channels buffered elements are not deep copied

## Usage

```sh
go get -u github.com/qdm12/reprint
```

You can check out [Golang Playground](https://play.golang.org/p/ukbIYl_gLqu) and activate **Imports** at the top, or read this:

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
    reprint.FromTo(&x, &y) // you can check the error returned also
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

## Development

1. Install [Docker](https://docs.docker.com/install/)
    - On Windows, share a drive with Docker Desktop and have the project on that partition
    - On OSX, share your project directory with Docker Desktop
1. With [Visual Studio Code](https://code.visualstudio.com/download), install the [remote containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
1. In Visual Studio Code, press on `F1` and select `Remote-Containers: Open Folder in Container...`
1. Your dev environment is ready to go!... and it's running in a container :+1:

## TODOs

- (Research) deep copy elements currently in channel
    - Race conditions
    - Pause channel?
- Polish `FromTo`
    - Initialize copy to copy's type if it's a typed nil
    - Initialize copy to original's type if it's an untyped nil
    - Returns typed nil instead of untyped nil if original is a nil pointer (typed)
- `forceCopyValue` might not be needed
