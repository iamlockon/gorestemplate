package cache_service

import (
	"strconv"
	"strings"

	"github.com/iamlockon/gorestemplate/pkg/error"
)

type Example struct {
	ID    int
	Name  string
	State int

	PageNum  int
	PageSize int
}

func (ex *Example) GetExamplesKey() string {
	keys := []string{
		error.CACHE_EXAMPLE,
		"LIST",
	}

	if ex.Name != "" {
		keys = append(keys, ex.Name)
	}

	if ex.State >= 0 {
		keys = append(keys, strconv.Itoa(ex.State))
	}

	if ex.PageNum > 0 {
		keys = append(keys, strconv.Itoa(ex.PageNum))
	}

	if ex.PageSize > 0 {
		keys = append(keys, strconv.Itoa(ex.PageSize))
	}

	return strings.Join(keys, "_")
}
