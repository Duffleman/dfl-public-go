package pkg1

import (
	"github.com/Duffleman/dfl-public-go/pkg2"
)

type Humanoid struct {
	Name  string      `json:"name"`
	Limbs []*Limb     `json:"limbs"`
	Pets  []*pkg2.Pet `json:"pets"`
}

type Limb struct {
	Name     string  `json:"name"`
	Parent   *Limb   `json:"parent"`
	Children []*Limb `json:"children"`
}
