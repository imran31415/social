package main

import (
	"fmt"
	"log"
	"net/http"
	"social/app/grpc"
	"social/repo"
	"time"
)

const (
	tenmplatesDir = "./frontend/src/static"
)

type Page struct {
	Title string
}

func main() {
	// Serve static files from the frontend/dist directory.
	var err error
	go func() {
		// TODO: replace with env variables
		r, err := repo.NewRepo(&repo.Config{
			DbPass: "13tg1t8bqfsa76u",
			DbUser: "root",
			DbName: "db",
			DbHost: "localhost",
			DbPort: "3306",
		})
		if err != nil {
			log.Fatal(err)
			return
		}
		err = grpc.Run("localhost:50053", grpc.NewServer(r))
	}()
	time.Sleep(time.Second * 2)
	if err != nil {
		log.Fatal(err)
		return
	}
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(tenmplatesDir))))

	// Start the server.
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
