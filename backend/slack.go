package backend

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"time"
)

var (
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
				MinVersion:         tls.VersionTLS12,
				// strongest ciphers
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				},

				// https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8
				CurvePreferences: []tls.CurveID{
					// secp521r1
					tls.CurveP521,
					// secp384r1
					tls.CurveP384,
					// secp256r1
					tls.CurveP256,
				},
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

func slackSendMessage(message string) {
	if !slackEnabled || len(slackWebHooks) == 0 {
		return
	}

	for _, hookURL := range slackWebHooks {
		body := map[string]string{"text": message}
		jsonBytes, _ := json.Marshal(body)
		res, err := slackClient.Post(hookURL, "application/json", bytes.NewBuffer(jsonBytes))
		if err != nil {
			log.Errorf("error sending slack message %s", err)
			continue
		}
		defer res.Body.Close()
		if res.StatusCode == 200 {
			log.Info("sent slack message successfully")
		} else {
			log.Errorf("slack message response code was: %v", res.StatusCode)
		}
	}
}
