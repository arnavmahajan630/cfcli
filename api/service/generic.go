package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/arnavmahajan630/cfcli/api/models"
)

func decodeCFResponse[T any](body []byte) ([]T, error) {
	var base models.BaseResponse
	if err := json.Unmarshal(body, &base); err != nil {
		return nil, fmt.Errorf("failed to decode base response: %v", err)
	}
	if base.Status != "OK" {
		return nil, fmt.Errorf("API error: %s", base.Comment)
	}

	var result []T
	if err := json.Unmarshal(base.Result, &result); err != nil {
		return nil, fmt.Errorf("failed to decode result: %v", err)
	}
	return result, nil
}

func FetchAndDecode[T any](url string) ([]T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("❌ request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("❌ read failed: %w", err)
	}

	data, err := decodeCFResponse[T](body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
