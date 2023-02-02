package session

import (
	"github.com/orchestd/dependencybundler/depBundler"
	"github.com/orchestd/dependencybundler/interfaces/contextData"
)

func NewContextData() contextData.ContextDataResolver {
	return depBundler.ContextData{}
}
