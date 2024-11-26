package net

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
)

func GetLocalIp() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ip = ipnet.IP.String()
			return
		}
	}

	return
}

func ToUrlQuery(value interface{}) string {
	b, _ := json.Marshal(value)
	m := make(map[string]interface{})
	if e := json.Unmarshal(b, &m); e != nil {
		log.Println(e)
	}

	querySlice := make([]string, 0)
	for k, v := range m {
		querySlice = append(querySlice, fmt.Sprintf("%s=%v", k, v))
	}

	return strings.Join(querySlice, "&")
}
