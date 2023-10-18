package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var ctx context.Context
var client *firestore.Client

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	opt := option.WithCredentialsFile("playbook-e092e-firebase-adminsdk-gf19j-5d351e19fd.json")
	ctx = context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing firestore: %v\n", err)
	}
	defer client.Close()

	log.Printf("Welcome to the todo app!\n")
	scanner := bufio.NewScanner(os.Stdin)
	log.Printf("List of commands: list, create, read, update, delete\n")
	scanner.Scan()
	log.Printf("Enter command: ")
	command := scanner.Text()
	for command != "exit" {
		log.Printf("command: %v\n", command)
		switch command {
		case "list":
			log.Printf("List of todos: \n")
			todos, err := GetAll()
			if err != nil {
				log.Fatalf("error getting all todos: %v\n", err)
			}
			for _, t := range todos {
				log.Printf("%v\n", t)
			}
		case "create":
			log.Printf("Enter title: ")
			scanner.Scan()
			title := scanner.Text()
			log.Printf("Enter done: ")
			scanner.Scan()
			done := scanner.Text()
			t := &Todo{
				Title: title,
				Done:  done == "true",
			}
			err := Create(t)
			if err != nil {
				log.Fatalf("error creating todo: %v\n", err)
			}
			log.Printf("Todo created: %v\n", t)
		case "read":
			log.Printf("Enter id: ")
			scanner.Scan()
			id := scanner.Text()
			t, err := Read(id)
			if err != nil {
				log.Fatalf("error reading todo: %v\n", err)
			}
			log.Printf("Todo read: %v\n", t)
		case "update":
			log.Printf("Enter id: ")
			scanner.Scan()
			id := scanner.Text()
			log.Printf("Enter title: ")
			scanner.Scan()
			title := scanner.Text()
			log.Printf("Enter done: ")
			scanner.Scan()
			done := scanner.Text()
			t := &Todo{
				Title: title,
				Done:  done == "true",
			}
			err := Update(id, t)
			if err != nil {
				log.Fatalf("error updating todo: %v\n", err)
			}
			log.Printf("Todo updated: %v\n", t)
		case "delete":
			log.Printf("Enter id: ")
			scanner.Scan()
			id := scanner.Text()
			err := Delete(id)
			if err != nil {
				log.Fatalf("error deleting todo: %v\n", err)
			}
		case "default":
			log.Printf("bye!\n")
		}
		log.Printf("List of commands: list, create, read, update, delete\n")
		scanner.Scan()
		log.Printf("Enter command: ")
		command = scanner.Text()
	}
}
