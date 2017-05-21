# canonical-gen
Generate html for use with canonical import paths of Go packages.
See [go doc][1] for details of canonical import paths and go packages.

Install:
```sh
go get gangleri.io/pkg/canonical-gen
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

|flag|description|
| :--- | :--- |
**-pkg**            |The package import name
**-url**        |The repository URL 
**-type**       |The repository type, defaults to *git*

[1]: https://golang.org/doc/go1.4#canonicalimports
