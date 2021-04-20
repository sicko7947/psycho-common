package psychohttp

import (
	"fmt"
	"net/http"
	"strings"
)

// GetParsedCookie : Handle Cookie From Response
func GetParsedCookie(cookie []*http.Cookie) string {
	cookieLen := len(cookie)
	result := []string{}

	for i := 0; i < cookieLen; i++ {
		cookie := fmt.Sprintf(`%s=%s;`, cookie[i].Name, cookie[i].Value)
		result = append(result, cookie)
	}
	return strings.Join(result[:], "")
}

//解析setCookie
func ParseCookiesFromHeader(cookies []string) map[string]http.Cookie {
	var mapper = map[string]http.Cookie{}
	for _, v := range cookies {
		var ck = ParseSetCookieSingle(v)
		mapper[ck.Name] = ck
	}
	return mapper
}

func ParseSetCookieSingle(str string) http.Cookie {
	var ss = strings.Split(str, ";")[0]
	var kv = strings.Split(ss, "=")
	var ck = http.Cookie{}
	ck.Name = kv[0]
	ck.Value = kv[1]
	return ck
}

func PackCookiesToString(cks map[string]http.Cookie) string {
	var str = ""
	for k, v := range cks {
		str += k + "=" + v.Value + ";"
	}
	return str
}
