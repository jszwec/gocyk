package cyk

// Executes Cocke–Younger–Kasami algorithm on given formal grammar in Chomsky reduced form
// and on the given word. The result is a table containing the results for both word
// and it's subwords.
func Cyk(g map[string][]string, w string) CykTable {
	ct := NewCykTable(len(w))
	for i := range ct[0] {
		ct[0][i], ct[1][i] = CykCell{string(w[i]): {}}, make(CykCell)
		searchNadd(g, ct[1][i], string(w[i]))
	}
	return cyk(g, ct, len(w))
}

func cyk(g map[string][]string, ct CykTable, n int) CykTable {
	for i := 2; i < n+1; i++ {
		for j := 0; j < n-i+1; j++ {
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
