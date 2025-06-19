package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/internal/application/auth/handler"
	"github.com/radityacandra/go-cms/internal/application/auth/service"
	"github.com/radityacandra/go-cms/internal/application/auth/types"
	mockService "github.com/radityacandra/go-cms/mocks/internal_/application/auth/service"
	"github.com/radityacandra/go-cms/pkg/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestAuthLoginPost(t *testing.T) {
	type fields struct {
		Service service.IService
	}
	type args struct {
		reqBody    api.AuthLoginRequest
		reqBodyStr string
	}

	type expected struct {
		statusCode int
		body       interface{}
	}

	type test struct {
		name     string
		fields   fields
		args     args
		mock     func(tt test) test
		expected expected
	}

	tests := []test{
		{
			name: "should return error if failed to bind request",
			args: args{
				reqBodyStr: `{"username": false, "password": "somepassword"}`,
			},
			mock: func(tt test) test {
				return tt
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				body: api.DefaultErrorResponse{
					Error: "code=400, message=Unmarshal type error: expected=string, got=bool, field=username, offset=18, internal=json: cannot unmarshal bool into Go struct field AuthLoginRequest.username of type string",
				},
			},
		},
		{
			name: "should return error if failed to validate request",
			args: args{
				reqBodyStr: `{"username": "", "password": ""}`,
			},
			mock: func(tt test) test {
				return tt
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				body: api.DefaultErrorResponse{
					Error: "password must not be empty.\nusername must not be empty.",
				},
			},
		},
		{
			name: "should return error if service throw an error",
			args: args{
				reqBodyStr: `{"username": "someuser", "password": "somepassword"}`,
			},
			mock: func(tt test) test {
				mockService := mockService.NewMockIService(t)
				tt.fields.Service = mockService

				mockService.EXPECT().Login(mock.Anything, types.LoginInput{
					Username: "someuser",
					Password: "somepassword",
				}).Return(types.LoginOutput{}, types.ErrPasswordMissmatch).Times(1)

				return tt
			},
			expected: expected{
				statusCode: http.StatusUnauthorized,
				body: api.DefaultErrorResponse{
					Error: "username or password is incorrect",
				},
			},
		},
		{
			name: "should return correct token",
			args: args{
				reqBodyStr: `{"username": "someuser", "password": "somepassword"}`,
			},
			mock: func(tt test) test {
				mockService := mockService.NewMockIService(t)
				tt.fields.Service = mockService

				mockService.EXPECT().Login(mock.Anything, types.LoginInput{
					Username: "someuser",
					Password: "somepassword",
				}).Return(types.LoginOutput{
					Token: "sometoken",
				}, nil).Times(1)

				return tt
			},
			expected: expected{
				statusCode: http.StatusOK,
				body: api.AuthLoginResponse{
					Token: "sometoken",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt = tt.mock(tt)

			e := echo.New()
			e.Validator = validator.NewValidator()

			var reqBody string
			if tt.args.reqBodyStr != "" {
				reqBody = tt.args.reqBodyStr
			} else {
				bytes, _ := json.Marshal(tt.args.reqBody)
				reqBody = string(bytes)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			h := &handler.Handler{
				Logger:  zap.NewNop(),
				Service: tt.fields.Service,
			}
			err := h.AuthLoginPost(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.statusCode, rec.Code)

			body, _ := io.ReadAll(rec.Result().Body)

			if tt.expected.statusCode != http.StatusOK {
				var bodyStruct api.DefaultErrorResponse
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.body, bodyStruct)
			} else {
				var bodyStruct api.AuthLoginResponse
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.body, bodyStruct)
			}
		})
	}
}
