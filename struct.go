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

type Filters struct {
	Limit  string `from:"limit" json:"limit"`
	Sex_eq string `form:"sex_eq" json:"sex_eq"`

	Status_eq  string `form:"status_eq" json:"status_eq"`
	Status_neq string `form:"status_neq" json:"status_neq"`

	Fname_eq   string `form:"fname_eq" json:"fname_eq"`
	Fname_any  string `form:"fname_any" json:"fname_any"`
	Fname_null string `form:"fname_null" json:"fname_null"`

	Email_domain string `form:"email_domain" json:"email_domain"`
	Email_lt     string `form:"email_lt" json:"email_lt"`
	Email_gt     string `form:"email_gt" json:"email_gt"`

	SName_eq     string `form:"sname_eq" json:"email_domain"`
	SName_starts string `form:"sname_starts" json:"email_domain"`
	SName_null   string `form:"sname_null" json:"email_domain"`

	Phone_code string `form:"sname_eq" json:"email_domain"`
	Phone_null string `form:"sname_eq" json:"email_domain"`

	Country_eq   string `form:"country_eq" json:"country_eq"`
	Country_null string `form:"country_null" json:"country_null"`

	City_eq   string `form:"city_eq" json:"city"`
	City_any  string `form:"city_any" json:"city_any"`
	City_null string `form:"city_null" json:"city_null"`

	//city string
	//birth string
	//interests string
	//likes string
	//premium string
}
