/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// These URLs are returned as part of node and task execution data requests.
type AdminFlyteUrLs struct {
	Inputs string `json:"inputs,omitempty"`
	Outputs string `json:"outputs,omitempty"`
	Deck string `json:"deck,omitempty"`
}
