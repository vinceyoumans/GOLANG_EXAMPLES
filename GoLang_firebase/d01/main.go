package main

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	ctx := context.Background()
	//projectID := "get"
	sa := option.WithCredentialsFile("./scai-qit-fb-adminsdk.json")
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	iter := client.Collection("COMPLEX_NONACS").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println("--------------------------")
		dd := doc.Ref
		// fmt.Println(doc.Ref)
		fmt.Println(doc.Data())

	}

}
