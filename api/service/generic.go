package service

import (
	"encoding/json"
	"fmt"

	"github.com/arnavmahajan630/cfcli/api/models"
)
func DecodeCFResponse[T any](body []byte) ([]T, error) {
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