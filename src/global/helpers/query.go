package helpers

import "github.com/lea55/BACKENDCANDITICKET/src/global/core"

func QueryCalc(page uint16) (uint16, uint16) {
	from := page * core.PagLimit
	limit := uint16(core.PagLimit)
	return from, limit
}
