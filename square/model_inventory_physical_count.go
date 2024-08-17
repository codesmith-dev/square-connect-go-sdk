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

// Represents the quantity of an item variation that is physically present at a specific location, verified by a seller or a seller's employee. For example, a physical count might come from an employee counting the item variations on hand or from syncing with an external system.
type InventoryPhysicalCount struct {
	// A unique Square-generated ID for the [InventoryPhysicalCount](https://developer.squareup.com/reference/square_2024-07-17/objects/InventoryPhysicalCount).
	Id string `json:"id,omitempty"`
	// An optional ID provided by the application to tie the [InventoryPhysicalCount](https://developer.squareup.com/reference/square_2024-07-17/objects/InventoryPhysicalCount) to an external system.
	ReferenceId string `json:"reference_id,omitempty"`
	// The Square-generated ID of the [CatalogObject](https://developer.squareup.com/reference/square_2024-07-17/objects/CatalogObject) being tracked.
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The [type](https://developer.squareup.com/reference/square_2024-07-17/enums/CatalogObjectType) of the [CatalogObject](https://developer.squareup.com/reference/square_2024-07-17/objects/CatalogObject) being tracked.   The Inventory API supports setting and reading the `\"catalog_object_type\": \"ITEM_VARIATION\"` field value.  In addition, it can also read the `\"catalog_object_type\": \"ITEM\"` field value that is set by the Square Restaurants app.
	CatalogObjectType string `json:"catalog_object_type,omitempty"`
	// The current [inventory state](https://developer.squareup.com/reference/square_2024-07-17/enums/InventoryState) for the related quantity of items.
	State string `json:"state,omitempty"`
	// The Square-generated ID of the [Location](https://developer.squareup.com/reference/square_2024-07-17/objects/Location) where the related quantity of items is being tracked.
	LocationId string `json:"location_id,omitempty"`
	// The number of items affected by the physical count as a decimal string. The number can support up to 5 digits after the decimal point.
	Quantity string             `json:"quantity,omitempty"`
	Source   *SourceApplication `json:"source,omitempty"`
	// The Square-generated ID of the [Employee](https://developer.squareup.com/reference/square_2024-07-17/objects/Employee) responsible for the physical count.
	EmployeeId string `json:"employee_id,omitempty"`
	// The Square-generated ID of the [Team Member](https://developer.squareup.com/reference/square_2024-07-17/objects/TeamMember) responsible for the physical count.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// A client-generated RFC 3339-formatted timestamp that indicates when the physical count was examined. For physical count updates, the `occurred_at` timestamp cannot be older than 24 hours or in the future relative to the time of the request.
	OccurredAt string `json:"occurred_at,omitempty"`
	// An RFC 3339-formatted timestamp that indicates when the physical count is received.
	CreatedAt string `json:"created_at,omitempty"`
}
