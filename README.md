# canonical-gen
[![Build Status](https://travis-ci.org/gangleri/canonical-gen.svg?branch=master)](https://travis-ci.org/gangleri/canonical-gen)
[![Report card](https://goreportcard.com/badge/gangleri.io/pkg/canonical-gen)](https://goreportcard.com/report/gangleri.io/pkg/canonical-gen)

Generate html for use with canonical import paths of Go packages.
See [go doc][1] for details of canonical import paths and go packages.

Install:
```sh
go get -u gangleri.io/pkg/canonical-gen
```

Usage:
From the root of the package you wish to publish run `canonical-gen`, 
this will write the HTML to stdout; you can redirect to any file 
location you like.

e.g.
```html
canonical-gen > index.html
```

The default behavior is to get the package import path and the repository 
URL from the root directory of the package code. However, you can change 
this using the following arguments:

|Flag|Description|
| --- | --- |
**-pkg**            |The package import name
**-repo**        |The repository URL 
**-type**       |The repository type, defaults to *git*

[1]: https://golang.org/doc/go1.4#canonicalimports
