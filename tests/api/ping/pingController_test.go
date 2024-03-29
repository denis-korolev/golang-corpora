package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"parser/config/route"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestHelloAction(t *testing.T) {
	router := route.SetupRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Body.String(), "{\"message\":\"Hello World\"}")
}
