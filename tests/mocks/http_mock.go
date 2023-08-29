package mocks

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	resp, ok := args.Get(0).(*http.Response)
	if !ok {
		return nil, args.Error(1)
	}
	return resp, args.Error(1)
}
