package resolvers

import (
	"context"
	"graphql-learning/autogen"

	"gorm.io/gorm"
)

type contextKey int

const (
	ContextKeyAuth contextKey = iota
	ContextKeyDB
)

type resolver struct{}

func NewConfig() autogen.Config {
	return autogen.Config{
		Resolvers: &resolver{},
	}
}

func (r *resolver) Query() autogen.QueryResolver {
	return &queryResolver{r}
}

func (r *resolver) Mutation() autogen.MutationResolver {
	return &mutationResolver{r}
}

func (r *resolver) db(ctx context.Context) *gorm.DB {
	if db, ok := ctx.Value(ContextKeyDB).(*gorm.DB); ok {
		return db
	}
	return nil
}
