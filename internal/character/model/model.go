package model

type Role struct {
	ID         int         `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Characters []Character `json:"characters,omitempty"`
}

type Character struct {
	ID          int     `json:"id,omitempty"`
	RoleName    string  `json:"role_name,omitempty"`
	Name        string  `json:"name,omitempty"`
	ImgUrl      string  `json:"img_url,omitempty"`
	Description string  `json:"description,omitempty"`
	Level       int     `json:"level,omitempty"`
	Skills      []Skill `json:"skills,omitempty"`
	Stats       []Stat  `json:"stats,omitempty"`
}

type Skill struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	ImgUrl      string    `json:"img_url,omitempty"`
	Description string    `json:"description,omitempty"`
	Level       int       `json:"level,omitempty"`
	Button      string    `json:"button,omitempty"`
	Formula     []Formula `json:"formula,omitempty"`
}

type Stat struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
	Scaling string `json:"scaling,omitempty"`
}

type Formula struct {
	ID       int    `json:"id,omitempty"`
	StatName string `json:"stat_name,omitempty"`
	Formula  string `json:"formula,omitempty"`
}
