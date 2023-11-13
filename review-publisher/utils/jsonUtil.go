package utils

import (
	"encoding/json"
	"log"
	"review-consumer/models"
)

func ObjectTojson(review *models.ReviewReq) string {

	var jsonData []byte
	jsonData, err := json.Marshal(review)
	if err != nil {
		log.Println(err)
	}

	return string(jsonData)
}

func JsonToObject(jsonData []byte) models.ReviewReq {

	var unreview models.ReviewReq
	_ = json.Unmarshal(jsonData, &unreview)

	return unreview
}
