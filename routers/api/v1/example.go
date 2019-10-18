package v1

import (
	// "template/pkg/logging"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/iamlockon/gorestemplate/pkg/error"
	"github.com/iamlockon/gorestemplate/pkg/setting"
	"github.com/iamlockon/gorestemplate/pkg/util"
	"github.com/iamlockon/gorestemplate/pkg/app"
	"github.com/iamlockon/gorestemplate/service/example_service"
)

// @Summary Get multiple examples
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/examples [get]
func GetExamples(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	exampleService := example_service.Example{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := exampleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error.ERROR_GET_EXAMPLE_FAIL, nil)
		return
	}

	count, err := exampleService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error.ERROR_COUNT_EXAMPLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error.SUCCESS, map[string]interface{}{
		"lists": tags,
		"total": count,
	})
}

type AddExampleForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Add example
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/examples [post]
func AddExample(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddExampleForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != error.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	exampleService := example_service.Example{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
		State:     form.State,
	}
	exists, err := exampleService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error.ERROR_EXIST_EXAMPLE_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, error.ERROR_EXAMPLE_EXIST, nil)
		return
	}

	err = exampleService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error.ERROR_ADD_EXAMPLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error.SUCCESS, nil)
}

type EditExampleForm struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Update example
// @Produce  json
// @Param id path int true "ID"
// @Param name body string true "ID"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/{id} [put]
func EditExample(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditExampleForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != error.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	exampleService := example_service.Example{
		ID:         form.ID,
		Name:       form.Name,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}

	exists, err := exampleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error.ERROR_EXIST_EXAMPLE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, error.ERROR_EXAMPLE_NOT_EXIST, nil)
		return
	}

	err = exampleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error.ERROR_EDIT_EXAMPLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error.SUCCESS, nil)
}

// @Summary Delete example
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/example/{id} [delete]
func DeleteExample(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID must be greater than 0.")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, error.INVALID_PARAMS, nil)
	}

	exampleService := example_service.Example{ID: id}
	exists, err := exampleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error.ERROR_EXIST_EXAMPLE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, error.ERROR_EXAMPLE_NOT_EXIST, nil)
		return
	}

	if err := exampleService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, error.ERROR_DELETE_EXAMPLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error.SUCCESS, nil)
}