package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/pacerino/pr0sauce_graphql/graph/model"
)

type Resolver struct {
	items    []*model.Item
	comments []*model.Comment
}
