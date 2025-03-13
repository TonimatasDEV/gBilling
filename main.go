package main

import (
	"context"
	"errors"
	"github.com/TonimatasDEV/BillingPanel/src/database"
	"github.com/TonimatasDEV/BillingPanel/src/pages"
	"github.com/TonimatasDEV/BillingPanel/src/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	utils.LoadEnvFile()
	utils.InitCloudflare()
	db := database.Connect()
	database.CreateTables()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	utils.HandleFunc("/", src.IndexHandler, false)
	utils.HandleFunc("/logout", src.LogoutHandler, true)
	utils.HandleFunc("/login", src.LoginHandler, true)

	server := &http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if os.Getenv("MODE") == "dev" {
			log.Println("Web: http://localhost:8080")

			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalln("Server crashed:", err)
			}
		} else if os.Getenv("MODE") == "production" {
			domainSSL := "/etc/letsencrypt/live/" + os.Getenv("DOMAIN") + "/"

			if err := server.ListenAndServeTLS(domainSSL+"fullchain.pem", domainSSL+"privkey.pem"); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalln("Server crashed:", err)
			}
		} else {
			log.Fatalln("Server crashed: MODE not supported.")
		}
	}()

	log.Println("Server listening on :" + os.Getenv("PORT") + ".")

	<-stop
	log.Println("Server shutting down.")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		err := db.Close()
		if err != nil {
			log.Println("Error occurred closing the database:", err)
		}

		log.Println("Error stopping the server:", err)
	}

	log.Println("Server stopped successfully.")
}
