package shop

import (
	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
)

func GetShopByOwnerID(ownerID string) ([]*models.Shop, error) {
	db := database.CreateCon()
	shopsResult, err := db.Query("select * from ico_shop where owner_id=" + ownerID + ";")
	if err != nil {
		return nil, err
	}

	shops := []*models.Shop{}
	for shopsResult.Next() {
		shop := models.Shop{}
		shopsResult.Scan(&shop.ID, &shop.OwnerID, &shop.CatID, &shop.ShopName, &shop.Slug, &shop.Status)
		shops = append(shops, &shop)
	}

	return shops, nil
}

func GetShopByShopID(shopID string) (*models.Shop, error) {
	var shop models.Shop
	db := database.CreateCon()
	err := db.QueryRow("select * from ico_shop where id="+shopID+";").Scan(&shop.ID, &shop.OwnerID, &shop.CatID, &shop.ShopName, &shop.Slug, &shop.Status)

	if err != nil {
		return &shop, nil
	}

	return &shop, nil
}

func GetShops() ([]*models.Shop, error) {
	db := database.CreateCon()
	shopsResult, err := db.Query("select * from ico_shop")
	if err != nil {
		return nil, err
	}

	var shops []*models.Shop
	for shopsResult.Next() {
		var shop models.Shop
		_ = shopsResult.Scan(&shop.ID, &shop.OwnerID, &shop.CatID, &shop.ShopName, &shop.Slug, &shop.Status)
		shops = append(shops, &shop)
	}

	return shops, nil
}
