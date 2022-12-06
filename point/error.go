package point

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// 多一个函数当作构造函数
func New(text string) error {
	return &errorString{text}
}
