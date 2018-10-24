package cache

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// DataBuffer is cacher for data
type DataBuffer struct {
	store          *MemoryStore
	dataList       *list.List
	dataIndex      map[string]map[string]*list.Element
	mutex          sync.RWMutex
	MaxElementSize int
	Expired        time.Duration
	GcInterval     time.Duration
}

var (
	defaultExpired    = time.Second * 60
	defaultGcInterval = time.Second * 60
	defaultMaxRemoved = 50
)

// NewDataBuffer create a DataBuffer
func NewDataBuffer(maxElementSize int) *DataBuffer {
	buffer := &DataBuffer{
		store: NewMemoryStore(), dataList: list.New(), MaxElementSize: maxElementSize,
		GcInterval: defaultGcInterval, Expired: defaultExpired,
		dataIndex: make(map[string]map[string]*list.Element),
	}
	buffer.RunGC()
	return buffer
}

// RunGC run once every m.GcInterval
func (m *DataBuffer) RunGC() {
	time.AfterFunc(m.GcInterval, func() {
		m.RunGC()
		m.GC()
	})
}

// GC check ids lit and sql list to remove all element expired
func (m *DataBuffer) GC() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var removedNum int
	for e := m.dataList.Front(); e != nil; {
		if removedNum <= defaultMaxRemoved &&
			time.Now().After(e.Value.(*dataNode).expiredAt) {
			removedNum++
			next := e.Next()
			node := e.Value.(*dataNode)
			m.delData(node.method, node.url)
			e = next
		} else {
			break
		}
	}
}

// GetData returns data according method and url from cache
func (m *DataBuffer) GetData(method string, url string) interface{} {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if _, ok := m.dataIndex[method]; !ok {
		m.dataIndex[method] = make(map[string]*list.Element)
	}
	tid := genID(method, url)
	if v, err := m.store.Get(tid); err == nil {
		if el, ok := m.dataIndex[method][url]; ok {
			// if expired, remove the node and return nil
			if time.Now().After(el.Value.(*dataNode).expiredAt) {
				m.DelData(method, url)
				return nil
			}
		} else {
			el = m.dataList.PushBack(newDataNode(method, url))
			m.dataIndex[method][url] = el
		}
		return v
	}

	m.delData(method, url)
	return nil
}

func (m *DataBuffer) clearDatas(method string) {
	if tis, ok := m.dataIndex[method]; ok {
		for url, v := range tis {
			m.dataList.Remove(v)
			tid := genID(method, url)
			m.store.Del(tid)
		}
	}
	m.dataIndex[method] = make(map[string]*list.Element)
}

// ClearDatas clears all method-url mapping on table tableName from cache
func (m *DataBuffer) ClearDatas(method string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.clearDatas(method)
}

// PutData puts data
func (m *DataBuffer) PutData(method string, url string, obj interface{}, expired ...time.Duration) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var el *list.Element
	var ok bool

	if el, ok = m.dataIndex[method][url]; !ok {
		el = m.dataList.PushBack(newDataNode(method, url))
		m.dataIndex[method][url] = el
	} else {
		var exp = m.Expired
		if len(expired) > 0 && expired[0] > m.Expired {
			exp = expired[0]
		}
		el.Value.(*dataNode).expiredAt = time.Now().Add(exp)
	}

	m.store.Put(genID(method, url), obj)
	if m.dataList.Len() > m.MaxElementSize {
		e := m.dataList.Front()
		node := e.Value.(*dataNode)
		m.delData(node.method, node.url)
	}
}

func (m *DataBuffer) delData(method string, url string) {
	if _, ok := m.dataIndex[method]; ok {
		if el, ok := m.dataIndex[method][url]; ok {
			delete(m.dataIndex[method], url)
			m.dataList.Remove(el)
		}
	}
	m.store.Del(genID(method, url))
}

// DelData deletes data
func (m *DataBuffer) DelData(method string, url string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.delData(method, url)
}

type dataNode struct {
	method    string
	url       string
	expiredAt time.Time
}

func newDataNode(method, url string) *dataNode {
	return &dataNode{method, url, time.Now()}
}

func genID(prefix string, id string) string {
	return fmt.Sprintf("%v-%v", prefix, id)
}
