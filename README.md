slug
====

Package `slug` generate slug from unicode string, URL-friendly slugify with
multiple languages support.

[Documentation online](http://godoc.org/bitbucket.org/gosimple/slug)

## Example

	package main

	import(
		"bitbucket.org/gosimple/slug"
	    "fmt"
	)

	func main () {
		text := slug.Make("Hellö Wörld хелло ворлд")
		fmt.Println(text) // Will print hello-world-khello-vorld

		someText := slug.Make("影師")
		fmt.Println(someText) // Will print: ying-shi

		enText := slug.MakeLang("This & that", "en")
		fmt.Println(enText) // Will print 'this-and-that'

		deText := slug.MakeLang("Diese & Dass", "de")
		fmt.Println(deText) // Will print 'diese-und-dass'

		slug.CustomSub = map[string]string{
			"water": "sand",
		}
		textSub := slug.Make("water is hot")
		fmt.Println(textSub) // Will print 'sand-is-hot'
	}

### Requests or bugs?
<https://bitbucket.org/gosimple/slug/issues>

## Installation

	go get -u bitbucket.org/gosimple/slug

## License

The source files are distributed under the
[Mozilla Public License, version 2.0](http://mozilla.org/MPL/2.0/),
unless otherwise noted.
Please read the [FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html)
if you have further questions regarding the license.
