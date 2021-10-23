package router

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/jphacks/A_2108/src/domain"
)

type htttpReq struct {
	method string
	url    string
	body   io.Reader
}

// エンドポイントのテスト
func TestRoute(t *testing.T) {
	router := Route()
	w := httptest.NewRecorder()
	tests := []struct {
		name         string
		req          htttpReq
		statusCode   int
		responseBody interface{}
	}{
		// ここにテストケースを追加してください
		{
			name: "/fire success test",
			req: htttpReq{
				method: "GET",
				url:    "/fire",
				body:   nil,
			},
			statusCode: 201,
			responseBody: domain.Fire1{
				A: 1,
				B: "mieruka?",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.req.method == "GET" || tt.req.method == "DELETE" {
				req, _ := http.NewRequest(tt.req.method, tt.req.url, tt.req.body)
				router.ServeHTTP(w, req)
			}

			if tt.statusCode != w.Code {
				t.Errorf("StatusCode = %v, want %v", w.Code, tt.statusCode)
				return
			}

			var resBody domain.Fire1
			err := json.Unmarshal(w.Body.Bytes(), &resBody)
			if err != nil {
				t.Errorf("Json unmarshal error")
				return
			}

			ok := reflect.DeepEqual(resBody, tt.responseBody)
			if !ok {
				t.Errorf("Body = %v, want %v", resBody, tt.responseBody)
				return
			}
		})
	}
}
