package filters

import (
	"github.com/deissh/HighLoad18/core/account"
	"github.com/pkg/errors"
)

type sexFilter struct {
	operation string
	value     string
}

func makeSexFilter(operation, value string) (Filter, error) {
	if value != "m" && value != "f" {
		return nil, errors.New("bad filter value")
	}

	return sexFilter{
		operation: operation,
		value:     value,
	}, nil
}

func (f sexFilter) Test(account account.Account) error {
	if account.Sex == f.value {
		return nil
	}
	return ErrTestFailed
}

func (f sexFilter) Fill(src account.Account, dst map[string]interface{}) {
	dst["sex"] = src.Sex
}
