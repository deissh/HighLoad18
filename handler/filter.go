package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/julienschmidt/httprouter"

	filtering "github.com/deissh/HighLoad18/core/filters"
)

type response struct {
	Accounts []map[string]interface{} `json:"accounts"`
}

type filterparam struct {
}

//
//func FilterHandler(c *gin.Context) {
//	var limit int
//	var filters []filtering.Filter
//
//	limit, err := strconv.Atoi(c.Param("limit"))
//	if err != nil {
//		c.JSON(400, gin.H{})
//		return
//	}
//}

func New(f filtering.Service) http.Handler {
	router := httprouter.New()

	router.GET("/accounts/filter/", filter(f))

	return router
}

func filter(f filtering.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Add("Content-Type", "application/json")

		if err := r.ParseForm(); err != nil {
			panic("can't parse form")
		}

		var limit int
		var filters []filtering.Filter
		for key, value := range r.Form {
			if key == "query_id" {
				continue
			}

			if key == "limit" {
				l, err := strconv.Atoi(value[0])
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				limit = l
				continue
			}

			filter, err := filtering.Make(key, value[0])
			if errors.Cause(err) == filtering.ErrFilterNotFount {
				w.WriteHeader(http.StatusBadRequest)
				return
			} else if err != nil {
				panic(err.Error())
			}

			filters = append(filters, filter)
		}

		accounts, err := f.Fetch(filters, limit)
		if err != nil {
			panic(err.Error())
		}

		resp := response{
			Accounts: make([]map[string]interface{}, 0, len(accounts)),
		}

		for _, acc := range accounts {
			accountMap := map[string]interface{}{
				"id":    acc.ID,
				"email": acc.Email,
			}

			for _, f := range filters {
				f.Fill(acc, accountMap)
			}

			resp.Accounts = append(resp.Accounts, accountMap)
		}

		bb, err := json.Marshal(resp)
		if err != nil {
			panic(err.Error())
		}

		w.Header().Add("Content-Length", strconv.Itoa(len(bb)))

		if _, err := w.Write(bb); err != nil {
			panic(err.Error())
		}
	}
}
