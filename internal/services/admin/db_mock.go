// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package admin

import (
	"github.com/dalot/go-skeleton-mid/pkg/resources"
	"sync"
)

// Ensure, that DatabaseMock does implement Database.
// If this is not the case, regenerate this file with moq.
var _ Database = &DatabaseMock{}

// DatabaseMock is a mock implementation of Database.
//
// 	func TestSomethingThatUsesDatabase(t *testing.T) {
//
// 		// make and configure a mocked Database
// 		mockedDatabase := &DatabaseMock{
// 			GetHashFunc: func(username string) ([]byte, error) {
// 				panic("mock out the GetHash method")
// 			},
// 			GetMessageFunc: func(id string) (*resources.Message, error) {
// 				panic("mock out the GetMessage method")
// 			},
// 			MessagesFunc: func() []*resources.Message {
// 				panic("mock out the Messages method")
// 			},
// 			UpdateTextFunc: func(id string, text string) (*resources.Message, error) {
// 				panic("mock out the UpdateText method")
// 			},
// 			UserExistsFunc: func(key string) bool {
// 				panic("mock out the UserExists method")
// 			},
// 		}
//
// 		// use mockedDatabase in code that requires Database
// 		// and then make assertions.
//
// 	}
type DatabaseMock struct {
	// GetHashFunc mocks the GetHash method.
	GetHashFunc func(username string) ([]byte, error)

	// GetMessageFunc mocks the GetMessage method.
	GetMessageFunc func(id string) (*resources.Message, error)

	// MessagesFunc mocks the Messages method.
	MessagesFunc func() []*resources.Message

	// UpdateTextFunc mocks the UpdateText method.
	UpdateTextFunc func(id string, text string) (*resources.Message, error)

	// UserExistsFunc mocks the UserExists method.
	UserExistsFunc func(key string) bool

	// calls tracks calls to the methods.
	calls struct {
		// GetHash holds details about calls to the GetHash method.
		GetHash []struct {
			// Username is the username argument value.
			Username string
		}
		// GetMessage holds details about calls to the GetMessage method.
		GetMessage []struct {
			// ID is the id argument value.
			ID string
		}
		// Messages holds details about calls to the Messages method.
		Messages []struct {
		}
		// UpdateText holds details about calls to the UpdateText method.
		UpdateText []struct {
			// ID is the id argument value.
			ID string
			// Text is the text argument value.
			Text string
		}
		// UserExists holds details about calls to the UserExists method.
		UserExists []struct {
			// Key is the key argument value.
			Key string
		}
	}
	lockGetHash    sync.RWMutex
	lockGetMessage sync.RWMutex
	lockMessages   sync.RWMutex
	lockUpdateText sync.RWMutex
	lockUserExists sync.RWMutex
}

// GetHash calls GetHashFunc.
func (mock *DatabaseMock) GetHash(username string) ([]byte, error) {
	if mock.GetHashFunc == nil {
		panic("DatabaseMock.GetHashFunc: method is nil but Database.GetHash was just called")
	}
	callInfo := struct {
		Username string
	}{
		Username: username,
	}
	mock.lockGetHash.Lock()
	mock.calls.GetHash = append(mock.calls.GetHash, callInfo)
	mock.lockGetHash.Unlock()
	return mock.GetHashFunc(username)
}

// GetHashCalls gets all the calls that were made to GetHash.
// Check the length with:
//     len(mockedDatabase.GetHashCalls())
func (mock *DatabaseMock) GetHashCalls() []struct {
	Username string
} {
	var calls []struct {
		Username string
	}
	mock.lockGetHash.RLock()
	calls = mock.calls.GetHash
	mock.lockGetHash.RUnlock()
	return calls
}

// GetMessage calls GetMessageFunc.
func (mock *DatabaseMock) GetMessage(id string) (*resources.Message, error) {
	if mock.GetMessageFunc == nil {
		panic("DatabaseMock.GetMessageFunc: method is nil but Database.GetMessage was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetMessage.Lock()
	mock.calls.GetMessage = append(mock.calls.GetMessage, callInfo)
	mock.lockGetMessage.Unlock()
	return mock.GetMessageFunc(id)
}

// GetMessageCalls gets all the calls that were made to GetMessage.
// Check the length with:
//     len(mockedDatabase.GetMessageCalls())
func (mock *DatabaseMock) GetMessageCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetMessage.RLock()
	calls = mock.calls.GetMessage
	mock.lockGetMessage.RUnlock()
	return calls
}

// Messages calls MessagesFunc.
func (mock *DatabaseMock) Messages() []*resources.Message {
	if mock.MessagesFunc == nil {
		panic("DatabaseMock.MessagesFunc: method is nil but Database.Messages was just called")
	}
	callInfo := struct {
	}{}
	mock.lockMessages.Lock()
	mock.calls.Messages = append(mock.calls.Messages, callInfo)
	mock.lockMessages.Unlock()
	return mock.MessagesFunc()
}

// MessagesCalls gets all the calls that were made to Messages.
// Check the length with:
//     len(mockedDatabase.MessagesCalls())
func (mock *DatabaseMock) MessagesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockMessages.RLock()
	calls = mock.calls.Messages
	mock.lockMessages.RUnlock()
	return calls
}

// UpdateText calls UpdateTextFunc.
func (mock *DatabaseMock) UpdateText(id string, text string) (*resources.Message, error) {
	if mock.UpdateTextFunc == nil {
		panic("DatabaseMock.UpdateTextFunc: method is nil but Database.UpdateText was just called")
	}
	callInfo := struct {
		ID   string
		Text string
	}{
		ID:   id,
		Text: text,
	}
	mock.lockUpdateText.Lock()
	mock.calls.UpdateText = append(mock.calls.UpdateText, callInfo)
	mock.lockUpdateText.Unlock()
	return mock.UpdateTextFunc(id, text)
}

// UpdateTextCalls gets all the calls that were made to UpdateText.
// Check the length with:
//     len(mockedDatabase.UpdateTextCalls())
func (mock *DatabaseMock) UpdateTextCalls() []struct {
	ID   string
	Text string
} {
	var calls []struct {
		ID   string
		Text string
	}
	mock.lockUpdateText.RLock()
	calls = mock.calls.UpdateText
	mock.lockUpdateText.RUnlock()
	return calls
}

// UserExists calls UserExistsFunc.
func (mock *DatabaseMock) UserExists(key string) bool {
	if mock.UserExistsFunc == nil {
		panic("DatabaseMock.UserExistsFunc: method is nil but Database.UserExists was just called")
	}
	callInfo := struct {
		Key string
	}{
		Key: key,
	}
	mock.lockUserExists.Lock()
	mock.calls.UserExists = append(mock.calls.UserExists, callInfo)
	mock.lockUserExists.Unlock()
	return mock.UserExistsFunc(key)
}

// UserExistsCalls gets all the calls that were made to UserExists.
// Check the length with:
//     len(mockedDatabase.UserExistsCalls())
func (mock *DatabaseMock) UserExistsCalls() []struct {
	Key string
} {
	var calls []struct {
		Key string
	}
	mock.lockUserExists.RLock()
	calls = mock.calls.UserExists
	mock.lockUserExists.RUnlock()
	return calls
}
