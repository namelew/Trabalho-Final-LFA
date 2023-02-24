package input

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var Names = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "K", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "k", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "/", "!", "@", "#","$", "%", "&", "*", "+", "=", "?",
}

type Rule struct {
	Name        string
	Productions []string
}

func pop (s *[]string) string{
	n := (*s)[0]
	*s = (*s)[1:]
	return n
}

func rm (sl []string,s string) []string {
	id := -1
	for i,n := range sl {
		if n == s {
			id = i
		}
	}

	if id != -1 {
		sl[id] = sl[len(sl) - 1]
		return sl[:len(sl)-1]
	}

	return sl
}

func readGR(unames *[]string,entry string) []Rule{
	var rules []Rule

	for _, line := range strings.Split(entry, "\n") {
		if len(line) < 2{
			continue
		}
		pline := strings.Replace(line, " ", "", -1)
		ruleLine := strings.Split(pline, "::=")
		rules = append(rules, Rule{ruleLine[0], strings.Split(ruleLine[1], "|")})
		n := strings.ReplaceAll(ruleLine[0], "<", "")
		n = strings.ReplaceAll(n, ">", "")
		*unames = rm(*unames, n)
	}

	return rules
}

func readTokens(unames *[]string,entry string) []Rule{
	var rules []Rule

	lines := strings.Split(entry,"\n")

	for _,line := range lines{
		if line == ""{
			continue
		}
		var afstate string
		for id,r := range line {
			created := func (s string) (bool, int, string) {
				for _,rl := range rules {
					if s == string(rl.Productions[0][0]) {
						return true,len(rl.Productions),rl.Productions[0][1:]
					}
				}
				return false,-1,""
			}

			yes,tam,nrl := created(string(r)) 

			if !yes && id < len(line) - 1{
				if afstate == "" {
					afstate = "<"+pop(unames)+">"
					rules = append(rules, Rule{"<"+pop(unames)+">", []string{string(r)+afstate}})
				} else {
					n := afstate
					afstate = "<"+pop(unames)+">"
					if id == len(line) - 2 {
						rules = append(rules, Rule{n, []string{string(r)+afstate, string(r)}})
					} else {
						rules = append(rules, Rule{n, []string{string(r)+afstate}})
					}
				}
			} else if yes && tam > 1 {
				afstate = nrl
			}
		}
	}

	return rules
}

func ReadRules(filename string) []Rule {
	data, err := os.ReadFile("rules.in")

	unames := Names

	if err != nil {
		log.Fatal(err.Error())
	}
	temp := strings.Split(string(data), "--")
	rtokens,gr := temp[0],temp[1] 

	rules := readGR(&unames, gr)

	fmt.Println(readTokens(&unames, rtokens))

	Names = nil

	return rules
}
