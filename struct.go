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
	Fname     string   `json:"fname"`
	Sname     string   `json:"sname"`
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
}

type Filters struct {
	Limit  string `from:"limit"`
	Sex_eq string `form:"sex_eq"`

	Status_eq  string `form:"status_eq"`
	Status_neq string `form:"status_neq"`

	Fname_eq string `form:"fname_eq"`
	//Fname_any  string `form:"fname_any"`
	Fname_null string `form:"fname_null"`

	Email_domain string `form:"email_domain"`
	Email_lt     string `form:"email_lt"`
	Email_gt     string `form:"email_gt"`

	SName_eq     string `form:"sname_eq"`
	SName_starts string `form:"sname_starts"`
	SName_null   string `form:"sname_null"`

	Phone_code string `form:"phone_code"`
	Phone_null string `form:"phone_null"`

	Country_eq   string `form:"country_eq"`
	Country_null string `form:"country_null"`

	City_eq   string `form:"city_eq"`
	City_any  string `form:"city_any"`
	City_null string `form:"city_null"`

	//Interests_contains string `form:"interests_contains"`
	//Interests_any      string `form:"interests_any"`

	//birth string
	Birth_lt   int    `from:"birth_lt"`
	Birth_gt   int    `from:"birth_gt"`
	Birth_year string `from:"birth_year"`

	//likes string
	//premium string
}

type GroupKey struct {
	sex       string
	status    string
	interests string
	country   string
	city      string
}
