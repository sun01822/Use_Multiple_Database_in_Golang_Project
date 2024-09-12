package methodutils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"time"
)

func LogAPICall(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("Time: %s, API called: %s %s\n", currentTime, c.Request().Method, c.Request().RequestURI)
		return next(c)
	}
}
