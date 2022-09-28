package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/droslean/99designs-gqlgen-filter/graph/generated"
	"github.com/droslean/99designs-gqlgen-filter/graph/model"
)

// QueryBug is the resolver for the queryBug field.
func (r *queryResolver) QueryBug(ctx context.Context, filter *model.BugFilter, first *int, offset *int) ([]*model.Bug, error) {
	data := []byte(`
[
  {
    "id": "0x12345",
    "name": "Test BUG",
    "products": [
      {
        "name": "product-1"
      },
      {
        "name": "product-2"
      },
      {
        "name": "product-3"
      },
      {
        "name": "product-4"
      }
    ]
  }
]`)

	var bugs []*model.Bug

	if err := json.Unmarshal(data, &bugs); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return bugs, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
