package igprotocol

import (
	"fmt"
)

// Node ...
type Node struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Market ...
type Market struct {
	Bid                      int    `json:"bid"`
	DelayTime                int    `json:"delayTime"`
	Epic                     string `json:"epic"`
	Expiry                   string `json:"expiry"`
	High                     int    `json:"high"`
	InstrumentName           string `json:"instrumentName"`
	InstrumentType           string `json:"instrumentType"`
	LotSize                  int    `json:"lotSize"`
	Low                      int    `json:"low"`
	MarketStatus             string `json:"marketStatus"`
	NetChange                int    `json:"netChange"`
	Offer                    int    `json:"offer"`
	OtcTradeable             bool   `json:"otcTradeable"`
	PercentageChange         int    `json:"percentageChange"`
	ScalingFactor            int    `json:"scalingFactor"`
	StreamingPricesAvailable bool   `json:"streamingPricesAvailable"`
	UpdateTime               string `json:"updateTime"`
	UpdateTimeUTC            string `json:"updateTimeUTC"`
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

// URLPostFix ...
func (request MarketNavigationRequest) URLPostFix() string {
	if request.NodeID != "" {
		return fmt.Sprintf("/%s", request.NodeID)
	}
	return ""
}

// MarketNavigationResponseResult ...
type MarketNavigationResponseResult struct {
	Header HTTPResponseHeader
	Data   MarketNavigationResponse
}
