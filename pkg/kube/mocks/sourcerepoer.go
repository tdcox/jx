// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/pkg/kube (interfaces: SourceRepoer)

package kube_test

import (
	v1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockSourceRepoer struct {
	fail func(message string, callerSkip ...int)
}

func NewMockSourceRepoer() *MockSourceRepoer {
	return &MockSourceRepoer{fail: pegomock.GlobalFailHandler}
}

func (mock *MockSourceRepoer) FailHandler() pegomock.FailHandler {
	return mock.fail
}

func (mock *MockSourceRepoer) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }

func (mock *MockSourceRepoer) CreateOrUpdateSourceRepository(_param0 string, _param1 string, _param2 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSourceRepoer().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CreateOrUpdateSourceRepository", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockSourceRepoer) CreateSourceRepository(_param0 string, _param1 string, _param2 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSourceRepoer().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CreateSourceRepository", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockSourceRepoer) DeleteSourceRepository(_param0 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSourceRepoer().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DeleteSourceRepository", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockSourceRepoer) GetSourceRepository(_param0 string) (*v1.SourceRepository, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSourceRepoer().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetSourceRepository", params, []reflect.Type{reflect.TypeOf((**v1.SourceRepository)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *v1.SourceRepository
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*v1.SourceRepository)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockSourceRepoer) ListSourceRepositories() (*v1.SourceRepositoryList, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSourceRepoer().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ListSourceRepositories", params, []reflect.Type{reflect.TypeOf((**v1.SourceRepositoryList)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *v1.SourceRepositoryList
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*v1.SourceRepositoryList)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockSourceRepoer) VerifyWasCalledOnce() *VerifierSourceRepoer {
	return &VerifierSourceRepoer{mock, pegomock.Times(1), nil}
}

func (mock *MockSourceRepoer) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierSourceRepoer {
	return &VerifierSourceRepoer{mock, invocationCountMatcher, nil}
}

func (mock *MockSourceRepoer) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierSourceRepoer {
	return &VerifierSourceRepoer{mock, invocationCountMatcher, inOrderContext}
}

type VerifierSourceRepoer struct {
	mock                   *MockSourceRepoer
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierSourceRepoer) CreateOrUpdateSourceRepository(_param0 string, _param1 string, _param2 string) *SourceRepoer_CreateOrUpdateSourceRepository_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CreateOrUpdateSourceRepository", params)
	return &SourceRepoer_CreateOrUpdateSourceRepository_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type SourceRepoer_CreateOrUpdateSourceRepository_OngoingVerification struct {
	mock              *MockSourceRepoer
	methodInvocations []pegomock.MethodInvocation
}

func (c *SourceRepoer_CreateOrUpdateSourceRepository_OngoingVerification) GetCapturedArguments() (string, string, string) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *SourceRepoer_CreateOrUpdateSourceRepository_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]string, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierSourceRepoer) CreateSourceRepository(_param0 string, _param1 string, _param2 string) *SourceRepoer_CreateSourceRepository_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CreateSourceRepository", params)
	return &SourceRepoer_CreateSourceRepository_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type SourceRepoer_CreateSourceRepository_OngoingVerification struct {
	mock              *MockSourceRepoer
	methodInvocations []pegomock.MethodInvocation
}

func (c *SourceRepoer_CreateSourceRepository_OngoingVerification) GetCapturedArguments() (string, string, string) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *SourceRepoer_CreateSourceRepository_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]string, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierSourceRepoer) DeleteSourceRepository(_param0 string) *SourceRepoer_DeleteSourceRepository_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DeleteSourceRepository", params)
	return &SourceRepoer_DeleteSourceRepository_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type SourceRepoer_DeleteSourceRepository_OngoingVerification struct {
	mock              *MockSourceRepoer
	methodInvocations []pegomock.MethodInvocation
}

func (c *SourceRepoer_DeleteSourceRepository_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *SourceRepoer_DeleteSourceRepository_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierSourceRepoer) GetSourceRepository(_param0 string) *SourceRepoer_GetSourceRepository_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetSourceRepository", params)
	return &SourceRepoer_GetSourceRepository_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type SourceRepoer_GetSourceRepository_OngoingVerification struct {
	mock              *MockSourceRepoer
	methodInvocations []pegomock.MethodInvocation
}

func (c *SourceRepoer_GetSourceRepository_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *SourceRepoer_GetSourceRepository_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierSourceRepoer) ListSourceRepositories() *SourceRepoer_ListSourceRepositories_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ListSourceRepositories", params)
	return &SourceRepoer_ListSourceRepositories_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type SourceRepoer_ListSourceRepositories_OngoingVerification struct {
	mock              *MockSourceRepoer
	methodInvocations []pegomock.MethodInvocation
}

func (c *SourceRepoer_ListSourceRepositories_OngoingVerification) GetCapturedArguments() {
}

func (c *SourceRepoer_ListSourceRepositories_OngoingVerification) GetAllCapturedArguments() {
}
