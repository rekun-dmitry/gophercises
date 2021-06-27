package provinces

import (
	"fmt"
	"sync"
)

type Province struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	AdmDev    int      `json:"admin_dev"`
	DipDev    int      `json:"dip_dev"`
	MilDev    int      `json:"mil_dev"`
	TradeGood string   `json:"trade_good"`
	TradeNode string   `json:"trade_node"`
	Modifiers []string `json:"modifiers"`
}

type ProvinceStore struct {
	sync.Mutex
	Provinces map[int]Province
}

func New() *ProvinceStore {
	ts := &ProvinceStore{}
	ts.Provinces = make(map[int]Province)
	return ts
}

// CreateProvince creates a new province in the store.
func (ps *ProvinceStore) CreateProvince(province Province) int {
	ps.Lock()
	defer ps.Unlock()

	provinceCreated := Province{
		Id:        province.Id,
		Name:      province.Name,
		AdmDev:    province.AdmDev,
		DipDev:    province.DipDev,
		MilDev:    province.MilDev,
		TradeGood: province.TradeGood,
		TradeNode: province.TradeNode,
		Modifiers: province.Modifiers,
	}

	ps.Provinces[province.Id] = provinceCreated
	return province.Id
}

// DeleteAllProvinces deletes all provinces in the store.
func (ps *ProvinceStore) DeleteAllProvinces() error {
	ps.Lock()
	defer ps.Unlock()

	ps.Provinces = make(map[int]Province)
	return nil
}

// GetAllProvinces returns all the provinces in the store, in arbitrary order.
func (ps *ProvinceStore) GetAllProvinces() []Province {
	ps.Lock()
	defer ps.Unlock()

	allProvinces := make([]Province, 0, len(ps.Provinces))
	for _, province := range ps.Provinces {
		allProvinces = append(allProvinces, province)
	}
	return allProvinces
}

// GetProvince retrieves a province from the store, by id. If no such id exists, an
// error is returned.
func (ps *ProvinceStore) GetProvince(id int) (Province, error) {
	ps.Lock()
	defer ps.Unlock()

	t, ok := ps.Provinces[id]
	if ok {
		return t, nil
	} else {
		return Province{}, fmt.Errorf("province with id=%d not found", id)
	}
}

// DeleteProvince deletes the province with the given id. If no such id exists, an error
// is returned.
func (ts *ProvinceStore) DeleteProvince(id int) error {
	ts.Lock()
	defer ts.Unlock()

	if _, ok := ts.Provinces[id]; !ok {
		return fmt.Errorf("province with id=%d not found", id)
	}

	delete(ts.Provinces, id)
	return nil
}
