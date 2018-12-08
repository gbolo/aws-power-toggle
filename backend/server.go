// +build go1.8

// enforce go 1.8+ just so we can support X25519 curve :)

package backend

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

var (
	// set timeouts to avoid Slowloris attacks.
	httpWriteTimeout = time.Second * 15
	httpReadTimeout  = time.Second * 15
	// the maximum amount of time to wait for the
	// next request when keep-alives are enabled
	httpIdleTimeout = time.Second * 60

	// PCI compliance as of Jun 30, 2018: anything under TLS 1.1 must be disabled
	// we bump this up to TLS 1.2 so we can support best possible ciphers
	tlsMinVersion = uint16(tls.VersionTLS12)
	// allowed ciphers when in hardened mode
	// disable CBC suites (Lucky13 attack) this means TLS 1.1 can't work (no GCM)
	// only use perfect forward secrecy ciphers
	tlsCiphers = []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		// these ciphers require go 1.8+
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
	}
	// EC curve preference when in hardened mode
	// curve reference: http://safecurves.cr.yp.to/
	tlsCurvePreferences = []tls.CurveID{
		// this curve is a non-NIST curve with no NSA influence. Prefer this over all others!
		// this curve required go 1.8+
		tls.X25519,
		// These curves are provided by NIST; prefer in descending order
		tls.CurveP521,
		tls.CurveP384,
		tls.CurveP256,
	}
)

func startHTTPServer() (err error) {

	// create routes
	mux := newRouter()

	// get server config
	srv := configureHTTPServer(mux)

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

func configureHTTPServer(mux *mux.Router) (httpServer *http.Server) {

	// apply standard http server settings
	address := fmt.Sprintf(
		"%s:%s",
		viper.GetString("server.bind_address"),
		viper.GetString("server.bind_port"),
	)

	httpServer = &http.Server{
		Addr: address,

		WriteTimeout: httpWriteTimeout,
		ReadTimeout:  httpReadTimeout,
		IdleTimeout:  httpIdleTimeout,
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

	// configure hardened TLS settings
	tlsConfig.Certificates = []tls.Certificate{cert}
	tlsConfig.MinVersion = tlsMinVersion
	tlsConfig.InsecureSkipVerify = false
	tlsConfig.PreferServerCipherSuites = true
	tlsConfig.CurvePreferences = tlsCurvePreferences
	tlsConfig.CipherSuites = tlsCiphers

	return
}
