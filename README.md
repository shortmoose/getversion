# getversion


[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/shortmoose/getversion)
[![Go Report Card](https://goreportcard.com/badge/shortmoose/getversion)](https://goreportcard.com/report/shortmoose/getversion)
[![Releases](https://img.shields.io/github/release-pre/shortmoose/getversion.svg?sort=semver)](https://github.com/shortmoose/getversion/releases)
[![LICENSE](https://img.shields.io/github/license/shortmoose/getversion.svg)](https://github.com/shortmoose/getversion/blob/master/LICENSE)


This tool is designed to help bridge the gap between which version of a
repository is needed and which version of the repository best matches that
need. For example if you want version 1.1 of a repository, that likely means you
want the latest version of the 1.1 branch of the repository, not necessarily
version 1.1.0.

This tool assumes projects are using [Semantic Versioning](https://semver.org/)
which is a structured way of creating version numbers.


## Usage

```
# Download and install the tool.
# Hey look - go get also uses semantic versioning!
go get -u github.com/shortmoose/getversion@v0.1

# Clone the repository. 
git clone https://github.com/<repo>

# Use getversion to checkout the correct state of the repository.
getversion <repo>@v1.1

# At some point in the future we would like to improve the workflow
# so getversion can also do the clone.
# For example: getversion github.com/<repo>@v1.1
```

Note that this tool often leaves the repository in a disconnected state.
This is okay since when using getversion we are normally not planning
to modify the given repository, we are just wanting to use it.
