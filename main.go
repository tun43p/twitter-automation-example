package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("TWITTER_ACCESS_TOKEN must be set")
	}

	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	if accessTokenSecret == "" {
		log.Fatal("TWITTER_ACCESS_TOKEN_SECRET must be set")
	}

	apiKey := os.Getenv("TWITTER_API_KEY")
	if apiKey == "" {
		log.Fatal("TWITTER_API_KEY must be set")
	}

	apiKeySecret := os.Getenv("TWITTER_API_KEY_SECRET")
	if apiKeySecret == "" {
		log.Fatal("TWITTER_API_KEY_SECRET must be set")
	}

	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           accessToken,
		OAuthTokenSecret:     accessTokenSecret,
		APIKey:               apiKey,
		APIKeySecret:         apiKeySecret,
	}

	c, err := gotwi.NewClient(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.CreateInput{
		Text: gotwi.String("This is a test tweet with poll."),
		Poll: &types.CreateInputPoll{
			DurationMinutes: gotwi.Int(5),
			Options: []string{
				"Cyan",
				"Magenta",
				"Yellow",
				"Key plate",
			},
		},
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
}
