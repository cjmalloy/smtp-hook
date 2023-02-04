package main

import (
	"flag"
	"os"
	"strconv"
)

var (
	flagServerName     = flag.String("name", getEnvString("NAME", "smtp2http"), "the server name")
	flagListenAddr     = flag.String("listen", getEnvString("LISTEN", ":smtp"), "the smtp address to listen on")
	flagWebhook        = flag.String("webhook", getEnvString("WEBHOOK", "http://localhost:8080/my/webhook"), "the webhook to send the data to")
	flagMaxMessageSize = flag.Int64("msglimit", getEnvInt64("MSGLIMIT", 1024*1024*2), "maximum incoming message size")
	flagReadTimeout    = flag.Int("timeout.read", getEnvInt("TIMEOUT_READ", 5), "the read timeout in seconds")
	flagWriteTimeout   = flag.Int("timeout.write", getEnvInt("TIMEOUT_WRITE", 5), "the write timeout in seconds")
	flagAuthUSER       = flag.String("user", getEnvString("USER", ""), "user for smtp client")
	flagAuthPASS       = flag.String("pass", getEnvString("PASS", ""), "pass for smtp client")
	flagDomain         = flag.String("domain", getEnvString("DOMAIN", ""), "domain for receiving mails")
)

func init() {
	flag.Parse()
}

func getEnvString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt64(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, _ := strconv.ParseInt(value, 10, 64)
		return i
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, _ := strconv.Atoi(value)
		return i
	}
	return fallback
}
