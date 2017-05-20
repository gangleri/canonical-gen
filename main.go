package main

import (
	"flag"
	"os"
	"text/template"
)

var html string = `
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<meta name="go-import" content="{{.pkg}} {{.type}} {{.repo}}">
		<meta http-equiv="refresh" content="0; url={{.repo}}">
	</head>
	<body>
		Nothing to see here; <a href="{{.repo}}">move along</a>.
	</body>
</html>
`

func main() {
	pkg := flag.String("pkg", "", "Package name")
	repo := flag.String("repoUrl", "", "repository url")
	repoType := flag.String("repoType", "git", "repository type, git is the default")
	flag.Parse()

	t, _ := template.New("html").Parse(html)

	t.Execute(os.Stdout, map[string]interface{}{
		"pkg":  *pkg,
		"repo": *repo,
		"type": *repoType,
	})
}
