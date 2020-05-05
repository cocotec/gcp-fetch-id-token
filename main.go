package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	idtoken "google.golang.org/api/idtoken"
)

func main() {
	audience := flag.String("audience", "", "The audience to generate the token for")
	credentials := flag.String("credentials", "", "A path to a JSON credentials file")

	flag.Parse()

	if *audience == "" {
		fmt.Print("--audience is required")
		os.Exit(255)
	}

	var options []idtoken.ClientOption
	if *credentials != "" {
		options = append(options, idtoken.WithCredentialsFile(*credentials))
	}

	tokenSource, err := idtoken.NewTokenSource(context.Background(), *audience, options...)
	if err != nil {
		fmt.Printf("Could not create token source: %s", err)
		os.Exit(1)
	}
	var token *oauth2.Token
	token, err = tokenSource.Token()
	if err != nil {
		fmt.Printf("Could not create token: %s", err)
		os.Exit(1)
	}
	fmt.Print(token.AccessToken)
	os.Exit(0)
}
