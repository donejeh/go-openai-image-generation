package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get API key from environment variable
	apiKey := os.Getenv("OPENAI_API")
	if apiKey == "" {
		log.Fatalf("OPENAI_API environment variable is not set")
	}

	c := openai.NewClient(apiKey)

	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your prompt: >")

	if !s.Scan() {
		panic("failed to get user input")
	}

	req := openai.ImageRequest{
		Prompt:         s.Text(),
		Size:           openai.CreateImageSize512x512,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	}

	resp, err := c.CreateImage(context.Background(), req)

	if err != nil {
		panic(err)
	}

	b, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)

	if err != nil {
		panic(err)
	}

	f, err := os.Create("image.png")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		panic(err)
	}
}
