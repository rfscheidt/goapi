package circus

type Animal interface {
	Speaks() string
}

func Speaks(a Animal) string {
	return a.Speaks()
}
