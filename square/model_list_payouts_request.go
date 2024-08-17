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

// A request to retrieve payout records.
type ListPayoutsRequest struct {
	// The ID of the location for which to list the payouts. By default, payouts are returned for the default (main) location associated with the seller.
	LocationId string `json:"location_id,omitempty"`
	// If provided, only payouts with the given status are returned.
	Status string `json:"status,omitempty"`
	// The timestamp for the beginning of the payout creation time, in RFC 3339 format. Inclusive. Default: The current time minus one year.
	BeginTime string `json:"begin_time,omitempty"`
	// The timestamp for the end of the payout creation time, in RFC 3339 format. Default: The current time.
	EndTime string `json:"end_time,omitempty"`
	// The order in which payouts are listed.
	SortOrder string `json:"sort_order,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). If request parameters change between requests, subsequent results may contain duplicates or missing records.
	Cursor string `json:"cursor,omitempty"`
	// The maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page. The default value of 100 is also the maximum allowed value. If the provided value is greater than 100, it is ignored and the default value is used instead. Default: `100`
	Limit int32 `json:"limit,omitempty"`
}
