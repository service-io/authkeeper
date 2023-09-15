// Package netx
// @author tabuyos
// @since 2023/8/9
// @description netx
package netx

import (
	"deepsea/config"
	"deepsea/helper/recorderx"
	"net"
)

// GetPrivateIp 获取本机 IP
func GetPrivateIp() string {
	recorder := recorderx.DefaultRecorder()
	ip := "127.0.0.1"
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		recorder.Errorf("无法获取网络地址接口信息: %v\n", err)
		return ip
	}

	for _, address := range interfaceAddr {
		addr, ok := address.(*net.IPNet)
		// 检查ip地址判断是否回环地址
		if ok && !addr.IP.IsLoopback() && addr.IP.IsPrivate() {
			if addr.IP.To4() != nil {
				return addr.IP.String()
			}
		}
	}
	return ip
}

// JoinHostPort 拼接主机和端口
func JoinHostPort(host, port string) string {
	return net.JoinHostPort(host, port)
}

// JoinPort 拼接本地主机和端口
func JoinPort(port string) string {
	return net.JoinHostPort(GetPrivateIp(), port)
}

// GetAddress 从配置获取地址
func GetAddress() string {
	serverConfig := config.TomlConfig().Server

	address := serverConfig.Address

	if address != "" {
		return address
	}

	return net.JoinHostPort(serverConfig.IP, serverConfig.Port)
}

// ParseAddress 解析主机和端口
func ParseAddress(address string) (string, string) {
	if address == "" {
		serverConfig := config.TomlConfig().Server
		return serverConfig.IP, serverConfig.Port
	}

	host, port, err := net.SplitHostPort(address)

	if err != nil {
		panic(err)
	}

	return host, port
}
