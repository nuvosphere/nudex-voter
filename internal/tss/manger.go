package tss

import (
	"sync"

	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/samber/lo"
)

type Manger struct {
	grw          sync.RWMutex
	groups       map[helper.GroupID]*helper.Group
	srw          sync.RWMutex
	sessions     map[helper.SessionID]*helper.Session
	sessionTasks map[helper.TaskID]*helper.Session
}

func NewManger() *Manger {
	return &Manger{
		srw:      sync.RWMutex{},
		grw:      sync.RWMutex{},
		groups:   make(map[helper.GroupID]*helper.Group),
		sessions: make(map[helper.SessionID]*helper.Session),
	}
}

func (m *Manger) AddGroup(group *helper.Group) {
	m.grw.Lock()
	defer m.grw.Unlock()
	m.groups[group.GroupID] = group
}

func (m *Manger) AddSession(session *helper.Session) bool {
	m.grw.Lock()
	_, ok := m.groups[session.Group.GroupID]
	m.grw.Unlock()

	if ok {
		m.srw.Lock()
		m.sessions[session.SessionID] = session
		m.sessionTasks[session.TaskID] = session
		m.srw.Unlock()
	}

	return ok
}

func (m *Manger) GetGroup(groupID helper.GroupID) *helper.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return m.groups[groupID]
}

func (m *Manger) GetSession(sessionID helper.SessionID) *helper.Session {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return m.sessions[sessionID]
}

func (m *Manger) IsMeeting(taskID helper.TaskID) bool {
	m.srw.RLock()
	defer m.srw.RUnlock()
	_, ok := m.sessionTasks[taskID]

	return ok
}

func (m *Manger) GetGroups() []*helper.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return lo.MapToSlice(m.groups, func(_ helper.GroupID, group *helper.Group) *helper.Group { return group })
}

func (m *Manger) GetSessions() []*helper.Session {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return lo.MapToSlice(m.sessions, func(_ helper.SessionID, group *helper.Session) *helper.Session { return group })
}

func (m *Manger) ReleaseGroup(groupID helper.GroupID) {
	m.grw.Lock()
	defer m.grw.Unlock()
	delete(m.groups, groupID)
}

func (m *Manger) ReleaseSession(session helper.SessionID) {
	m.srw.Lock()
	defer m.srw.Unlock()

	s, ok := m.sessions[session]
	if ok {
		delete(m.sessions, session)
		delete(m.sessionTasks, s.TaskID)
	}
}

func (m *Manger) ReleaseAll() {
	m.grw.Lock()
	m.groups = make(map[helper.GroupID]*helper.Group)
	m.grw.Unlock()
	m.srw.Lock()
	m.sessions = make(map[helper.SessionID]*helper.Session)
	m.sessionTasks = make(map[helper.TaskID]*helper.Session)
	m.srw.Unlock()
}
