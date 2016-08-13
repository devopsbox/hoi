// Copyright 2016 Atelier Disko. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package project

import (
	"errors"
	"fmt"
	"path/filepath"
)

const (
	// Advices to keep the www prefix. This will not deploy
	// any redirects and just leave the two domains untouched.
	WWW_KEEP = "keep"
	// Advices to drop the www prefix and always redirect
	// to the naked domain.
	WWW_DROP = "drop"
	// Advices to add the www prefix and redirect to the
	// prefixed domain.
	WWW_ADD = "add"
)

type DomainDirective struct {
	// The naked domain name; required.
	FQDN string
	// Configures how the www prefix is handled/normalized; optional; either "keep",
	// "drop" or "add"; defaults to "drop".
	WWW string
	// Optionally configures SSL for this domain; by default not enabled.
	SSL SSLDirective
	// Allows to protect the domain with authentication; optional; by default not enabled.
	Auth AuthDirective
	// A list of domains this domain is also served under; optional; by default empty.
	Aliases []string
	// A list of domains that should redirect to this domain; optional; by default empty.
	Redirects []string
}

const (
	// TODO Allows special value "derived" where the User is taken
	// from the project name. acme_stage -> acme, acme -> acme, acme_shop -> acme
	USER_DERIVED = "!derived"
	// TODO Allows special value "autogenerate" where an unsafe password
	// is generated and mailed to the administrator.
	PASSWORD_AUTOGEN = "!autogenerated"
)

// Auth is considered enabled, once a value for User is given. Empty
// passwords are not allowed.
type AuthDirective struct {
	User     string
	Password string
}

func (drv AuthDirective) IsEnabled() bool {
	return drv.User != ""
}

const (
	// Will use letsencrypt to get a valid cert and renew it automatically.
	CERT_ACME = "!acme"
	// Will generate a self-signed corp cert on the fly.
	CERT_OWNCA = "!own-ca"
	// Will generate a self-signed cert on the fly.
	CERT_SELFSIGNED = "!self-signed"
)

// SSL is considered enabled, once a value for Certificate is given.
type SSLDirective struct {
	// Paths to certificate and certificate key. Paths must be relative to
	// project root i.e. config/ssl/example.org.crt.
	Certificate    string
	CertificateKey string
}

func (drv SSLDirective) IsEnabled() bool {
	return drv.Certificate != ""
}

func (drv SSLDirective) GetCertificate() (string, error) {
	switch drv.Certificate {
	case CERT_ACME:
		return "", errors.New("unimplemented")
	case CERT_OWNCA:
		return "", errors.New("unimplemented")
	case CERT_SELFSIGNED:
		return "", errors.New("unimplemented")
	default:
		if filepath.IsAbs(drv.Certificate) {
			return drv.Certificate, fmt.Errorf("cert has absolute path: %s", drv.Certificate)
		}
		return drv.Certificate, nil
	}
}

func (drv SSLDirective) GetCertificateKey() (string, error) {
	switch drv.CertificateKey {
	case CERT_ACME:
		return "", errors.New("unimplemented")
	case CERT_OWNCA:
		return "", errors.New("unimplemented")
	case CERT_SELFSIGNED:
		return "", errors.New("unimplemented")
	default:
		if filepath.IsAbs(drv.CertificateKey) {
			return drv.CertificateKey, fmt.Errorf("cert key has absolute path: %s", drv.CertificateKey)
		}
		return drv.CertificateKey, nil
	}
}
