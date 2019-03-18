package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, strconv.Itoa(p002()))
	})
	e.Logger.Fatal(e.Start(":8082"))
}

func fibo(n int) int {
	switch {
	case n < 1:
		return 0

	case n == 1:
		return 1
	case n == 2:
		return 2
	default:
		return fibo(n-1) + fibo(n-2)
	}
}

func p002() int {
	sum := 0
	for i := 0; i < 100; i++ {
		n := fibo(i)
		if n > 4000000 {
			break
		}
		if n%2 == 0 {
			sum += n
		}
	}
	return sum
}
