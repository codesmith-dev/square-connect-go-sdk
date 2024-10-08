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

import "encoding/json"

// Represents a [BulkUpsertMerchantCustomAttributes](https://developer.squareup.com/reference/square_2024-07-17/merchant-custom-attributes-api/bulk-upsert-merchant-custom-attributes) request.
type BulkUpsertMerchantCustomAttributesRequest struct {
	// A map containing 1 to 25 individual upsert requests. For each request, provide an arbitrary ID that is unique for this `BulkUpsertMerchantCustomAttributes` request and the information needed to create or update a custom attribute.
	Values map[string]json.RawMessage `json:"values"`
}
