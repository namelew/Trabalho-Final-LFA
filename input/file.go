package input

import (
	"log"
	"os"
	"strings"
)

type Rule struct {
	Name        string
	Productions []string
}

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
