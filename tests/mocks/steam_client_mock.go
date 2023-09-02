package mocks

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockSteamClient struct {
	mock.Mock
}

func (m *MockSteamClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	resp, ok := args.Get(0).(*http.Response)
	if !ok {
		return nil, args.Error(1)
	}
	return resp, args.Error(1)
}

func (m *MockSteamClient) GetAppDetails(id int) ([]byte, error) {
	args := m.Called(id)
	return args.Get(0).([]byte), args.Error(1)
}
