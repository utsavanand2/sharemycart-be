package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	api "github.com/utsavanand2/sharemycart/gohandlers/go"
)

var (
	// Version of the build
	Version = "dev"
	// GitCommit of the build
	GitCommit = ""

	port               = os.Getenv("PORT")
	credentialFilePath = os.Getenv("SVC")
)

func main() {
	ctx := context.Background()
	serviceAccount := option.WithCredentialsFile(credentialFilePath)
	app, err := firebase.NewApp(ctx, nil, serviceAccount)
	if err != nil {
		logrus.Fatalf("error initializing app: %v\n", err)
	}

	firebaseApp := api.NewFirebaseApp(*app)
	router := firebaseApp.NewRouter()

	server := NewServer(router, port)
	err = server.ListenAndServe()
	if err != nil {
		logrus.Fatalf("error starting server on port %s: %v", port, err)
	}
}

// NewServer returns a new *http.Server
func NewServer(router *mux.Router, port string) *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      router,
	}
	return srv
}
