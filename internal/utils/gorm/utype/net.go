package utype

import (
	"database/sql/driver"
	"net"
)

type IP net.IP

func (ip *IP) Scan(value interface{}) error {
	if value == nil {
		*ip = nil
		return nil
	}
	addr := net.ParseIP(value.(string))
	*ip = IP(addr)
	return nil
}

func (ip IP) Value() (driver.Value, error) {
	return net.IP(ip).String(), nil
}
