package session

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/session"
	"bitbucket.org/HeilaSystems/session/sessionmock"
)

type ActiveOrder sessionmock.ActiveOrder
func NewSessionMock(sessionToken string , customerId string, activeOrderId string , fakeNow *string, cacheVersions map[string]string,  order ActiveOrder) (session.SessionResolver,error){
	return  sessionmock.NewSessionMockWrapper(sessionToken,customerId,activeOrderId,fakeNow,order.StoreId,order.TimeTo,order.Tags) , nil
}

