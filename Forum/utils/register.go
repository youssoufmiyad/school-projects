package forum

//ANCHOR - STRUCTURE USER
type EntryUser struct {
	Email      string
	Emails     []string
	Username   string
	Users      []string
	Password   string
	Passwords  []string
	ID         int
	IDs        []int
	IsLogged   bool
	IsSignedIn bool
}

//ANCHOR - InitialisationUsers

// crée des arrays à la bonne taille
func InitialisationUsers(U *EntryUser, ID int) {
	U.Users = make([]string, ID+1)
	U.Emails = make([]string, ID+1)
	U.Passwords = make([]string, ID+1)
	U.IDs = make([]int, ID+1)
}

//ANCHOR - AddUser
func (post *Post) AddUser(username, email, password string, ID int) {
	stmt, _ := post.DB.Prepare(`INSERT INTO "users" (email, username, password,ID) values (?, ?, ?,?)`)
	stmt.Exec(email, username, password, ID)
}

//ANCHOR - MajUsers

// Passe par tous les users dans la BDD et les stocke dans la struct
func MajUsers(U *EntryUser, update *Post, ID int, email, password, username string) {
	for i := 1; i <= ID; i++ {
		err := update.DB.QueryRow("SELECT email,username,password from users WHERE ID = ?", i).Scan(&email, &username, &password)
		U.Emails[i-1] = email
		U.Users[i-1] = username
		U.Passwords[i-1] = password
		U.IDs[i-1] = i
		if err != nil {
			panic(err)
		}
	}
}
