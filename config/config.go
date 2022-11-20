package config

import (
	"fmt"
	"os"
)

var (
	host    string = "0.0.0.0"
	port    string = "8080"
	baseURL string = "http://localhost"
	user    string = "test-user"
	keyPem  string = ""
)

func Host() string { return host }
func Port() string { return port }
func BindAddress() string {
	return fmt.Sprintf("%s:%s", host, port)
}

func BaseURL() string { return baseURL }
func User() string    { return user }
func KeyPem() string  { return keyPem }

func SetHost(newHost string) error   { host = newHost; return nil }
func SetPort(newPort string) error   { port = newPort; return nil }
func SetBaseURL(newUrl string) error { baseURL = newUrl; return nil }
func SetUser(newUser string) error   { user = newUser; return nil }
func SetKeyPem(pemPath string) error {
	keyPemRaw, err := os.ReadFile(pemPath)
	if err != nil {
		return err
	}
	keyPem = string(keyPemRaw)
	return nil
}

type OptionFn func(string) error

var mapping = map[string]OptionFn{
	"PUBD_BIND_HOST": SetHost,
	"PUBD_BIND_PORT": SetPort,
	"PUBD_BASE_URL":  SetBaseURL,
	"PUBD_USER":      SetUser,
	"PUBD_KEY_PATH":  SetKeyPem,
}

func ParseEnv() map[string]error {
	errors := make(map[string]error)
	for key, fn := range mapping {
		if val := os.Getenv(key); val != "" {
			if err := fn(val); err != nil {
				errors[key] = err
			}
		}
	}
	return errors
}
