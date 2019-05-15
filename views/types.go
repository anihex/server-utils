package views

import (
	"github.com/anihex/server-utils/tools"
)

type ctxID int

const toSkip ctxID = 0
const toTime tools.TimeType = 1
const gotError ctxID = 2
