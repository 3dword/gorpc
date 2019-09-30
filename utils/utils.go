package utils

import (
	"github.com/diubrother/gorpc/codes"
	"strings"
)

// 解析 target 地址，例如 ip://127.0.0.1:6379
func ParseAddress(target string) (string, string, error){
	if target == "" {
		return "","",codes.ConfigError
	}
	strs := strings.Split(target, "://")
	if len(strs) <= 1 {
		return "","",codes.ConfigError
	}
	ipAndPort := strings.Split(strs[1],":")
	if len(ipAndPort) <= 1 {
		return "","",codes.ConfigError
	}
	return ipAndPort[0], ipAndPort[1],nil
}