package main

import "log"

func main() {
	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8081", store)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
