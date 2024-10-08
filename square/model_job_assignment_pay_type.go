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

// JobAssignmentPayType : Enumerates the possible pay types that a job can be assigned.
type JobAssignmentPayType string

// List of JobAssignmentPayType
const (
	NONE_JobAssignmentPayType   JobAssignmentPayType = "NONE"
	HOURLY_JobAssignmentPayType JobAssignmentPayType = "HOURLY"
	SALARY_JobAssignmentPayType JobAssignmentPayType = "SALARY"
)
