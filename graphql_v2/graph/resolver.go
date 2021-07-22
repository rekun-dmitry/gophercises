package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"graphql_v2/graph/internal/provincestore"
)

type Resolver struct {
	Store *provincestore.ProvinceStore
}
