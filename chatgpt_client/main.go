package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	//request := gpt3.CompletionRequest{
	//	Prompt: []string{"How many coffees should I drink per day?"},
	//}

	for true {
		fmt.Print("\n\n> ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		complete(ctx, client, line)
	}
	//resp, err := client.Completion(ctx, request)
	//resp, err := client.CompletionWithEngine(ctx, "gpt-3.5-turbo-instruct", request)
	//if err != nil {
	//	fmt.Printf("%s\n", err)
	//} else {
	//	fmt.Printf("Answer:\n %s\n", resp.Choices[0].Text)
	//}
}

func makeRequest(question string) gpt3.CompletionRequest {

	maxToken, _ := strconv.Atoi(os.Getenv("MAX_TOKEN"))
	temperature, _ := strconv.ParseFloat(os.Getenv("TEMPERATURE"), 32)

	questions := []string{question}
	return gpt3.CompletionRequest{
		Prompt:      questions,
		MaxTokens:   gpt3.IntPtr(maxToken),
		Temperature: gpt3.Float32Ptr(float32(temperature)),
	}
}

func complete(ctx context.Context, client gpt3.Client, question string) {

	model := os.Getenv("MODEL")
	request := makeRequest(question)
	//resp, _ := client.CompletionWithEngine(ctx, model, request)
	_ = client.CompletionStreamWithEngine(ctx, model, request, func(resp *gpt3.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)
	})
	//fmt.Print(resp.Choices[0].Text)
}

// Querying the Available Models to Use with ChatGPT
//func main() {
//	godotenv.Load()
//	apiKey := os.Getenv("API_KEY")
//	if apiKey == "" {
//		log.Fatalln("Missing API KEY")
//	}
//	ctx := context.Background()
//	client := gpt3.NewClient(apiKey)
//	engines, err := client.Engines(ctx)
//	if err != nil {
//		return
//	}
//	for _, engine := range engines.Data {
//		fmt.Printf("Engine ID: %s, Name: %s, Ready: %t\n", engine.ID, engine.
//			Owner, engine.Ready)
//	}
//}
