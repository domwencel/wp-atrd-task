package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/pawmart/wp-atrd-task/internal/http/util"
	"net/http"
)

func (a *App) CreateSecretHandler() gin.HandlerFunc {
	type request struct {
		Secret           string  `form:"secret" binding:"required"`
		ExpireAfterViews uint32  `form:"expireAfterViews" binding:"required,gt=0"`
		ExpireAfter      *uint32 `form:"expireAfter" binding:"required,gte=0"`
	}

	return func(c *gin.Context) {
		var req request

		if err := c.ShouldBindWith(&req, binding.FormPost); err != nil {
			// TODO add errors to response
			c.Status(http.StatusMethodNotAllowed)
			return
		}

		secret := a.AddSecret(req.Secret, req.ExpireAfterViews, *req.ExpireAfter)
		util.PrepareResponse(http.StatusOK, secret, c)
	}
}


func (a *App) GetSecretHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		secret, exist := a.GetSecret(id)
		if !exist {
			c.Status(http.StatusNotFound)
			return
		}

		util.PrepareResponse(http.StatusOK, secret, c)
	}
}