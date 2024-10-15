package pojo

type Hero struct {
	Name   string
	Age    int
	Skill  string
	Weapon string
}

func (h *Hero) GetName() string {
	return h.Name
}
func (h *Hero) SetName(name string) {
	h.Name = name
}
