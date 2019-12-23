package common

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func GetCerts(certFile, keyFile string) (keyPair *tls.Certificate, certPool *x509.CertPool) {
	certData, err := ioutil.ReadFile(certFile)
	if err != nil {
		panic(err)
	}
	keyData, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic(err)
	}

	pair, err := tls.X509KeyPair(certData, keyData)
	keyPair = &pair

	certPool = x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(certData)
	if !ok {
		panic(err)
	}
	return
}
