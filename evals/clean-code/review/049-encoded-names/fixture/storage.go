package fixture

type IStorage interface {
	Save(strKey string, arrBytes []byte) error
}

type MemoryStorage struct {
	m_entries map[string][]byte
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{m_entries: map[string][]byte{}}
}

func (s *MemoryStorage) Save(strKey string, arrBytes []byte) error {
	s.m_entries[strKey] = append([]byte(nil), arrBytes...)
	return nil
}

func (s *MemoryStorage) Count() int {
	return len(s.m_entries)
}
