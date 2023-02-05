package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Rule struct {
	Name        string
	Productions []string
}

type AF map[string]map[string][]string

var terminals []string

func ReadRules(filename string) []Rule {
	rules := make([]Rule, 0, 10)

	data, err := os.ReadFile("rules.in")

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, line := range strings.Split(string(data), "\n") {
		pline := strings.Replace(line, " ", "", -1)
		ruleLine := strings.Split(pline, "::=")
		rules = append(rules, Rule{ruleLine[0], strings.Split(ruleLine[1], "|")})
	}

	return rules
}

func removeUnterminals(production string) string {
	var start int = 0
	var end int = 0

	for id, char := range production {
		if char == '<' {
			start = id
		}
		if char == '>' {
			end = id
		}
	}

	if start == end {
		return production
	}

	return production[:start] + production[end+1:]
}

func removeTerminals(production string) string {
	var start int = 0
	var end int = 0

	for id, char := range production {
		if char == '<' {
			start = id
		}
		if char == '>' {
			end = id
		}
	}

	if start == end {
		return production
	}

	return production[start : end+1]
}

func isUnterminal(production string) bool {
	var start int = 0
	var end int = 0

	for id, char := range production {
		if char == '<' {
			start = id
		}
		if char == '>' {
			end = id
		}
	}

	return start != end
}

// impossivel dar append num array dentro de um map
// maps são imutáveis
// recriar a key com um novo dado pode ser uma solução

func BuildAF(rules []Rule) AF {
	af := make(AF)

	for _, rule := range rules {
		af[rule.Name] = make(map[string][]string)
		for _, production := range rule.Productions {
			rstate := removeTerminals(production)

			if rstate == production {
				terminals = append(terminals, rule.Name)
			} else {
				simbol := removeUnterminals(production)
				af[rule.Name][simbol] = make([]string, 0)
			}
		}
	}

	for _, rule := range rules {
		af[rule.Name] = make(map[string][]string)
		for _, production := range rule.Productions {
			rstate := removeTerminals(production)
			simbol := removeUnterminals(production)
			af[rule.Name][simbol] = append(af[rule.Name][simbol], rstate)
		}
	}

	return af
}

func main() {
	rules := ReadRules("rules.in")
	af := BuildAF(rules)
	fmt.Println(af["<S>"]["a"][1])
}
