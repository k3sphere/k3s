//go:build !linux
// +build !linux

package flannel

import (
	"errors"
	"net"
)

func GetGatewayIP(iface *net.Interface) (net.IP, error) {
    // Placeholder for non-Linux OS logic
    return nil, errors.New("gateway IP retrieval is not implemented for this OS")
}