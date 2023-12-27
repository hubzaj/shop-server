package utils

import (
	"encoding/json"
	"fmt"
	"io"
)

func UnmarshalArrayResponseBody[T interface{}](body io.ReadCloser, t []*T) []*T {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		panic(fmt.Errorf("error during reading response body: %s", err))
	}
	if err := json.Unmarshal(bodyBytes, &t); err != nil {
		panic(fmt.Errorf("error during unmarshalling response body: %s", err))
	}
	return t
}

func UnmarshalResponseBody[T interface{}](body io.ReadCloser, t *T) *T {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		panic(fmt.Errorf("error during reading response body: %s", err))
	}
	if err := json.Unmarshal(bodyBytes, &t); err != nil {
		panic(fmt.Errorf("error during unmarshalling response body: %s", err))
	}
	return t
}
