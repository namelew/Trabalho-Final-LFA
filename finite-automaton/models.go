package af

type Beam struct {
	Simbol string
	State  string
}

type State struct {
	Name       string
	Ind 	   string
	Production []Beam
}

type Indetermination struct {
	Simbol      string
	States      string
	Parent      *State
	Productions []Beam
}

type AF []State
