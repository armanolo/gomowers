package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApplicationManage(t *testing.T) {

	var tests = []struct {
		name     string
		request  string
		response string
		code     int
		err      string
	}{
		{"valid", "55\n12N\nLMLMLMLMM\n33E\nMMRMMRMRRM",
			`"13N\n51E\n"`, http.StatusOK, ""},
		{"not valid", "55\n12N", "", http.StatusBadRequest,
			`{"error":{"message":"bad format input data"}}`},
		{"valid", "55\n12N\nLM\n12N\nLM", "null", http.StatusBadRequest,
			`{"error":{"message":"there is another mower in the position: 1:2"}}`},
	}

	for _, test := range tests {
		rb := strings.NewReader(test.request)
		request := httptest.NewRequest("POST", "/manage", rb)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Manage)

		handler.ServeHTTP(rr, request)

		if rr.Code != test.code {
			t.Errorf("%q returned wrong status code: expected 200 but got %d", test.name, rr.Code)
		}

		bb, _ := io.ReadAll(rr.Body)
		br := string(bb)

		if test.err != "" {
			if test.err != br {
				t.Errorf("%q returned wrong response: expected %s but got %s", test.name, br, test.response)
			}

		} else {

			if br != test.response {
				t.Errorf("%q returned wrong response: expected %s but got %s", test.name, br, test.response)
			}
		}

	}

}
