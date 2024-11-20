package utils

import (
	"errors"
	"net"
	"os"
	"strings"
)

func GetEnv() string {
	env := strings.ToLower(os.Getenv("ENV"))
	if env == "" {
		env = "dev"
		os.Setenv("ENV", env)
	}

	return env
}

func IsDevEnv() bool {
	env := GetEnv()
	return env == "dev" || env == "local"
}

func GetLocalIP() (string, error) {
	ip := strings.ToLower(os.Getenv("POD_IP"))
	if ip != "" {
		return ip, nil
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}

	return "", errors.New("no local ip found")
}
