package main

type AccountsData struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	ID        int           `json:"id" binding:"required"`
	Fname     string        `json:"fname" binding:"required"`
	Email     string        `json:"email" binding:"required"`
	Interests []interface{} `json:"interests" binding:"required"`
	Status    string        `json:"status" binding:"required"`
	Premium   struct {
		Start  int `json:"start"`
		Finish int `json:"finish"`
	} `json:"premium" binding:"required"`
	Sex   string `json:"sex" binding:"required"`
	Phone string `json:"phone"`
	Likes []struct {
		Dt int `json:"dt"`
		ID int `json:"id"`
	} `json:"likes" binding:"required"`
	Birth   int    `json:"birth" binding:"required"`
	City    string `json:"city"`
	Country string `json:"country" binding:"required"`
	Joined  int    `json:"joined" binding:"required"`
}
