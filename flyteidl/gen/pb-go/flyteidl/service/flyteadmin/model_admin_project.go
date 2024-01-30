/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Top-level namespace used to classify different entities like workflows and executions.
type AdminProject struct {
	// Globally unique project name.
	Id string `json:"id,omitempty"`
	// Display name.
	Name string `json:"name,omitempty"`
	Domains []AdminDomain `json:"domains,omitempty"`
	Description string `json:"description,omitempty"`
	// Leverage Labels from flyteidl.admin.common.proto to tag projects with ownership information.
	Labels *AdminLabels `json:"labels,omitempty"`
	State *ProjectProjectState `json:"state,omitempty"`
	// Optional, org key applied to the resource.
	Org string `json:"org,omitempty"`
}
