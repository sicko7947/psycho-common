package psyhttp

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func GetProxyList(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	rawData, _ := ioutil.ReadAll(file)
	proxyList := strings.Split(strings.Replace(string(rawData), "\r\n", "\n", -1), "\n")
	return proxyList
}

func GetProxy(p string) *Proxy {
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
	if p.NeedAuth {
		return fmt.Sprintf("%v://%v:%v@%v:%v", p.Protocol, p.Username, p.Password, p.Host, p.Port)
	}
	return fmt.Sprintf("%v://%v:%v", p.Protocol, p.Host, p.Port)
}

func GetProxyTypes() []ProxyType {
	var types = []ProxyType{ProxySocks5, ProxySocks4, ProxyHTTP, ProxyHTTPS}
	return types
}
