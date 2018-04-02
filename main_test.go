package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHello(t *testing.T){

   req, err := http.NewRequest("GET", "/", nil)
   if err != nil {
       t.Fatal(err)
   }

   responseRecorder := httptest.NewRecorder()
   handler := http.HandlerFunc(Hello)
   handler.ServeHTTP(responseRecorder, req)

   if status := responseRecorder.Code; status != http.StatusOK {
        t.Errorf("Got wrong status code! We expected: %v, but got: %v",
            http.StatusOK, status)
    }

   expected := "hello world"
   if responseRecorder.Body.String() != expected {
       t.Errorf("Something wrong! We expected: %v, but got: %v", expected, responseRecorder.Body.String())
    }

}
