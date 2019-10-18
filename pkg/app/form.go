package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/iamlockon/gorestemplate/pkg/error"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, error.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, error.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, error.INVALID_PARAMS
	}

	return http.StatusOK, error.SUCCESS
}
