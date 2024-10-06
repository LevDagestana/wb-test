package validator

import (
	"errors"
	"fmt"
	"regexp"
	"wb/models"
)

func ValidateOrder(order models.Order) error {
	if order.OrderUID == "" {
		return errors.New("order_uid не может быть пустым")
	}
	if order.TrackNumber == "" {
		return errors.New("track_number не может быть пустым")
	}
	if order.Entry == "" {
		return errors.New("entry не может быть пустым")
	}

	if err := validateDelivery(order.Delivery); err != nil {
		return fmt.Errorf("ошибка в delivery: %v", err)
	}

	if err := validatePayment(order.Payment); err != nil {
		return fmt.Errorf("ошибка в payment: %v", err)
	}

	if len(order.Items) == 0 {
		return errors.New("список items не может быть пустым")
	}
	for _, item := range order.Items {
		if err := validateItem(item); err != nil {
			return fmt.Errorf("ошибка в item: %v", err)
		}
	}

	if order.Locale == "" {
		return errors.New("locale не может быть пустым")
	}
	if order.CustomerID == "" {
		return errors.New("customer_id не может быть пустым")
	}
	if order.DeliveryService == "" {
		return errors.New("delivery_service не может быть пустым")
	}
	if order.ShardKey == "" {
		return errors.New("shardkey не может быть пустым")
	}
	if order.SmID == 0 {
		return errors.New("sm_id не может быть равен 0")
	}

	if order.OofShard == "" {
		return errors.New("oof_shard не может быть пустым")
	}

	return nil
}

func validateDelivery(delivery models.Delivery) error {
	if delivery.Name == "" {
		return errors.New("имя не может быть пустым")
	}
	if !validatePhoneNumber(delivery.Phone) {
		return errors.New("некорректный номер телефона")
	}
	if delivery.Zip == "" {
		return errors.New("zip не может быть пустым")
	}
	if delivery.City == "" {
		return errors.New("город не может быть пустым")
	}
	if delivery.Address == "" {
		return errors.New("адрес не может быть пустым")
	}
	if delivery.Region == "" {
		return errors.New("регион не может быть пустым")
	}
	if !validateEmail(delivery.Email) {
		return errors.New("некорректный email")
	}
	return nil
}

func validatePayment(payment models.Payment) error {
	if payment.Transaction == "" {
		return errors.New("transaction не может быть пустым")
	}
	if payment.Currency == "" {
		return errors.New("currency не может быть пустым")
	}
	if payment.Provider == "" {
		return errors.New("provider не может быть пустым")
	}
	if payment.Amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}

	if payment.Bank == "" {
		return errors.New("bank не может быть пустым")
	}
	return nil
}

func validateItem(item models.Item) error {
	if item.ChrtID == 0 {
		return errors.New("chrt_id не может быть 0")
	}
	if item.TrackNumber == "" {
		return errors.New("track_number не может быть пустым")
	}
	if item.Price <= 0 {
		return errors.New("цена должна быть положительной")
	}
	if item.Rid == "" {
		return errors.New("rid не может быть пустым")
	}
	if item.Name == "" {
		return errors.New("name не может быть пустым")
	}
	if item.TotalPrice <= 0 {
		return errors.New("total_price должен быть положительным")
	}
	if item.Brand == "" {
		return errors.New("brand не может быть пустым")
	}
	return nil
}

func validatePhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+\d{1,3}\d{7,}$`)
	return re.MatchString(phone)
}

func validateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
