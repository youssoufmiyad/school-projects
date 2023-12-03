package forum

func CheckUserExist(U *EntryUser, E *EntryPost, C *EntryComment, UserID int, username string, password string) {
	U.IsSignedIn = false
	U.IsLogged = false
	E.IsLogged = false
	C.IsLogged = false
	for i := 0; i < UserID; i++ {
		if U.Users[i] == username && U.Passwords[i] == password {
			U.IsLogged = true
			E.IsLogged = true
			C.IsLogged = true
		} else if U.Emails[i] == username && U.Passwords[i] == password {
			U.IsLogged = true
			E.IsLogged = true
			C.IsLogged = true
		}
	}
}
