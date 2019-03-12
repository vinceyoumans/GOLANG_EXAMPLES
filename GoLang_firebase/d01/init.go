import (
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

opt := option.WithCredentialsFile("./scai-qit-firebase-adminsdk-y8bcu-a6fec9f53d")
app, err := firebase.NewApp(context.Background(), nil, opt)
if err != nil {
	log.Fatalf("error initializing app: %v\n", err)
}