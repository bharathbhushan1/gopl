package main

import (
	"ch4/github/github"
	"html/template"
	"log"
	"os"
	"sort"
	"time"
)

const htmlTemplate = `
<h1>{{.TotalCount}} issues</h1>
     <table>
     <tr style='text-align: left'>
       <th>#</th>
       <th>State</th>
       <th>User</th>
       <th>Title</th>
     </tr>
     {{range .Items}}
     <tr>
       <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
       <td>{{.State}}</td>
       <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
       <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
     </tr>
     {{end}}
     </table>
`

const templ = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(
	template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

var htmlReport = template.Must(
	template.New("htmlReport").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(htmlTemplate))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if result == nil {
		log.Fatal("No results from github")
	}
	/*
		fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	*/

	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt)
	})
	/*
		if err := report.Execute(os.Stdout, result); err != nil {
			log.Fatal(err)
		}
	*/
	if err := htmlReport.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
