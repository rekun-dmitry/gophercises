package provincestore

import (
	"fmt"
	"graphql_v2/graph/model"
	"sync"
)

// TaskStore is a simple in-memory database of tasks; TaskStore methods are
// safe to call concurrently.
type ProvinceStore struct {
	sync.Mutex
	Provinces map[string]*model.Province
}

func New() *ProvinceStore {
	ps := &ProvinceStore{}
	ps.Provinces = make(map[string]*model.Province)
	return ps
}

// CreateProvince creates a new province in the store.
func (ps *ProvinceStore) CreateProvince(id string, name string, modifiers []*model.Modifier) string {
	ps.Lock()
	defer ps.Unlock()

	province := model.Province{
		ID:        id,
		Name:      name,
		Modifiers: modifiers}
	ps.Provinces[id] = &province
	return id
}

// GetProvince retrieves a province from the store, by id. If no such id exists, an
// error is returned.
func (ps *ProvinceStore) GetProvince(id string) (*model.Province, error) {
	ps.Lock()
	defer ps.Unlock()

	t, ok := ps.Provinces[id]
	if ok {
		return t, nil
	} else {
		return nil, fmt.Errorf("province with id=%d not found", id)
	}
}

// DeleteProvince deletes the province with the given id. If no such id exists, an error
// is returned.
func (ps *ProvinceStore) DeleteProvince(id string) error {
	ps.Lock()
	defer ps.Unlock()

	if _, ok := ps.Provinces[id]; !ok {
		return fmt.Errorf("province with id=%d not found", id)
	}

	delete(ps.Provinces, id)
	return nil
}

// DeleteAllProvinces deletes all provinces in the store.
func (ps *ProvinceStore) DeleteAllProvinces() error {
	ps.Lock()
	defer ps.Unlock()
	ps.Provinces = make(map[string]*model.Province)
	return nil

}

// GetAllProvinces returns all the provinces in the store, in arbitrary order.
func (ps *ProvinceStore) GetAllProvinces() []*model.Province {
	ps.Lock()
	defer ps.Unlock()

	allProvinces := make([]*model.Province, 0, len(ps.Provinces))
	for _, province := range ps.Provinces {
		allProvinces = append(allProvinces, province)
	}
	return allProvinces
}
