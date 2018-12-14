package main

import (
	"encoding/json"
	"log"
)

type AccountsData struct {
	Accounts []struct {
		ID        int      `json:"id"`
		Fname     string   `json:"fname"`
		Email     string   `json:"email"`
		Interests []string `json:"interests"`
		Status    string   `json:"status"`
		Premium   struct {
			Start  int `json:"start"`
			Finish int `json:"finish"`
		} `json:"premium"`
		Sex   string `json:"sex"`
		Phone string `json:"phone"`
		Likes []struct {
			Dt int `json:"dt"`
			ID int `json:"id"`
		} `json:"likes"`
		Birth   int    `json:"birth"`
		City    string `json:"city"`
		Country string `json:"country"`
		Joined  int    `json:"joined"`
	} `json:"accounts"`
}

func parseFile(f []byte) {
	log.Println("converting to json ...")
	result := AccountsData{}
	err := json.Unmarshal(f, &result)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("file converted to json")
}
