package logs

type Encoding uint8

const (
	EncodingEmergency Encoding = iota
	EncodingConsole
	EncodingJson
)

func NewEncoding(encoding Encoding) string {
	var e string

	switch encoding {
	case EncodingJson:
		e = "json"
	default: // default console mode
		e = "console"
	}

	return e
}
