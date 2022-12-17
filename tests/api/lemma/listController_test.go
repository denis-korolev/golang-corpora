package lemma

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"parser/app/api/lemma"
	"parser/app/lemma/repository"
	"parser/clients"
	"parser/config/route"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestListAction(t *testing.T) {
	//teardownSuite := setupSuite(t)
	//defer teardownSuite(t)

	router := route.SetupRoutes()

	w := httptest.NewRecorder()

	uri := "/list?"

	params := make(url.Values)
	params["T"] = []string{"муха"}
	params["V"] = []string{"Name,NOUN"}

	req, _ := http.NewRequest(http.MethodGet, uri+params.Encode(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := new(lemma.ListResponse)

	json.Unmarshal(w.Body.Bytes(), &result)

	assert.Equal(t, len(result.Data), 1)

}

/** метод чтобы перед тестом можно было что - то создать и потом удалить. Сейчас он не подключен */
func setupSuite(tb testing.TB) func(tb testing.TB) {
	fmt.Println("Создаем записи в таблице")

	es, err := clients.CreateElasticClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	repository.IndexLemmaData("dd", []byte("{\"message\":\"Hello World\"}"), es)

	// Return a function to teardown the test
	return func(tb testing.TB) {
		fmt.Println("Очищаем записи в таблице")
		repository.DeleteLemmaDataByID("dd", es)
	}
}
