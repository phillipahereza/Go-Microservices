package dbclient

import (
	"github.com/phillipahereza/go_microservices/model"
	"github.com/stretchr/testify/mock"
)

type MockBoltClient struct {
	mock.Mock
}

func (mc *MockBoltClient) OpenBoltDb() {

}

func (mc *MockBoltClient) Seed() {

}

func (mc *MockBoltClient) QueryAccount(accountId string) (model.Account, error) {
	args := mc.Mock.Called(accountId)
	return args.Get(0).(model.Account), args.Error(1)
}


