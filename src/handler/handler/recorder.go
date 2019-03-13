package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func record(t *testing.T, method, path string, params map[string]string, body interface{}, handlerFunc func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	t.Helper()

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		t.Error(err)
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(recorder, req)
	return recorder
}
