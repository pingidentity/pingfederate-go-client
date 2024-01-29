package configurationapi

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	_ oauth2.TokenSource = OAuthValues{}
)

type OAuthValues struct {
	Transport    *http.Transport
	TokenUrl     string
	ClientId     string
	ClientSecret string
	Scopes       []string
}

func (oav OAuthValues) Token() (*oauth2.Token, error) {
	tokenVal, tokenErr := getToken(oav, context.Background())
	return tokenVal, tokenErr
}

func getToken(oav OAuthValues, ctx context.Context) (*oauth2.Token, error) {
	oauthConfig := &clientcredentials.Config{
		TokenURL:     oav.TokenUrl,
		ClientID:     oav.ClientId,
		ClientSecret: oav.ClientSecret,
		Scopes:       oav.Scopes,
	}

	httpClient := &http.Client{Transport: oav.Transport}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	return oAuthExponentialBackOffRetry(func() (*oauth2.Token, error) {
		return oAuthProcessResponse(func() (*oauth2.Token, error) {
			return oauthConfig.Token(ctx)
		})
	})
}

func oAuthProcessResponse(f func() (*oauth2.Token, error)) (*oauth2.Token, error) {
	return oAuthExponentialBackOffRetry(f)
}

func oAuthExponentialBackOffRetry(f func() (*oauth2.Token, error)) (*oauth2.Token, error) {
	var token *oauth2.Token
	var err error
	backOffTime := time.Second
	var isRetryable bool

	for i := 0; i < maxRetries; i++ {
		token, err = f()

		if err != nil {
			backOffTime, isRetryable = oAuthTestForRetryable(err, backOffTime)

			if isRetryable {
				log.Printf("Attempt %d failed: %v, backing off by %s.", i+1, err, backOffTime.String())
				time.Sleep(backOffTime)
				continue
			}
		}

		return token, err
	}

	log.Printf("Request failed after %d attempts", maxRetries)

	return token, err // output the final error
}

func oAuthTestForRetryable(err error, currentBackoff time.Duration) (time.Duration, bool) {

	backoff := currentBackoff * 2

	retryAbleCodes := []int{
		429,
		500,
		502,
		503,
		504,
	}

	for _, v := range retryAbleCodes {
		if m, mErr := regexp.MatchString(fmt.Sprintf("%d ", v), err.Error()); mErr == nil && m {
			log.Printf("HTTP status code %d detected, available for retry", v)
			return backoff, true
		}
	}

	return backoff, false
}
