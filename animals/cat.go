package animals

type Cat struct {
	Name string
}

func (a Cat) Speaks() string {
	return a.Name + " miau miau"
}
