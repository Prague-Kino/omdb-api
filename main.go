package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Prague-Kino/omdb-api/api"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	key := getAPIKey()
	omdb, err := api.NewOMDb(key)

	movie, err := omdb.FetchMovie("Die my love")
	if err != nil {
		log.Fatal(err)
	}

	prettyPrint(movie)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .ENV:  ", err)
	}
}

func getAPIKey() string {
	apiKey := os.Getenv("OMDB_API_KEY")
	if apiKey == "" {
		log.Fatal("OMDB_API_KEY not set in .env file")
	}
	return apiKey
}

func prettyPrint(i any) {
	s, _ := json.MarshalIndent(i, "", "    ")
	fmt.Println(string(s))
}
