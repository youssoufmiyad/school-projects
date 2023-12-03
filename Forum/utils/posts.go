package forum

import (
	"database/sql"
	"fmt"
)

// ANCHOR - STRUCTURE BDD
type Post struct {
	DB *sql.DB
}

// ANCHOR - AddPost
func (post *Post) AddPost(id int, text string, user string, filtre string) {
	stmt, _ := post.DB.Prepare(`INSERT INTO "posts"(ID, Text, User, Filter) values(?, ?, ?,?)`)
	_, err := stmt.Exec(id, text, user, filtre)
	fmt.Println(err)
}

// ANCHOR - InitialisationPost

// crée des arrays à la bonne taille
func InitialisationPost(E *EntryPost, ID int) {
	E.ID = make([]int, ID+1)
	E.Texts = make([]string, ID+1)
	E.Users = make([]string, ID+1)
	E.Filters = make([]string, ID+1)

	E.Filter = make([]string, 1)
	E.User = make([]string, 1)
}

// ANCHOR - MajPost

// Passe par tous les postes dans la BDD et les stock dans la struct
func MajPost(E *EntryPost, update *Post, ID int, User string, Filtre string) {
	for i := 1; i <= ID; i++ {
		err := update.DB.QueryRow("SELECT Text,User,Filter from posts WHERE ID = ?", i).Scan(&E.Text, &User, &Filtre)
		E.Texts[i-1] = E.Text
		E.Users[i-1] = User
		E.ID[i-1] = i
		E.Filters[i-1] = Filtre
		if err != nil {
			panic(err)
		}
	}
}

func (post *Post) IsLiked(E *EntryPost, username string, ID int) {
	var l string
	post.DB.QueryRow("SELECT like from likes WHERE postid = ? AND username = ?", ID, username).Scan(&l)

	if l == "true" {
		E.IsLiked = true
	} else {
		E.IsLiked = false
	}
}
