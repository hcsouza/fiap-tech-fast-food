package mocks

import "github.com/stretchr/testify/mock"

type MockDatabaseAdapter struct {
	mock.Mock
}

func (m *MockDatabaseAdapter) FindOne(id string) (interface{}, error) {
	mockedFunctionParams := m.Called(id)

	var mockedFirstReturn interface{}
	if mockedFunction, ok := mockedFunctionParams.Get(0).(func(string) interface{}); ok {
		mockedFirstReturn = mockedFunction(id)
	} else {
		mockedFirstReturn = mockedFunctionParams.Get(0)
	}

	var mockedSecondReturn error
	if mockedFunction, ok := mockedFunctionParams.Get(1).(func(string) error); ok {
		mockedSecondReturn = mockedFunction(id)
	} else {
		mockedSecondReturn = mockedFunctionParams.Error(1)
	}

	return mockedFirstReturn, mockedSecondReturn
}

func (m *MockDatabaseAdapter) Save(id string, data interface{}) error {
	mockedFunctionParams := m.Called(id, data)

	var mockedReturn error
	if mockedFunction, ok := mockedFunctionParams.Get(0).(func(string, interface{}) error); ok {
		mockedReturn = mockedFunction(id, data)
	} else {
		mockedReturn = mockedFunctionParams.Error(0)
	}

	return mockedReturn
}
