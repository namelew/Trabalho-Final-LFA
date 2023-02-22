package input

import (
	"log"
	"os"
	"strings"
)

// var Names = []string{
// 	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "K", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
// 	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "k", "r", "s", "t", "u", "v", "w", "x", "y", "z",
// 	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "/", "!", "@", "#","$", "%", "&", "*", "+", "=", "?", "|",
// }

type Rule struct {
	Name        string
	Productions []string
}

// func strDiff(a string, b string) string {
// 	la := len(a)
// 	lb := len(b)
// 	diff := ""

// 	if lb > la {
// 		return strDiff(b, a)
// 	}

// 	for id, r := range b {
// 		if string(r) == string(a[id]) {
// 			diff += string(r)
// 		}
// 	}

// 	return diff
// }

func ReadRules(filename string) []Rule {
	rules := make([]Rule, 0, 10)

	data, err := os.ReadFile("rules.in")

	// unames := Names
	// pop := func () string{
	// 	n := unames[0]
	// 	unames = unames[1:]
	// 	return n
	// }

	// rm := func (s string) []string {
	// 	id := -1
	// 	for i,n := range unames {
	// 		if n == s {
	// 			id = i
	// 		}
	// 	}

	// 	if id != -1 {
	// 		unames[id] = unames[len(unames) - 1]
	// 		return unames[:len(unames)-1]
	// 	}

	// 	return unames
	// }

	if err != nil {
		log.Fatal(err.Error())
	}
	temp := strings.Split(string(data), "--")
	rtokens,gr := temp[0],temp[1] 

	rtokens += ""

	for _, line := range strings.Split(gr, "\n") {
		if len(line) < 2{
			continue
		}
		pline := strings.Replace(line, " ", "", -1)
		ruleLine := strings.Split(pline, "::=")
		rules = append(rules, Rule{ruleLine[0], strings.Split(ruleLine[1], "|")})
		//n := strings.ReplaceAll(ruleLine[0], "<", "")
		//n = strings.ReplaceAll(n, ">", "")
		//unames = rm(n)
	}

	// lines := strings.Split(rtokens,"\n")
	// var rw []Rule

	// for _,line := range lines{
	// 	if line == ""{
	// 		continue
	// 	}
	// 	rw = append(rw, Rule{"<"+pop()+">", []string{line}})
	// }

	// var prev Rule = Rule{"", []string{}}
	// for id, v := range rw {
	// 	if prev.Name != "" {
	// 		diff := strDiff(prev.Productions[0], v.Productions[0])
	// 		if diff != "" {
	// 			r := Rule{"<"+pop()+">", []string{string(v.Productions[0][0]) + v.Name}}
	// 			rw[id].Productions[0] = rw[id].Productions[0][1:] 
	// 			rw = append(rw, r)
	// 		}
	// 	}
	// 	prev = v
	// }

	// Names = nil

	return rules
}
