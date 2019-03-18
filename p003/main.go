package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, strconv.Itoa(p003()))
	})
	e.Logger.Fatal(e.Start(":8083"))
}

func prime_decompo(n int) []int {
	table := []int{1}
	i := 2
	for i*i <= n {
		for n%i == 0 {
			n /= i
			table = append(table, i)
		}
		i++
	}
	if n > 1 {
		table = append(table, n)
	}
	return table
}

func p003() int {
	pl := prime_decompo(600851475143)
	return pl[len(pl)-1]
}
