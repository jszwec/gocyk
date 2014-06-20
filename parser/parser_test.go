package parser

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseSuccess(t *testing.T) {
	data := strings.NewReader(
		"S -> AB|a|BB|c|d\n" +
			"  X -> a\n" +
			"\n" +
			"  \t\n" +
			"Y -> XY\n")
	expected := map[string][]string{
		"S": []string{"AB", "a", "BB", "c", "d"},
		"X": []string{"a"},
		"Y": []string{"XY"},
	}
	res, err := parse(data)
	if err != nil {
		t.Fatalf("expected nil, got %s", err.Error())
	}
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestParseNoMatch(t *testing.T) {
	data := []string{
		"S -> aa\n",
		"S -> ABC\n",
		"S -> Aa\n",
		"S -> aA\n",
		"S -> \n",
		"AB -> a\n",
	}
	for i := range data {
		_, err := parse(strings.NewReader(data[i]))
		if err == nil {
			t.Fatalf("expected '%v', got nil", errNoMatch(data[i]))
		}
		if err.Error() != errNoMatch(strings.TrimRight(data[i], " \t\n")).Error() {
			t.Errorf("expected '%v', got '%v'", errNoMatch(data[i]), err)
		}
	}
}

func TestParseDuplicateEntry(t *testing.T) {
	data := strings.NewReader("S -> a\nS -> b\n")
	_, err := parse(data)
	if err == nil {
		t.Fatalf("expected '%v', got nil", errDuplicateEntry("S"))
	}
	if err.Error() != errDuplicateEntry("S").Error() {
		t.Errorf("expected '%v', got '%v'", errDuplicateEntry("S"), err)
	}
}
