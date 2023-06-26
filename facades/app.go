package facades

import (
	foundationcontract "github.com/chenyuIT/framework2/contracts/foundation"
	"github.com/chenyuIT/framework2/foundation"
)

func App() foundationcontract.Application {
	return foundation.App
}
