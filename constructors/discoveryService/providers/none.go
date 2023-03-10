package providers

import "github.com/orchestd/servicereply"

type noDSP struct {
}

func NewNoDSP() noDSP {
	return noDSP{}
}

func (dsp noDSP) Register() (sr servicereply.ServiceReply) {
	return servicereply.NewNil()
}
func (dsp noDSP) GetAddress(serviceName string) (sr servicereply.ServiceReply) {
	return servicereply.NewNil().WithReplyValues(servicereply.ValuesMap{address: serviceName})
}
