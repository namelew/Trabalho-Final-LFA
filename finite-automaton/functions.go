package af

import (
	"fmt"

	"github.com/namelew/automato-finito/input"
)

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

func isTerminalState(tstate string) bool {
	for _, state := range terminals {
		if state == tstate {
			return true
		}
	}
	return false
}

func getState(finiteAutomaton AF, id string) *State {
	for i := 0; i < len(finiteAutomaton); i++ {
		if finiteAutomaton[i].Name == id {
			return &finiteAutomaton[i]
		}
	}
	return nil
}

func Build(rules []input.Rule) AF {
	finiteAutomaton := make(AF, 0)

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
		finiteAutomaton = append(finiteAutomaton, state)
	}

	for _, rule := range rules {
		state := getState(finiteAutomaton, rule.Name)
		for _, production := range rule.Productions {
			simbol := removeUnterminals(production)
			rstate := removeTerminals(production)
			if rstate == production {
				state.Production = append(state.Production, Beam{simbol, emptyState})
			} else {
				state.Production = append(state.Production, Beam{simbol, rstate})
			}
		}
	}

	return finiteAutomaton
}

func Print(finiteAutomaton AF) {
	for _, state := range finiteAutomaton {
		if isTerminalState(state.Name) {
			fmt.Printf("*%s: ", state.Name)
		} else {
			fmt.Printf("%s: ", state.Name)
		}
		npd := len(state.Production)
		for id, production := range state.Production {
			if id != npd-1 {
				fmt.Printf("(%s, %s), ", production.Simbol, production.State)
			} else {
				fmt.Printf("(%s, %s)", production.Simbol, production.State)
			}
		}
		fmt.Println()
	}
}
