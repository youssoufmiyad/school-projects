package forum

import (
	"fmt"
	"strconv"
)

type EntryComment struct {
	ID          []string
	IDs         []string
	CommentID   []int
	Texts       []string
	Text        string
	User        []string
	PostText    string
	PostUser    string
	PostIsLiked bool
	PostLikes   int
	IsLogged    bool
}

func (post *Post) AddComment(id string, commentID int, text string, user string) {
	stmt, _ := post.DB.Prepare(`INSERT INTO "comment"(ID,CommentID, Text, User) values(?,?, ?, ?)`)
	_, err := stmt.Exec(id, commentID, text, user)
	fmt.Println(err)
}

func InitialisationComment(C *EntryComment, CommentID int) {
	C.CommentID = make([]int, CommentID+1)
	C.Texts = make([]string, CommentID+1)
	C.User = make([]string, CommentID+1)
	C.IDs = make([]string, CommentID+1)
	C.ID = make([]string, 1)
}

func MajComment(C *EntryComment, update *Post, CommentID int, User string) {
	InitialisationComment(C, CommentID)
	for i := 1; i <= CommentID; i++ {
		err := update.DB.QueryRow("SELECT Text,User,ID from comment WHERE CommentID = ?", i).Scan(&C.Text, &User, &C.IDs[i-1])
		if err == nil {
			C.Texts[i-1] = C.Text
			C.User[i-1] = User
			C.CommentID[i-1] = i
		}
	}
}

func LienCommentPost(C *EntryComment, E *EntryPost, ID int) *EntryComment {
	C.ID[0] = strconv.Itoa(ID)
	if ID != 0 {
		C.PostText = E.Texts[ID-1]
		C.PostUser = E.Users[ID-1]
	} else {
		C.PostText = E.Texts[ID]
		C.PostUser = E.Users[ID]
	}
	return C
}
