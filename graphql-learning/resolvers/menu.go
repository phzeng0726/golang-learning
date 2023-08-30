package resolvers

import (
	"context"

	"graphql-learning/autogen"
	"graphql-learning/orm"
)

func (r *queryResolver) Menus(ctx context.Context, category *autogen.MenuCategory) ([]orm.Menu, error) {
	conn := r.db(ctx)
	var results []orm.Menu
	if category != nil {
		if err := conn.Where("category = ?", category.String()).Find(&results).Error; err != nil {
			return nil, err
		}
	} else {
		if err := conn.Find(&results).Error; err != nil {
			return nil, err
		}
	}
	return results, nil
}
