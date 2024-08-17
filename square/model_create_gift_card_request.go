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

// Represents a [CreateGiftCard](https://developer.squareup.com/reference/square_2024-07-17/gift-cards-api/create-gift-card) request.
type CreateGiftCardRequest struct {
	// A unique identifier for this request, used to ensure idempotency. For more information,  see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey string `json:"idempotency_key"`
	// The ID of the [location](https://developer.squareup.com/reference/square_2024-07-17/objects/Location) where the gift card should be registered for  reporting purposes. Gift cards can be redeemed at any of the seller's locations.
	LocationId string    `json:"location_id"`
	GiftCard   *GiftCard `json:"gift_card"`
}
