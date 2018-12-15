package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"strconv"
	"strings"
)

func FilterUser(c *gin.Context) {
	var filter Filters
	err := c.Bind(&filter)
	if err != nil {
		c.Status(400)
		return
	}

	i := 0
	limit, err := strconv.Atoi(c.DefaultQuery("limit", ""))

	if err != nil {
		c.JSON(200, gin.H{
			"accounts": Acc,
		})
		return
	}

	result := funk.Filter(Acc, func(ac Account) bool {
		f := false
		if i >= limit {
			return false
		} else {
			i++
		}

		//SEx
		if filter.Sex_eq == ac.Sex {
			f = true
		} else {
			if filter.Sex_eq != "" {
				f = false
			}
		}

		//Status
		if filter.Status_eq == ac.Status {
			f = true
		} else {
			if filter.Status_eq != "" {
				f = false
			}
		}
		if filter.Status_neq != "" && filter.Status_neq != ac.Status {
			f = true
		} else {
			if filter.Status_neq != "" {
				f = false
			}
		}

		// Fname
		if filter.Fname_eq == ac.Fname {
			f = true
		} else {
			if filter.Fname_eq != "" {
				f = false
			}
		}
		if SliceExists(filter.Fname_any, ac.Fname) {
			f = true
		} else {
			if len(filter.Fname_any) > 0 {
				f = false
			}
		}
		if filter.Fname_null == "1" && ac.Fname == "" {
			f = true
		} else {
			if filter.Fname_null != "" {
				f = false
			}
		}

		//Email
		if filter.Email_domain == strings.Split(ac.Email, "@")[1] {
			f = true
		} else {
			if filter.Email_domain != "" {
				f = false
			}
		}
		//todo

		//SName
		if filter.SName_eq == ac.Sname {
			f = true
		} else {
			if filter.SName_eq != "" {
				f = false
			}
		}
		if filter.SName_starts != "" && strings.Contains(ac.Sname, filter.SName_starts) {
			f = true
		} else {
			if filter.SName_starts != "" {
				f = false
			}
		}
		if filter.SName_null == "1" && ac.Sname == "" {
			f = true
		} else {
			if filter.SName_null != "" {
				f = false
			}
		}

		//Phone
		if strings.Contains(ac.Phone, "("+filter.Phone_code+")") {
			f = true
		} else {
			if filter.Phone_code != "" {
				f = false
			}
		}
		if filter.Phone_null == "1" && ac.Phone == "" {
			f = true
		} else {
			if filter.Phone_null != "" {
				f = false
			}
		}

		//Country
		if filter.Country_eq == ac.Country {
			f = true
		} else {
			if filter.Country_eq != "" {
				f = false
			}
		}
		if filter.Country_null == "1" && ac.Country == "" {
			f = true
		} else {
			if filter.Country_null != "" {
				f = false
			}
		}

		//City
		if filter.City_eq == ac.City {
			f = true
		} else {
			if filter.City_eq != "" {
				f = false
			}
		}
		if SliceExists(filter.Fname_any, ac.Fname) {
			f = true
		} else {
			if len(filter.Fname_any) > 0 {
				f = false
			}
		}
		if filter.City_null == "1" && ac.City == "" {
			f = true
		} else {
			if filter.City_null != "" {
				f = false
			}
		}

		//Interests
		if len(filter.Interests_any) > 0 && funk.Contains(filter.Interests_any, ac.Interests) {
			f = true
		} else {
			if len(filter.Interests_any) > 0 {
				f = false
			}
		}

		if len(filter.Interests_contains) > 0 && SliceExists(filter.Interests_contains, ac.Interests) {
			f = true
		} else {
			if len(filter.Interests_contains) > 0 {
				f = false
			}
		}

		return f
	})

	c.JSON(200, gin.H{
		"accounts": result,
	})
}
