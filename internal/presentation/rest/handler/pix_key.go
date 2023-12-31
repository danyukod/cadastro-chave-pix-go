package handler

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/application/commands"
	req "github.com/danyukod/cadastro-chave-pix-go/internal/presentation/rest/handler/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/rest/handler/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PixKeyHandlerInterface interface {
	FindPixKeyByKey(c *gin.Context)
	FindPixKey(c *gin.Context)
	RegisterPixKey(c *gin.Context)
}

func NewPixKeyHandlerInterface(
	registerPixKeyUsecase commands.RegisterPixKeyUsecase,
	findPixKeyByKeyUsecase commands.FindPixKeyByKeyUsecase,
	findPixKeyUsecase commands.FindPixKeyUsecase,
) PixKeyHandlerInterface {
	return &handler{
		registerPixKeyUsecase:  registerPixKeyUsecase,
		findPixKeyByKeyUsecase: findPixKeyByKeyUsecase,
		findPixKeyUsecase:      findPixKeyUsecase,
	}
}

type handler struct {
	registerPixKeyUsecase  commands.RegisterPixKeyUsecase
	findPixKeyByKeyUsecase commands.FindPixKeyByKeyUsecase
	findPixKeyUsecase      commands.FindPixKeyUsecase
}

// RegisterPixKey godoc
// @Summary Register Pix Key
// @Description Register Pix Key
// @Tags pix-keys
// @Accept json
// @Produce json
// @Param request body request.RegisterPixKeyRequest true "Pix Key Request"
// @Success 201 {object} response.RegisterPixKeyResponse
// @Failure 500 {object} ErrorsResponse
// @Failure 400 {object} ErrorsResponse
// @Router /pix-keys [post]
// @Security ApiKeyAuth
func (p *handler) RegisterPixKey(c *gin.Context) {
	var request req.RegisterPixKeyRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	pixKeyDomain, err := p.registerPixKeyUsecase.Execute(request.ToDTO())
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.PixKeyDomainToRegisterWebResponse(pixKeyDomain))
}

// FindPixKeyByKey godoc
// @Summary Find PixKey By Key
// @Description Find PixKey By Key
// @Tags pix-keys
// @Accept json
// @Produce json
// @Param key path string true "Pix Key"
// @Success 200 {object} response.FindPixKeyResponse
// @Failure 500 {object} ErrorsResponse
// @Failure 400 {object} ErrorsResponse
// @Router /pix-keys/{key} [get]
// @Security ApiKeyAuth
func (p *handler) FindPixKeyByKey(c *gin.Context) {
	var request req.FindPixKeyRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	pixKeyDomain, err := p.findPixKeyByKeyUsecase.Execute(request.ToDTO())
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response.PixKeyDomainToFindWebResponse(pixKeyDomain))
}

// FindPixKey godoc
// @Summary Find PixKey
// @Description Find PixKey
// @Tags pix-keys
// @Accept json
// @Produce json
// @Success 200 {object} []response.FindPixKeyResponse
// @Failure 500 {object} ErrorsResponse
// @Failure 400 {object} ErrorsResponse
// @Router /pix-keys [get]
// @Security ApiKeyAuth
func (p *handler) FindPixKey(c *gin.Context) {
	pixKeyDomain, err := p.findPixKeyUsecase.Execute()
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	var responseList []response.FindPixKeyResponse

	for _, pixKey := range pixKeyDomain {
		resp := response.PixKeyDomainToFindWebResponse(pixKey)
		responseList = append(responseList, *resp)
	}

	c.JSON(http.StatusOK, responseList)
}
