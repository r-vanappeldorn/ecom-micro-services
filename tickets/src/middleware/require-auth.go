package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAuth(ctx *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://ingress-nginx-controller.ingress-nginx/api/users/currentuser", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Host", "ticketing.io")
	req.Header.Set("Cookie", ctx.Request.Header.Get("Cookie"))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	ctx.Next()
}
