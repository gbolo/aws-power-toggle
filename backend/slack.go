package backend

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

var (
	// values are set by ConfigInit
	slackClient   *http.Client
	slackEnabled  bool
	slackWebHooks []string
)

func init() {
	slackClient = createHTTPClient()
}

// createHTTPClient creates an http client which we can reuse
func createHTTPClient() *http.Client {
	return &http.Client{
		// http request timeout
		Timeout: 5 * time.Second,

		// transport settings
		Transport: &http.Transport{
			// TLS Config
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
				MinVersion:         tlsMinVersion,
				CipherSuites:       tlsCiphers,
				CurvePreferences:   tlsCurvePreferences,
			},

			// sane timeouts
			ExpectContinueTimeout: 1 * time.Second,
			IdleConnTimeout:       90 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
			MaxIdleConns:          10,
			MaxIdleConnsPerHost:   5,
			DisableCompression:    true,

			// dialer timeouts
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
		},
	}
}

// slackSendMessage boradcasts a text message to all configured webhooks
func slackSendMessage(message string) (errs []error) {
	if !slackEnabled || len(slackWebHooks) == 0 || message == "" {
		return
	}

	for i, hookURL := range slackWebHooks {
		body := map[string]string{"text": message}
		jsonBytes, _ := json.Marshal(body)
		res, err := slackClient.Post(hookURL, "application/json", bytes.NewBuffer(jsonBytes))
		if err != nil {
			log.Errorf("error sending slack message (%d/%d) %s", i+1, len(slackWebHooks), err)
			errs = append(errs, err)
			continue
		}
		defer res.Body.Close()
		if res.StatusCode == 200 {
			log.Infof("sent slack message successfully (%d/%d)", i+1, len(slackWebHooks))
		} else {
			log.Errorf("slack API response code was not successful (%d/%d): %v", i+1, len(slackWebHooks), res.StatusCode)
			errs = append(errs, fmt.Errorf("slack API response code was not successful (%d/%d): %v", i+1, len(slackWebHooks), res.StatusCode))
		}
	}

	return
}
