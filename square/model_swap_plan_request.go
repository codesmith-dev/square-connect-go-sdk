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

// Defines input parameters in a call to the [SwapPlan](https://developer.squareup.com/reference/square_2024-07-17/subscriptions-api/swap-plan) endpoint.
type SwapPlanRequest struct {
	// The ID of the new subscription plan variation.  This field is required.
	NewPlanVariationId string `json:"new_plan_variation_id,omitempty"`
	// A list of PhaseInputs, to pass phase-specific information used in the swap.
	Phases []PhaseInput `json:"phases,omitempty"`
}
