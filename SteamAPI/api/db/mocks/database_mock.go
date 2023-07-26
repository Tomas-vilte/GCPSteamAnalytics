package mocks

import "errors"

type MockDatabase struct {
	Connected  bool
	ShouldFail bool // Nuevo campo para indicar si debe simular un error
}

func (m *MockDatabase) Connect() error {
	// Simulate a successful connection unless ShouldFail is true
	if m.ShouldFail {
		return errors.New("error de conexión simulado")
	}

	m.Connected = true
	return nil
}

func (m *MockDatabase) Close() error {
	// Simulate closing the connection
	m.Connected = false
	return nil
}
