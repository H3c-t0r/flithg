/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Encapsulates a resource name and value.
type ResourcesResourceEntry struct {
	// Resource name.
	Name  *ResourcesResourceName `json:"name,omitempty"`
	Value string                 `json:"value,omitempty"`
}
