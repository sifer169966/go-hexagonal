package validators

type Validators interface {
	ValidateSomeField(inf interface{}) error
}
type validators struct{}

func New() Validators {
	return &validators{}
}

func (vld *validators) ValidateSomeField(inf interface{}) error {
	return nil
}
