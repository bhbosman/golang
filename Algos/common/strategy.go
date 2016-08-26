package common

import "github.com/bhbosman/golang/Algos/common"

type StrategyContext interface {
	MarketData(instrument common.Instrument) common.MarketData
}

type Strategy interface {
	Execute(context StrategyContext)
}
