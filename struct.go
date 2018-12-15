package main

type AccountsData struct {
	Accounts []Account `json:"accounts"`
}

type Likes struct {
	Likes []struct {
		Likee int `json:"likee"  binding:"required"`
		Ts    int `json:"ts" binding:"required"`
		Liker int `json:"liker" binding:"required"`
	} `json:"likes" binding:"required"`
}

type Account struct {
	ID        int      `json:"id"`
	Fname     string   `json:"fname" binding:"required"`
	Sname     string   `json:"sname"`
	Email     string   `json:"email" binding:"required"`
	Interests []string `json:"interests"`
	Status    string   `json:"status"`
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
	Country string `json:"country"`
	Joined  int    `json:"joined" binding:"required"`
}

type Filters struct {
	Limit  string `from:"limit" json:"limit"`
	Sex_eq string `form:"sex_eq" json:"sex_eq"`

	Status_eq  string `form:"status_eq" json:"status_eq"`
	Status_neq string `form:"status_neq" json:"status_neq"`

	Fname_eq   string   `form:"fname_eq" json:"fname_eq"`
	Fname_any  []string `form:"fname_any" json:"fname_any"`
	Fname_null string   `form:"fname_null" json:"fname_null"`

	Email_domain string `form:"email_domain" json:"email_domain"`
	Email_lt     string `form:"email_lt" json:"email_lt"`
	Email_gt     string `form:"email_gt" json:"email_gt"`

	SName_eq     string `form:"sname_eq" json:"sname_eq"`
	SName_starts string `form:"sname_starts" json:"sname_starts"`
	SName_null   string `form:"sname_null" json:"sname_null"`

	Phone_code string `form:"phone_code" json:"phone_code"`
	Phone_null string `form:"phone_null" json:"phone_null"`

	Country_eq   string `form:"country_eq" json:"country_eq"`
	Country_null string `form:"country_null" json:"country_null"`

	City_eq   string   `form:"city_eq" json:"city"`
	City_any  []string `form:"city_any" json:"city_any"`
	City_null string   `form:"city_null" json:"city_null"`

	Interests_contains []string `form:"interests_contains" json:"interests_contains"`
	Interests_any      []string `form:"interests_any" json:"interests_any"`

	//birth string
	//likes string
	//premium string
}
