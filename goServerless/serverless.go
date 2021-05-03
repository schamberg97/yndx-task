package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Входной JSON-документ будет автоматически преобразован в объект данного типа
type Request struct {
	Message string `json:"message"`
	Number  int    `json:"number"`
}

type JSONString string

func (j JSONString) MarshalJSON() ([]byte, error) {
	return []byte(j), nil
}

func Handler(ctx context.Context, request *Request) ([]byte, error) {

	unixTime := time.Now().Unix()
	unixTimeStr := strconv.FormatInt(unixTime, 10)
	s := `{"body":` + unixTimeStr + `}`
	fmt.Println(s)
	res, err := json.Marshal(JSONString(s))

	if err != nil {
		return nil, err
	}
	return res, nil
}
