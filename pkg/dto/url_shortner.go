package dto

import (
	"errors"
	"github.com/agrison/go-commons-lang/stringUtils"
	"regexp"
)

type GetUrl struct {
	LongUrl string
}

func (r *GetUrl) ValidateRequest() error {
	if stringUtils.IsBlank(r.LongUrl) {
		return errors.New("Empty url. ")
	} else if matched, _ := regexp.MatchString("[^A-Za-z0-9_.]", r.LongUrl); !matched {
		return errors.New(" Should Contain Characters in A-Z r a-z, 0-9, _, . ")
	}
	return nil
}
