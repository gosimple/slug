# slug

Package `slug` generate slug from unicode string, URL-friendly slugify with
multiple languages support.

[![Go Reference](https://pkg.go.dev/badge/github.com/gosimple/slug.svg)](https://pkg.go.dev/github.com/gosimple/slug)
[![Tests](https://github.com/gosimple/slug/actions/workflows/tests.yml/badge.svg)](https://github.com/gosimple/slug/actions/workflows/tests.yml)
[![codecov](https://codecov.io/gh/gosimple/slug/branch/master/graph/badge.svg?token=FT2kEZHQW7)](https://codecov.io/gh/gosimple/slug)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/gosimple/slug?logo=github&sort=semver)](https://github.com/gosimple/slug/releases)

## Example

```go
package main

import (
	"fmt"
	"github.com/gosimple/slug"
)

func main() {
	text := slug.Make("Hellö Wörld хелло ворлд")
	fmt.Println(text) // Will print: "hello-world-khello-vorld"

	someText := slug.Make("影師")
	fmt.Println(someText) // Will print: "ying-shi"

	enText := slug.MakeLang("This & that", "en")
	fmt.Println(enText) // Will print: "this-and-that"

	deText := slug.MakeLang("Diese & Dass", "de")
	fmt.Println(deText) // Will print: "diese-und-dass"

	slug.Lowercase = false // Keep uppercase characters
	deUppercaseText := slug.MakeLang("Diese & Dass", "de")
        fmt.Println(deUppercaseText) // Will print: "Diese-und-Dass"

	slug.CustomSub = map[string]string{
		"water": "sand",
	}
	textSub := slug.Make("water is hot")
	fmt.Println(textSub) // Will print: "sand-is-hot"
}
```

### Requests or bugs?

<https://github.com/gosimple/slug/issues>

## Installation

```shell
go get -u github.com/gosimple/slug
```

## Benchmarking

```shell
go get -u golang.org/x/perf/cmd/benchstat

go test -run=NONE -bench=. ./... > old.txt
# make changes
go test -run=NONE -bench=. ./... > new.txt

benchstat old.txt new.txt
```

## License

The source files are distributed under the
[Mozilla Public License, version 2.0](http://mozilla.org/MPL/2.0/),
unless otherwise noted.
Please read the [FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html)
if you have further questions regarding the license.
