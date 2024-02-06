package endpoints

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	feather_commons_validation "github.com/guidomantilla/go-feather-commons/pkg/validation"
	feather_web_rest "github.com/guidomantilla/go-feather-web/pkg/rest"

	"multipliers/pkg/models"
	"multipliers/pkg/services"
)

type DefaultNumbersEndpoint struct {
	numbersService services.NumbersService
}

func NewDefaultNumbersEndpoint(numbersService services.NumbersService) *DefaultNumbersEndpoint {
	return &DefaultNumbersEndpoint{
		numbersService: numbersService,
	}
}

func (endpoint *DefaultNumbersEndpoint) Save(ctx *gin.Context) {

	var err error
	var numberToSave *models.Number
	if err = ctx.ShouldBindJSON(&numberToSave); err != nil {
		ex := feather_web_rest.BadRequestException("error unmarshalling request json to object")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if err = feather_commons_validation.ValidateFieldIsRequired("this", "number", numberToSave.Number); err != nil {
		ex := feather_web_rest.BadRequestException("error validating the object", err)
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if err = endpoint.numbersService.Save(ctx.Request.Context(), numberToSave); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	ctx.JSON(http.StatusOK, numberToSave)
}

func (endpoint *DefaultNumbersEndpoint) Get(ctx *gin.Context) {

	var err error

	var body []byte
	if body, err = io.ReadAll(ctx.Request.Body); err != nil {
		ex := feather_web_rest.BadRequestException("error reading body")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if len(body) != 0 {
		ex := feather_web_rest.BadRequestException("body is not empty")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	number := ctx.Params.ByName("number")
	if number == "" {
		ex := feather_web_rest.BadRequestException("object id not defined in url path")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	numberToFind := &models.Number{
		Number: &number,
	}

	if err = feather_commons_validation.ValidateFieldIsRequired("this", "number", numberToFind.Number); err != nil {
		ex := feather_web_rest.BadRequestException("error validating the object", err)
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	var numberFound *models.Number
	if numberFound, err = endpoint.numbersService.Find(ctx.Request.Context(), numberToFind); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	ctx.JSON(http.StatusOK, numberFound)
}

func (endpoint *DefaultNumbersEndpoint) GetAll(ctx *gin.Context) {

	var err error

	var body []byte
	if body, err = io.ReadAll(ctx.Request.Body); err != nil {
		ex := feather_web_rest.BadRequestException("error reading body")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if len(body) != 0 {
		ex := feather_web_rest.BadRequestException("body is not empty")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	var numbers []string
	if numbers, err = endpoint.numbersService.FindAll(ctx.Request.Context()); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	ctx.JSON(http.StatusOK, numbers)
}
