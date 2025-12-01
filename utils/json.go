package utils

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bayuf/Golang_Authentication-System/model"
)

// Decode File Json menjadi slice
func DecoderTask(filePath string, v any) error {

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New("gagal membaca file json")
	}

	return json.Unmarshal(jsonData, v)

}

// Streaming Json File dan menulis
func EncoderTask(filePath string, user []model.User) error {

	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("cant create file json")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(user)

	return nil

}
