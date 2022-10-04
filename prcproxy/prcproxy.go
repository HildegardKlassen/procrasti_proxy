package prcproxy

import (
	"time"
)

type PrcProxy struct {
	Port               uint
	BlockRequests      bool
	ProxyMakesRequests bool
	BlockList          []string
	OfficeTimeSlot     time.Duration
}

func NewPrcProxy(port uint, blockRequests bool, proxyMakesRequests bool, blockList []string, officeTimeSlot time.Duration) *PrcProxy {
	return &PrcProxy{
		Port:               port,
		BlockRequests:      blockRequests,
		ProxyMakesRequests: proxyMakesRequests,
		BlockList:          blockList,
		OfficeTimeSlot:     officeTimeSlot,
	}

}
