package af

type Beam struct {
	Simbol string
	State  string
}

type State struct {
	Name       string
	Production []Beam
}

type AF []State
