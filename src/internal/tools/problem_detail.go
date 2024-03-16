package tools

type ProblemDetail struct {
	Field  string `json:"field"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Value  any    `json:"value"`
	Type   string `json:"type"`
}
