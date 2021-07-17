package graph

import (
	"fmt"
	"graphql/graph/model"
	"sync"
)

type ProvinceStore struct {
	sync.Mutex
	Provinces map[string]*model.Province
}

func New() *ProvinceStore {
	ts := &ProvinceStore{}
	ts.Provinces = make(map[string]*model.Province)
	return ts
}

// CreateProvince creates a new province in the store.
func (ps *ProvinceStore) CreateProvince(province model.NewProvince) string {
	ps.Lock()
	defer ps.Unlock()

	provinceCreated := model.Province{
		ID:        province.ID,
		Name:      province.Name,
		AdmDev:    province.AdmDev,
		DipDev:    province.DipDev,
		MilDev:    province.MilDev,
		TradeGood: province.TradeGood,
		TradeNode: province.TradeNode,
		Modifiers: province.Modifiers,
	}

	ps.Provinces[province.ID] = &provinceCreated
	return province.ID
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

// GetProvince retrieves a province from the store, by id. If no such id exists, an
// error is returned.
func (ps *ProvinceStore) GetProvince(id string) (*model.Province, error) {
	ps.Lock()
	defer ps.Unlock()

	t, ok := ps.Provinces[id]
	if ok {
		return t, nil
	} else {
		return &model.Province{}, fmt.Errorf("province with id=%d not found", id)
	}
}

// DeleteProvince deletes the province with the given id. If no such id exists, an error
// is returned.
func (ts *ProvinceStore) DeleteProvince(id string) error {
	ts.Lock()
	defer ts.Unlock()

	if _, ok := ts.Provinces[id]; !ok {
		return fmt.Errorf("province with id=%d not found", id)
	}

	delete(ts.Provinces, id)
	return nil
}
