package provinces

import "sync"

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
	nextId    int
}

func New() *ProvinceStore {
	ts := &ProvinceStore{}
	ts.Provinces = make(map[int]Province)
	ts.nextId = 0
	return ts
}
