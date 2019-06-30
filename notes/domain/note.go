package domain

type Note struct {
	ID      string
	Title   string
	Content string
}

var EmptyNote = Note{}
