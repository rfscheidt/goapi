package animals

type Dog struct {
	Name string
}

func (a Dog) Speaks() string {
	return a.Name + " au au"
}
