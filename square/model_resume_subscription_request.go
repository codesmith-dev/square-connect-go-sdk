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

// Defines input parameters in a request to the [ResumeSubscription](https://developer.squareup.com/reference/square_2024-07-17/subscriptions-api/resume-subscription) endpoint.
type ResumeSubscriptionRequest struct {
	// The `YYYY-MM-DD`-formatted date when the subscription reactivated.
	ResumeEffectiveDate string `json:"resume_effective_date,omitempty"`
	// The timing to resume a subscription, relative to the specified `resume_effective_date` attribute value.
	ResumeChangeTiming string `json:"resume_change_timing,omitempty"`
}
