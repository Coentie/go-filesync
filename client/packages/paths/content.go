package paths

type CONTENT struct {
	Paths []string `json:"paths"`
}

func NewContent() CONTENT {
	return CONTENT{
		Paths: make([]string, 0),
	}
}
