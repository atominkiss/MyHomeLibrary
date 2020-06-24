package main

type Book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  *Author `json:"author"`
}

type Author struct {
	Firstname   string  `json:"firstname"`
	Lastname    string  `json:"lastname"`
}

func main() {
	
}