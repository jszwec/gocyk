package cyk

import (
	"html/template"
	"os"
)

const (
	tname         = "CykTable"
	tableTemplate = `<table border=1>
{{range $i := .}}  <tr>
{{range $j := $i}}   <td>{{if $j}}{{range $key, $val := $j}}{{$key}}{{end}}{{else}}-{{end}}</td>
{{end}}  </tr>
{{end}}</table>`
)

type CykCell map[string]struct{}
type CykTable [][]CykCell

// Creates new CykTable with size <n+1, n> - elements are not initialized
func NewCykTable(n int) CykTable {
	ct := make(CykTable, n+1)
	for i := 0; i < n+1; i++ {
		ct[i] = make([]CykCell, n)
	}
	return ct
}

// Dumps CykTable type to file in html table style
func (ct *CykTable) ToFile(fn string) error {
	var t = template.Must(template.New(tname).Parse(tableTemplate))
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(f, tname, ct)
}
