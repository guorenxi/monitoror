// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/models"
	mock "github.com/stretchr/testify/mock"

	registry "github.com/monitoror/monitoror/registry"

	versions "github.com/monitoror/monitoror/api/config/versions"
)

// Registry is an autogenerated mock type for the Registry type
type Registry struct {
	mock.Mock
}

// GetMonitorables provides a mock function with given fields:
func (_m *Registry) GetMonitorables() []*registry.MonitorableMetadata {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMonitorables")
	}

	var r0 []*registry.MonitorableMetadata
	if rf, ok := ret.Get(0).(func() []*registry.MonitorableMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*registry.MonitorableMetadata)
		}
	}

	return r0
}

// RegisterGenerator provides a mock function with given fields: generatedTileType, minimalVersion, variantNames
func (_m *Registry) RegisterGenerator(generatedTileType models.TileType, minimalVersion versions.RawVersion, variantNames []models.VariantName) registry.GeneratorEnabler {
	ret := _m.Called(generatedTileType, minimalVersion, variantNames)

	if len(ret) == 0 {
		panic("no return value specified for RegisterGenerator")
	}

	var r0 registry.GeneratorEnabler
	if rf, ok := ret.Get(0).(func(models.TileType, versions.RawVersion, []models.VariantName) registry.GeneratorEnabler); ok {
		r0 = rf(generatedTileType, minimalVersion, variantNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(registry.GeneratorEnabler)
		}
	}

	return r0
}

// RegisterMonitorable provides a mock function with given fields: monitorable
func (_m *Registry) RegisterMonitorable(monitorable models.Monitorable) {
	_m.Called(monitorable)
}

// RegisterTile provides a mock function with given fields: tileType, minimalVersion, variantNames
func (_m *Registry) RegisterTile(tileType models.TileType, minimalVersion versions.RawVersion, variantNames []models.VariantName) registry.TileEnabler {
	ret := _m.Called(tileType, minimalVersion, variantNames)

	if len(ret) == 0 {
		panic("no return value specified for RegisterTile")
	}

	var r0 registry.TileEnabler
	if rf, ok := ret.Get(0).(func(models.TileType, versions.RawVersion, []models.VariantName) registry.TileEnabler); ok {
		r0 = rf(tileType, minimalVersion, variantNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(registry.TileEnabler)
		}
	}

	return r0
}

// NewRegistry creates a new instance of Registry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRegistry(t interface {
	mock.TestingT
	Cleanup(func())
}) *Registry {
	mock := &Registry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
