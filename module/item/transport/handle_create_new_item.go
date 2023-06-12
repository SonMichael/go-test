package transport

import (
	business "go-test/module/item/business"
	model "go-test/module/item/model"
	storage "go-test/module/item/storage"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleCreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.Item
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Name = strings.TrimSpace(data.Name)

		storage := storage.NewMysqlStorage(db)
		biz := business.NewCreateNewBiz(storage)
		if err := biz.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data.Id})
	}
}