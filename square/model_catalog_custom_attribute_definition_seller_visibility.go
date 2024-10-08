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

// CatalogCustomAttributeDefinitionSellerVisibility : Defines the visibility of a custom attribute to sellers in Square client applications, Square APIs or in Square UIs (including Square Point of Sale applications and Square Dashboard).
type CatalogCustomAttributeDefinitionSellerVisibility string

// List of CatalogCustomAttributeDefinitionSellerVisibility
const (
	HIDDEN_CatalogCustomAttributeDefinitionSellerVisibility            CatalogCustomAttributeDefinitionSellerVisibility = "SELLER_VISIBILITY_HIDDEN"
	READ_WRITE_VALUES_CatalogCustomAttributeDefinitionSellerVisibility CatalogCustomAttributeDefinitionSellerVisibility = "SELLER_VISIBILITY_READ_WRITE_VALUES"
)
