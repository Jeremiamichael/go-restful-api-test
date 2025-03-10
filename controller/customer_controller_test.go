package controller

import (
	"bytes"
	"encoding/json"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/service/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestApp1(mockService *mocks.MockCustomerService) *fiber.App {
	app := fiber.New()
	customerController := NewCustomerController(mockService)

	api := app.Group("/api")
	customers := api.Group("/customers")
	customers.Post("/", customerController.Create)
	customers.Put("/:customerId", customerController.Update)
	customers.Delete("/:customerId", customerController.Delete)
	customers.Get("/:customerId", customerController.FindById)
	customers.Get("/", customerController.FindAll)

	return app
}

func TestCustomerController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockCustomerService(ctrl)
	app := setupTestApp1(mockService)

	tests := []struct {
		name           string
		method         string
		url            string
		body           interface{}
		setupMock      func()
		expectedStatus int
		expectedBody   web.WebResponse
	}{
		{
			name:   "Update customer - success",
			method: "PUT",
			url:    "/api/customers/1",
			body:   web.CustomerUpdateRequest{CustomerID: 1, Name: "Updated", Email: "updated@example.com", Phone: "123456789", Address: "Updated Address", LoyaltyPts: 100},
			setupMock: func() {
				mockService.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(web.CustomerResponse{CustomerID: 1, Name: "Updated", Email: "updated@example.com", Phone: "123456789", Address: "Updated Address", LoyaltyPts: 100}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: web.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   web.CustomerResponse{CustomerID: 1, Name: "Updated", Email: "updated@example.com", Phone: "123456789", Address: "Updated Address", LoyaltyPts: 100},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			var reqBody []byte
			if tt.body != nil {
				reqBody, _ = json.Marshal(tt.body)
			}

			req := httptest.NewRequest(tt.method, tt.url, bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var respBody web.WebResponse
			json.NewDecoder(resp.Body).Decode(&respBody)

			if dataMap, ok := respBody.Data.(map[string]interface{}); ok {
				respBody.Data = web.CustomerResponse{
					CustomerID: uint64(dataMap["customer_id"].(float64)),
					Name:       dataMap["name"].(string),
					Email:      dataMap["email"].(string),
					Phone:      dataMap["phone"].(string),
					Address:    dataMap["address"].(string),
					LoyaltyPts: int(dataMap["loyalty_pts"].(float64)),
				}
			}

			assert.Equal(t, tt.expectedBody, respBody)
		})
	}
}
