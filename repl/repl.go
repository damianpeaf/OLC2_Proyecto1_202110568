package repl

type ReplContext struct {
	ScopeTrace *ScopeTrace
	// ... error table ...
}

func NewReplContext() *ReplContext {

	return &ReplContext{
		ScopeTrace: NewScopeTrace(),
	}
}
