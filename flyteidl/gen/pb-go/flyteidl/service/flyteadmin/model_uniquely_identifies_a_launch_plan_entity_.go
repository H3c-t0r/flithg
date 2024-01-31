/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Uniquely identifies a launch plan entity.
type UniquelyIdentifiesALaunchPlanEntity_ struct {
	// Identifies the specific type of resource that this identifier corresponds to.
	ResourceType *CoreResourceType `json:"resource_type,omitempty"`
	// Name of the project the resource belongs to.
	Project string `json:"project,omitempty"`
	// Name of the domain the resource belongs to. A domain can be considered as a subset within a specific project.
	Domain string `json:"domain,omitempty"`
	// User provided value for the resource.
	Name string `json:"name,omitempty"`
	// Specific version of the resource.
	Version string `json:"version,omitempty"`
}
