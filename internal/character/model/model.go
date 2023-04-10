package model

type Role struct {
	Name       string      `json:"name,omitempty"`
	Characters []Character `json:"characters,omitempty"`
}

type Character struct {
	RoleName    string
	Name        string
	ImgUrl      string
	Description string
	Level       int
	Skills      []Skill
	Stats       []Stat
}

type Skill struct {
	Name        string
	ImgUrl      string
	Description string
	Level       int
	Button      string
	Formula     []Formula
}

type Stat struct {
	Name    string
	Value   string
	Scaling string
}

type Formula struct {
	StatName string
	Formula  string
}
