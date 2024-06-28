package helper

import (
	"encoding/json"
	"fmt"
)

func StructToJsonToString(v any) (string, error) {
	parsed, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Failed marshalling: ", err)
		return "", err
	}
	return string(parsed), err
}
