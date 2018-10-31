# Verse of the day (VOTD) Command line app

A command line interface (CLI) to fetch the current verse of the day from [YouVersion](https://www.youversion.com) public votd [API](https://developers.youversion.com)

The app has been developed using the [Go programming language](https://golang.org) and currently requires to be installed from source. If [Go](https://golang.org/doc/install#install) has been installed than installation is as easy as running:

```bash
go get -u github.com/jyksnw/yv-votd
go install github.com/jyksnw/yv-votd
```

## Environment Variables

There are two environment variables that may be set prior to running the application:

| Variable  | Required  | Description  |
|---|:---:|---|
| YOUVERSION_VOTD_TOKEN   | ✅ |  Your YouVersion Developer Token|
| YOUVERSION_VOTD_VERSION | | The YouVersion Bible Version ID|

A YouVersion Developer Token can be obtained by creating an account on the [YouVersion Developer Portal](https://developers.youversion.com)

The YouVersion Bible Version ID can be obtained by calling the [YouVersion Versions API](https://yv-public-api-docs.netlify.com/api/versions.html). If the `YOUVERSION_VOTD_VERSION` environment variable is not set than a default version_id of 1 will be used which maps to the KJV.

## Caching

The application caches each day's verse of the day in a dated file found in the `$GOPATH/bin/.votd` folder. For example, if the application was executed for the first time on 10/30/2018 the resulting response would be cached in the file `$GOPATH/bin/.votd/20181030`. Removing this file or this folder will result in the application to re-fetch the current verse of the day.

## Example Usage

The application could be used to print the current verse of the day to any new terminal windows by setting up your `.bashrc` or `.bash_profile` with the following (assuming that the `$GOPATH/bin` directory is on your `$PATH`):

```bash
export YOUVERSION_VOTD_TOKEN={your_developer_token}
export YOUVERSION_VOTD_VERSION=1 #1=KJV, 12=ASV, 206=WEB
yv-votd | cowsay -f stegosaurus
```

![Stegosaurus verse of the day](images/yv_votd_stegasaurus.png)

## TODO

- [ ] Add support to pass in command line arguments
- [ ] Fail back to prior votd on error
- [ ] Add support to fetch supported versions and their Id's
- [ ] Add support for converting VOTD image to ASCII art
- [ ] Cache the votd based on the date and version_id (currently the cache ignores the version id and will load the last successfully cached votd)
