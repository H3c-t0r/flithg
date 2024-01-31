/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Identifier of launch plan for which to change state. +required.
type IdentifierOfLaunchPlanForWhichToChangeStateRequired_ struct {
	// Identifies the specific type of resource that this identifier corresponds to.
	ResourceType *CoreResourceType `json:"resource_type,omitempty"`
	// Optional, org key applied to the resource.
	Org string `json:"org,omitempty"`
}
