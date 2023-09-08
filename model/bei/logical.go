// Package bei
// @author tabuyos
// @since 2023/9/8
// @description bei
package bei

type Logical struct {
	// enable logical deleted
	enable bool
	// deleted key
	key string
	// done deleted values
	ddval string
	// undone deleted values
	udval string
	// current deleted values
	cdval string
}

func OfLogical() *Logical {
	return &Logical{
		enable: true,
		key:    "deleted",
		ddval:  "1",
		udval:  "0",
		cdval:  "0",
	}
}

func (l *Logical) CurrentDelVal(val string) *Logical {
	var logical = l
	if logical == nil {
		logical = OfLogical()
	}
	logical.cdval = val
	return logical
}

func (l *Logical) Enable() *Logical {
	var logical = l
	if logical == nil {
		logical = OfLogical()
	}
	logical.enable = true
	return logical
}

func (l *Logical) Disable() *Logical {
	var logical = l
	if logical == nil {
		logical = OfLogical()
	}
	logical.enable = false
	return logical
}
