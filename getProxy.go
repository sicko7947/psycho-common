package common

import (
	"fmt"
	"math/rand"
	"strings"
)

type ProxyType string

type Proxy struct {
	Host     string    `json:"host"`
	Port     string    `json:"port"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
	NeedAuth bool      `json:"need_auth,omitempty"`
	Protocol ProxyType `json:"protocol,omitempty"`
}

var (
	ProxySocks5 ProxyType = "socks5"
	ProxySocks4 ProxyType = "socks4"
	ProxyHTTPS  ProxyType = "https"
	ProxyHTTP   ProxyType = "http"
)

// GetRawProxy : GetRawProxy
func GetRawProxy(proxylist []string) string {
	if len(proxylist) <= 0 {
		return ""
	}

	return strings.Replace(proxylist[rand.Intn(len(proxylist))], "\r\n", "\n", -1)
}

// GetProxy : GetProxy
func GetProxy(proxylist []string) *Proxy {
	if len(proxylist) <= 0 {
		return &Proxy{}
	}
	p := strings.Replace(proxylist[rand.Intn(len(proxylist))], "\r\n", "\n", -1)

	raw := strings.Split(p, ":")

	proxy := &Proxy{
		Host:     raw[0],
		Port:     raw[1],
		Username: raw[2],
		Password: raw[3],
		NeedAuth: true,
		Protocol: ProxyHTTP,
	}
	return proxy
}

func (p Proxy) String() string {
	if len(p.Host) <= 0 {
		return ""
	}
	if p.NeedAuth {
		return fmt.Sprintf("%v://%v:%v@%v:%v", p.Protocol, p.Username, p.Password, p.Host, p.Port)
	}
	return fmt.Sprintf("%v://%v:%v", p.Protocol, p.Host, p.Port)
}

// GetProxyTypes : GetProxyTypes
func GetProxyTypes() []ProxyType {
	var types = []ProxyType{ProxySocks5, ProxySocks4, ProxyHTTP, ProxyHTTPS}
	return types
}
