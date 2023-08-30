package resolvers

import (
	"context"
	"graphql-learning/autogen"
	"graphql-learning/orm"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *queryResolver) Users(ctx context.Context) ([]orm.UserData, error) {
	conn := r.db(ctx)
	var results []orm.UserData
	sql := "SELECT A.user_id, A.name, A.email, (SELECT SUM(Z.quantity * Y.price) FROM user_purchase Z)" +
		" INNER JOIN menu Y ON Z.menu_id = Y.id WHERE Z.user_id = A.user_id) total_price FROM user A"

	if err := conn.Raw(sql).Find(&results).Error; err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return results, nil
	}

	for _, cf := range graphql.CollectFieldsCtx(ctx, nil) {
		if cf.Name == "purchases" {
			var dests []orm.UserPurchaseData
			sql := "SELECT A.user_id, A.quantity, B.category, B.name B.price, (A.quantity * B.price) sub_total" +
				" FROM user_purchase A INNER JOIN menu B ON A.menu_id = B.id"

			if err := conn.Raw(sql).Find(&dests).Error; err != nil {
				return nil, err
			}

			if len(dests) != 0 {
				m := make(map[string]*orm.UserData, len(results))
				for i := 0; i < len(results); i++ {
					m[results[i].UserId] = &results[i]
				}
				for _, d := range dests {
					m[d.UserId].Purchases = append(m[d.UserId].Purchases, d)
				}
			}
		}
	}
	return results, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *autogen.NewUserInput) (*string, error) {
	conn := r.db(ctx)
	userId := uuid.New().String()
	now := time.Now().UTC()
	if err := conn.Transaction(func(tx *gorm.DB) error {
		email := strings.ToLower(input.Email)
		if err := tx.Create(&orm.User{UserId: userId, Name: input.Name, Email: email, CreatedAt: now}).Error; err != nil {
			return err
		}
		if len(input.Purchases) != 0 {
			for _, p := range input.Purchases {
				if err := tx.Create(&orm.UserPurchase{UserId: userId, MenuId: p.MenuID, Quantity: p.Quantity, CreatedAt: now}).Error; err != nil {
					return nil
				}
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &userId, nil
}

func (r *mutationResolver) AppendPurchase(ctx context.Context, input autogen.AppendPurchaseInput) (bool, error) {
	conn := r.db(ctx)
	now := time.Now().UTC()
	if err := conn.Transaction(func(tx *gorm.DB) error {
		for _, p := range input.Purchases {
			if err := tx.Create(&orm.UserPurchase{UserId: input.UserID, MenuId: p.MenuID, Quantity: p.Quantity, CreatedAt: now}).Error; err != nil {
				return nil
			}
		}
		return nil
	}); err != nil {
		return false, err
	}

	return true, nil
}
