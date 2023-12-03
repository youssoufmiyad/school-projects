package main

//ANCHOR - IMPORTS
import (
	"database/sql"
	"fmt"
	forum "forum/utils"
	"net/http"
	"strconv"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

// ANCHOR - VARIABLES

var update *forum.Post
var ID int
var CommentID int
var com string
var txt string = "test"
var User string
var UserID int
var filtre string
var username string
var password string
var U *forum.EntryUser
var E *forum.EntryPost
var C *forum.EntryComment
var L *forum.Like

const port = ":8080"

// ANCHOR - PAGE D'ACCUEIL
func Home(w http.ResponseWriter, r *http.Request) {
	U = &forum.EntryUser{}
	E = &forum.EntryPost{}
	C = &forum.EntryComment{}

	db, _ := sql.Open("sqlite3", "test.db")
	update = forum.NewDB(db)

	update.DB.QueryRow("SELECT MAX(ID) from users").Scan(&UserID)
	forum.InitialisationUsers(U, UserID)
	forum.MajUsers(U, update, UserID, U.Email, password, username)
	forum.CheckUserExist(U, E, C, UserID, username, password)

	update.DB.QueryRow("SELECT MAX(ID) from posts").Scan(&ID)
	forum.InitialisationPost(E, ID)
	forum.MajPost(E, update, ID, User, filtre)
	update.LikeUser(U, E)

	templates, _ := template.New("home.html").ParseFiles("templates/home.html")
	// TODO - Gérer le bouton disconnect
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "home.html", U)
	case "POST":
		return
	}
}

// ANCHOR - PAGE INSCRIPTION
func Register(w http.ResponseWriter, r *http.Request) {
	// Ouverture / création de la bdd //
	db, _ := sql.Open("sqlite3", "test.db")
	update = forum.NewDB(db)

	templates, _ := template.New("register.html").ParseFiles("templates/register.html")

	// Synchronisation struct/BDD (import from BDD)
	update.DB.QueryRow("SELECT MAX(ID) from users").Scan(&UserID)
	forum.InitialisationUsers(U, UserID)
	forum.MajUsers(U, update, UserID, U.Email, U.Password, U.Username)

	update.DB.QueryRow("SELECT MAX(ID) from posts").Scan(&ID)
	forum.InitialisationPost(E, ID)
	forum.MajPost(E, update, ID, User, filtre)

	forum.CheckUserExist(U, E, C, UserID, username, password)

	//TODO - faire que le bouton inscription renvoi les data en plus d'envoyer sur la page d'accueil (il renvoi sur une page blanche)
	switch r.Method {
	case "GET":
		// Affichage basique de la page
		templates.ExecuteTemplate(w, "register.html", U)
	case "POST":
		// Récupération des entrées
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		//TODO - implémenter une fonction de hashage pour la sécurité (BONUS)

		// Synchronisation struct/BDD (import entrées to Struct & export entrées to BDD)
		if username != "" {
			forum.StockageStructUser(U, UserID, username, email, password)
			UserID += 1
			update.AddUser(username, email, password, UserID)
			fmt.Println(U)
			update.LikeNewUser(username, E)
			U.IsSignedIn = true
		}
		templates.ExecuteTemplate(w, "register.html", U)
	}

}

// ANCHOR - LOGIN
func Login(w http.ResponseWriter, r *http.Request) {
	// Initialisation de la struct //
	U = &forum.EntryUser{IsLogged: false}

	// Ouverture / création de la bdd //
	db, _ := sql.Open("sqlite3", "test.db")
	update = forum.NewDB(db)

	// Synchronisation struct/BDD (import from BDD)
	update.DB.QueryRow("SELECT MAX(ID) from users").Scan(&UserID)
	forum.InitialisationUsers(U, UserID)
	forum.MajUsers(U, update, UserID, U.Email, U.Password, U.Username)

	// Initialisation du template //
	templates, _ := template.New("room.html").ParseFiles("templates/login.html")

	switch r.Method {
	case "GET":
		// Affichage de la page
		templates.ExecuteTemplate(w, "login.html", U)
	case "POST":
		// Récupération des entrées
		username = r.FormValue("username/mail")
		password = r.FormValue("loginPassword")
		forum.CheckUserExist(U, E, C, UserID, username, password)

		fmt.Println(U)
		fmt.Println("user = ", username, " password = ", password)
		fmt.Println("isLogged = ", U.IsLogged)
		// MAJ de la page
		templates.ExecuteTemplate(w, "login.html", U)
	}

}

