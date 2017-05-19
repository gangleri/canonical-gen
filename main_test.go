package main

import (
	"bytes"
	"testing"
)

func TestGetRepoUrlFrom(t *testing.T) {
	r := getRepoUrl("")
	if r != "https://github.com/gangleri/canonical-gen" {
		t.Errorf("Should have determined repo URL. Got [%s]", r)
	}
}

func TestGetRepoUrlFromArg(t *testing.T) {
	r := getRepoUrl("from_arg")
	if r != "from_arg" {
		t.Error("Should have used the supplied repo url")
	}
}

func TestGetPkgName(t *testing.T) {
	p := getPkgName("")
	if p != "gangleri.io/pkg/canonical-gen" {
		t.Errorf("Should have used the supplied package name. Got [%s]", p)
	}
}

func TestGetPkgNameFromArg(t *testing.T) {
	p := getPkgName("from_arg")
	if p != "from_arg" {
		t.Errorf("Should have used the supplied package name")
	}
}

func TestGenerateHtml(t *testing.T) {
	const expectedHtml = `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="gangleri.io/pkg/canonical-gen git https://github.com/gangleri/canonical-gen">
<meta http-equiv="refresh" content="0; url=https://github.com/gangleri/canonical-gen">
<title>gangleri.io/pkg/canonical-gen</title>
</head>
<body>
Click <a href="https://github.com/gangleri/canonical-gen">here</a> to visit actual repository
</body>
</html>
`

	var tpl bytes.Buffer
	writeHtml(&tpl, &pkgInfo{
		Name: "gangleri.io/pkg/canonical-gen",
		Url:  "https://github.com/gangleri/canonical-gen",
		Type: "git",
	})

	if tpl.String() != expectedHtml {
		t.Errorf("Generated Html does not match expected")
		t.Logf("Expected:\n%s\n---------------\nActual\n%s\n", expectedHtml, tpl.String())
	}

}
