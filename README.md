gocyk
=====

Cocke–Younger–Kasami algorithm implementation in GO.

Installation
------------

go get github.com/jszwec/gocyk

Compilation
------------

run 'go build' in cmd/gocyk folder 

Usage
-----

  gocyk -input grammar.txt -word baaba -output file.html

  grammar example : 
  
  S->AB|BC
  
  A->BA|a
  
  B->CC|b
  
  C->AB|a
