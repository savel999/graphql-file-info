package graph

import "graphql-file-info/graph/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ItemsRepository *repository.ListItemsRepository
}
