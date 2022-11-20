package lemma

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"parser/app/lemma/repository"
	"parser/clients"
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

	es, err := clients.CreateElasticClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	result := repository.SearchLemmaData("lemma", "T=муха", es)

	//тут надо сделать вывод даных

	resp := new(ListResponse)

	item := new(ListItem)
	item.Text = string(result.Hits.Total.Value)

	resp.Data = append(resp.Data, *item)

	c.JSON(http.StatusOK, resp)
}
