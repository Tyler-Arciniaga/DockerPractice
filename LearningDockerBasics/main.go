package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Item struct {
	ID         int
	Name       string
	Created_at time.Time
}

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading environment variables")
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("Error connecting to database: DockerPractice", "err", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Was Pinged!")
		newItem := fmt.Sprintf("I Love God - %s", time.Now().Format(time.DateTime))
		_, err := conn.Exec(context.Background(), "INSERT INTO ITEMS (name, created_at) VALUES ($1, $2)", newItem, time.Now())
		if err != nil {
			slog.Error("Error writing to DB", "err", err)
			os.Exit(1)
		}

		rows, err := conn.Query(context.Background(), "SELECT * FROM ITEMS")
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

	slog.Info("Server listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
