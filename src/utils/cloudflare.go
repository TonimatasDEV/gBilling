package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var cloudflareIPs []net.IPNet
var proxyEnabled bool

func InitCloudflare() {
	proxyEnabled, _ = strconv.ParseBool(os.Getenv("CLOUDFLARE_PROXY"))

	if proxyEnabled {
		readIpRanges(4)
		readIpRanges(6)
		log.Println("Cloudflare ips loaded.")
	}
}

func readIpRanges(rangeType int) {
	url := fmt.Sprintf("https://www.cloudflare.com/ips-v%v", rangeType)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error getting cloudflare ips:", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing cloudflare ips body:", err)
		}
	}(resp.Body)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()

		_, i, err := net.ParseCIDR(line)

		if err != nil {
			log.Fatalln("Error parsing", line+":", err)
		}

		cloudflareIPs = append(cloudflareIPs, *i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error reading cloudflare ips:", err)
	}
}

func IsCloudflareIP(ip string) bool {
	for _, ipNet := range cloudflareIPs {
		if ipNet.Contains(net.ParseIP(ip)) {
			return true
		}
	}
	return false
}
