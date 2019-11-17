# Reprint

*Reprint is a Go library to deep copy any object THE RIGHT WAY TM*

[![reprint](https://github.com/qdm12/reprint/raw/master/title.png)](https://hub.docker.com/r/qmcgaw/REPONAME_DOCKER)

[![Build Status](https://travis-ci.org/qdm12/reprint.svg?branch=master)](https://travis-ci.org/qdm12/reprint)

[![GitHub last commit](https://img.shields.io/github/last-commit/qdm12/reprint.svg)](https://github.com/qdm12/reprint/issues)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/qdm12/reprint.svg)](https://github.com/qdm12/reprint/issues)
[![GitHub issues](https://img.shields.io/github/issues/qdm12/reprint.svg)](https://github.com/qdm12/reprint/issues)

## Setup

```sh
go get -u github.com/qdm12/reprint
```

## Usage


## Development

1. Install [Docker](https://docs.docker.com/install/)
    - On Windows, share a drive with Docker Desktop and have the project on that partition
    - On OSX, share your project directory with Docker Desktop
1. With [Visual Studio Code](https://code.visualstudio.com/download), install the [remote containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
1. In Visual Studio Code, press on `F1` and select `Remote-Containers: Open Folder in Container...`
1. Your dev environment is ready to go!... and it's running in a container :+1:

## TODOs

- Write usage readme
- Finish `FromTo` corner cases (nil pointers etc.)
- `forceCopyValue` might not be needed, test with func, channels etc.
- Verify no dereferencing is needed for types: Func, Chan, UintPtr, UnsafePointer, Array
