package main

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

func fetchUserEvent(username string) error {
	url := fmt.Sprintf("https://api.github.com/users/%s/events/public", username)

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("unable to get URL: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed reading body: %w", err)
	}

	var events []map[string]interface{}
	if err := json.Unmarshal(body, &events); err != nil{
		return fmt.Errorf("failed unmarshalling: %w", err)
	}

	for i, event := range events {
		fmt.Printf("%d. Type: %s\n", i+1, event["type"])
	}

	return nil
}
