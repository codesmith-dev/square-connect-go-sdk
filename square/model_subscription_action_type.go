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

// SubscriptionActionType : Supported types of an action as a pending change to a subscription.
type SubscriptionActionType string

// List of SubscriptionActionType
const (
	CANCEL_SubscriptionActionType                     SubscriptionActionType = "CANCEL"
	PAUSE_SubscriptionActionType                      SubscriptionActionType = "PAUSE"
	RESUME_SubscriptionActionType                     SubscriptionActionType = "RESUME"
	SWAP_PLAN_SubscriptionActionType                  SubscriptionActionType = "SWAP_PLAN"
	CHANGE_BILLING_ANCHOR_DATE_SubscriptionActionType SubscriptionActionType = "CHANGE_BILLING_ANCHOR_DATE"
)
