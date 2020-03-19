package main

import (
	"context"
	"net/http"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	"github.com/sirupsen/logrus"
	api "github.com/utsavanand2/sharemycart/gohandlers/go"
)

var (
	// Version of the build
	Version = "dev"
	// GitCommit of the build
	GitCommit = ""
)

func main() {

	ctx := context.Background()
	serviceAccount := option.WithCredentialsFile("")
	_, err := firebase.NewApp(ctx, nil, serviceAccount)
	if err != nil {
		logrus.Fatalf("error initializing app: %v\n", err)
	}

	router := api.NewRouter()
	http.ListenAndServe(":8080", router)
}
