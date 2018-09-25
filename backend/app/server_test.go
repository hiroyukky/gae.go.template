package app

import (
	"io"
	"os"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"google.golang.org/appengine/aetest"
)

type TestingRequestParam struct {
	T        *testing.T
	Path     string
	Method   string
	Header   map[string]string
	Param    io.Reader
}

var instance aetest.Instance = nil

func TestMain(m *testing.M) {
	opt := aetest.Options{StronglyConsistentDatastore: true}
	instance, _ = aetest.NewInstance(&opt)
	defer instance.Close()

	ret := m.Run()
	if ret != 0 {
		os.Exit(ret)
	}
}

func requestHTTPTest(p TestingRequestParam) *httptest.ResponseRecorder {
	req, _ := instance.NewRequest(p.Method, p.Path, p.Param)
	if len(p.Header) > 0 {
		for k, v := range p.Header {
			p.T.Log(k,v)
			req.Header.Add(k, v)
		}
	}
	res := httptest.NewRecorder()

	router := Router()
	router.ServeHTTP(res, req)

	return res
}

func TestSample(t *testing.T) {
	res := requestHTTPTest(TestingRequestParam{
		T:       t, 
		Path:    "/v1/sample",
		Method:  "GET",
	})
	t.Log("res : ", res.Code, res.Body)

	if res.Code != http.StatusOK {
		t.Fatalf("httpStatus: %d", res.Code)
	}
}
