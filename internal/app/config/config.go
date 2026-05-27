package config

import (
	"net"
	"os"
	"strings"
)

type Config struct {
	ListenAddr string
	Port       string
}

func Load() Config {
	addr := os.Getenv("ADDR")
	if addr != "" {
		if strings.Trim(addr, "0123456789") == "" {
			addr = ":" + addr
		}
		_, port, err := net.SplitHostPort(addr)
		if err != nil {
			port = ""
		}
		return Config{ListenAddr: addr, Port: port}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9194"
	}
	return Config{ListenAddr: ":" + port, Port: port}
}

func (c Config) Addr() string {
	return c.ListenAddr
}
