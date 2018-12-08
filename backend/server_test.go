package backend

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

const (
	// these constants may need to change when testdata content changes
	caRootCertSubjectCN    = "linuxctl ECC Root Certification Authority (Test)"
	caIntCertSubjectCN     = "linuxctl ECC Intermediate Certification Authority (Test)"
	serverCertSubjectCN    = "power-toggle"
	serverCertFileLocation = "../testdata/tls/server_power-toggle-chain.pem"
	serverKeyFileLocation  = "../testdata/tls/server_power-toggle-key.pem"
)

// helper function to create TLS config while handling error
func createTestTLSConfig(t *testing.T) *tls.Config {
	testTLSConfig, err := configureTLS()
	if err != nil {
		t.Fatalf("we got an unexpected error while calling createTLSConfig: %s", err)
	}

	return &testTLSConfig
}

// TestTLSConfig should test the behaviour of configureTLS
func TestTLSConfig(t *testing.T) {

	// configure TLS options
	os.Setenv("POWER_TOGGLE_SERVER_TLS_ENABLED", "true")
	os.Setenv("POWER_TOGGLE_SERVER_TLS_CERT_CHAIN", serverCertFileLocation)
	os.Setenv("POWER_TOGGLE_SERVER_TLS_PRIVATE_KEY", serverKeyFileLocation)
	testTLSConfig := createTestTLSConfig(t)

	if testTLSConfig.MinVersion != tlsMinVersion {
		t.Error("MinVersion is not set to expected value")
	}
	if testTLSConfig.InsecureSkipVerify {
		t.Error("InsecureSkipVerify is not set to false")
	}
	if !testTLSConfig.PreferServerCipherSuites {
		t.Errorf("PreferServerCipherSuites is not set to true")
	}
	if len(testTLSConfig.CipherSuites) == 0 {
		t.Error("CipherSuites is not set")
	}
	if len(testTLSConfig.CurvePreferences) == 0 {
		t.Error("CurvePreferences is not set")
	}
	for i := range testTLSConfig.CipherSuites {
		if testTLSConfig.CipherSuites[i] != tlsCiphers[i] {
			t.Error("discrepancy found in CipherSuites")
		}
	}
	for i := range testTLSConfig.CurvePreferences {
		if testTLSConfig.CurvePreferences[i] != tlsCurvePreferences[i] {
			t.Error("discrepancy found in CurvePreferences")
		}
	}

	// test that the expected certs are loaded
	if len(testTLSConfig.Certificates) == 0 {
		t.Fatal("certificate file was not loaded")
	} else if len(testTLSConfig.Certificates) > 1 {
		t.Fatalf("more than 1 certficate file was loaded: %v", len(testTLSConfig.Certificates))
	}

	// our test chain cert file should have 3 certs (server > intermediate ca > root ca)
	if len(testTLSConfig.Certificates[0].Certificate) != 3 {
		t.Fatalf("expected to have 3 x509 certificates loaded, but found: %v", len(testTLSConfig.Certificates[0].Certificate))
	}

	// confirm the correct 3 certificates are loaded (we use subjects here, instead of SKI)
	for _, cert := range testTLSConfig.Certificates[0].Certificate {
		cert, err := x509.ParseCertificate(cert)
		if err != nil {
			t.Fatalf("failed to parse certficate: %v", err)
		}

		switch {
		case bytes.Contains(cert.RawSubject, []byte(caRootCertSubjectCN)):
			t.Log("found root ca cert cn")
		case bytes.Contains(cert.RawSubject, []byte(caIntCertSubjectCN)):
			t.Log("found intermediate ca cert cn")
		case bytes.Contains(cert.RawSubject, []byte(serverCertSubjectCN)):
			t.Log("found server cert cn")
		default:
			t.Fatal("expected CN not found in certificate subject")
		}
	}

	// confirm that we can disable TLS
	os.Unsetenv("POWER_TOGGLE_SERVER_TLS_ENABLED")
	testTLSConfig = createTestTLSConfig(t)
	if len(testTLSConfig.Certificates) != 0 {
		t.Fatal("failed to disable TLS")
	}
}

func TestHTTPConfig(t *testing.T) {
	httpServerConfig := configureHTTPServer(&mux.Router{})

	if httpServerConfig.WriteTimeout != httpWriteTimeout {
		t.Error("WriteTimeout is not set to correct value")
	}
	if httpServerConfig.ReadTimeout != httpReadTimeout {
		t.Error("ReadTimeout is not set to correct value")
	}
	if httpServerConfig.IdleTimeout != httpIdleTimeout {
		t.Error("IdleTimeout is not set to correct value")
	}
}

// TODO: test startHTTPServer() somehow
