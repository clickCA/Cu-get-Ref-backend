package utils

import (
	"encoding/json"
	"log"
	"review-consumer/models"
)

func ObjectTojson(review *models.Review) string {

	var jsonData []byte
	jsonData, err := json.Marshal(review)
	if err != nil {
		log.Println(err)
	}

	return string(jsonData)
}

func JsonToObject(jsonData []byte) models.Review {

	var unreview models.Review
	_ = json.Unmarshal(jsonData, &unreview)

	return unreview
}
