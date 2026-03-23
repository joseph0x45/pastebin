package models

type Paste struct {
	ID      string `db:"id"`
	Title   string `db:"title"`
	Preview string `db:"preview"`
	Content string `db:"content"`
}
