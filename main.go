package main // import "gangleri.io/pkg/canonical-gen"

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"golang.org/x/tools/go/vcs"
)

var html string = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<meta name="go-import" content="{{.pkg}} {{.type}} {{.repo}}">
		<meta http-equiv="refresh" content="0; url={{.repo}}">
		<title>{{.pkg}}</title>
	</head>
	<body>
		Click <a href="{{.repo}}">here</a> to visit actual repository
	</body>
</html>
`

func main() {
	pkg := flag.String("pkg", "", "Package name")
	repo := flag.String("url", "", "repository url")
	repoType := flag.String("type", "git", "repository type, git is the default")
	flag.Parse()

	if *pkg == "" {
		dir, _ := os.Getwd()
		*pkg = strings.Split(dir, "src/")[1]
	}

	if *repo == "" {
		repoRoot, err := vcs.RepoRootForImportPath(*pkg, false)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to obtain Repository URL, try specifying it with the -url flag")
			os.Exit(1)
		}
		*repo = repoRoot.Repo
	}

	t, _ := template.New("html").Parse(html)

	t.Execute(os.Stdout, map[string]interface{}{
		"pkg":  *pkg,
		"repo": *repo,
		"type": *repoType,
	})
}
