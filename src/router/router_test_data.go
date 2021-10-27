package router

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
	"io/ioutil"

	"github.com/jphacks/A_2108/src/domain"
)

type htttpReq struct {
	method string
	url    string
	body   io.Reader
}

type testData []struct {
	name                string
	req                 htttpReq
	statusCode          int
	responseBody        interface{}
	isCheckResponseBody bool
}

func convertToIoReader(it interface{}) io.Reader {
	jsonIt, _ := json.Marshal(it)
	return bytes.NewReader(jsonIt)
}

func routerTestData() testData {
	/* image, _ := ioutil.ReadFile("test.jpeg") */

	return testData{
		// ここにテストケースを追加してください
		{
			name: "/fire success test",
			req: htttpReq{
				method: "GET",
				url:    "/fire",
				body:   nil,
			},
			statusCode: 200,
			responseBody: domain.Fire1{
				A: 1,
				B: "mieruka?",
			},
			isCheckResponseBody: true,
		},
		{
			name: "/fire pathpara success test",
			req: htttpReq{
				method: "GET",
				url:    "/fire/pathpara/100",
				body:   nil,
			},
			statusCode: 200,
			responseBody: domain.Fire1{
				A: 1,
				B: "100",
			},
			isCheckResponseBody: true,
		},
		{
			name: "/user GET success test",
			req: htttpReq{
				method: "GET",
				url:    "/user",
				body:   nil,
			},
			statusCode:          200,
			isCheckResponseBody: false,
		},
		{
			name: "/plan GET success test",
			req: htttpReq{
				method: "GET",
				url:    "/plan",
				body:   nil,
			},
			statusCode:          200,
			isCheckResponseBody: false,
		},
		{
			name: "/plan GET pathparam success test",
			req: htttpReq{
				method: "GET",
				url:    "/plan/100",
				body:   nil,
			},
			statusCode:          200,
			isCheckResponseBody: false,
		},
		{
			name: "/plan POST success test",
			req: htttpReq{
				method: "POST",
				url:    "/plan",
				body: convertToIoReader(
					domain.Plan{},
				),
			},
			statusCode:          200,
			isCheckResponseBody: false,
		},
		{
			name: "/plan DELETE success test",
			req: htttpReq{
				method: "DELETE",
				url:    "/plan/100",
				body:   nil,
			},
			statusCode:          200,
			isCheckResponseBody: false,
		},
		/* 		{
			name: "/image POST success test",
			req: htttpReq{
				method: "POST",
				url:    "/image",
				body:   convertToIoReader(image),
			},
			statusCode:          200,
			isCheckResponseBody: false,
		}, */
	}
}
