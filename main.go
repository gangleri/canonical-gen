// package caninical-gen is a cli tool that generates the HTML needed to publish
// a Golang package using a Canonical import path
package main // import "gangleri.io/pkg/canonical-gen"

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

// Termplate used to generate redirecting HTML
var html string = `<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="{{.Name}} {{.Type}} {{.Url}}">
<meta http-equiv="refresh" content="0; url={{.Url}}">
<title>{{.Name}}</title>
</head>
<body>
Click <a href="{{.Url}}">here</a> to visit actual repository
</body>
</html>
`

// Stuct holds package related information
type PkgInfo struct {
	Name string
	Url  string
	Type string
}

// If a repository url is supplied as an argumet that will be returned otherwise
// if will attempt to get the the url from repository's remote info
func getRepoUrl(arg string) (url string) {
	url = arg
	if url == "" {
		o, err := exec.Command("git", "remote", "-v").Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to obtain Repository URL, try specifying it with the -repo flag")
			os.Exit(1)
		}
		re := regexp.MustCompile(`.+(github\.com.+)\ .+`)
		m := re.FindAllStringSubmatch(string(o), -1)

		if len(m) < 1 || len(m[0]) < 1 {
			fmt.Fprintln(os.Stderr, "Unable to obtain Repository URL, try specifying it with the -repo flag")
			os.Exit(1)
		}
		url = "https://" + m[0][1]
	}
	return
}

// If a package name has been supplied as an argument this value will be used
// otherwise it will get the package name based on the current working directory
func getPkgName(arg string) (pkg string) {
	pkg = arg
	if pkg == "" {
		dir, _ := os.Getwd()
		pkg = strings.Split(dir, "src/")[1]
	}
	return
}

// Generate the Html using the info supplied in PkgInfo ad write the Html to the
// specified io.Writer
func writeHtml(out io.Writer, pkg *PkgInfo) {
	t, _ := template.New("html").Parse(html)
	t.Execute(out, pkg)
}

func main() {
	pkg := flag.String("pkg", "", "Package name")
	repo := flag.String("repo", "", "repository url")
	repoType := flag.String("type", "git", "repository type, git is the default")
	flag.Parse()

	*pkg = getPkgName(*pkg)
	*repo = getRepoUrl(*repo)

	writeHtml(os.Stdout, &PkgInfo{
		Name: *pkg,
		Url:  *repo,
		Type: *repoType,
	})
}
