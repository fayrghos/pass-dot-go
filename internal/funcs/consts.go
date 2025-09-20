package funcs

const (
	MinPassLen = 8
	MaxPassLen = 128
)

const (
	CharsLetter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsNum    = "0123456789"
	CharsSymbol = "@$&!#%&*"
)

const (
	PassLetter = iota
	PassNum
	PassSymbol
)
