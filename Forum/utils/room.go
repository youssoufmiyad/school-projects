package forum

import (
	"database/sql"
	"strconv"
)

// ANCHOR - STRUCTURE POSTS
type EntryPost struct {
	ID       []int
	Texts    []string
	Text     string
	Users    []string
	User     []string
	Filters  []string
	Filter   []string
	IsLogged bool
	IsLiked  bool
}

// ANCHOR - NewDB
// crée les structure des postes,des commentaires,et des utilisateurs dans la BDD
func NewDB(db *sql.DB) *Post {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "posts" (
			"ID"	INT,
			"Text"	TEXT,
			"User"	TEXT,
			"Filter" TEXT,
			PRIMARY KEY("ID")
		);
	`)
	stmt.Exec()
	stmt, _ = db.Prepare(`
	CREATE TABLE IF NOT EXISTS "comment"(
		"ID" INT,
		"CommentID" INT,
		"Text" TEXT,
		"User" TEXT,
		PRIMARY KEY("CommentID")
	  );)
	  `)

	stmt.Exec()
	stmt, _ = db.Prepare(`
		CREATE TABLE IF NOT EXISTS "users" (
			"ID" INT,
			"email"	TEXT UNIQUE,
			"username"	TEXT,
			"password"	TEXT,
			PRIMARY KEY("ID")
		);)
		`)
	stmt.Exec()
	stmt, _ = db.Prepare(`
		CREATE TABLE IF NOT EXISTS "likes" (
			"postid"	TEXT,
			"username"	TEXT,
			"like"	TEXT
		);
	`)
	stmt.Exec()

	return &Post{
		DB: db,
	}
}

// ANCHOR - StockageStructUser
func StockageStructUser(U *EntryUser, ID int, username, email, password string) {
	// Stockage dans la struct //

	// username //
	U.Username = username
	U.Users[ID-1] = username

	// ID //
	U.IDs[ID-1] = ID

	// email //
	U.Email = email
	U.Emails[ID-1] = email

	// password //
	U.Password = password
	U.Passwords[ID-1] = password

	// permet de se connecter automatiquement au compte qu'on viens de créer
	U.IsLogged = true
}

// ANCHOR - StockageStructPost
func StockageStructPost(E *EntryPost, ID int, txt string, User string, Filter string) {
	// Stockage dans la struct //

	// Text //
	E.Texts[ID] = txt

	// ID //
	E.ID[ID] = ID + 1

	// Users //
	E.Users[ID] = User

	// Filtres
	E.Filters[ID] = Filter
}

// ANCHOR - StockageStructComment
func StockageStructComment(C *EntryComment, ID, CommentID int, txt string, User string) {
	// Stockage dans la struct //

	// Text //
	C.Texts[CommentID] = txt

	// ID //
	C.ID[0] = strconv.Itoa(ID)

	// CommentID //
	C.CommentID[CommentID] = CommentID + 1

	// Users //
	C.User[CommentID] = User
}

// ANCHOR - OuvertureBDD
func OuvertureBDD(E *EntryPost, ID int, User string, filtre string) *Post {
	// Ouverture / création de la bdd //
	db, _ := sql.Open("sqlite3", "test.db")
	update := NewDB(db)

	// Synchronisation struct/BDD (import from BDD)
	update.DB.QueryRow("SELECT MAX(id) from posts").Scan(&ID)
	InitialisationPost(E, ID)
	MajPost(E, update, ID, User, filtre)

	return &Post{
		DB: update.DB,
	}
}
