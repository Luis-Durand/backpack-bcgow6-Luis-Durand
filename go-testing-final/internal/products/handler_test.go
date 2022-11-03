package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var errGet error = errors.New("Error in get")

func createServer() *gin.Engine {

	/* repo := NewRepository() */

	NewProds := []Product{{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	},
		{
			ID:          "stub",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		}}

	serv := MockService{ErrorService: nil, Prods: NewProds}

	p := NewHandler(&serv)

	r := gin.Default()

	pr := r.Group("/products")

	pr.GET("", p.GetProducts)

	return r
}

func createServerFail() *gin.Engine {
	serv := MockService{ErrorService: errGet}

	p := NewHandler(&serv)

	r := gin.Default()

	pr := r.Group("/products")

	pr.GET("", p.GetProducts)

	return r

}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {

	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	return req, httptest.NewRecorder()

}

func TestGetProduct(t *testing.T) {

	var resp []Product
	r := createServer()
	route := fmt.Sprintf("/products?seller_id=%s", "pepe")
	req, rr := createRequestTest(http.MethodGet, route, "")

	r.ServeHTTP(rr, req)

	err := json.Unmarshal(rr.Body.Bytes(), &resp)

	assert.Nil(t, err)
	assert.Equal(t, 200, rr.Code)

}

func TestGetProductFail(t *testing.T) {

	var resp []Product
	r := createServer()

	req, rr := createRequestTest(http.MethodGet, "/products", "")

	r.ServeHTTP(rr, req)

	err := json.Unmarshal(rr.Body.Bytes(), &resp)

	assert.NotNil(t, err)
	assert.Equal(t, 400, rr.Code)

	var respFail []Product
	r = createServerFail()
	req, rr = createRequestTest(http.MethodGet, "/products?seller_id=pepe", "")

	r.ServeHTTP(rr, req)

	err = json.Unmarshal(rr.Body.Bytes(), &respFail)

	assert.NotNil(t, err)
	assert.Equal(t, 500, rr.Code)

}
