package controllers

import (
	"github.com/revel/revel"
	"encoding/json"
	"google.golang.org/api/oauth2/v2"
	"net/http"
	"strconv"
	"fmt"
	"time"

	"Fire-Dragon/app/models"
	"Fire-Dragon/app"
)

type Auth struct {
	*revel.Controller
}

func persistAuth(token string, email string, iss string, exp int64, c Auth) (user *models.UserAuth) {
	// save validated token info into db.
	user = &models.UserAuth{JWTSub: iss, Email: email, JWT: token, JWTExp: exp, AuthType: "google"}

	sql := user.UpsertSQL()
	_, err := app.DB.Query(sql)

	if err != nil {
		c.Log.Error(err.Error())
		panic(err)
	}

	return nil
}

var httpClient = &http.Client{}

func verifyIdToken(idToken string, c Auth) (*oauth2.Tokeninfo, error) {
	// call google's api to verify user's id token.
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func (c Auth) VerifyGoogleIdToken() revel.Result {
	// handle the app's request to verify user's google id token. response
	// 200 status if the token is valid.
	var m map[string]interface{}
	if err := json.Unmarshal(c.Params.JSON, &m); err == nil {
		token := m["idToken"].(string)

		info, err := verifyIdToken(token, c)
		if err == nil {
			httpStatusCode := strconv.Itoa(info.ServerResponse.HTTPStatusCode)
			if httpStatusCode != "200" {
				c.Response.Status = http.StatusBadRequest
				return c.RenderText("Google Auth Failed. Status Code: %s", httpStatusCode)
			}
			// save user login info into db.
			now := time.Now()
			persistAuth(token, info.Email, info.Audience, info.ExpiresIn + now.Unix(), c)
		} else {
			c.Log.Error(err.Error())
			c.Response.Status = http.StatusBadRequest
			return c.RenderText("Google Auth Failed. Error: %s", err.Error())
		}
	} else {
		// somthing wrong with the request parameter
		c.Response.Status = http.StatusBadRequest
		return c.RenderText(fmt.Sprintf("Bad Google Token. Err: %s", err))
	}

	return c.RenderText("OK")
}
