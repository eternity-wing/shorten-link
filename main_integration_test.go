package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestIndexRoute(t *testing.T) {

	os.Setenv("APP_ENV", "test")
	app := setup()

	tests := []struct {
		description string

		// Test input
		route  string
		method string
		body   *strings.Reader

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "Valid Shorten request",
			route:         "/api/v1/links",
			method:        "POST",
			body:          strings.NewReader(`{"URL": "https://google.com","userid" : 22}`),
			expectedError: false,
			expectedCode:  fiber.StatusOK,
			expectedBody:  "",
		},
		{
			description:   "Invalid Request body",
			route:         "/api/v1/links",
			method:        "POST",
			body:          strings.NewReader(`{"URL": "https://google.com"`),
			expectedError: false,
			expectedCode:  fiber.StatusBadRequest,
			expectedBody:  `{"error": "cannot parse JSON", "title": "Invalid request format","type": "Invalid request"}`,
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest(
			test.method,
			test.route,
			test.body,
		)

		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		// Verify, that the reponse body equals the expected body
		if test.expectedBody != "" {
			assert.Equal(t, strings.Replace(test.expectedBody, " ", "", -1), strings.Replace(string(body), " ", "", -1), test.description)
		}

	}

}
