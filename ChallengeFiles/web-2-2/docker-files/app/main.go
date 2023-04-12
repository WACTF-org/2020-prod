package main

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

var database *sql.DB
var err error

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func setup() {
	_, err = copy("./trumpme.db", "./trumpme_temp.db")
	if err != nil {
		// file busted
		log.Fatal(err)
	}
	database, err = sql.Open("sqlite3", "./trumpme_temp.db")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/voterportal", VoterPortal)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/vote", Vote)
	http.HandleFunc("/results", Results)
	http.HandleFunc("/robots.txt", Robots)
	http.HandleFunc("/reset", Reset)
	http.HandleFunc("/", Home)

	setup()
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var epoch = time.Unix(0, 0).Format(time.RFC1123)
var noCacheHeaders = map[string]string{
	"Expires":         epoch,
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

type Page struct {
	Filename string
	Body     []byte
}

func loadPage(filename string) (*Page, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Filename: filename, Body: body}, nil
}

// Vote endpoint - this should be the final stage after the vote has logged in and their vote validated. This submits and updates the db table

type HomePageVars struct {
	Items          []Item
	AllVotes       int
	MaxVotes       int
	RemainingVotes int
	PollStatus     string
}

type Item struct {
	Name  string
	Count int
}

var targetCandidate = "LA Ice Cola"
var MAX_VOTES = 500

func Home(w http.ResponseWriter, r *http.Request) {
	for _, v := range etagHeaders {
		if r.Header.Get(v) != "" {
			r.Header.Del(v)
		}
	}

	// Set our NoCache headers
	for k, v := range noCacheHeaders {
		w.Header().Set(k, v)
	}

	var name string
	var count int
	var stuff = HomePageVars{Items: []Item{}, AllVotes: 0, MaxVotes: MAX_VOTES, RemainingVotes: 0}

	stmt, err := database.Prepare("select count(candidate) from votes")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&stuff.AllVotes)
		if err != nil {
			// log.Println(err)
			rows.Close()
			return
		}
	}
	rows.Close()

	stuff.RemainingVotes = stuff.MaxVotes - stuff.AllVotes // this should be 1 but we want it to underflow :)
	// Home page - show current voter tally, and a login form to vote
	stmt, err = database.Prepare("select candidate, count(candidate) as c from votes GROUP BY candidate ORDER BY c DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement
	rows, err = stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name, &count)
		if err != nil {
			// log.Println(err)
		}
		// fmt.Println(name, count)
		stuff.Items = append(stuff.Items, Item{Name: name, Count: count})
		// if count > winnerCount {
		// 	winnerCount = count
		// }
	}
	rows.Close()
	stuff.PollStatus = fmt.Sprintf("<h2>Polling is still open</h2><p>There are <b>[%d] (out of [%d])</b> votes to be cast before the election closes!</p>", stuff.RemainingVotes, MAX_VOTES)

	if stuff.RemainingVotes <= 0 {
		stuff.PollStatus = fmt.Sprintf("<h2>Polling has closed for this election!</h2><p>All <b>[%d] (out of [%d])</b> votes were cast and the results are in. Head to the results page for the outcome</p>", stuff.AllVotes, MAX_VOTES)

	}

	// log.Print(stuff)
	t, err := template.ParseFiles("html/home.html") //parse the html file homepage.html
	if err != nil {                                 // if there is an error
		// log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, stuff) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {           // if there is an error
		// log.Print("template executing error: ", err) //log it
	}
	// p, _ := loadPage("html/home.html")
	// fmt.Fprintf(w, string(p.Body))
}

type DataBoi struct {
	Data string
}

