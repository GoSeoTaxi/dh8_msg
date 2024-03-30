package client

import (
	"net/url"
	"time"

	"github.com/go-ping/ping"
)

type PingResult struct {
	IPAddr                string
	PacketsRecv           int
	PacketsSent           int
	PacketsRecvDuplicates int
	PacketLoss            float64
	Addr                  string
	Rtts                  []time.Duration
	MinRtt                time.Duration
	MaxRtt                time.Duration
	AvgRtt                time.Duration
	StdDevRtt             time.Duration
}

func PingUrlWithTimeOutAndCount(urlS string, count int, maxTimeOut time.Duration) (*PingResult, error) {

	parsedURL, err := url.Parse(urlS)
	if err != nil {
		return nil, err
	}

	pinger, err := ping.NewPinger(parsedURL.Hostname())
	if err != nil {
		return nil, err
	}
	pinger.Timeout = maxTimeOut // на все пинги, не зависимо от количества
	pinger.Count = count
	err = pinger.Run()
	if err != nil {
		return nil, err
	}
	stats := pinger.Statistics()

	result := &PingResult{
		IPAddr:                stats.IPAddr.String(),
		PacketsRecv:           stats.PacketsRecv,
		PacketsSent:           stats.PacketsSent,
		PacketsRecvDuplicates: stats.PacketsRecvDuplicates,
		PacketLoss:            stats.PacketLoss,
		Addr:                  stats.Addr,
		Rtts:                  stats.Rtts,
		MinRtt:                stats.MinRtt,
		MaxRtt:                stats.MaxRtt,
		AvgRtt:                stats.AvgRtt,
		StdDevRtt:             stats.StdDevRtt,
	}

	return result, nil
}
