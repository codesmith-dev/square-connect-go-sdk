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

// InventoryChangeType : Indicates how the inventory change was applied to a tracked product quantity.
type InventoryChangeType string

// List of InventoryChangeType
const (
	PHYSICAL_COUNT_InventoryChangeType InventoryChangeType = "PHYSICAL_COUNT"
	ADJUSTMENT_InventoryChangeType     InventoryChangeType = "ADJUSTMENT"
	TRANSFER_InventoryChangeType       InventoryChangeType = "TRANSFER"
)
