package initializers

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var OAuthConfig *oauth2.Config

func SetupOAuth2() {
    OAuthConfig = &oauth2.Config{
        ClientID:     "your-client-id",
        ClientSecret: "your-client-secret",
        RedirectURL:  "http://localhost:8080/oauth/callback",
        Scopes:       []string{"openid", "profile", "email"},
        Endpoint:     google.Endpoint,
    }
}