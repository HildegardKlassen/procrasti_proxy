package prcproxy

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	startBlockTime = "9:00"
	endBlockTime   = "17:00"
)

type PrcProxy struct {
	Port             string
	BlockAllRequests bool
	BlockList        []string
	OfficeTimeSlot   time.Duration
	BlockStartTime   time.Time
	BlockEndTime     time.Time
}

func NewPrcProxy(port string, blockAllRequests bool, blockList []string, blockStartTime time.Time, blockEndTime time.Time) *PrcProxy {
	return &PrcProxy{
		Port:             port,
		BlockAllRequests: blockAllRequests,
		BlockList:        blockList,
		BlockStartTime:   blockStartTime,
		BlockEndTime:     blockEndTime,
	}
}

func Run() error {
	blockAll := flag.Bool("blockall", false, "Set to block all requests always.")
	blockList := flag.String("blocksites", "", "Sites to block seperated by comma. Exampe: --blocksites google.com,github.com")
	port := flag.String("port", "1994", "Port to listen on. Example: --port 1994")
	startBlockTime := flag.String("startblocktime", "09:00", "Time the blocking requests window is aktive. Example: --startblocktime 09:00")
	endBlockTime := flag.String("endblocktime", "17:00", "Time the blocking requests window is aktive. Example: --endblocktime 17:00")

	list, err := parseBlockList(*blockList)
	if err != nil {
		log.Fatal(err)
	}

	st, err := parseTime(*startBlockTime)
	if err != nil {
		return err
	}
	et, err := parseTime(*endBlockTime)
	if err != nil {
		return err
	}

	proxy := NewPrcProxy(*port, *blockAll, list, st, et)
	router := proxy.CreateRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", proxy.Port), router))
	return nil
}

func parseBlockList(blocklist string) ([]string, error) {
	blockList := strings.Split(blocklist, ",")

	if len(blockList) > 0 {
		return blockList, nil
	}
	return nil, fmt.Errorf("could not get valid host to block")
}

func parseTime(t string) (time.Time, error) {
	stringToParse := t + ":00"
	pt, err := time.Parse("15:04:05", stringToParse)
	if err != nil {
		return time.Time{}, err
	}
	return pt, nil
}
