// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	product "github.com/vikramyadav1/weaver/sampleApp/models/product"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// All provides a mock function with given fields:
func (_m *ProductRepository) All() ([]product.Product, error) {
	ret := _m.Called()

	var r0 []product.Product
	if rf, ok := ret.Get(0).(func() []product.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: p
func (_m *ProductRepository) Create(p product.Product) (*product.Product, error) {
	ret := _m.Called(p)

	var r0 *product.Product
	if rf, ok := ret.Get(0).(func(product.Product) *product.Product); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(product.Product) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *ProductRepository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *ProductRepository) Get(id int) (*product.Product, error) {
	ret := _m.Called(id)

	var r0 *product.Product
	if rf, ok := ret.Get(0).(func(int) *product.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, p
func (_m *ProductRepository) Update(id int, p product.Product) (*product.Product, error) {
	ret := _m.Called(id, p)

	var r0 *product.Product
	if rf, ok := ret.Get(0).(func(int, product.Product) *product.Product); ok {
		r0 = rf(id, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, product.Product) error); ok {
		r1 = rf(id, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
