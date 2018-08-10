package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

type foo struct {
	Name string
}

func main() {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	addSampleData(ctx, client)
}

func addSampleData(ctx context.Context, client *firestore.Client) {
	foo := foo{Name: "sample"}
	ref, res, err := client.Collection("sample").Add(ctx, foo)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Ref: %#v\n", ref)
	log.Printf("Res: %#v\n", res)
}
