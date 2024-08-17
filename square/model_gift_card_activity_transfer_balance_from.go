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

// Represents details about a `TRANSFER_BALANCE_FROM` [gift card activity type](https://developer.squareup.com/reference/square_2024-07-17/objects/GiftCardActivityType).
type GiftCardActivityTransferBalanceFrom struct {
	// The ID of the gift card to which the specified amount was transferred.
	TransferToGiftCardId string `json:"transfer_to_gift_card_id"`
	AmountMoney          *Money `json:"amount_money"`
}
