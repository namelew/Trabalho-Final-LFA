package main

import (
	af "github.com/namelew/automato-finito/finite-automaton"
	"github.com/namelew/automato-finito/input"
)

func main() {
	rules := input.ReadRules("rules.in")
	automatoFinito := af.Build(rules)

	af.Print("afnd.out", &automatoFinito)

	det := af.Determining(automatoFinito)

	af.Print("afd.out", &det)
}
