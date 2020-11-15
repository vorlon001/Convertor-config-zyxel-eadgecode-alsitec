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
	list := []myType{{"id1", "name1", "test1"}, {"i2", "n2", "t2"}}

	tmpl := `
<table>{{range $y, $x := . }}
  <tr>
    <td>{{ $x.ID }}</td>
    <td>{{ $x.Name }}</td>
    <td>{{ $x.Test }}</td>
  </tr>{{end}}
</table>
`

	t := template.Must(template.New("tmpl").Parse(tmpl))

	err := t.Execute(os.Stdout, list)
	if err != nil {
		fmt.Println("executing template:", err)
	}
}
