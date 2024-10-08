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

type ListEmployeesRequest struct {
	LocationId string `json:"location_id,omitempty"`
	// Specifies the EmployeeStatus to filter the employee by.
	Status string `json:"status,omitempty"`
	// The number of employees to be returned on each page.
	Limit int32 `json:"limit,omitempty"`
	// The token required to retrieve the specified page of results.
	Cursor string `json:"cursor,omitempty"`
}
