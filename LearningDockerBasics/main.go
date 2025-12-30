package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Item struct {
	ID         int
	Name       string
	Created_at time.Time
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		slog.Error("Error getting container hostname", "err", err)
	}

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM, syscall.SIGINT)

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("Error connecting to database: DockerPractice", "err", err)
		os.Exit(1)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprint(hostname, "health Check Pinged!"))
		if err := dbpool.Ping(context.Background()); err != nil {
			slog.Error("DB ping failed!", "err", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprint(hostname, "was Pinged!"))
		newItem := fmt.Sprintf("I Love God - %s", time.Now().Format(time.DateTime))
		_, err := dbpool.Exec(context.Background(), "INSERT INTO ITEMS (name, created_at) VALUES ($1, $2)", newItem, time.Now())
		if err != nil {
			slog.Error("Error writing to DB", "err", err)
			os.Exit(1)
		}

		rows, err := dbpool.Query(context.Background(), "SELECT * FROM ITEMS")
		if err != nil {
			slog.Error("Error reading from DB", "err", err)
			os.Exit(1)
		}

		var items []Item

		for rows.Next() {
			var item Item
			err := rows.Scan(&item.ID, &item.Name, &item.Created_at)
			if err != nil {
				slog.Error("Error scanning row", "err", err)
				os.Exit(1)
			}

			items = append(items, item)
		}

		marshalledItems, err := json.MarshalIndent(items, "", " ")
		if err != nil {
			slog.Error("Error marshalling items slice", "err", err)
			os.Exit(1)
		}

		w.WriteHeader(http.StatusOK)

		w.Write(marshalledItems)
	})

	server := http.Server{Addr: ":8080", Handler: http.DefaultServeMux}

	go func() {
		slog.Info("Server listening on port :8080")
		server.ListenAndServe()
	}()

	// Wait for signal to clean up and exit gracefully
	<-gracefulShutdown

	slog.Info("Exit signal recieved!")

	// exit within a 10 second time limit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// close database connection
	dbpool.Close()

	// shutdown server gracefully
	server.Shutdown(ctx)

	slog.Info(fmt.Sprint(hostname, "has been shutdown gracefully"))
}
