package af

import "github.com/namelew/automato-finito/input"

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
			state.Production = append(state.Production, Beam{simbol, rstate})
		}
	}

	return finiteAutomaton
}
