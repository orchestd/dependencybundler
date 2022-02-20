package contextData

import (
	"bitbucket.org/HeilaSystems/session/models"
	"context"
	"time"
)

type ContextDataResolver interface {
	GetVersionsFromContext(c context.Context) (models.Versions, bool, error)
	GetVersionForCollectionFromContext(c context.Context, collectionName string) (string, error)
	GetDateNow(c context.Context) (time.Time, bool, error)
}
