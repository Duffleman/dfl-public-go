package pkg1

type Humanoid struct {
	Name  string `json:"name"`
	Limbs []Limb `json:"limbs"`
}

type Limb struct {
	Name     string  `json:"name"`
	Parent   *Limb   `json:"parent"`
	Children []*Limb `json:"children"`
}
