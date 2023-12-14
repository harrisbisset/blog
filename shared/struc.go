package shared

type Post struct {
	Name    string
	Title   string
	Summary string
	Date    string
	Ref     string
}

type Pages []Page

type Page struct {
	Name string
	Ref  string
	Dir  string
}
