package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	board       [6][7]string
	turn        = "R"
	winner      string
	player1Name = "Joueur Rouge"
	player2Name = "Joueur Jaune"
)

var templates = template.Must(template.ParseFiles("templates/home.html", "templates/game.html"))

func renderBoard() template.HTML {
	var sb strings.Builder
	sb.WriteString("<table>")

	sb.WriteString("<tr>")
	for c := 0; c < 7; c++ {
		disabled := winner != "" || board[0][c] != ""
		if disabled {
			sb.WriteString("<td></td>")
		} else {
			sb.WriteString(fmt.Sprintf("<td><form action='/play' method='POST'><button class='arrow' name='col' value='%d'>↓</button></form></td>", c))
		}
	}
	sb.WriteString("</tr>")

	for r := 0; r < 6; r++ {
		sb.WriteString("<tr>")
		for c := 0; c < 7; c++ {
			cls := ""
			switch board[r][c] {
			case "R":
				cls = "red"
			case "J":
				cls = "yellow"
			}
			sb.WriteString(fmt.Sprintf("<td><div class='cell %s'></div></td>", cls))
		}
		sb.WriteString("</tr>")
	}
	sb.WriteString("</table>")

	if winner != "" {
		sb.WriteString(fmt.Sprintf("<h2>%s</h2>", winner))
		sb.WriteString(`<form action="/" method="GET"><button style="margin-top:20px;">Rejouer</button></form>`)
	} else {
		currentPlayer := player1Name
		if turn == "J" {
			currentPlayer = player2Name
		}
		sb.WriteString(fmt.Sprintf("<p>Tour de : %s</p>", currentPlayer))
	}
	return template.HTML(sb.String())
}

func resetBoard() {
	board = [6][7]string{}
	turn = "R"
	winner = ""
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if winner != "" || isBoardFull() {
		resetBoard()
	}
	templates.ExecuteTemplate(w, "home.html", nil)
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		p1 := r.FormValue("p1")
		p2 := r.FormValue("p2")
		if p1 != "" {
			player1Name = p1
		} else {
			player1Name = "Joueur Rouge"
		}
		if p2 != "" {
			player2Name = p2
		} else {
			player2Name = "Joueur Jaune"
		}
	}
	resetBoard()
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	data := struct{ BoardHTML template.HTML }{BoardHTML: renderBoard()}
	templates.ExecuteTemplate(w, "game.html", data)
}

func checkWin(p string) bool {
	for r := 0; r < 6; r++ {
		for c := 0; c <= 3; c++ {
			if board[r][c] == p && board[r][c+1] == p && board[r][c+2] == p && board[r][c+3] == p {
				return true
			}
		}
	}
	for c := 0; c < 7; c++ {
		for r := 0; r <= 2; r++ {
			if board[r][c] == p && board[r+1][c] == p && board[r+2][c] == p && board[r+3][c] == p {
				return true
			}
		}
	}
	for r := 0; r <= 2; r++ {
		for c := 0; c <= 3; c++ {
			if board[r][c] == p && board[r+1][c+1] == p && board[r+2][c+2] == p && board[r+3][c+3] == p {
				return true
			}
		}
	}
	for r := 3; r < 6; r++ {
		for c := 0; c <= 3; c++ {
			if board[r][c] == p && board[r-1][c+1] == p && board[r-2][c+2] == p && board[r-3][c+3] == p {
				return true
			}
		}
	}
	return false
}

func isBoardFull() bool {
	for r := 0; r < 6; r++ {
		for c := 0; c < 7; c++ {
			if board[r][c] == "" {
				return false
			}
		}
	}
	return true
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || winner != "" {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}
	colStr := r.FormValue("col")
	col, err := strconv.Atoi(colStr)
	if err != nil || col < 0 || col > 6 {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}
	placed := false
	for r := 5; r >= 0; r-- {
		if board[r][col] == "" {
			board[r][col] = turn
			placed = true
			break
		}
	}
	if !placed {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}
	if checkWin(turn) {
		if turn == "R" {
			winner = player1Name + " a gagné ! 🎉"
		} else {
			winner = player2Name + " a gagné ! 🎉"
		}
	} else if isBoardFull() {
		winner = "Match nul : la grille est remplie 🎯"
	} else {
		if turn == "R" {
			turn = "J"
		} else {
			turn = "R"
		}
	}
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func main() {
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/play", playHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Serveur lancé sur http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