// ANCHOR - ROOM (postes)
func Room(w http.ResponseWriter, r *http.Request) {
	// Initialisation de la struct //
	E = &forum.EntryPost{}
	C = &forum.EntryComment{}
	forum.CheckUserExist(U, E, C, UserID, username, password)
	// Ouverture / création de la bdd //
	db, _ := sql.Open("sqlite3", "test.db")
	update = forum.NewDB(db)

	// Synchronisation struct/BDD (import from BDD)
	update.DB.QueryRow("SELECT MAX(id) from posts").Scan(&ID)
	forum.InitialisationPost(E, ID)
	forum.MajPost(E, update, ID, User, filtre)

	// Initialisation du template //
	templates, _ := template.New("room.html").ParseFiles("templates/room.html")
	fmt.Println("E ", E)
	fmt.Println("U ", U)
	switch r.Method {
	case "GET":
		// affichage de base //
		templates.Execute(w, E)
	case "POST":
		// affichage en cas de création de postes //

		// récupération du texte (uniquement si user il y a)
		txt = r.FormValue("poste")
		if txt != "" {
			filtre = r.FormValue("filtre")
			postID, _ := strconv.Atoi(r.PostFormValue("commentRedirect"))
			println(postID)
			C = &forum.EntryComment{}

			forum.InitialisationComment(C, CommentID)
			C.ID[0] = r.FormValue("commentRedirect")

			E.Text = txt

			// Stockage dans la struct //
			forum.StockageStructPost(E, ID, txt, username, filtre)
			update.LikeNewPost(ID, U)

			ID += 1
			E.User[0] = username
			// Stockage dans la BDD //
			update.AddPost(ID, txt, username, filtre)
			update.LikeNewPost(ID, U)

			// vérification des entrées
			fmt.Println("E.Text : ", E.Text,
				"E.Texts : ", E.Texts,
				"E.ID : ", E.ID,
				"E.User : ", E.Users,
				"E.Filter : ", E.Filters)
		}

		//ANCHOR - FILTRE PAR SUJET
		E.Filter[0] = r.FormValue("searchbar")

		// Mise à jour de la page //
		error := templates.ExecuteTemplate(w, "room.html", E)
		if error != nil {
			panic(error)
		}
	}
}

// ANCHOR - ROOM (commentaires)
func RoomComment(w http.ResponseWriter, r *http.Request) {
	L = &forum.Like{}
	// Ouverture / création de la bdd //
	db, _ := sql.Open("sqlite3", "test.db")
	update = forum.NewDB(db)
	U = &forum.EntryUser{}

	update.DB.QueryRow("SELECT MAX(ID) from users").Scan(&UserID)
	forum.InitialisationUsers(U, UserID)
	forum.MajUsers(U, update, UserID, U.Email, U.Password, U.Username)

	// Synchronisation struct/BDD (import from BDD)
	update.DB.QueryRow("SELECT MAX(id) from posts").Scan(&ID)

	forum.InitialisationPost(E, ID)
	forum.MajPost(E, update, ID, User, filtre)

	update.DB.QueryRow("SELECT MAX(CommentID) from comment").Scan(&CommentID)
	forum.InitialisationComment(C, CommentID)
	forum.MajComment(C, update, CommentID, User)

	ID, _ = strconv.Atoi(r.FormValue("commentRedirect"))

	if ID != 0 {
		C = forum.LienCommentPost(C, E, ID)
	}
	forum.GetLikesFromPost(L, update, ID)
	// forum.MajLikes(L, update)

	update.IsLiked(E, username, ID)
	C.PostIsLiked = E.IsLiked

	templates, _ := template.New("roomComment.html").ParseFiles("templates/roomComment.html")
	switch r.Method {
	case "GET":
		C.PostLikes = forum.NbLikesPost(ID, U, update)
		templates.Execute(w, C)
	case "POST":

		update.DB.QueryRow("SELECT ID FROM from posts WHERE text = ?", C.PostText).Scan(&ID)

		for i := 0; i < len(L.User); i++ {
			if L.User[i] == C.PostUser && L.Like[i] == "true" {
				E.IsLiked = true
				C.PostIsLiked = true
			} else {
				E.IsLiked = false
				C.PostIsLiked = false
			}
		}
		// C.PostIsLiked = E.IsLiked
		like := r.FormValue("like")

		// affichage en cas de création de postes //
		// récupération du texte
		com = r.FormValue("commentaire")
		E.Text = com
		if com != "" {
			forum.StockageStructComment(C, ID, CommentID, com, username)
			CommentID += 1
			update.AddComment(C.ID[0], CommentID, com, username)
			fmt.Println("C.ID = ", C.ID)
			fmt.Println("E.ID = ", E.ID)
			fmt.Println("C.POSTTEXT = ", C.PostText)

		} else {
			forum.IsLiked(U, E, C, update, ID, username, like)
			templates.Execute(w, C)

		}
		println("like : ", like)

	}
	fmt.Println("C.ID = ", C.ID)
	println("is liked : ", E.IsLiked)

	// Stockage dans la BDD //

}

// ANCHOR - MAIN
func main() {
	// Handler (affecte une page à une adresse locale)
	http.HandleFunc("/", Home)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/room", Room)
	http.HandleFunc("/comment", RoomComment)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	// Démarrer le serveur HTTP
	fmt.Println("Accéder au site : http://localhost:8080")
	http.ListenAndServe(port, nil)
}
