package tss

import (
	"sync"

	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/samber/lo"
)

// Releaser is the interface that wraps the basic Release method.
type Releaser interface {
	// Release releases associated resources. Release should always success
	// and can be called multiple times without causing error.
	Release()
}

type SessionReleaser interface {
	SessionRelease(session helper.SessionID)
}

type Manager[T comparable] struct {
	grw          sync.RWMutex
	groups       map[helper.GroupID]*helper.Group
	srw          sync.RWMutex
	sessions     map[helper.SessionID]Session[T]
	sessionTasks map[T]Session[T]
	inToOut      chan *SignResult[T]
}

func NewManager[T comparable]() *Manager[T] {
	return &Manager[T]{
		srw:          sync.RWMutex{},
		grw:          sync.RWMutex{},
		groups:       make(map[helper.GroupID]*helper.Group),
		sessions:     make(map[helper.SessionID]Session[T]),
		sessionTasks: make(map[T]Session[T]),
		inToOut:      make(chan *SignResult[T], 1024),
	}
}

func (m *Manager[T]) AddGroup(group *helper.Group) {
	m.grw.Lock()
	defer m.grw.Unlock()
	m.groups[group.GroupID] = group
}

func (m *Manager[T]) AddSession(session Session[T]) bool {
	m.grw.Lock()
	_, ok := m.groups[session.GroupID()]
	m.grw.Unlock()

	if ok {
		m.srw.Lock()
		m.sessions[session.SessionID()] = session
		m.sessionTasks[session.TaskID()] = session
		m.srw.Unlock()
	}

	return ok
}

func (m *Manager[T]) GetGroup(groupID helper.GroupID) *helper.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return m.groups[groupID]
}

func (m *Manager[T]) GetSession(sessionID helper.SessionID) Session[T] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return m.sessions[sessionID]
}

func (m *Manager[T]) IsMeeting(taskID T) bool {
	m.srw.RLock()
	defer m.srw.RUnlock()
	_, ok := m.sessionTasks[taskID]

	return ok
}

func (m *Manager[T]) GetGroups() []*helper.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return lo.MapToSlice(m.groups, func(_ helper.GroupID, group *helper.Group) *helper.Group { return group })
}

func (m *Manager[T]) GetSessions() []Session[T] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return lo.MapToSlice(m.sessions, func(_ helper.SessionID, session Session[T]) Session[T] { return session })
}

func (m *Manager[T]) ReleaseGroup(groupID helper.GroupID) {
	m.grw.Lock()
	defer m.grw.Unlock()
	delete(m.groups, groupID)
}

func (m *Manager[T]) SessionRelease(session helper.SessionID) {
	m.srw.Lock()
	defer m.srw.Unlock()

	s, ok := m.sessions[session]
	if ok {
		delete(m.sessions, session)
		delete(m.sessionTasks, s.TaskID())
	}
}

func (m *Manager[T]) Release() {
	m.grw.Lock()
	m.groups = make(map[helper.GroupID]*helper.Group)
	m.grw.Unlock()
	m.srw.Lock()
	m.sessions = make(map[helper.SessionID]Session[T])
	m.sessionTasks = make(map[T]Session[T])
	m.srw.Unlock()
	close(m.inToOut)
}
