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

type Beam struct {
	Simbol string
	State  string
}

type State struct {
	Name       string
	Production []Beam
}

type AF []State // substituir o map

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

func getState(af AF, id string) *State {
	for i := 0; i < len(af); i++ {
		if af[i].Name == id {
			return &af[i]
		}
	}
	return nil
}

// impossivel dar append num array dentro de um map
// maps são imutáveis
// recriar a key com um novo dado pode ser uma solução

func BuildAF(rules []Rule) AF {
	af := make(AF, 0)

	for _, rule := range rules {
		state := State{rule.Name, nil}
		for _, production := range rule.Productions {
			rstate := removeTerminals(production)

			if rstate == production {
				terminals = append(terminals, rule.Name)
			} else {
				state.Production = make([]Beam, 0)
			}
		}
		af = append(af, state)
	}

	for _, rule := range rules {
		state := getState(af, rule.Name)
		for _, production := range rule.Productions {
			rstate := removeTerminals(production)
			simbol := removeUnterminals(production)
			state.Production = append(state.Production, Beam{simbol, rstate})
		}
	}

	return af
}

func main() {
	rules := ReadRules("rules.in")
	af := BuildAF(rules)
	fmt.Println(af[0].Production[5]) // o simbolo da ultima não está fazendo
}
