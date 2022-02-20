package depBundler

import (
	"bitbucket.org/HeilaSystems/session/models"
	"bitbucket.org/HeilaSystems/session/sessionresolver"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

var DataVersionsKey = sessionresolver.DataVersionsKey
var DataNowKey = sessionresolver.DataNowKey

type ContextData struct {
}

func (cd ContextData) GetVersionsFromContext(c context.Context) (models.Versions, bool, error) {
	v := c.Value(DataVersionsKey)
	if v == nil {
		return nil, false, nil
	}

	versions := make(models.Versions)
	err := json.Unmarshal([]byte(v.(string)), &versions)
	if err != nil {
		return nil, false, err
	}

	return versions, true, nil
}

func (cd ContextData) GetVersionForCollectionFromContext(c context.Context, collectionName string) (string, error) {
	versions, _, err := cd.GetVersionsFromContext(c)
	if err != nil {
		return "", err
	}
	ver, ok := versions[collectionName]
	if !ok {
		return "", fmt.Errorf("latest version for collection %v not  found", collectionName)
	}
	return ver, nil
}

func (cd ContextData) GetDateNow(c context.Context) (time.Time, bool, error) {
	t := time.Time{}
	v := c.Value(DataNowKey)
	if v == nil {
		return t, false, nil
	}
	err := json.Unmarshal([]byte(v.(string)), &t)
	if err != nil {
		return t, false, err
	}

	return t, true, nil
}
