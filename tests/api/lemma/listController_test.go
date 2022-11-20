package lemma

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"log"
	"net/http"
	"net/http/httptest"
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

	//todo сделать объект запроса с параметрами
	//todo сравнить ответ
	//todo сделать описание граммем, чтобы можно было по ним фильтровать

	router := route.SetupRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/list", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Body.String(), "{\"message\":\"Hello World\"}")
}

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
