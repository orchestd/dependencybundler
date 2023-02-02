package contextData

import (
	"context"
	"github.com/orchestd/session/models"
	"time"
)

type ContextDataResolver interface {
	GetVersionsFromContext(c context.Context) (models.Versions, bool, error)
	GetVersionForCollectionFromContext(c context.Context, collectionName string) (string, error)
	GetDateNow(c context.Context) (time.Time, bool, error)
}
