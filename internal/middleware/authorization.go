package middleware

import(
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"xs-coin-checker/internal/httpx"
	"xs-coin-checker/internal/tools"
)

var UnAuthorizedError = errors.New("Invalid CRedentials or Tokens")

func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(res http.ResponseWriter , req *http.Request){

		var username string = req.URL.Query().Get("username")
		var token = req.Header.Get("Authorization")
		var err error

		if username == "" || token == ""{
			log.Error(UnAuthorizedError)
			httpx.RequestErrorHandler(res,UnAuthorizedError)
			return
		}

		var database  *tools.DatabaseInterface
		database , err = tools.NewDatabase()
		if err != nil{
			httpx.InternalErrorHandler(res)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token !=  (*loginDetails).AuthToken)){
			log.Error(UnAuthorizedError)
			httpx.RequestErrorHandler(res,UnAuthorizedError)
			return
		}

		next.ServeHTTP(res,req)
	})
}