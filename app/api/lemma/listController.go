package lemma

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /

// LemmaList godoc
// @Summary Вывод списка лемм по запросу
// @Schemes
// @Description Метод для вывода списка лемм
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} lemma.ListResponse
// @Router /lemma [get]
func ListAction(c *gin.Context) {

	resp := new(ListResponse)

	item := new(ListItem)
	item.Text = "sfsfsf"

	resp.Data = append(resp.Data, *item)

	c.JSON(http.StatusOK, resp)
}
