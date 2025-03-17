package dao

import (
	"context"

	"gorm.io/gorm"
)
type FavoritesDao struct {
	*gorm.DB
}
func NewFavoritesDao(ctx context.Context) *FavoritesDao {
	return &FavoritesDao{NewDBclient(ctx)}
}