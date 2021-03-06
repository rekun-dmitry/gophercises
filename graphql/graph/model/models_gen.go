// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Modifier struct {
	Name     string `json:"Name"`
	Contents string `json:"Contents"`
}

type NewModifier struct {
	Name     string `json:"Name"`
	Contents string `json:"Contents"`
}

type NewProvince struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	AdmDev    int            `json:"admin_dev"`
	DipDev    int            `json:"dip_dev"`
	MilDev    int            `json:"mil_dev"`
	TradeGood string         `json:"trade_good"`
	TradeNode string         `json:"trade_node"`
	Modifiers []*NewModifier `json:"modifiers"`
}

type Province struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	AdmDev    int            `json:"admin_dev"`
	DipDev    int            `json:"dip_dev"`
	MilDev    int            `json:"mil_dev"`
	TradeGood string         `json:"trade_good"`
	TradeNode string         `json:"trade_node"`
	Modifiers []*NewModifier `json:"modifiers"`
}
