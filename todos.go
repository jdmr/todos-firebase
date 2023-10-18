package main

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Todo struct {
	ID    string `firestore:"-"`
	Title string `firestore:"title"`
	Done  bool   `firestore:"done"`
}

func Create(t *Todo) error {
	_, _, err := client.Collection("todos").Add(ctx, t)
	return err
}

func Read(id string) (*Todo, error) {
	doc, err := client.Collection("todos").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	t := &Todo{}
	doc.DataTo(t)
	t.ID = doc.Ref.ID
	return t, nil
}

func Update(id string, t *Todo) error {
	_, err := client.Collection("todos").Doc(id).Set(ctx, t)
	return err
}

func UpdateProperties(id string, t *Todo) error {
	_, err := client.Collection("todos").Doc(id).Update(ctx, []firestore.Update{
		{Path: "title", Value: t.Title},
		{Path: "done", Value: t.Done},
	})
	return err
}

func Delete(id string) error {
	_, err := client.Collection("todos").Doc(id).Delete(ctx)
	return err
}

func GetAll() ([]*Todo, error) {
	iter := client.Collection("todos").Documents(ctx)
	todos := make([]*Todo, 0)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		t := &Todo{}
		doc.DataTo(t)
		t.ID = doc.Ref.ID
		todos = append(todos, t)
	}
	return todos, nil
}
