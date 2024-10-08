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

// The filtering criteria. If the request specifies multiple filters,  the endpoint uses a logical AND to evaluate them.
type LoyaltyEventFilter struct {
	LoyaltyAccountFilter *LoyaltyEventLoyaltyAccountFilter `json:"loyalty_account_filter,omitempty"`
	TypeFilter           *LoyaltyEventTypeFilter           `json:"type_filter,omitempty"`
	DateTimeFilter       *LoyaltyEventDateTimeFilter       `json:"date_time_filter,omitempty"`
	LocationFilter       *LoyaltyEventLocationFilter       `json:"location_filter,omitempty"`
	OrderFilter          *LoyaltyEventOrderFilter          `json:"order_filter,omitempty"`
}
