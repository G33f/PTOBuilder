package character

type Role struct {
	Name       string
	Characters []Character
}

type Character struct {
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
