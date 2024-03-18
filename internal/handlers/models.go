package handlers

type Film struct {
	Id          int
	Name        string
	Year        int
	Description string
	Rating      float64
	Actors      []string
}

type Actor struct {
	Id       int
	Name     string
	Gender   string
	Birthday string
	Films    []string
}
