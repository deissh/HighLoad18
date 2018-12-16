package main

import (
	"github.com/gin-gonic/gin"
)

func FilterUser(c *gin.Context) {
	//var filter Filters
	//err := c.Bind(&filter)
	//if err != nil {
	//	c.Status(400)
	//	return
	//}
	//
	//i := 0
	//limit, err := strconv.Atoi(c.DefaultQuery("limit", ""))

	//if err != nil {
	//	c.JSON(200, gin.H{
	//		"accounts": Acc,
	//	})
	//	return
	//}

	var result []Account
	//for _, ac := range Acc {
	//	status := func(ac Account) bool {
	//		f := false
	//		//SEx
	//		if filter.Sex_eq != "" && filter.Sex_eq == ac.Sex {
	//			f = true
	//		} else {
	//			if filter.Sex_eq != "" {
	//				return false
	//			}
	//		}
	//
	//		//Status
	//		if filter.Status_eq != "" && filter.Status_eq == ac.Status {
	//			f = true
	//		} else {
	//			if filter.Status_eq != "" {
	//				return false
	//			}
	//		}
	//		if filter.Status_neq != "" && filter.Status_neq != ac.Status {
	//			f = true
	//		} else {
	//			if filter.Status_neq != "" {
	//				return false
	//			}
	//		}
	//
	//		// Fname
	//		if filter.Fname_eq != "" && filter.Fname_eq == ac.Fname {
	//			f = true
	//		} else {
	//			if filter.Fname_eq != "" {
	//				return false
	//			}
	//		}
	//		//if len(filter.Fname_any) > 0 && funk.Contains(filter.Fname_any, ac.Fname) {
	//		//	f = true
	//		//} else {
	//		//	if len(filter.Fname_any) > 0 {
	//		//		return false
	//		//	}
	//		//}
	//		if filter.Fname_null == "1" && ac.Fname == "" {
	//			f = true
	//		} else {
	//			if filter.Fname_null != "" {
	//				return false
	//			}
	//		}
	//
	//		//Email
	//		if filter.Email_domain != "" && filter.Email_domain == strings.Split(ac.Email, "@")[1] {
	//			f = true
	//		} else {
	//			if filter.Email_domain != "" {
	//				return false
	//			}
	//		}
	//		//todo
	//
	//		//SName
	//		if filter.SName_eq != "" && filter.SName_eq == ac.Sname {
	//			f = true
	//		} else {
	//			if filter.SName_eq != "" {
	//				return false
	//			}
	//		}
	//		if filter.SName_starts != "" && strings.Contains(ac.Sname, filter.SName_starts) {
	//			f = true
	//		} else {
	//			if filter.SName_starts != "" {
	//				return false
	//			}
	//		}
	//		if filter.SName_null == "1" && ac.Sname == "" {
	//			f = true
	//		} else {
	//			if filter.SName_null != "" {
	//				return false
	//			}
	//		}
	//
	//		//Phone
	//		if filter.Phone_code != "" && strings.Contains(ac.Phone, "("+filter.Phone_code+")") {
	//			f = true
	//		} else {
	//			if filter.Phone_code != "" {
	//				return false
	//			}
	//		}
	//		if filter.Phone_null == "1" && ac.Phone == "" {
	//			f = true
	//		} else {
	//			if filter.Phone_null != "" {
	//				return false
	//			}
	//		}
	//
	//		//Country
	//		if filter.Country_eq != "" && filter.Country_eq == ac.Country {
	//			f = true
	//		} else {
	//			if filter.Country_eq != "" {
	//				return false
	//			}
	//		}
	//		if filter.Country_null == "1" && ac.Country == "" {
	//			f = true
	//		} else {
	//			if filter.Country_null != "" {
	//				return false
	//			}
	//		}
	//		if filter.Country_null == "0" && ac.Country != "" {
	//			f = true
	//		} else {
	//			if filter.Country_null != "" {
	//				return false
	//			}
	//		}
	//
	//		//City
	//		if filter.City_eq != "" && filter.City_eq == ac.City {
	//			f = true
	//		} else {
	//			if filter.City_eq != "" {
	//				return false
	//			}
	//		}
	//		//if len(filter.Fname_any) > 0 && funk.Contains(filter.Fname_any, ac.Fname) {
	//		//	f = true
	//		//} else {
	//		//	if len(filter.Fname_any) > 0 {
	//		//		return false
	//		//	}
	//		//}
	//		if filter.City_null == "1" && ac.City == "" {
	//			f = true
	//		} else {
	//			if filter.City_null != "" {
	//				return false
	//			}
	//		}
	//
	//		//Interests
	//		//if len(filter.Interests_any) > 0 && SliceExists(filter.Interests_any, ac.Interests) {
	//		//	f = true
	//		//} else {
	//		//	if len(filter.Interests_any) > 0 {
	//		//		return false
	//		//	}
	//		//}
	//		//
	//		//if len(filter.Interests_contains) > 0 && SliceExists(filter.Interests_contains, ac.Interests) {
	//		//	f = true
	//		//} else {
	//		//	if len(filter.Interests_contains) > 0 {
	//		//		return false
	//		//	}
	//		//}
	//
	//		// Birth
	//		if filter.Birth_gt != 0 && filter.Birth_gt >= ac.Birth {
	//			f = true
	//		} else {
	//			if filter.Birth_gt != 0 {
	//				return false
	//			}
	//		}
	//		if filter.Birth_lt != 0 && filter.Birth_lt <= ac.Birth {
	//			f = true
	//		} else {
	//			if filter.Birth_lt != 0 {
	//				return false
	//			}
	//		}
	//		if filter.Birth_year != "" && filter.Birth_year == strconv.Itoa(time.Unix(int64(ac.Birth), 0).Year()) {
	//			f = true
	//		} else {
	//			if filter.Birth_year != "" {
	//				return false
	//			}
	//		}
	//
	//		// если нету условий
	//		if !f {
	//			return true
	//		}
	//
	//		return f
	//	}(ac)
	//
	//	if status {
	//		if i >= limit {
	//			break
	//		} else {
	//			if status {
	//				i++
	//			}
	//		}
	//
	//		result = append(result, ac)
	//	}
	//}

	//var result
	//result = funk.Filter(Acc, )

	if result == nil {
		c.JSON(200, gin.H{
			"accounts": []string{},
		})
		return
	}

	c.Writer.Header().Set("Transfer-Encoding", "identity")
	c.JSON(200, gin.H{
		"accounts": result,
	})
}
