package af

import (
	"fmt"
	"log"
	"os"
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

func getState(a AF, state string) State {
	for _, s := range a {
		if s.Name == state {
			return s
		}
	}
	return State{}
}

func solvedState(a AF, inder string) string {
	for _,s := range a {
		if s.Ind == inder {
			return s.Name
		}
	}
	return emptyState
}

func rmName(n string) {
	rm := func(s []string, i int) []string {
		s[i] = s[len(s)-1]
		return s[:len(s)-1]
	}

	for id, tn := range Names {
		if n == tn {
			Names = rm(Names, id)
		}
	}
}

func getIdeterminations(s *State) []Indetermination {
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
					states.Parent = s
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
		state := State{sanitaze(rule.Name), "",nil}
		sname := strings.ReplaceAll(rule.Name, "<", "")
		sname = strings.ReplaceAll(sname, ">", "")
		rmName(sname)
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

					simbol = sanitaze(simbol)

					if func(s string) bool{
						for _,ts := range simbols {
							if s == ts {
								return false
							}
						}
						return true
					} (simbol) && simbol != "epsi"{
						simbols = append(simbols, simbol)
					}

					if rstate == production {
						state.Production = append(state.Production, Beam{simbol, sanitaze(emptyState)})
					} else {
						state.Production = append(state.Production, Beam{simbol, sanitaze(rstate)})
					}
				}
			}
		}
	}

	return finiteAutomaton
}

func Print(fname string, finiteAutomaton *AF) {
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err.Error())
	}

	output := "\t\t"

	ns := len(simbols)

	for id,s := range simbols {
		if id != ns-1 {
			output += fmt.Sprintf("|       %s       ", s)
		} else {
			output += fmt.Sprintf("|       %s       |\n", s)
		}
	}

	emptyStr := 0
	for emptyStr != -1 {
		emptyStr = -1
		for id,simbol := range simbols {
			if sanitaze(simbol) == "" {
				emptyStr = id
			}
		}
		if emptyStr != -1 {
			simbols = append(simbols[:emptyStr], simbols[emptyStr+1:]...)
		}
	}

	for id := range *finiteAutomaton {
		state := &(*finiteAutomaton)[id]
		if isTerminalState(state.Name) {
			output += fmt.Sprintf("| *%s  |", state.Name)
		} else {
			output += fmt.Sprintf("|  %s  |", state.Name)
		}

		for _,simbol := range simbols {
			sts := "      "
			for _, production := range state.Production {
				if production.Simbol == simbol && production.State != emptyState{
					sts += production.State
				}
			}
			if len(sts) < 9 {
				sts += "   "
			}
			sts += "      |"
			lsts := len(sts)

			if lsts > 16 {
				overlap := lsts - 16

				if overlap%2 == 0 {
					overlap = overlap/2
					sts = sts[overlap:]
					sts = sts[:lsts - overlap]
				} else {
					sts = sts[overlap:]
				}
			}

			output += sts
		}
		output += "\n"
	}
	_, err = f.Write([]byte(output))

	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()
}

func Determining(finiteAutomaton AF) AF {
	ti := 0

	Determinded := finiteAutomaton
	var indeterminations []Indetermination
	var log AF

	for id := range Determinded {
		state := &Determinded[id]
		indeterminations = append(indeterminations, getIdeterminations(state)...)
	}

	ti += len(indeterminations)

	for ti > 0 {
		enqueu := func (a AF, s State) AF{
			in := func(a AF, key State) (bool, int) {
				for id, st := range a {
					if st.Name == key.Name {
						return true, id
					}
				}
				return false,-1
			}

			exists,id := in(a, s)

			if exists {
				a[id] = s
			} else {
				return append(a, s)
			}

			return a
		}
		// criar novo estado
		for _, ind := range indeterminations {
			var state State
			sname := strings.ReplaceAll(ind.States, "<", "")
			sname = strings.ReplaceAll(sname, ">", "")

			inder := sanitaze(ind.Simbol + " " + sname)

			removeIndetermination := func(s *State, simbol string, states string, new string) {
				sLefts := len(states)
				rm := func(s []Beam, i int) []Beam {
					s[i] = s[len(s)-1]
					return s[:len(s)-1]
				}

				for sLefts > 0 {
					for id, pd := range s.Production {
						if pd.Simbol == simbol {
							br := false
							for _, r := range states {
								if pd.State == sanitaze(emptyState) || pd.State == "<"+string(r)+">" {
									s.Production = rm(s.Production, id)
									br = true
									sLefts--
									break
								}
							}
							if br || sLefts <= 0 {
								break
							}
						}
					}
				}

				s.Production = append(s.Production, Beam{simbol, new})
			}

			if sn := solvedState(Determinded, inder); sn == emptyState {
				state = State{"<" + Names[currentName] + ">", sanitaze(ind.Simbol + " " + sname), nil}

				rmName(Names[currentName])

				isIn := func(p []Beam, key Beam) bool {
					for _, i := range p {
						if i == key {
							return true
						}
					}
					return false
				}

				// novo estado herda a combinação das produções dos estados que antes gerava a interdeminização
				for _, s := range sname {
					if s != '-' {
						for _, pd := range getState(Determinded, "<"+string(s)+">").Production {
							if !isIn(state.Production, Beam{sanitaze(pd.Simbol), sanitaze(pd.State)}){
								state.Production = append(state.Production, pd)
							}
						}
					}
				}

				for _,s := range sname {
					if s == '-' {
						haveTerminal := false
						
						for _,pd := range state.Production {
							if pd.State == "-" {
								haveTerminal = true
							}
						}

						if !haveTerminal {
							state.Production = append(state.Production, Beam{sanitaze("epsi"), sanitaze("-")})
						}
					}
				}
			} else {
				state = getState(Determinded, sn)
			}

			// indeterminização é removida
			removeIndetermination(ind.Parent, ind.Simbol, sname, state.Name)
			log = enqueu(log, *ind.Parent)
			ti--

			// se um ou mais estados que geraram o novo estado for terminal, ele também será
			for _, r := range sname {
				if isTerminalState("<" + string(r) + ">") {
					terminals = append(terminals, state.Name)
					break
				}
			}
			Determinded = enqueu(Determinded, state)
		}
		// corrigir erro de estados com transições duplicadas
		for _,state := range log {
			Determinded = enqueu(Determinded, state)
		}
		// repetir esses processo para cada estado referenciado
		// problema com estados duplicados após remoção de indeterminizações
		indeterminations = nil
		log = nil
		for id := range Determinded {
			state := &Determinded[id]
			indeterminations = append(indeterminations, getIdeterminations(state)...)
		}
		ti += len(indeterminations)
	}

	return Determinded
}

func RemovingDeathStates(finiteAutomaton AF) AF {
	return nil
}

func RemovingUnreachebleStates(finiteAutomaton AF) AF {
	return nil
}
