package session

import (
	"github.com/orchestd/dependencybundler/interfaces/session"
	"github.com/orchestd/session/sessionmock"
)

type ActiveOrder sessionmock.ActiveOrder

func NewSessionMock(sessionToken string, customerId string, activeOrderId string, fakeNow *string, cacheVersions map[string]string, order ActiveOrder) (session.SessionResolver, error) {
	return sessionmock.NewSessionMockWrapper(sessionToken, customerId, activeOrderId, fakeNow, order.StoreId, order.TimeTo, order.Tags), nil
}
