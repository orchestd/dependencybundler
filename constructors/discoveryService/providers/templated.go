package providers

import (
	"context"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/log"
	"github.com/orchestd/servicereply"
	"os"
	"strings"
)

type templatedDSP struct {
	conf     configuration.Config
	template string
}

const serviceKeyword = "{servicename}"

func NewTemplatedDSP(conf configuration.Config, lg log.Logger) templatedDSP {

	t := templatedDSP{}

	if confTemplate, err := conf.Get("dspTemplate").String(); err != nil {
		t.template = "http://" + serviceKeyword
		lg.Info(context.Background(), "dspTemplate missing from configuration, defaulting to:dspTemplate="+t.template)
	} else {
		t.template = confTemplate
	}

	return t
}

func (dsp templatedDSP) Register() (sr servicereply.ServiceReply) {
	return servicereply.NewNil()
}

func (dsp templatedDSP) GetAddress(serviceName string) (sr servicereply.ServiceReply) {
	if overrideHost := os.Getenv(serviceName + urlKeyword); len(overrideHost) > 0 {
		return servicereply.NewNil().WithReplyValues(servicereply.ValuesMap{address: overrideHost})
	}

	host := strings.ReplaceAll(dsp.template, serviceKeyword, serviceName)
	return servicereply.NewNil().WithReplyValues(servicereply.ValuesMap{address: host})
}
