package question

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"os"
	"strings"

	"github.com/qdrant/go-client/qdrant"
)


type OpenRouterEmbeddingRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

type OpenRouterEmbeddingResponse struct {
	Data []struct {
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
}


func Question() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan pertanyaan: ")
	question, _ := reader.ReadString('\n')
	question = strings.TrimSpace(question)
	fmt.Println("Pertanyaan:", question)


	queryVector, err := embedQuestion(ctx, question)
	if err != nil {
		panic(err)
	}

	limit := uint64(5)
	resp, err := client.Query(ctx, &qdrant.QueryPoints{
		CollectionName: "bps_knowledge",
		Query:          qdrant.NewQuery(queryVector...),
		Limit:          &limit,
		WithPayload:    qdrant.NewWithPayload(true),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("\nHasil dari Qdrant:")
	for _, point := range resp {
		payload := point.Payload
		fmt.Println("--------------------")
		fmt.Println("Text   :", payload["text"])
		fmt.Println("Source :", payload["source"])
	}
}


func embedQuestion(ctx context.Context, question string) ([]float32, error) {
	apiKey := os.Getenv("OPENROUTER_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENROUTER_KEY")
	}

	reqBody := OpenRouterEmbeddingRequest{
		Model: "openai/text-embedding-3-small",
		Input: []string{question},
	}

	bodyBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequestWithContext(ctx, "POST", "https://openrouter.ai/v1/embeddings", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, _ :=io.ReadAll(resp.Body)
	var orResp OpenRouterEmbeddingResponse
	err = json.Unmarshal(respBytes, &orResp)
	if err != nil {
		return nil, err
	}

	if len(orResp.Data) == 0 {
		return nil, fmt.Errorf("embedding kosong dari OpenRouter")
	}

	vec := make([]float32, len(orResp.Data[0].Embedding))
	for i, v := range orResp.Data[0].Embedding {
		vec[i] = float32(v)
	}

	return vec, nil
}
