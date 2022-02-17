package contextData

import (
	"bitbucket.org/HeilaSystems/session/models"
	"context"
	"time"
)

type ContextDataResolver interface {
	GetVersions(c context.Context) (models.Versions, bool, error)
	GetVersionForCollection(c context.Context, collectionName string) (string, error)
	GetDateNow(c context.Context) (time.Time, error)
}
