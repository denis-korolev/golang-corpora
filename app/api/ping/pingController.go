package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description Метод для вывода тестового сообщения
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} ping.Response
// @Router /ping [get]
func PingAction(c *gin.Context) {

	resp := new(Response)

	resp.Message = "Hello World"

	c.JSON(http.StatusOK, resp)
}
