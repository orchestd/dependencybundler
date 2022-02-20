package session

import (
	"bitbucket.org/HeilaSystems/dependencybundler/depBundler"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/contextData"
)

func NewContextData() contextData.ContextDataResolver {
	return depBundler.ContextData{}
}
