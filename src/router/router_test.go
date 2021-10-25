package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/jphacks/A_2108/src/domain"
)

// エンドポイントのテスト
func TestRoute(t *testing.T) {
	router := Route()
	tests := routerTestData()
	for _, tt := range tests {
		w := httptest.NewRecorder()
		t.Run(tt.name, func(t *testing.T) {
			if tt.req.method == "GET" || tt.req.method == "DELETE" {
				req, _ := http.NewRequest(tt.req.method, tt.req.url, tt.req.body)
				router.ServeHTTP(w, req)
			}

			if tt.statusCode != w.Code {
				t.Errorf("StatusCode = %v, want %v", w.Code, tt.statusCode)
				return
			}

			if !tt.isCheckResponseBody {
				return
			}

			var resBody domain.Fire1
			err := json.Unmarshal(w.Body.Bytes(), &resBody)
			if err != nil {
				t.Errorf("Json unmarshal error: %v", err)
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
