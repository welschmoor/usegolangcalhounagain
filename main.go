package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/welschmoor/usegolangcalhounagain/models"

	"github.com/jackc/pgx/v5"
)

type PostgresConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func (pc *PostgresConfig) String() string {
	return fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		pc.host,
		pc.port,
		pc.user,
		pc.password,
		pc.dbname,
		pc.sslmode,
	)
}

func main() {
	// cnnstring := `host=localhost port=5432 user=rootuser password=passwort dbname=anzeigen sslmode=disable`
	// cnnstring2 := "postgres://rootuser:passwort@localhost:5432/anzeigen?sslmode=disable"
	cnnstring := PostgresConfig{
		host:     "localhost",
		port:     "5432",
		user:     "rootuser",
		password: "passwort",
		dbname:   "anzeigen",
		sslmode:  "disable",
	}

	conn, err := pgx.Connect(
		context.Background(),
		cnnstring.String(),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println("err:::", err)
		os.Exit(1)
	} else {
		fmt.Println("Ping success")
	}

	// _, err = conn.Exec(context.Background(), `
	// CREATE TABLE IF NOT EXISTS users (
	// 	id SERIAL PRIMARY KEY,
	// 	first_name TEXT,
	// 	last_name TEXT,
	// 	email TEXT UNIQUE NOT NULL
	// );

	// CREATE TABLE IF NOT EXISTS orders (
	// 	id SERIAL PRIMARY KEY,
	// 	user_id INT NOT NULL,
	// 	amount INT,
	// 	description TEXT
	// );`,
	// )
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Tables created:::")

	userService := models.UserService{
		DB: conn,
	}
	userService.CreateUser(
		"1111@awdm.com",
		"1111",
		"11111",
		"111pass",
		time.Now(),
	)

	r := chi.NewRouter()
	r.Get("/", homeHandler)
	http.ListenAndServe(":4000", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello, Bitches!")
	w.Write([]byte("Hello"))
}
