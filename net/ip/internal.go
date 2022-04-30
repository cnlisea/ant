package ip

import (
	"net"
	"strings"
)

// InternalIP get internal ip
func InternalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for i := range inters {
		if inters[i].Flags&net.FlagUp == net.FlagUp && !strings.HasPrefix(inters[i].Name, "lo") {
			addrs, err := inters[i].Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}
