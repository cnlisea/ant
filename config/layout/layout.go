package layout

type Layout interface {
	ExtName() string
	Parse(data []byte, obj interface{}) error
}
