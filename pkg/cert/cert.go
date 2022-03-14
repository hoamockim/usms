package cert

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"time"
	"usms/pkg/configs"
)

type CA struct {
	*x509.Certificate
	Expired   bool
	CARevoked bool
	key       *ecdsa.PrivateKey
	parent    *CA
}

type Cert struct {
	*x509.Certificate
	Key       *ecdsa.PrivateKey
	Rawkey    []byte
	IsServer  bool
	IsCA      bool
	Expired   bool
	CARevoked bool

	Additional []byte
}

func cksum(pk *ecdsa.PublicKey) []byte {
	pm := elliptic.Marshal(pk.Curve, pk.X, pk.Y)
	return hash(pm)
}

func hash(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	return h.Sum(nil)
}

func randSerial() *big.Int {
	min := big.NewInt(1)
	min.Lsh(min, 120)

	max := big.NewInt(1)
	max.Lsh(max, 130)

	for {
		serial, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(fmt.Errorf("ca: can't generate serial#: %w", err))
		}

		if serial.Cmp(min) > 0 {
			return serial
		}
	}
	panic("can't gen new CA serial")
}

func createCA(config configs.Cert) (*Cert, error) {
	eckey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("ca: can't generate ECC P256 key: %s", err)
	}
	pubkey := eckey.Public().(*ecdsa.PublicKey)
	skid := cksum(pubkey)
	serial := randSerial()

	subject := pkix.Name{
		CommonName:         config.CommonName,
		Country:            []string{config.Country},
		Province:           []string{config.Province},
		Locality:           []string{config.Locality},
		Organization:       []string{config.Organization},
		OrganizationalUnit: []string{config.OrganizationalUnit},
	}
	now := time.Now().UTC()
	template := x509.Certificate{
		SignatureAlgorithm:    x509.ECDSAWithSHA512,
		PublicKeyAlgorithm:    x509.ECDSA,
		SerialNumber:          serial,
		Issuer:                subject,
		Subject:               subject,
		NotBefore:             now,
		NotAfter:              now.Add(config.Validity),
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            -1,
		SubjectKeyId:          skid,
		AuthorityKeyId:        skid,

		KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}
	der, err := x509.CreateCertificate(rand.Reader, &template, &template, pubkey, eckey)
	if err != nil {
		return nil, fmt.Errorf("ca: can't sign root CA cert: %s", err)
	}
	crt, err := x509.ParseCertificate(der)
	if err != nil {
		panic(err)
	}

	ck := &Cert{
		Certificate: crt,
		Key:         eckey,
		IsCA:        true,
	}

	return ck, nil
}

func saveCA(cert *Cert) error {
	return nil
}
