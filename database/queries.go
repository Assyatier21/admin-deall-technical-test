package database

const (
	// CRUD User
	InsertUser            = "INSERT INTO users VALUES('%d','%s','%s','%d','%s')"
	GetRegisteredUser     = "SELECT id, username, role_id FROM users WHERE role_id = 1"
	UpdateRegisteredUser  = "UPDATE users SET username='%s',password='%s',role_id='%d',token='%s' WHERE id='%d'"
	GetUserByUsernamePass = "SELECT * FROM users WHERE username='%s' AND password='%s' AND token='%s' AND role_id='%d'"
	DeleteRegisteredUser  = "DELETE FROM users WHERE id='%d'"

	// CRUD User Points
	GetUsersPoints      = "SELECT created_by, SUM(Points) as points FROM articles GROUP BY created_by"
	UpdateArticlePoints = "UPDATE articles SET points='%d' WHERE id='%d'"
	ResetUserPoints     = "UPDATE articles SET points=0 WHERE created_by='%d'"

	// CRUD Articles
	InsertArticle     = "INSERT INTO articles VALUES('%d','%s','%s','%s','%d','%d')"
	GetArticleById    = "SELECT * FROM articles WHERE id='%d'"
	UpdateArticleById = "UPDATE articles SET title='%s',content='%s',points='%d' WHERE id='%d'"
	DeleteArticleById = "DELETE FROM articles WHERE id='%d'"
)
