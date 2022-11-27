package entity

type Article struct {
	Id        int    `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	CreatedAt string `json:"created_at" form:"created_at"`
	CreatedBy int    `json:"created_by" form:"created_by"`
	Points    int    `json:"points" form:"points"`
}
type Points struct {
	Points int `json:"points"`
}

type ArticlePoints struct {
	Id     int `json:"id" form:"id"`
	Points int `json:"points" form:"points"`
}
type UpdatedArticle struct {
	Id      int    `json:"id" form:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Points  int    `json:"points" form:"points"`
}
