package main

import (
	"fmt"
	"os"
	"text/template"
)

type myType struct {
	ID   string
	Name string
	Test string
}

func main() {
	list := myType{"id1", "name1", "test1"}

	tmpl := `
<table>
  <tr>
    <td>{{ .ID }}</td>
    <td>{{ .Name }}</td>
    <td>{{ .Test }}</td>
  </tr>
</table>
`

	t := template.Must(template.New("tmpl").Parse(tmpl))

	err := t.Execute(os.Stdout, list)
	if err != nil {
		fmt.Println("executing template:", err)
	}
}
