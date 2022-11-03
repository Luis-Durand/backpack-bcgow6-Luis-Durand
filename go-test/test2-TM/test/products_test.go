package products

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-test/test2-TM/internal/products"
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-test/test2-TM/server/handler"
	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-test/test2-TM/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockstore mocks.MockStorage) *gin.Engine {

	repo := products.NewRepository(&mockstore)
	serv := products.NewService(repo)
	p := handler.NewHandler(serv)

	r := gin.Default()

	pr := r.Group("/products")

	pr.PUT("/:id", p.Store())
	pr.DELETE("/:id", p.GetAll())

	return r

}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	return req, httptest.NewRecorder()

}

func TestUpdateProduct(t *testing.T) {

	/* UpdatedProd := products.Products{
		Id:     2,
		Name:   "Baggio",
		Type:   "jugo",
		Price:  120,
		Amount: 1,
	} */

	mockstorage := mocks.MockStorage{
		AllProds: []products.Products{
			{Id: 1,
				Name:   "Block",
				Type:   "Chocolate",
				Price:  200,
				Amount: 1,
			}, {
				Id:     2,
				Name:   "Cofler",
				Type:   "Chocolate",
				Price:  150,
				Amount: 1,
			},
		},
	}

	var resp products.Products
	r := createServer(mockstorage)
	route := fmt.Sprintf("/products/%d", 2)
	req, rr := createRequestTest(http.MethodPut, route, `{
	"name":"Baggio","type":"jugo","price":120,"Amount":1
}`)

	// act
	r.ServeHTTP(rr, req)

	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 200, rr.Code)

}

func TestDeleteProduct(t *testing.T) {

	/* UpdatedProd := products.Products{
		Id:     2,
		Name:   "Baggio",
		Type:   "jugo",
		Price:  120,
		Amount: 1,
	} */

	mockstorage := mocks.MockStorage{
		AllProds: []products.Products{
			{Id: 1,
				Name:   "Block",
				Type:   "Chocolate",
				Price:  200,
				Amount: 1,
			}, {
				Id:     2,
				Name:   "Cofler",
				Type:   "Chocolate",
				Price:  150,
				Amount: 1,
			},
		},
	}

	r := createServer(mockstorage)
	route := fmt.Sprintf("/products/%d", 2)
	req, rr := createRequestTest(http.MethodDelete, route, "")

	// act
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

}
