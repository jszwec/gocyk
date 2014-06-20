package cyk

import "strings"

// Executes Cocke–Younger–Kasami algorithm on given formal grammar in Chomsky reduced form
// and on the given word. The result is a table containing the results for both word
// and it's subwords.
func Cyk(g map[string][]string, w string) CykTable {
	c, ct := strings.Split(w, ""), NewCykTable(len(w))
	for i := range ct[0] {
		ct[0][i], ct[1][i] = CykCell{c[i]: {}}, make(CykCell)
		searchNadd(g, ct[1][i], c[i])
	}
	return cyk(g, ct, c)
}

func cyk(g map[string][]string, ct CykTable, c []string) CykTable {
	for i := 2; i < len(c)+1; i++ {
		for j := 0; j < len(c)-i+1; j++ {
			ct[i][j] = make(CykCell)
			for k := 1; k < i; k++ {
				splitNsearch(g, ct[i][j], ct[k][j], ct[i-k][j+k])
			}
		}
	}
	return ct
}

func splitNsearch(g map[string][]string, cc, c1, c2 CykCell) {
	for key1 := range c1 {
		for key2 := range c2 {
			searchNadd(g, cc, key1+key2)
		}
	}
}

func searchNadd(g map[string][]string, cc CykCell, p string) {
	for key, val := range g {
		for i := range val {
			if val[i] == p {
				cc[key] = struct{}{}
			}
		}
	}
}
