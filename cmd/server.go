package cmd

import (
	"context"
	"net/http"

	"aws-lambda-in-go-lang/models"

	"github.com/labstack/echo"
)

type RequestInput struct {
	PayLoad string `json:"payload"`
}

func RunAPIServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ping")
	})

	e.POST("/invoke", invoke)
	e.Logger.Fatal(e.Start(":8080"))
}

func invoke(c echo.Context) error {
	input := new(RequestInput)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	inp := []byte(`{
		"payLoad" : "` + input.PayLoad + `"
		}`)

	resp, err := Lambdahandler(context.Background(), inp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.LambdaResponse{
			Message: "error occured",
			Success: false,
			Error: map[string]string{
				"err": err.Err().Error(),
				"msg": err.Msg(),
			},
		})
	}
	return c.JSON(http.StatusOK, resp)

}
