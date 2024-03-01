// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	domain "github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockProductRepository is an autogenerated mock type for the ProductRepository type
type MockProductRepository struct {
	mock.Mock
}

type MockProductRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProductRepository) EXPECT() *MockProductRepository_Expecter {
	return &MockProductRepository_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: id
func (_m *MockProductRepository) Delete(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProductRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockProductRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id string
func (_e *MockProductRepository_Expecter) Delete(id interface{}) *MockProductRepository_Delete_Call {
	return &MockProductRepository_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *MockProductRepository_Delete_Call) Run(run func(id string)) *MockProductRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockProductRepository_Delete_Call) Return(_a0 error) *MockProductRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProductRepository_Delete_Call) RunAndReturn(run func(string) error) *MockProductRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields:
func (_m *MockProductRepository) FindAll() ([]domain.Product, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Product, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProductRepository_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockProductRepository_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *MockProductRepository_Expecter) FindAll() *MockProductRepository_FindAll_Call {
	return &MockProductRepository_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *MockProductRepository_FindAll_Call) Run(run func()) *MockProductRepository_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProductRepository_FindAll_Call) Return(_a0 []domain.Product, _a1 error) *MockProductRepository_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProductRepository_FindAll_Call) RunAndReturn(run func() ([]domain.Product, error)) *MockProductRepository_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllByCategory provides a mock function with given fields: _a0
func (_m *MockProductRepository) FindAllByCategory(_a0 string) ([]domain.Product, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindAllByCategory")
	}

	var r0 []domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]domain.Product, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProductRepository_FindAllByCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllByCategory'
type MockProductRepository_FindAllByCategory_Call struct {
	*mock.Call
}

// FindAllByCategory is a helper method to define mock.On call
//   - _a0 category.Category
func (_e *MockProductRepository_Expecter) FindAllByCategory(_a0 interface{}) *MockProductRepository_FindAllByCategory_Call {
	return &MockProductRepository_FindAllByCategory_Call{Call: _e.mock.On("FindAllByCategory", _a0)}
}

func (_c *MockProductRepository_FindAllByCategory_Call) Run(run func(_a0 string)) *MockProductRepository_FindAllByCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockProductRepository_FindAllByCategory_Call) Return(_a0 []domain.Product, _a1 error) *MockProductRepository_FindAllByCategory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProductRepository_FindAllByCategory_Call) RunAndReturn(run func(string) ([]domain.Product, error)) *MockProductRepository_FindAllByCategory_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: id
func (_m *MockProductRepository) FindById(id string) (*domain.Product, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Product, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProductRepository_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type MockProductRepository_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - id string
func (_e *MockProductRepository_Expecter) FindById(id interface{}) *MockProductRepository_FindById_Call {
	return &MockProductRepository_FindById_Call{Call: _e.mock.On("FindById", id)}
}

func (_c *MockProductRepository_FindById_Call) Run(run func(id string)) *MockProductRepository_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockProductRepository_FindById_Call) Return(_a0 *domain.Product, _a1 error) *MockProductRepository_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProductRepository_FindById_Call) RunAndReturn(run func(string) (*domain.Product, error)) *MockProductRepository_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: product
func (_m *MockProductRepository) Save(product *domain.Product) error {
	ret := _m.Called(product)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProductRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockProductRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - product *domain.Product
func (_e *MockProductRepository_Expecter) Save(product interface{}) *MockProductRepository_Save_Call {
	return &MockProductRepository_Save_Call{Call: _e.mock.On("Save", product)}
}

func (_c *MockProductRepository_Save_Call) Run(run func(product *domain.Product)) *MockProductRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Product))
	})
	return _c
}

func (_c *MockProductRepository_Save_Call) Return(_a0 error) *MockProductRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProductRepository_Save_Call) RunAndReturn(run func(*domain.Product) error) *MockProductRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: product
func (_m *MockProductRepository) Update(product *domain.Product) error {
	ret := _m.Called(product)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProductRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockProductRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - product *domain.Product
func (_e *MockProductRepository_Expecter) Update(product interface{}) *MockProductRepository_Update_Call {
	return &MockProductRepository_Update_Call{Call: _e.mock.On("Update", product)}
}

func (_c *MockProductRepository_Update_Call) Run(run func(product *domain.Product)) *MockProductRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Product))
	})
	return _c
}

func (_c *MockProductRepository_Update_Call) Return(_a0 error) *MockProductRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProductRepository_Update_Call) RunAndReturn(run func(*domain.Product) error) *MockProductRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProductRepository creates a new instance of MockProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProductRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProductRepository {
	mock := &MockProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
