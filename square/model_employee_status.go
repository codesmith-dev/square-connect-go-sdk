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

// EmployeeStatus : The status of the Employee being retrieved.  DEPRECATED at version 2020-08-26. Replaced by [TeamMemberStatus](https://developer.squareup.com/reference/square_2024-07-17/objects/TeamMemberStatus).
type EmployeeStatus string

// List of EmployeeStatus
const (
	ACTIVE_EmployeeStatus   EmployeeStatus = "ACTIVE"
	INACTIVE_EmployeeStatus EmployeeStatus = "INACTIVE"
)
