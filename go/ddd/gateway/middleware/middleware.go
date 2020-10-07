package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	libCtx "github.com/muhfaris/adsrobot/internal/context"
	"github.com/muhfaris/request"
)

type (
	// Exception is wrap error auth
	Exception struct {
		Error string `json:"Error"`
	}

	// Token is wrap token data
	Token struct {
		Data Claim `json:"data"`
	}

	// Claim is token data
	Claim struct {
		jwt.StandardClaims
		Identity         IdentityData `json:"identity"`
		MyIdentity       IdentityData `json:"my_identity"`
		CampaignIdentity IdentityData `json:"campaign_identity"`
	}

	// IdentityData is identity data
	IdentityData struct {
		AccessToken string    `json:"access_token"`
		SessionID   uuid.UUID `json:"session_id"`
		Role        string    `json:"role"`
		Username    string    `json:"username"`
		Email       string    `json:"email"`
	}
)

// APIsConfig is api config
type APIsConfig struct {
	User string
	Cron string
}

// Middleware is wrap config for middleware
type Middleware struct {
	APIs               APIsConfig
	ManagixInternalKey string
}

// AuthenticationMiddleware is to validation user
func (mw *Middleware) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Exception{Error: "Invalid authorization token"})
			return
		}

		url := fmt.Sprintf("%s/verify", mw.APIs.User)
		req := request.ReqApp{
			URL:           url,
			Authorization: header,
		}

		response, err := req.GET()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Exception{Error: "Invalid authorization token,can not reach verify API"})
			return
		}

		var token Token
		err = json.Unmarshal(response.Body, &token)
		if err != nil {
			json.NewEncoder(w).Encode(Exception{Error: "error unmarshal token"})
			return
		}

		if token.Data.Identity.AccessToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Exception{Error: "Token not found"})
			return
		}

		ctx := r.Context()
		ctx = libCtx.SetContext(ctx, libCtx.FBAccessToken, token.Data.Identity.AccessToken)
		ctx = libCtx.SetContext(ctx, libCtx.RoleContextKey, token.Data.Identity.Role)
		ctx = libCtx.SetContext(ctx, libCtx.UsernameContextKey, token.Data.Identity.Username)
		ctx = libCtx.SetContext(ctx, libCtx.ManagixToken, header)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

//InternalAuthenticationMiddleware is middleware to check incoming request from internal Managix Microservices
func (mw *Middleware) InternalAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("x-api-key")

		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Exception{Error: "Invalid authorization token"})
			return
		}

		if header != mw.ManagixInternalKey {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Exception{Error: "Invalid authorization token"})
			return
		}

		ctx := r.Context()
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
