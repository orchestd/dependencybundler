package providers

import (
	"bitbucket.org/HeilaSystems/servicereply"
	"os"
)

const urlKeyword = "Url"

type confDSP struct {
}

func NewConfDSP() confDSP {
	return confDSP{}
}

func (dsp confDSP) Register() (sr servicereply.ServiceReply) {
	return servicereply.NewNil()
}

func (dsp confDSP) GetAddress(serviceName string) (sr servicereply.ServiceReply) {
	if overrideHost := os.Getenv(serviceName + urlKeyword); len(overrideHost) > 0 {
		serviceName = overrideHost
	}
	return servicereply.NewNil().WithReplyValues(servicereply.ValuesMap{address: serviceName})
}