func Results(w http.ResponseWriter, r *http.Request) {
	for _, v := range etagHeaders {
		if r.Header.Get(v) != "" {
			r.Header.Del(v)
		}
	}

	// Set our NoCache headers
	for k, v := range noCacheHeaders {
		w.Header().Set(k, v)
	}

	var allVotes int
	var name string
	var count int
	stmt, err := database.Prepare("select count(candidate) from votes")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement
	rows, err := stmt.Query()
	for rows.Next() {
		err := rows.Scan(&allVotes)
		if err != nil {
			// log.Println(err)
			rows.Close()
			return
		}
	}
	rows.Close()

	remainingVotes := MAX_VOTES - allVotes // this should be 1 but we want it to underflow :)
	stmt, err = database.Prepare("select candidate, count(candidate) as c from votes GROUP BY candidate ORDER BY c DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement
	rows, err = stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var winnerCount int = 0
	var winnerName string = ""
	for rows.Next() {
		err := rows.Scan(&name, &count)
		if err != nil {
			// log.Println(err)
		}
		// fmt.Println(name, count)
		if count > winnerCount {
			winnerName = name
			winnerCount = count
		}
	}
	rows.Close()
	data := DataBoi{Data: ""}
	if remainingVotes <= 0 {
		// log.Println(winnerName, targetCandidate)
		if winnerName == targetCandidate {
			// log.Println("THE WINNER IS: ", winnerName)
			data.Data = WinCondition()

		} else {
			data.Data = LoseCondition()
		}
	} else {
		data.Data = fmt.Sprintf("Voting is still open: [%d] votes are still outstanding.\n", remainingVotes)
	}
	t, err := template.ParseFiles("html/results.html") //parse the html file homepage.html
	if err != nil {                                    // if there is an error
		// log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, data) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {          // if there is an error
		// log.Print("template executing error: ", err) //log it
	}
	return
}

// Reset flushes the db, once the user has sufficiently failed the challenge
func Reset(w http.ResponseWriter, r *http.Request) {
	for _, v := range etagHeaders {
		if r.Header.Get(v) != "" {
			r.Header.Del(v)
		}
	}

	// Set our NoCache headers
	for k, v := range noCacheHeaders {
		w.Header().Set(k, v)
	}

	// Do the thing to reset the db - call the sql init script?
	// Using sqlite - probs just rm the current db and cp the original across
	// Finally redirect to home
	database.Close()
	setup()
	http.Redirect(w, r, "/home", 301)
}

// CheckAuthenticated validates the user's session cookie to determine if they're authorised to access the admin pages
func CheckAuthenticated(r *http.Request) bool {

	database.Ping()

	count := 0
	stmt, err := database.Prepare("select count(*) from sessions where val = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement

	cookie, err := r.Cookie("session")
	if err != nil {
		return false
	}
	rows, err := stmt.Query(cookie.Value)
	if err != nil {
		// log.Println(err)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&count)
		if count < 1 {
			// not authenticated
			rows.Close()
			return false
		}
	}
	rows.Close()
	stmt.Close()
	stmt, err = database.Prepare("select userid, val, expires from sessions where val = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement

	rows, err = stmt.Query(cookie.Value)
	if err != nil {
		// log.Println(err)
		rows.Close()
		return false
	}
	currTime := time.Now()
	defer rows.Close()

	var user string
	var val string
	var expire string
	layout := "2006-01-02T15:04:05.000Z"

	for rows.Next() {
		err := rows.Scan(&user, &val, &expire)
		if err != nil {
			// log.Println(err)
			rows.Close()

			return false
		}

		if x, err := time.Parse(layout, expire); currTime.After(x) {
			if err != nil {
				// log.Println(err)
			}
			rows.Close()

			return false
		}

	}
	rows.Close()

	return true
}

// Login authenticates a user if they're not already authenticated
func Login(w http.ResponseWriter, r *http.Request) {
	for _, v := range etagHeaders {
		if r.Header.Get(v) != "" {
			r.Header.Del(v)
		}
	}

	// Set our NoCache headers
	for k, v := range noCacheHeaders {
		w.Header().Set(k, v)
	}

	if CheckAuthenticated(r) && r.Method == "GET" {
		http.Redirect(w, r, "/voterportal", 301)
		return
	}
	r.ParseForm()

	err = database.Ping()
	if err != nil {
		// log.Println("Error:", err)
		http.Redirect(w, r, "/home", 301)
		return
	}
	stmt, err := database.Prepare("select count(id) from users where username = ? and password = ?")
	if err != nil {
		// log.Println("Error:", err)
		http.Redirect(w, r, "/home", 301)
		return
	}
	defer stmt.Close() // closing the statement
	rows, err := stmt.Query(r.FormValue("username"), r.FormValue("password"))

	if err != nil {
		// log.Println("Error:", err)
		http.Redirect(w, r, "/home", 301)
		return
	}
	defer rows.Close()

	var ck http.Cookie
	var count int
	for rows.Next() {
		rows.Scan(&count)
		if count == 0 {
			ck.Name = "msg"
			ck.Value = "PASSWORD_WRONG"
			http.SetCookie(w, &ck)
			http.Redirect(w, r, "/home", 301)
			rows.Close()

			return
		}
	}
	rows.Close()
	stmt.Close()

	stmt, err = database.Prepare("select id from users where username = ? and password = ?")
	if err != nil {
		// log.Println("Error:", err)
		http.Redirect(w, r, "/home", 301)
		return
	}
	defer stmt.Close() // closing the statement
	rows, err = stmt.Query(r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		// log.Println("Error:", err)
		http.Redirect(w, r, "/home", 301)
		return
	}
	stmt.Close()
	defer rows.Close()
	var id int
	var nonce uuid.UUID
	var expires time.Time

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

		expires = time.Now().AddDate(1, 0, 0)
		ck = http.Cookie{
			Name:    "session",
			Domain:  "",
			Path:    "/",
			Expires: expires,
		}
		// value of cookie
		u, err := uuid.NewV4()
		if err != nil {
			log.Fatal(err)
		}
		ck.Value = fmt.Sprintf("%s", u)
		// write the cookie to response
		// // log.Print("Valid user!")

		nonce, err = uuid.NewV4() // Generate voter nonce
		if err != nil {
			log.Println(err)
		}
	}
	rows.Close()
	stmt, err = database.Prepare("INSERT INTO sessions VALUES(NULL, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement

	// log.Println(id, "session", ck.Value, expires, nonce.String())

	expire_s := expires.Format("2006-01-02T15:04:05.000Z")
	res, err := stmt.Exec(id, "session", ck.Value, expire_s, nonce.String())
	if err != nil {
		log.Println(err)
	}
	stmt.Close()
	if num, err := res.RowsAffected(); num > 0 {
		if err != nil {
			log.Fatal(err)
		}
		// log.Println("hooray")
	}
	stmt.Close()
	http.SetCookie(w, &ck)
	http.Redirect(w, r, "/voterportal", 301)
	return
}

type PageVariables struct {
	Candidates []string
	Nonce      string
	Name       string
}

func WinCondition() string {
	// User should be logged in to see this :O
	return "<h1>Congratulations!</h1><p>Your candidate won!\n<b>WACTF{democracy_wins_everytime}</b>"

}

func LoseCondition() string {
	// User should be logged in to see this :O
	return "<h1>Too bad :(</h1><p>Your vote didn't matter after all. <b>Your candidate lost.</b></p><p>I guess democracy isn't all that it's cracked up to be?</p>"

}

func VoterPortal(w http.ResponseWriter, r *http.Request) {
	for _, v := range etagHeaders {
		if r.Header.Get(v) != "" {
			r.Header.Del(v)
		}
	}

	// Set our NoCache headers
	for k, v := range noCacheHeaders {
		w.Header().Set(k, v)
	}

	// User should be logged in to see this :O
	if !CheckAuthenticated(r) {
		http.Redirect(w, r, "/home", 301)
		return
	}

	// Need to fix this
	err = database.Ping()
	if err != nil {
		// log.Println(err)
	}
	stmt, err := database.Prepare("select count(`votes`.`voter`) from `sessions` INNER JOIN `votes` ON `sessions`.`userid` = `votes`.`voter` WHERE `sessions`.`val` = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement
	cookie, err := r.Cookie("session")
	if err != nil {
		// log.Println(err)
		return
	} // closing the statement

	rows, err := stmt.Query(cookie.Value)
	if err != nil {
		// log.Println(err)
		return
	} // c
	votecount := 0
	for rows.Next() {
		err := rows.Scan(&votecount)
		if err != nil {
			log.Fatal(err)
		}
		if votecount > 0 {

			rows.Close() // Already voted
			fmt.Fprint(w, "<html> <link rel=\"stylesheet\" href=\"/assets/w3.css\"> <link rel=\"stylesheet\" href=\"/assets/raleway.css\"> <body class=\"w3-light-grey\"> <style> table, th, td {border: 0px solid black; border-collapse: collapse; } table.center {margin-left: auto; margin-right: auto; } </style> <div class=\"w3-center\"> <header class=\"w3-container w3-center w3-padding-32\"> <h1>Sorry</h1> </header> </div> <div class=\"w3-card-4 w3-margin w3-white w3-center\"> You've already voted in this election. Be a real democracy next time :\\ <p>Check the <a href=\"/results\">Results?</a> </p> </div> </body> </html>")
			// http.Redirect(w, r, "/home", 301)
			return
		}
	}
	rows.Close()

	stmt.Close()
	stmt, err = database.Prepare("select distinct(candidate) from votes")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement

	rows, err = stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	var candidate string
	candidates := []string{}

	for rows.Next() {
		err := rows.Scan(&candidate)
		if err != nil {
			log.Fatal(err)
		}
		candidates = append(candidates, candidate)
	}
	rows.Close()

	stuff := PageVariables{Candidates: candidates}
	stmt, err = database.Prepare("select `users`.`name`, `sessions`.`nonce` from `sessions` INNER JOIN `users` ON `sessions`.`userid` = `users`.`id` WHERE `sessions`.`val` = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err = stmt.Query(cookie.Value) // Supplied nonce and cookie form the query
	if err != nil {
		// log.Println(err)
	}
	var nonce string
	var name string
	for rows.Next() {
		err := rows.Scan(&name, &nonce)
		if err != nil {
			log.Fatal(err)
		}
		stuff.Nonce = nonce
		stuff.Name = name
	}
	rows.Close()

	// log.Println("lol", stuff)

	t, err := template.ParseFiles("html/voterportal.html") //parse the html file homepage.html
	if err != nil {                                        // if there is an error
		// log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, stuff) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {           // if there is an error
		// log.Print("template executing error: ", err) //log it
	}
	return

}

func Robots(w http.ResponseWriter, r *http.Request) {
	// Troll hacker with robots.txt :D
	fmt.Fprintf(w, "Usr-agent: *\nDisallw: /lol\n")
}

type Thanks struct {
	Candidate string
}

func Vote(w http.ResponseWriter, r *http.Request) {
	for _, v := range etagHeaders {
		if r.Header.Get(v) != "" {
			r.Header.Del(v)
		}
	}

	// Set our NoCache headers
	for k, v := range noCacheHeaders {
		w.Header().Set(k, v)
	}

	if !CheckAuthenticated(r) {
		// log.Println("not authed")
		http.Redirect(w, r, "/home", 301)
		return
	}

	if r.Method != "POST" {
		http.Redirect(w, r, "/voterportal", 301)
		return
	}
	// p, _ := loadPage("html/thanks.html")
	r.ParseForm()

	stmt, err := database.Prepare("select `users`.`username`, `users`.`id` from `users` INNER JOIN `sessions` ON `users`.`id` = `sessions`.`userid` WHERE `sessions`.`val` = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement
	cookie, err := r.Cookie("session")
	if err != nil {
		// log.Println(err)

		return
	}
	rows, err := stmt.Query(cookie.Value)
	if err != nil {
		// log.Println(err)
		return
	}
	defer rows.Close()
	var username string
	var userid int
	for rows.Next() {
		err := rows.Scan(&username, &userid)
		if err != nil {
			log.Fatal(err)
		}
	}
	rows.Close()
	// log.Println("Vote for:", username, userid, r.FormValue("candidate"))

	// Need to check that the nonce is not only valid, but has not been used previously
	// First we check that the supplied nonce is tied to the user's actual session
	stmt, err = database.Prepare("select count(*) from `sessions` WHERE `sessions`.`nonce` = ? AND `sessions`.`val` = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()                                         // closing the statement
	rows, err = stmt.Query(r.FormValue("nonce"), cookie.Value) // Supplied nonce and cookie form the query
	if err != nil {
		// log.Println(err)
	}
	num := 0
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&num)
		if num <= 0 { // Nonce isn't valid
			// InvalidateSession(cookie.Value, userid)
			var ck http.Cookie
			ck.Name = "msg"
			ck.Value = "TAMPERING_DETECTED"
			http.SetCookie(w, &ck)
			var ck2 http.Cookie
			ck2.Name = "session"
			ck2.Value = ""
			http.SetCookie(w, &ck)
			http.Redirect(w, r, "/home", 301)
			// log.Println("invalid nonce:", r.FormValue("nonce"), cookie.Value)
			rows.Close()
			return

		}

	}
	rows.Close()
	// Now we check that the nonce hasn't been tied to a vote in the DB already
	stmt, err = database.Prepare("select count(*) from votes WHERE nonce = ? AND voter = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement
	// log.Println("here")
	rows, err = stmt.Query(r.FormValue("nonce"), userid)
	if err != nil {
		return
		// log.Println(err)
	}
	count := 0
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&count)
		if count != 0 {
			// Invalidate session
			// InvalidateSession(cookie.Value, userid)
			var ck http.Cookie
			ck.Name = "msg"
			ck.Value = "TAMPERING_DETECTED"
			http.SetCookie(w, &ck)
			var ck2 http.Cookie
			ck2.Name = "session"
			ck2.Value = ""
			http.SetCookie(w, &ck)
			rows.Close()
			http.Redirect(w, r, "/home", 301)
			// log.Println("nonce reuse:", r.FormValue("nonce"), cookie.Value)
			rows.Close()
			return
		}
	}
	rows.Close()
	// log.Println(username, r.FormValue("nonce"), cookie.Value, " TESTING ")

	stmt, err = database.Prepare("INSERT INTO votes VALUES(NULL, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // closing the statement
	res, err := stmt.Exec(r.FormValue("nonce"), userid, r.FormValue("candidate"))
	if err != nil {

		// log.Println(err)
		return
	}
	if num, _ := res.RowsAffected(); num != 1 {
		// log.Println(num, "rows affected")
	}
	t, err := template.ParseFiles("html/thanks.html") //parse the html file homepage.html
	if err != nil {                                   // if there is an error
		// log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, Thanks{Candidate: r.FormValue("candidate")}) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                                                 // if there is an error
		// log.Print("template executing error: ", err) //log it
	}
}
