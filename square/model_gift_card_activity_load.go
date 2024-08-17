/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package square

// Represents details about a `LOAD` [gift card activity type](https://developer.squareup.com/reference/square_2024-07-17/objects/GiftCardActivityType).
type GiftCardActivityLoad struct {
	AmountMoney *Money `json:"amount_money,omitempty"`
	// The ID of the [order](https://developer.squareup.com/reference/square_2024-07-17/objects/Order) that contains the `GIFT_CARD` line item.  Applications that use the Square Orders API to process orders must specify the order ID in the  [CreateGiftCardActivity](https://developer.squareup.com/reference/square_2024-07-17/gift-card-activities-api/create-gift-card-activity) request.
	OrderId string `json:"order_id,omitempty"`
	// The UID of the `GIFT_CARD` line item in the order that represents the additional funds for the gift card.  Applications that use the Square Orders API to process orders must specify the line item UID in the [CreateGiftCardActivity](https://developer.squareup.com/reference/square_2024-07-17/gift-card-activities-api/create-gift-card-activity) request.
	LineItemUid string `json:"line_item_uid,omitempty"`
	// A client-specified ID that associates the gift card activity with an entity in another system.   Applications that use a custom order processing system can use this field to track information related to  an order or payment.
	ReferenceId string `json:"reference_id,omitempty"`
	// The payment instrument IDs used to process the order for the additional funds, such as a credit card ID  or bank account ID.   Applications that use a custom order processing system must specify payment instrument IDs in  the [CreateGiftCardActivity](https://developer.squareup.com/reference/square_2024-07-17/gift-card-activities-api/create-gift-card-activity) request. Square uses this information to perform compliance checks.   For applications that use the Square Orders API to process payments, Square has the necessary  instrument IDs to perform compliance checks.
	BuyerPaymentInstrumentIds []string `json:"buyer_payment_instrument_ids,omitempty"`
}
