package x509

import (
	"crypto/x509"
	"reflect"
)

var IsEncryptedPEMBlock = reflect.ValueOf(x509.IsEncryptedPEMBlock)
var DecryptPEMBlock = reflect.ValueOf(x509.DecryptPEMBlock)
var EncryptPEMBlock = reflect.ValueOf(x509.EncryptPEMBlock)
var SystemCertPool = reflect.ValueOf(x509.SystemCertPool)
var NewCertPool = reflect.ValueOf(x509.NewCertPool)
var ParsePKCS8PrivateKey = reflect.ValueOf(x509.ParsePKCS8PrivateKey)
var ParseCertificates = reflect.ValueOf(x509.ParseCertificates)
var ParseDERCRL = reflect.ValueOf(x509.ParseDERCRL)
var ParseCRL = reflect.ValueOf(x509.ParseCRL)
var CreateCertificateRequest = reflect.ValueOf(x509.CreateCertificateRequest)
var CreateCertificate = reflect.ValueOf(x509.CreateCertificate)
var MarshalPKIXPublicKey = reflect.ValueOf(x509.MarshalPKIXPublicKey)
var ParsePKIXPublicKey = reflect.ValueOf(x509.ParsePKIXPublicKey)
var ParseCertificate = reflect.ValueOf(x509.ParseCertificate)
var ParseCertificateRequest = reflect.ValueOf(x509.ParseCertificateRequest)
var ParseECPrivateKey = reflect.ValueOf(x509.ParseECPrivateKey)
var MarshalECPrivateKey = reflect.ValueOf(x509.MarshalECPrivateKey)
var ParsePKCS1PrivateKey = reflect.ValueOf(x509.ParsePKCS1PrivateKey)
var MarshalPKCS1PrivateKey = reflect.ValueOf(x509.MarshalPKCS1PrivateKey)
