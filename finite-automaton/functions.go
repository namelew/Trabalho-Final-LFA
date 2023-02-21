package af

import (
	"fmt"
	"strings"

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

	return sanitaze(production[:start] + production[end+1:])
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
	return sanitaze(production[start : end+1])
}

func numStates(states string) int {
	nstate := 0

	for id, state := range states {
		if state == '<' {
			for j := id + 1; j < len(states); j++ {
				if states[j] == '>' {
					nstate++
				}
			}
		}
	}

	return nstate
}

func getIdeterminations(s State) []Indetermination {
	var indeterminations = []Indetermination{}
	var states Indetermination

	isIndetermination := func(a Beam, b Beam) bool {
		return b.Simbol == a.Simbol && a.State != b.State
	}

	isIn := func(i []Indetermination, a string) bool {
		for _, ind := range i {
			if ind.Simbol == a {
				return true
			}
		}
		return false
	}

	for id, prod := range s.Production {
		states.States += prod.State
		for j := id + 1; j < len(s.Production); j++ {
			if isIndetermination(prod, s.Production[j]) {
				if !isIn(indeterminations, prod.Simbol) {
					states.Simbol = prod.Simbol
					states.States += s.Production[j].State
				}
			}
		}

		if numStates(states.States) > 1 {
			indeterminations = append(indeterminations, states)
		}
		states.States = ""
	}

	return indeterminations
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

func sanitaze(s string) string {
	schars := []string{"\n", "\t", "\r", "\a", "\f", "\v", "\b"}

	for _, char := range schars {
		s = strings.ReplaceAll(s, char, "")
	}

	return s
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

	for id := range finiteAutomaton {
		state := &finiteAutomaton[id]
		for _, rule := range rules {
			if state.Name == rule.Name {
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

func Determining(finiteAutomaton AF) AF {
	var Determinded AF
	var indeterminations []Indetermination

	for _, state := range finiteAutomaton {
		indeterminations = append(indeterminations, getIdeterminations(state)...)
	}

	fmt.Println(indeterminations)

	// criar novo estado

	// novo estado herda a combinação das produções dos estados que antes gerava a interdeminização
	// se um ou mais estados que geraram o novo estado for terminal, ele também será
	// repetir esses processo para cada estado referenciado
	return Determinded
}

func RemovingDeathStates(finiteAutomaton AF) AF {
	return nil
}

func RemovingUnreachebleStates(finiteAutomaton AF) AF {
	return nil
}
