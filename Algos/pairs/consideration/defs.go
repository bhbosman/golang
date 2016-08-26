package consideration

import (
	"github.com/bhbosman/golang/Algos/common"
)



type Consideration interface {
	BuyInstrument() common.Instrument
	SellInstrument() common.Instrument
	Ratio() float64
	BuyConsideration() float64
	SellConsideration() float64
}
