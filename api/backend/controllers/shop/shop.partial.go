package shop

import (
	"github.com/khorasany/coffee/api/backend/models"
	"strconv"
	"time"
)

func shopMapToMapParmaToModel(param map[string]string) models.Shop {
	shopID, _ := strconv.Atoi(param["shop_id"])
	ownerID, _ := strconv.Atoi(param["owner_id"])
	catID, _ := strconv.Atoi(param["cat_id"])
	shopStatus, _ := strconv.Atoi(param["shop_status"])
	numberOfEmployers, _ := strconv.Atoi(param["shop_number_of_employers"])
	return models.Shop{
		ID:       int64(shopID),
		ShopName: param["shop_name"],
		OwnerID:  int64(ownerID),
		CatID:    int64(catID),
		Slug:     param["shop_slug"],
		Status:   int64(shopStatus),
		Meta: models.ShopMeta{
			TelePhone:         param["shop_phone"],
			Email:             param["shop_email"],
			Address:           param["shop_address"],
			Location:          param["shop_location"],
			NumberOfEmployers: int64(numberOfEmployers),
		},
	}
}

func productMapToMapParamToModel(param map[string]string) models.Product {
	productID, _ := strconv.Atoi(param["product_id"])
	shopID, _ := strconv.Atoi(param["shop_id"])
	status, _ := strconv.Atoi(param["status"])
	return models.Product{
		ID:          int64(productID),
		ProductName: param["product_name"],
		ShopID:      int64(shopID),
		Status:      int64(status),
		Meta:        models.ProductMeta{},
	}
}

func orderMapToMapParamToModel(param map[string]string) models.Order {
	orderID, _ := strconv.Atoi(param["order_id"])
	customerID, _ := strconv.Atoi(param["customer_id"])
	return models.Order{
		ID:              int64(orderID),
		CustomerID:      int64(customerID),
		Products:        param["products"],
		RefID:           param["ref_id"],
		OrderToken:      param["order_token"],
		Authority:       param["authority"],
		TransportPrice:  param["transport_price"],
		TotalPrice:      param["total_price"],
		DiscountAmount:  param["discount_amount"],
		DiscountCoupon:  param["discount_coupon"],
		WalletReduction: param["wallet_reduction"],
		CreatedAt:       time.Now(),
		Description:     param["description"],
		Status:          param["order_status"],
	}

}

func cardMapToMapParamToModel(param map[string][]map[string]string) []*models.Card {
	cards := []*models.Card{}
	for _, item := range param["card"] {
		card := models.Card{}
		card.ProductID = item["product_id"]
		card.ProductName = item["product_name"]
		card.Price = item["price"]
		card.Quantity = item["qty"]
		card.Discount = item["discount"]
		card.TotalPrice = item["total_price"]
		cards = append(cards, &card)
	}

	return cards
}
