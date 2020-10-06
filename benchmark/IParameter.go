package benchmark

type IParameter interface {
	GetValue() string
	SetValue(value string)
}
