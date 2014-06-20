package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	errNoMatch = func(line string) error {
		return fmt.Errorf("invalid line: %s", line)
	}
	errDuplicateEntry = func(p string) error {
		return fmt.Errorf("duplicate entry for: %s", p)
	}
)

var productRegex = regexp.MustCompile(`^([A-Z])\s*->\s*((?:[A-Z]{2}|[a-z])(?:\|[A-Z]{2}|\|[a-z])*)$`)

// Parses Grammar in the given file. It must be in Chomsky reduced form
// A -> BC
// A -> a, where 'A','B','C' are nonterminal symbols and 'a' is a terminal one.
// Terminal symbol cannot be longer than 1 character at the most.
// Alternatives are possible eg. 'A -> BC|a', 'A -> BA|a|BC' etc.
func Parse(file string) (map[string][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return parse(f)
}

func parse(file io.Reader) (map[string][]string, error) {
	g, sc := make(map[string][]string), bufio.NewScanner(file)
	for sc.Scan() {
		line := strings.Trim(sc.Text(), " \t")
		if line == "" {
			continue
		}
		if err := parseLine(line, g); err != nil {
			return nil, err
		}
	}
	return g, nil
}

func parseLine(line string, g map[string][]string) error {
	match := productRegex.FindStringSubmatch(line)
	if match == nil {
		return errNoMatch(line)
	}
	if _, ok := g[match[1]]; ok {
		return errDuplicateEntry(match[1])
	}
	g[match[1]] = strings.Split(match[2], "|")
	return nil
}
