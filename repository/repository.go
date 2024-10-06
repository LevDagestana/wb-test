package repository

import (
	"wb/cache"
	"wb/db"
	"wb/logger"
	"wb/models"
)

func InsertOrder(order models.Order) {
	var err error
	query := `
            INSERT INTO orders (
                order_uid, track_number, entry, locale, internal_signature, customer_id, 
                delivery_service, shardkey, sm_id, date_created, oof_shard
            ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        `
	_, err = db.Db.Exec(query,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale,
		order.InternalSignature, order.CustomerID, order.DeliveryService,
		order.ShardKey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		logger.Log.WithError(err).Error("Ошибка при вставке данных:")
	}

	deliveryQuery := `
            INSERT INTO delivery (id, order_uid, name, phone, zip, city, address, region, email)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = db.Db.Exec(deliveryQuery,
		1, order.OrderUID, order.Delivery.Name, order.Delivery.Phone,
		order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		logger.Log.WithError(err).Error("Ошибка при вставке данных:")
	}

	paymentQuery := `
            INSERT INTO payment (id, order_uid,transaction,request_id,currency,provider,amount,payment_dt,bank,delivery_cost,goods_total,custom_fee )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10,$11,$12)`
	_, err = db.Db.Exec(paymentQuery,
		1, order.OrderUID, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency,
		order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		logger.Log.WithError(err).Error("Ошибка при вставке данных:")
	}
	itemsQuery := `
            INSERT INTO items (id, order_uid,chrt_id,track_number,price,rid,name,sale,size,total_price,nm_id,brand,status )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10,$11,$12,$13)`
	item := order.Items[0]
	_, err = db.Db.Exec(itemsQuery,
		1, order.OrderUID, item.ChrtID, item.TrackNumber, item.Price,
		item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
	if err != nil {
		logger.Log.WithError(err).Error("Ошибка при вставке данных:")
	}
	cache.Cache.SetCache(order)

}

func LoadCache() {
	rows, err := db.Db.Query(`SELECT 
        o.order_uid,
        o.track_number,
        o.entry,
        o.locale,
        o.internal_signature,
        o.customer_id,
        o.delivery_service,
        o.shardkey,
        o.sm_id,
        o.date_created,
        o.oof_shard,
        d.name AS delivery_name,
        d.phone AS delivery_phone,
        d.zip AS delivery_zip,
        d.city AS delivery_city,
        d.address AS delivery_address,
        d.region AS delivery_region,
        d.email AS delivery_email,
        p.transaction AS payment_transaction,
        p.currency AS payment_currency,
        p.provider AS payment_provider,
        p.amount AS payment_amount,
        p.payment_dt AS payment_date,
        p.bank AS payment_bank,
        p.delivery_cost AS payment_delivery_cost,
        p.goods_total AS payment_goods_total,
        p.custom_fee AS payment_custom_fee,
        i.chrt_id AS item_chrt_id,
        i.track_number AS item_track_number,
        i.price AS item_price,
        i.rid AS item_rid,
        i.name AS item_name,
        i.sale AS item_sale,
        i.size AS item_size,
        i.total_price AS item_total_price,
        i.nm_id AS item_nm_id,
        i.brand AS item_brand,
        i.status AS item_status
    FROM orders o
    LEFT JOIN delivery d ON o.order_uid = d.order_uid
    LEFT JOIN payment p ON o.order_uid = p.order_uid
    LEFT JOIN items i ON o.order_uid = i.order_uid`)
	if err != nil {
		logger.Log.WithError(err).Fatal("Ошибка загрузки кеша")
	}
	defer rows.Close()

	orderMap := make(map[string]*models.Order)

	for rows.Next() {
		var order models.Order
		var delivery models.Delivery
		var payment models.Payment
		var item models.Item

		err := rows.Scan(
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerID,
			&order.DeliveryService,
			&order.ShardKey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard,
			&delivery.Name,
			&delivery.Phone,
			&delivery.Zip,
			&delivery.City,
			&delivery.Address,
			&delivery.Region,
			&delivery.Email,
			&payment.Transaction,
			&payment.Currency,
			&payment.Provider,
			&payment.Amount,
			&payment.PaymentDT,
			&payment.Bank,
			&payment.DeliveryCost,
			&payment.GoodsTotal,
			&payment.CustomFee,
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmID,
			&item.Brand,
			&item.Status,
		)
		if err != nil {
			logger.Log.WithError(err).Fatal("Ошибка сканирования")
		}

		if existingOrder, exists := orderMap[order.OrderUID]; exists {
			if item.ChrtID != 0 {
				existingOrder.Items = append(existingOrder.Items, item)
			}
		} else {
			newOrder := models.Order{
				OrderUID:          order.OrderUID,
				TrackNumber:       order.TrackNumber,
				Entry:             order.Entry,
				Locale:            order.Locale,
				InternalSignature: order.InternalSignature,
				CustomerID:        order.CustomerID,
				DeliveryService:   order.DeliveryService,
				ShardKey:          order.ShardKey,
				SmID:              order.SmID,
				DateCreated:       order.DateCreated,
				OofShard:          order.OofShard,
				Delivery:          delivery,
				Payment:           payment,
				Items:             []models.Item{},
			}
			orderMap[order.OrderUID] = &newOrder

			if item.ChrtID != 0 {
				orderMap[order.OrderUID].Items = append(orderMap[order.OrderUID].Items, item)
			}
		}
	}

	if err := rows.Err(); err != nil {
		logger.Log.WithError(err).Fatal("Ошибка чтения записи из бд")
	}
	for _, order := range orderMap {
		cache.Cache.SetCache(*order)
	}
}
