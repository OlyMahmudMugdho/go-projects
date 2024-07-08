package hello

type Hello struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func NewHello(ok bool, message string) *Hello {
	return &Hello{
		Ok:      ok,
		Message: message,
	}
}
