package packages

import (
	"sync"
)

var _ Package = &PackageMock{}

type PackageMock struct {
	// IsLocalFunc mocks the IsLocal method.
	IsLocalFunc func() bool

	// NameFunc mocks the Name method.
	NameFunc func() string

	// PathFunc mocks the Path method.
	PathFunc func() string

	// URLFunc mocks the URL method.
	URLFunc func() string

	// UpdateAvailableFunc mocks the UpdateAvailable method.
	UpdateAvailableFunc func() bool

	// calls tracks calls to the methods.
	calls struct {
		// IsLocal holds details about calls to the IsLocal method.
		IsLocal []struct {}
		// Name holds details about calls to the Name method.
		Name []struct {}
		// Path holds details about calls to the Path method.
		Path []struct {}
		// URL holds details about calls to the URL method.
		URL []struct {}
		// UpdateAvailable holds details about calls to the UpdateAvailable method.
		UpdateAvailable []struct {}
	}

	lockIsLocal         sync.RWMutex
	lockName            sync.RWMutex
	lockPath            sync.RWMutex
	lockURL             sync.RWMutex
	lockUpdateAvailable sync.RWMutex
}

// IsLocal calls IsLocalFunc.
func (mock *PackageMock) IsLocal() bool {
	if mock.IsLocalFunc == nil {
		panic("PackageMock.IsLocalFunc: method is nil but Package.IsLocal was just called")
	}

	callInfo := struct {}{}

	mock.lockIsLocal.Lock()
	mock.calls.IsLocal = append(mock.calls.IsLocal, callInfo)
	mock.lockIsLocal.Unlock()
	return mock.IsLocalFunc()
}

// IsLocalCalls gets all the calls that were made to IsLocal.
// Check the length with:
//     len(mockedPackage.IsLocalCalls())
func (mock *PackageMock) IsLocalCalls() []struct {} {
	var calls []struct {}

	mock.lockIsLocal.RLock()
	calls = mock.calls.IsLocal
	mock.lockIsLocal.RUnlock()
	return calls
}

// Name calls NameFunc.
func (mock *PackageMock) Name() string {
	if mock.NameFunc == nil {
		panic("PackageMock.NameFunc: method is nil but Package.Name was just called")
	}

	callInfo := struct {}{}

	mock.lockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	mock.lockName.Unlock()
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//     len(mockedPackage.NameCalls())
func (mock *PackageMock) NameCalls() []struct {} {
	var calls []struct {}

	mock.lockName.RLock()
	calls = mock.calls.Name
	mock.lockName.RUnlock()
	return calls
}

// Path calls PathFunc.
func (mock *PackageMock) Path() string {
	if mock.PathFunc == nil {
		panic("PackageMock.PathFunc: method is nil but Package.Path was just called")
	}

	callInfo := struct {}{}

	mock.lockPath.Lock()
	mock.calls.Path = append(mock.calls.Path, callInfo)
	mock.lockPath.Unlock()
	return mock.PathFunc()
}

// PathCalls gets all the calls that were made to Path.
// Check the length with:
//     len(mockedPackage.PathCalls())
func (mock *PackageMock) PathCalls() []struct {} {
	var calls []struct {}

	mock.lockPath.RLock()
	calls = mock.calls.Path
	mock.lockPath.RUnlock()
	return calls
}

// URL calls URLFunc.
func (mock *PackageMock) URL() string {
	if mock.URLFunc == nil {
		panic("PackageMock.URLFunc: method is nil but Package.URL was just called")
	}

	callInfo := struct {}{}

	mock.lockURL.Lock()
	mock.calls.URL = append(mock.calls.URL, callInfo)
	mock.lockURL.Unlock()
	return mock.URLFunc()
}

// URLCalls gets all the calls that were made to URL.
// Check the length with:
//     len(mockedPackage.URLCalls())
func (mock *PackageMock) URLCalls() []struct {} {
	var calls []struct {}

	mock.lockURL.RLock()
	calls = mock.calls.URL
	mock.lockURL.RUnlock()
	return calls
}

// UpdateAvailable calls UpdateAvailableFunc.
func (mock *PackageMock) UpdateAvailable() bool {
	if mock.UpdateAvailableFunc == nil {
		panic("PackageMock.UpdateAvailableFunc: method is nil but Package.UpdateAvailable was just called")
	}

	callInfo := struct {}{}

	mock.lockUpdateAvailable.Lock()
	mock.calls.UpdateAvailable = append(mock.calls.UpdateAvailable, callInfo)
	mock.lockUpdateAvailable.Unlock()
	return mock.UpdateAvailableFunc()
}

// UpdateAvailableCalls gets all the calls that were made to UpdateAvailable.
// Check the length with:
//     len(mockedPackage.UpdateAvailableCalls())
func (mock *PackageMock) UpdateAvailableCalls() []struct {} {
	var calls []struct {}
	mock.lockUpdateAvailable.RLock()
	calls = mock.calls.UpdateAvailable
	mock.lockUpdateAvailable.RUnlock()
	return calls
}
