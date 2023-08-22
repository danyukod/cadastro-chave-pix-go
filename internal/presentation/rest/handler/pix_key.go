package handler

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/application/commands"
	request2 "github.com/danyukod/cadastro-chave-pix-go/internal/presentation/rest/handler/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/rest/handler/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PixKeyHandlerInterface interface {
	FindPixKeyByKey(c *gin.Context)
	RegisterPixKey(c *gin.Context)
}

func NewPixKeyHandlerInterface(
	registerPixKeyUsecase commands.RegisterPixKeyUsecase,
	findPixKeyUsecase commands.FindPixKeyUsecase,
) PixKeyHandlerInterface {
	return &handler{
		registerPixKeyUsecase: registerPixKeyUsecase,
		findPixKeyUsecase:     findPixKeyUsecase,
	}
}

type handler struct {
	registerPixKeyUsecase commands.RegisterPixKeyUsecase
	findPixKeyUsecase     commands.FindPixKeyUsecase
}

// @BasePath /api/v1

// RegisterPixKey godoc
// @Summary Register Pix Key
// @Schemes
// @Description Register Pix Key
// @Tags pix-keys
// @Accept json
// @Produce json
// @Param request body request.RegisterPixKeyRequest true "register-pix-keys request"
// @Param Authorization header string true "Bearer token"
// @Success 201 {object} response.RegisterPixKeyResponse
// @Failure 500 {object} ErrorsResponse
// @Failure 400 {object} ErrorsResponse
// @Router /pix-keys [post]
// @Security ApiKeyAuth
func (p *handler) RegisterPixKey(c *gin.Context) {
	var request request2.RegisterPixKeyRequest

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

// @BasePath /api/v1

// FindPixKeyByKey godoc
// @Summary Find Pix By Key
// @Schemes
// @Description Find Pix By Key
// @Tags pix-keys/{key}
// @Accept json
// @Produce json
// @Param key path string true "pix-key parameter"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.RegisterPixKeyResponse
// @Failure 500 {object} ErrorsResponse
// @Failure 400 {object} ErrorsResponse
// @Router /pix-keys/{key} [get]
// @Security ApiKeyAuth
func (p *handler) FindPixKeyByKey(c *gin.Context) {
	var request request2.FindPixKeyRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	pixKeyDomain, err := p.findPixKeyUsecase.Execute(request.ToDTO())
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response.PixKeyDomainToFindWebResponse(pixKeyDomain))
}
