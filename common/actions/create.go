package actions

import (
	"goconf/common/dto"
	"goconf/common/models"
	"goconf/core/sdk/api"
	"goconf/core/sdk/pkg"
	"goconf/core/sdk/pkg/jwtauth/user"
	"goconf/core/sdk/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAction(control dto.Control) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := api.GetRequestLogger(c)
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}
		req := control.Generate()
		err = req.Bind(c)
		if err != nil {
			response.Error(c, http.StatusUnprocessableEntity, err, err.Error())
			return
		}

		var object models.ActiveRecord
		object, err = req.GenerateM()
		if err != nil {
			response.Error(c, 500, err, "模型生成失败")
			return
		}
		object.SetCreateBy(user.GetUserId(c))
		err = db.WithContext(c).Create(object).Error
		if err != nil {
			log.Errorf("Create error: %s", err)
			response.Error(c, 500, err, "创建失败")
			return
		}
		response.OK(c, object.GetId(), "创建成功")
		c.Next()
	}

}
