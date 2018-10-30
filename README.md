# YouVersion VOTD CLI

A command line interface (CLI) to fetch the current verse of the day from [YouVersion](https://www.youversion.com) public votd [API](https://developers.youversion.com)

The YouVersion VOTD CLI has been developed using the [Go programming language](https://golang.org) and currently requires to be installed from source. If go has been installed than than installation is as easy as:

```bash
go get -u github.com/jyksnw/yv-votd
go install github.com/jyksnw/yv-votd
```

There are two environment variable that must be set prior to running the YouVersion VOTD CLI:

| Variable  |    |
|---|---|
| YOUVERSION_VOTD_TOKEN   |   Your YouVersion Developer Token|
| YOUVERSION_VOTD_VERSION |   The YouVersion Bible Version ID|

A YouVersion Developer Token can be obtained by creating an account on the [YouVersion Developer Portal](https://developers.youversion.com)