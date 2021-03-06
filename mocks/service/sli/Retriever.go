// Code generated by mockery v1.0.0. DO NOT EDIT.

package sli

import mock "github.com/stretchr/testify/mock"
import sli "github.com/Medium/service-level-operator/pkg/service/sli"
import v1alpha1 "github.com/Medium/service-level-operator/pkg/apis/monitoring/v1alpha1"

// Retriever is an autogenerated mock type for the Retriever type
type Retriever struct {
	mock.Mock
}

// Retrieve provides a mock function with given fields: _a0
func (_m *Retriever) Retrieve(_a0 *v1alpha1.SLI) (sli.Result, error) {
	ret := _m.Called(_a0)

	var r0 sli.Result
	if rf, ok := ret.Get(0).(func(*v1alpha1.SLI) sli.Result); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(sli.Result)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.SLI) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
