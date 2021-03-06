// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

// Add provides a mock function with given fields: key, value
func (_m *Cache) Add(key interface{}, value interface{}) {
	_m.Called(key, value)
}

// AddWithDefaultExpires provides a mock function with given fields: key, value
func (_m *Cache) AddWithDefaultExpires(key interface{}, value interface{}) {
	_m.Called(key, value)
}

// AddWithExpiresInSecs provides a mock function with given fields: key, value, expireAtSecs
func (_m *Cache) AddWithExpiresInSecs(key interface{}, value interface{}, expireAtSecs int64) {
	_m.Called(key, value, expireAtSecs)
}

// Get provides a mock function with given fields: key
func (_m *Cache) Get(key interface{}) (interface{}, bool) {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(interface{}) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetInvalidateClusterEvent provides a mock function with given fields:
func (_m *Cache) GetInvalidateClusterEvent() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetOrAdd provides a mock function with given fields: key, value, ttl
func (_m *Cache) GetOrAdd(key interface{}, value interface{}, ttl time.Duration) (interface{}, bool) {
	ret := _m.Called(key, value, ttl)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, time.Duration) interface{}); ok {
		r0 = rf(key, value, ttl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(interface{}, interface{}, time.Duration) bool); ok {
		r1 = rf(key, value, ttl)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Keys provides a mock function with given fields:
func (_m *Cache) Keys() []interface{} {
	ret := _m.Called()

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func() []interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	return r0
}

// Len provides a mock function with given fields:
func (_m *Cache) Len() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *Cache) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Purge provides a mock function with given fields:
func (_m *Cache) Purge() {
	_m.Called()
}

// Remove provides a mock function with given fields: key
func (_m *Cache) Remove(key interface{}) {
	_m.Called(key)
}
