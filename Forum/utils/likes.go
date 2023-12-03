package forum

import (
	"fmt"
)

type Like struct {
	ID   int
	User []string
	Like []string
}

//ANCHOR - InitialisationUsers

// crée des arrays à la bonne taille
func InitialisationLikes(L *Like, ID int) {
	L.User = make([]string, ID+1)
	L.Like = make([]string, ID+1)
}

// ANCHOR - LikeNewPost
func (post *Post) LikeNewPost(postID int, U *EntryUser) {
	var p string
	post.DB.QueryRow("SELECT postid from likes WHERE postid = ?", postID).Scan(&p)
	fmt.Println("p = ", p)
	if len(p) == 0 {
		for i := 0; i < len(U.Users); i++ {
			stmt, _ := post.DB.Prepare(`INSERT INTO "likes" (postid, username, like) values (?, ?, ?)`)
			stmt.Exec(postID, U.Users[i], "0")
		}
	}
}

func (post *Post) LikeNewUser(username string, E *EntryPost) {
	for i := 1; i < len(E.Texts); i++ {
		stmt, err := post.DB.Prepare(`INSERT INTO "likes" (postid, username, like) values (?, ?, ?)`)
		if err != nil {
			panic(err)
		}
		stmt.Exec(i, username, "0")
	}
}

func (post *Post) LikeUser(U *EntryUser, E *EntryPost) {
	for i := 0; i < len(U.Users); i++ {
		for j := 1; j < len(E.Texts); j++ {
			exist := ""
			post.DB.QueryRow("SELECT username from likes WHERE postid = ? AND username = ?", j, U.Users[i]).Scan(&exist)
			if len(exist) == 0 {
				stmt, _ := post.DB.Prepare(`INSERT INTO "likes" (postid, username, like) values (?, ?, ?)`)
				stmt.Exec(j, U.Users[i], "0")
			} else {
			}

		}
	}
	stmt, _ := post.DB.Prepare(`DELETE from likes WHERE username = "";`)
	stmt.Exec()

}

// ANCHOR - AddLike
func (post *Post) AddLike(postID any, username string) {
	stmt, _ := post.DB.Prepare(`UPDATE likes SET like = 1 WHERE postid = ? AND username = ?`)
	stmt.Exec(postID, username)
}

// ANCHOR - RemoveLike
func (post *Post) RemoveLike(postID any, username string) {
	stmt, _ := post.DB.Prepare(`UPDATE likes SET like = 0 WHERE postid = ? AND username = ?`)
	stmt.Exec(postID, username)
}

// ANCHOR - IsLiked
func IsLiked(U *EntryUser, E *EntryPost, C *EntryComment, update *Post, ID int, username, like string) {
	if like == "false" {
		E.IsLiked = false
		C.PostIsLiked = false
		update.RemoveLike(ID, username)
		C.PostLikes = NbLikesPost(ID, U, update)
	} else if like == "true" {
		E.IsLiked = true
		C.PostIsLiked = true
		update.AddLike(ID, username)
		C.PostLikes = NbLikesPost(ID, U, update)
	}
}

// ANCHOR - GetLikesFromPost
// Passe par tous les users dans la BDD et les stocke dans la str
func GetLikesFromPost(L *Like, update *Post, postID int) {
	L.ID = postID
	var users string
	var likes string
	err := update.DB.QueryRow("SELECT username,like from likes WHERE postid = ?", postID).Scan(&users, &likes)
	if err != nil {
		println("empty")
	}
	InitialisationLikes(L, len(users))
	for i := 0; i < len(users); i++ {
		L.User[i] = users
		L.Like[i] = likes
	}

}

func NbLikesPost(postID int, U *EntryUser, update *Post) int {
	var isLiked int
	err := update.DB.QueryRow(`SELECT SUM(like) FROM likes WHERE postid = ?`, postID).Scan(&isLiked)
	if err != nil {
		panic(err)
	}
	println("is l : ", isLiked)

	return isLiked
}

func MajLikes(L *Like, update *Post) {
	var username string
	var like string
	reste := true
	i := 0
	for {

		rows, err := update.DB.Query("SELECT username from likes")
		rows.Scan(&username)
		println("like ouuuu : ", like, " p ", username, " i ", i)
		L.User[i] = username
		L.Like[i] = like
		if err != nil {
			panic(err)
		}
		i++
		if !reste {
			break
		}
	}
}
