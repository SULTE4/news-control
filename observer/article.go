package observer

type Article struct {
	Title   string
	Content string
	InStock bool
}

func NewArticle(title, content string) Article {
	return Article{Title: title, Content: content, InStock: true}
}
