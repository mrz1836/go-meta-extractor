# go-meta-extractor
> Robust starter template for a new Go library

[![Release](https://img.shields.io/github/release-pre/mrz1836/go-meta-extractor.svg?logo=github&style=flat&v=1)](https://github.com/mrz1836/go-meta-extractor/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/mrz1836/go-meta-extractor/run-tests.yml?branch=master&logo=github&v=2)](https://github.com/mrz1836/go-meta-extractor/actions)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-meta-extractor?style=flat&v=1)](https://goreportcard.com/report/github.com/mrz1836/go-meta-extractor)
[![codecov](https://codecov.io/gh/mrz1836/go-meta-extractor/branch/master/graph/badge.svg?v=1)](https://codecov.io/gh/mrz1836/go-meta-extractor)
[![Mergify Status](https://img.shields.io/endpoint.svg?url=https://api.mergify.com/v1/badges/mrz1836/go-meta-extractor&style=flat&v=1)](https://mergify.io)
[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-meta-extractor?v=1)](https://golang.org/)
<br>
[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/mrz1836/go-meta-extractor)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Sponsor](https://img.shields.io/badge/sponsor-mrz1836-181717.svg?logo=github&style=flat&v=1)](https://github.com/sponsors/mrz1836)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat&v=1)](https://gobitcoinsv.com/#sponsor?utm_source=github&utm_medium=sponsor-link&utm_campaign=go-meta-extractor&utm_term=go-meta-extractor&utm_content=go-meta-extractor)

<br/>

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

<br/>

## Installation

**go-meta-extractor** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-meta-extractor
```

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-meta-extractor)

[![GoDoc](https://godoc.org/github.com/mrz1836/go-meta-extractor?status.svg&style=flat&v=1)](https://pkg.go.dev/github.com/mrz1836/go-meta-extractor)

<br/>

<details>
<summary><strong><code>Repository Features</code></strong></summary>
<br/>

This repository was created using [MrZ's `go-template`](https://github.com/mrz1836/go-template#about)

#### Built-in Features
- Continuous integration via [GitHub Actions](https://github.com/features/actions)
- Build automation via [Make](https://www.gnu.org/software/make)
- Dependency management using [Go Modules](https://github.com/golang/go/wiki/Modules)
- Code formatting using [gofumpt](https://github.com/mvdan/gofumpt) and linting with [golangci-lint](https://github.com/golangci/golangci-lint) and [yamllint](https://yamllint.readthedocs.io/en/stable/index.html)
- Unit testing with [testify](https://github.com/stretchr/testify), [race detector](https://blog.golang.org/race-detector), code coverage [HTML report](https://blog.golang.org/cover) and [Codecov report](https://codecov.io/)
- Releasing using [GoReleaser](https://github.com/goreleaser/goreleaser) on [new Tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging)
- Dependency scanning and updating thanks to [Dependabot](https://dependabot.com) and [Nancy](https://github.com/sonatype-nexus-community/nancy)
- Security code analysis using [CodeQL Action](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning)
- Automatic syndication to [pkg.go.dev](https://pkg.go.dev/) on every release
- Basic templates for [Issues and Pull Requests](https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests/configuring-issue-templates-for-your-repository) in GitHub
- Code [ownership configuration](.github/CODEOWNERS) for GitHub
- Automatic sync for [labels](.github/labels.yml) into GitHub using a pre-defined [configuration](.github/labels.yml)
- Built-in powerful merging rules using [Mergify](https://mergify.io/)
- Welcome [new contributors](.github/mergify.yml) on their first Pull-Request
- Follows the [standard-readme](https://github.com/RichardLitt/standard-readme/blob/master/spec.md) specification
- [Visual Studio Code](https://code.visualstudio.com) configuration with [Go](https://code.visualstudio.com/docs/languages/go)
- (Optional) [Slack](https://slack.com), [Discord](https://discord.com) or [Twitter](https://twitter.com) announcements on new GitHub Releases
- (Optional) Easily add [contributors](https://allcontributors.org/docs/en/bot/installation) in any Issue or Pull-Request

</details>

<details>
<summary><strong><code>Package Dependencies</code></strong></summary>
<br/>

- [stretchr/testify](https://github.com/stretchr/testify)
</details>

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

Releases are automatically created when you create a new [git tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging)!

If you want to manually make releases, please install GoReleaser:

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to GitHub and can be installed:
- **using make:** `make install-releaser`
- **using brew:** `brew install goreleaser`

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

<br/>

### Automatic Releases on Tag Creation (recommended)
Automatic releases via [GitHub Actions](.github/workflows/release.yml) from creating a new tag:
```shell
make tag version=1.2.3
```

<br/>

### Manual Releases (optional)
Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production (manually).

<br/>

</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                           Runs multiple commands
clean                         Remove previous builds and any cached data
clean-mods                    Remove all the Go mod cache
coverage                      Shows the test coverage
diff                          Show the git diff
generate                      Runs the go generate command in the base of the repo
godocs                        Sync the latest tag with GoDocs
help                          Show this help message
install                       Install the application
install-all-contributors      Installs all contributors locally
install-go                    Install the application (Using Native Go)
install-releaser              Install the GoReleaser application
lint                          Run the golangci-lint application (install if not found)
release                       Full production release (creates release in Github)
release                       Runs common.release then runs godocs
release-snap                  Test the full release (build binaries)
release-test                  Full production test release (everything except deploy)
replace-version               Replaces the version in HTML/JS (pre-deploy)
tag                           Generate a new tag and push (tag version=0.0.0)
tag-remove                    Remove a tag if found (tag-remove version=0.0.0)
tag-update                    Update an existing tag to current commit (tag-update version=0.0.0)
test                          Runs lint and ALL tests
test-ci                       Runs all tests via CI (exports coverage)
test-ci-no-race               Runs all tests via CI (no race) (exports coverage)
test-ci-short                 Runs unit tests via CI (exports coverage)
test-no-lint                  Runs just tests
test-short                    Runs vet, lint and tests (excludes integration tests)
test-unit                     Runs tests and outputs coverage
uninstall                     Uninstall the application (and remove files)
update-contributors           Regenerates the contributors html/list
update-linter                 Update the golangci-lint package (macOS only)
vet                           Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and [examples](examples) run via [GitHub Actions](https://github.com/mrz1836/go-meta-extractor/actions) and
uses [Go version 1.22.x](https://golang.org/doc/go1.22). View the [configuration file](.github/workflows/run-tests.yml).

<br/>

Run all tests (including integration tests)
```shell script
make test
```

<br/>

Run tests (excluding integration tests)
```shell script
make test-short
```

<br/>

## Benchmarks
Run the Go benchmarks:
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## Usage
Checkout all the [examples](examples)!

<br/>

## Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

<br/>

### How can I help?
All kinds of contributions are welcome :raised_hands:!
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:.
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap:
or by making a [**bitcoin donation**](https://gobitcoinsv.com/#sponsor?utm_source=github&utm_medium=sponsor-link&utm_campaign=go-meta-extractor&utm_term=go-meta-extractor&utm_content=go-meta-extractor) to ensure this journey continues indefinitely! :rocket:

[![Stars](https://img.shields.io/github/stars/mrz1836/go-meta-extractor?label=Please%20like%20us&style=social)](https://github.com/mrz1836/go-meta-extractor/stargazers)

<br/>

### Contributors ✨
Thank you to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://mrz1818.com"><img src="https://avatars.githubusercontent.com/u/3743002?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Mr. Z</b></sub></a><br /><a href="#infra-mrz1836" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/mrz1836/go-meta-extractor/commits?author=mrz1836" title="Code">💻</a> <a href="#maintenance-mrz1836" title="Maintenance">🚧</a> <a href="#security-mrz1836" title="Security">🛡️</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

> This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification.

<br/>

## License

[![License](https://img.shields.io/github/license/mrz1836/go-meta-extractor.svg?style=flat&v=1)](LICENSE)
