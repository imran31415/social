package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"social/app/grpc"
	"social/repo"
	"time"
)

const (
	tenmplatesDir = "./frontend/src/static"
)

func main() {
	// Serve static files from the frontend/dist directory.
	var err error
	go func() {
		// TODO: replace with env
		host := os.Getenv("DB_HOST")
		dbname := os.Getenv("DB_NAME")
		pw := os.Getenv("DB_PASS")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		r, rErr := repo.NewRepo(&repo.Config{
			DbPass: pw,
			DbUser: user,
			DbName: dbname,
			DbHost: host,
			DbPort: port,
		})
		if rErr != nil {
			log.Fatal(rErr)
			return
		}
		err = grpc.Run(":50053", grpc.NewServer(r))
	}()
	time.Sleep(time.Second * 2)
	if err != nil {
		log.Fatal(err)
		return
	}
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(tenmplatesDir))))

	// Start the server.
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
