package lemma

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"parser/app/lemma/repository"
	"parser/clients"
	"strings"
)

// @BasePath /

// LemmaList godoc
// @Summary Вывод списка лемм по запросу
// @Schemes
// @Description Метод для вывода списка лемм. http://opencorpora.org/dict.php?act=gram
// @Tags list
// @Produce json
// @Param   T     query     string     true  "Слово по которому будет поиск"     example(муха)
// @Param   V     query     []string     false  "Часть речи"     Enums(Name, NOUN, ADJF, ADJS, COMP, VERB, INFN, PRTF, PRTS, GRND, NUMR, ADVB, NPRO, PRED, PREP, CONJ, PRCL, INTJ)
// @Success 200 {object} lemma.ListResponse
// @Router /list [get]
func ListAction(c *gin.Context) {

	var request LemmaSearchRequest
	es, err := clients.CreateElasticClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	if errBind := c.ShouldBindQuery(&request); errBind != nil {
		c.String(http.StatusOK, errBind.Error())
		return
	}

	query := "T=" + request.T

	if request.V != "" {
		query = query + " AND ("
		words := strings.Split(request.V, ",")
		for idx, word := range words {
			log.Println(idx)
			if idx == 0 {
				query = query + "V=" + word
			} else {
				query = query + " OR V=" + word
			}
		}
		query = query + ")"
	}

	log.Println(query)

	result, errorRequest := repository.SearchLemmaData("lemma", query, es)
	if errorRequest != nil {
		c.JSON(http.StatusOK, errorRequest.Error())
		return
	}

	resp := new(ListResponse)
	for _, hit := range result.Hits.Hits {
		resp.Data = append(resp.Data, hit.Source)
	}

	c.JSON(http.StatusOK, resp)
}

type LemmaSearchRequest struct {
	T string `form:"T"  binding:"required"`
	V string `form:"V"`
}
