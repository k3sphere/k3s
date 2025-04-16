//go:build linux
// +build linux

package flannel

import (
	"net"

	"github.com/pkg/errors"
	"github.com/vishvananda/netlink"
)

func GetGatewayIP(iface *net.Interface) (net.IP, error) {
    routes, err := netlink.RouteList(nil, netlink.FAMILY_ALL)
    if err != nil {
        return nil, errors.WithMessage(err, "failed to list routes")
    }
    for _, route := range routes {
        if route.LinkIndex == iface.Index && route.Gw != nil {
            return route.Gw, nil
        }
    }
    return nil, nil
}