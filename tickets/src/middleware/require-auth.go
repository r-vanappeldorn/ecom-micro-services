package middleware

import (
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"

	"ticketing.io/src/customerrors"
)

func RequireAuth(ctx *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://auth-srv:3000/api/users/currentuser", nil)
	if err != nil {
		customerrors.SendInternalServerError(ctx, err)
		return
	}

	req.Header.Set("Host", "ticketing.io")
	req.Header.Set("Cookie", ctx.Request.Header.Get("Cookie"))

	resp, err := client.Do(req)
	if err != nil {
		customerrors.SendInternalServerError(ctx, err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		customerrors.SendInternalServerError(ctx, err)
		return
	}

	var currentUser struct {
		CurrentUser interface{} `json:"currentUser"`
	}

	err = json.Unmarshal(body, &currentUser)
	if err != nil {
		customerrors.SendInternalServerError(ctx, err)
		return
	}

	if currentUser.CurrentUser == nil {
		customerrors.SendBadRequestError(ctx, "Not authorized")
		return
	}

	ctx.Next()
}
