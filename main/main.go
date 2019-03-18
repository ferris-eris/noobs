package main

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	var result string
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		resp, err := http.Get("http://localhost:8081")
		if err != nil {
			result += err.Error()
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				result += err.Error()
			} else {
				result += string(body) + "\n"
			}
		}
		resp, err = http.Get("http://localhost:8082")
		if err != nil {
			result += err.Error()
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				result += err.Error()
			} else {
				result += string(body) + "\n"
			}
		}
		resp, err = http.Get("http://localhost:8083")
		if err != nil {
			result += err.Error()
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				result += err.Error()
			} else {
				result += string(body) + "\n"
			}
		}
		return c.String(http.StatusOK, result)
	})
	e.Logger.Fatal(e.Start(":5000"))
}
