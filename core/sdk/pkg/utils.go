package pkg

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	TrafficKey = "X-Request-Id"
	LoggerKey  = "_go-admin-logger-request"
)

func CompareHashAndPassword(e, p string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}

	return true, nil

}

func GenerateMsgIDFromContext(c *gin.Context) string {
	requestId := c.GetHeader(TrafficKey)
	if requestId == "" {
		requestId = uuid.New().String()
		c.Header(TrafficKey, requestId)
	}
	return requestId
}

func GetOrm(c *gin.Context) (*gorm.DB, error) {
	idb, exist := c.Get("db")
	if !exist {
		return nil, errors.New("db connect not exist")
	}
	switch idb.(type) {
	case *gorm.DB:
		//新增操作
		return idb.(*gorm.DB), nil
	default:
		return nil, errors.New("db connect not exist")
	}
}
