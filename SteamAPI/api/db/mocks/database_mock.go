package mocks

import "errors"

type MockDatabase struct {
	Connected  bool
	ShouldFail bool // Nuevo campo para indicar si debe simular un error
}

func (m *MockDatabase) Connect() error {
	// Simula una conexión exitosa a menos que ShouldFail sea verdadero
	if m.ShouldFail {
		return errors.New("error de conexión simulado")
	}

	m.Connected = true
	return nil
}

func (m *MockDatabase) Close() error {
	// Simular el cierre de la conexión
	m.Connected = false
	return nil
}
