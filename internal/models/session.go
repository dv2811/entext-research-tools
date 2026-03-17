package models

import (
	"errors"
	"time"
)

var (
	ErrSessionTypeNotMatch = errors.New("session type not match")
)

type Session interface {
	Load(data []byte) error
	Save() ([]byte, error)
	UpdatedAt() time.Time
	Copy(source any) error
	// method inherited from mutex
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

type SessionID struct {
	UserID      string
	ServiceName string
}

func CreateSessionID(userID, servicename string) SessionID {
	return SessionID{
		UserID:      userID,
		ServiceName: servicename,
	}
}
