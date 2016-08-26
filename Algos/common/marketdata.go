package common

type Instrument interface {
	LongName() string
	Symbol(exchange string) string
}

// MarketData ...

type PriceAndVolume interface {
	Price() float64
	Volume() int
}

type TopLinePriceAndVolume interface {
	PriceAndVolume
	Ordercount() int
}

type TopLine interface {
	BestBid() TopLinePriceAndVolume
	BestOffer() TopLinePriceAndVolume
}

type BidOfferDepthPriceAndVolume interface {
	TopLinePriceAndVolume
	CummVolume() int
	CummVWAP() float64
}

type BidOfferDepth interface {
	Bid(index int) TopLinePriceAndVolume
	Offer(index int) TopLinePriceAndVolume
}

type MarketData interface {
	Instrument
	TopLine
	BidOfferDepth
	LastTrade() PriceAndVolume
}

// MarketOrder ...
type MarketOrder interface {
}
