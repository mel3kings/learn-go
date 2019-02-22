package functions

type HTTPCode int

const (
	Ok       HTTPCode = 200
	NotFound HTTPCode = 404
	ErrorCode = iota + 900000 // iota is resolved in the number of constants so it will be 900003
	ErrorCodeTest
	ErrorCodeFour
)

const(
	Number0 = iota + 80000
	Number1 // 800001
	Number2
)

