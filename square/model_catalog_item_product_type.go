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

// CatalogItemProductType : The type of a CatalogItem. Connect V2 only allows the creation of `REGULAR` or `APPOINTMENTS_SERVICE` items.
type CatalogItemProductType string

// List of CatalogItemProductType
const (
	REGULAR_CatalogItemProductType                         CatalogItemProductType = "REGULAR"
	GIFT_CARD_CatalogItemProductType                       CatalogItemProductType = "GIFT_CARD"
	APPOINTMENTS_SERVICE_CatalogItemProductType            CatalogItemProductType = "APPOINTMENTS_SERVICE"
	FOOD_AND_BEV_CatalogItemProductType                    CatalogItemProductType = "FOOD_AND_BEV"
	EVENT_CatalogItemProductType                           CatalogItemProductType = "EVENT"
	DIGITAL_CatalogItemProductType                         CatalogItemProductType = "DIGITAL"
	DONATION_CatalogItemProductType                        CatalogItemProductType = "DONATION"
	LEGACY_SQUARE_ONLINE_SERVICE_CatalogItemProductType    CatalogItemProductType = "LEGACY_SQUARE_ONLINE_SERVICE"
	LEGACY_SQUARE_ONLINE_MEMBERSHIP_CatalogItemProductType CatalogItemProductType = "LEGACY_SQUARE_ONLINE_MEMBERSHIP"
)
