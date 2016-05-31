package igprotocol

import (
	"fmt"
)

type Node struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Market struct {
	bid                      int    `json:"bid"`
	delayTime                int    `json:"delayTime"`
	epic                     string `json:"epic"`
	expiry                   string `json:"expiry"`
	high                     int    `json:"high"`
	instrumentName           string `json:"instrumentName"`
	instrumentType           string `json:"instrumentType"`
	lotSize                  int    `json:"lotSize"`
	low                      int    `json:"low"`
	marketStatus             string `json:"marketStatus"`
	netChange                int    `json:"netChange"`
	offer                    int    `json:"offer"`
	otcTradeable             bool   `json:"otcTradeable"`
	percentageChange         int    `json:"percentageChange"`
	scalingFactor            int    `json:"scalingFactor"`
	streamingPricesAvailable bool   `json:"streamingPricesAvailable"`
	updateTime               string `json:"updateTime"`
	updateTimeUTC            string `json:"updateTimeUTC"`
}

// MarketNavigationResponse ...
type MarketNavigationResponse struct {
	ErrorCode string   `json:"errorCode"`
	Nodes     []Node   `json:"nodes"`
	Markets   []Market `json:"markets"`
}

// MarketNavigationRequest ...
type MarketNavigationRequest struct {
	NodeID string `json:"nodeId"`
}

func (self MarketNavigationRequest) URLPostFix() string {
	if self.NodeID != "" {
		return fmt.Sprintf("/%s", self.NodeID)
	}
	return ""
}

// MarketNavigationResponseResult ...
type MarketNavigationResponseResult struct {
	Header HTTPResponseHeader
	Data   MarketNavigationResponse
}
