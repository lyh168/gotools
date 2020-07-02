package codec

type (
	Unmarshal func([]byte, interface{}) error
	Marshal   func(interface{}) ([]byte, error)
)
