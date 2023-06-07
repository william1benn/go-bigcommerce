package bigcommerce

import (
	"testing"
)

func TestOrderCoupon_TypeName(t *testing.T) {
	coupon := OrderCoupon{
		ID:       1,
		CouponID: 123,
		OrderID:  456,
		Code:     "ABC123",
		Amount:   10,
		Type:     2,
		Discount: 5.0,
	}

	typeName := coupon.TypeName()
	expectedTypeName := "per_total_discount"

	if typeName != expectedTypeName {
		t.Errorf("Expected type name: %s, but got: %s", expectedTypeName, typeName)
	}
}

func TestOrderCoupon_TypeName_Unknown(t *testing.T) {
	coupon := OrderCoupon{
		ID:       2,
		CouponID: 789,
		OrderID:  987,
		Code:     "XYZ789",
		Amount:   20,
		Type:     10, // Unknown type
		Discount: 0.0,
	}

	typeName := coupon.TypeName()
	expectedTypeName := "unknown"

	if typeName != expectedTypeName {
		t.Errorf("Expected type name: %s, but got: %s", expectedTypeName, typeName)
	}
}
