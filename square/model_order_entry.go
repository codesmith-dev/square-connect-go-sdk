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

// A lightweight description of an [order](https://developer.squareup.com/reference/square_2024-07-17/objects/Order) that is returned when `returned_entries` is `true` on a [SearchOrdersRequest](https://developer.squareup.com/reference/square_2024-07-17/orders-api/search-orders).
type OrderEntry struct {
	// The ID of the order.
	OrderId string `json:"order_id,omitempty"`
	// The version number, which is incremented each time an update is committed to the order. Orders that were not created through the API do not include a version number and therefore cannot be updated.  [Read more about working with versions.](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders)
	Version int32 `json:"version,omitempty"`
	// The location ID the order belongs to.
	LocationId string `json:"location_id,omitempty"`
}
