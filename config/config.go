package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Holder struct {
	DbConnectionStr string
	DbName string
	RsaPrivateKeyPath string
	CookieDomainName string
	SessionTime time.Duration
	DefaultRoles []string
}

var Config = LoadConfig()

func LoadConfig() *Holder {
	return &Holder {
		DbConnectionStr: configStrReq("DB_CONNECTION_STRING"),
		DbName: configStrReq("DB_NAME"),
		RsaPrivateKeyPath: configStrReq("RSA_PRIVATE_KEY"),
		CookieDomainName: configStrOpt("COOKIE_DOMAIN_NAME","hoppercloud.net"),
		SessionTime: configDurationOpt("SESSION_TIME", 4 * time.Hour),
		DefaultRoles: configStrAryReq("DEFAULT_ROLES"),
	}
}

func configStrOpt(envName string, defaultValue string) string {
	val := os.Getenv(envName)
	if val == "" {
		return defaultValue
	}
	return val
}

func configStrReq(envName string) string {
	val := os.Getenv(envName)
	if val == "" {
		log.Fatalf("Required env %s not specified", envName)
	}
	return val
}

func configStrAryReq(envName string) []string {
	val := os.Getenv(envName)
	if val == "" {
		log.Fatalf("Required env %s not specified", envName)
	}

	stringAry := strings.Split(val, ",")
	for i, v := range stringAry {
		stringAry[i] = strings.TrimSpace(v)
	}

	return stringAry
}

func configIntOpt(envName string, defaultValue int) int {
	val := os.Getenv(envName)
	if val == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Cannot convert %s to int in env %s", val, envName)
	}
	return i
}

func configDurationOpt(envName string, defaultValue time.Duration) time.Duration {
	val := os.Getenv(envName)
	if val == "" {
		return defaultValue
	}
	i, err := time.ParseDuration(val)
	if err != nil {
		log.Fatalf("Cannot parse %s as duration in env %s", val, envName)
	}
	return i
}
