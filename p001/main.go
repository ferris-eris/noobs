package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, strconv.Itoa(p001()))
	})
	e.Logger.Fatal(e.Start(":8081"))
}

func filter(n int) bool {
	if n%3 == 0 || n%5 == 0 {
		return true
	}
	return false
}

func fn(val int, sum int) int {
	return sum + val
}

func p001() int {
	sum := 0
	for i := 0; i < 1000; i++ {
		if filter(i) {
			sum = fn(i, sum)
		}
	}
	return sum
}
