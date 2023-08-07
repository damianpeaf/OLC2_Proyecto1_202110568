package repl

type Param struct {
	ExternName      string
	InnerName       string
	Type            string
	PassByReference bool
}

// 3 types of parameters
// 1. extern and inner name declared 			[args are specified with the extern name]
// 2. extern name = "_", inner name declared 	[args are specified with the order]
// 3. extern name = "", inner name declared	[args are specified with the inner name]
