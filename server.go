// +build go1.8

// enforce go 1.8+ just so we can support X25519 curve :)

package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func StartServer() (err error) {

	// create routes
	mux := newRouter()

	// get server config
	srv := configureHttpServer(mux)

	// get TLS config
	tlsConifig, err := configureTLS()
	if err != nil {
		log.Fatalf("error configuring TLS: %s", err)
		return
	}
	srv.TLSConfig = &tlsConifig

	// start the server
	if viper.GetBool("server.tls.enabled") {
		// cert and key should already be configured
		log.Info("starting HTTP server with TLS")
		err = srv.ListenAndServeTLS("", "")
	} else {
		err = srv.ListenAndServe()
	}

	if err != nil {
		log.Info("starting HTTP server")
		log.Fatalf("failed to start server: %s", err)
	}

	return
}

func configureHttpServer(mux *mux.Router) (httpServer *http.Server) {

	// apply standard http server settings
	address := fmt.Sprintf(
		"%s:%s",
		viper.GetString("server.bind_address"),
		viper.GetString("server.bind_port"),
	)

	httpServer = &http.Server{
		Addr: address,
		// set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,

		// the maximum amount of time to wait for the
		// next request when keep-alives are enabled
		IdleTimeout: time.Second * 60,
	}

	// explicitly enable keep-alives
	httpServer.SetKeepAlivesEnabled(true)

	// stdout access log enable/disable
	if viper.GetBool("server.access_log") {
		httpServer.Handler = handlers.CombinedLoggingHandler(os.Stdout, mux)
	} else {
		httpServer.Handler = mux
	}

	return
}

// configure TLS as defined in configuration
func configureTLS() (tlsConfig tls.Config, err error) {

	if !viper.GetBool("server.tls.enabled") {
		log.Debug("TLS not enabled, skipping TLS config")
		return
	}

	// attempt to load configured cert/key
	log.Info("TLS enabled, loading cert and key")
	log.Debugf("loading TLS cert and key: %s %s", viper.GetString("server.tls.cert_chain"), viper.GetString("server.tls.private_key"))
	cert, err := tls.LoadX509KeyPair(viper.GetString("server.tls.cert_chain"), viper.GetString("server.tls.private_key"))
	if err != nil {
		return
	}

	// configure tls settings
	tlsConfig.Certificates = []tls.Certificate{cert}
	tlsConfig.MinVersion = tls.VersionTLS12
	tlsConfig.InsecureSkipVerify = false
	tlsConfig.PreferServerCipherSuites = true

	// https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8
	// http://safecurves.cr.yp.to/
	tlsConfig.CurvePreferences = []tls.CurveID{
		// this curve is a non-NIST curve with no NSA influence. Prefer this over all others!
		tls.X25519,
		// The following curves are provided by NIST
		// secp521r1
		tls.CurveP521,
		// secp384r1
		tls.CurveP384,
		// secp256r1
		tls.CurveP256,
	}

	// allowed ciphers. Disable CBC suites (Lucky13)
	tlsConfig.CipherSuites = []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	}

	return
}
