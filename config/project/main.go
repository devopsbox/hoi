// Copyright 2016 Atelier Disko. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package project

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/hashicorp/hcl"
)

type Config struct {
	ProjectDirective `hcl:",squash"`
	Domain           map[string]DomainDirective
	Cron             map[string]CronDirective
	Worker           map[string]WorkerDirective
}

// Extracts username/password pairs from domain configuration. Keeps map unique
// and validates data.
func (c Config) GetCreds() (map[string]string, error) {
	creds := make(map[string]string)

	for k, v := range c.Domain {
		if !v.Auth.isEnabled() {
			continue
		}
		if v.Auth.Password == "" {
			return creds, fmt.Errorf("auth user %s given but empty password for domain %s", v.Auth.User, v.FQDN)
		}
		if _, hasKey := creds[k]; hasKey {
			if creds[k] == v.Auth.Password {
				return creds, fmt.Errorf("auth user %s given multiple times but with differing passwords for domain %s", v.Auth.User, v.FQDN)
			}
		}
		creds[v.Auth.User] = v.Auth.Password
	}
	return creds, nil
}

func New() (*Config, error) {
	cfg := &Config{}
	return cfg, nil
}
func NewFromFile(f string) (*Config, error) {
	cfg := &Config{}

	b, err := ioutil.ReadFile(f)
	if err != nil {
		return cfg, err
	}

	cfg, err = decodeInto(cfg, string(b))
	if err != nil {
		return cfg, fmt.Errorf("Failed to parse config file %s: %s", f, err)

	}
	cfg.Path = filepath.Dir(f)
	return cfg, nil
}
func NewFromString(s string) (*Config, error) {
	cfg := &Config{}
	return decodeInto(cfg, s)
}
func decodeInto(cfg *Config, s string) (*Config, error) {
	if err := hcl.Decode(cfg, s); err != nil {
		return cfg, err
	}
	for k, _ := range cfg.Domain {
		e := cfg.Domain[k]
		e.FQDN = k
		cfg.Domain[k] = e
	}
	for k, _ := range cfg.Cron {
		e := cfg.Cron[k]
		e.Name = k
		cfg.Cron[k] = e
	}
	for k, _ := range cfg.Worker {
		e := cfg.Worker[k]
		e.Name = k
		cfg.Worker[k] = e
	}
	for k, _ := range cfg.Database {
		e := cfg.Database[k]
		e.Name = k
		cfg.Database[k] = e
	}
	return cfg, nil
}
