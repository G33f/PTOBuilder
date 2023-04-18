package model

type Role struct {
	ID         int64       `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Characters []Character `json:"characters,omitempty"`
}

type Character struct {
	ID          int64            `json:"id,omitempty"`
	RoleID      int64            `json:"role_ID,omitempty"`
	Name        string           `json:"name,omitempty"`
	ImgUrl      string           `json:"img_url,omitempty"`
	Description string           `json:"description,omitempty"`
	Level       int              `json:"level,omitempty"`
	Skills      []Skill          `json:"skills,omitempty"`
	Stats       map[string]*Stat `json:"stats,omitempty"`
}

type Skill struct {
	ID          int64     `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	ImgUrl      string    `json:"img_url,omitempty"`
	Description string    `json:"description,omitempty"`
	Button      string    `json:"button,omitempty"`
	Level       int       `json:"level,omitempty"`
	Formula     []Formula `json:"formula,omitempty"`
}

type Stat struct {
	ID      int64 `json:"id,omitempty"`
	Value   int   `json:"value,omitempty"`
	Scaling int   `json:"scaling,omitempty"`
}

type Formula struct {
	ID      int64    `json:"id,omitempty"`
	Level   int      `json:"level,omitempty"`
	Formula string   `json:"formula,omitempty"`
	Stats   []string `json:"stats,omitempty"`
}
