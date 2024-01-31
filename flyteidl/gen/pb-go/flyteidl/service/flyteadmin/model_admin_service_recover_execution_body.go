/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Request to recover the referenced execution.
type AdminServiceRecoverExecutionBody struct {
	Id *IdentifierOfTheWorkflowExecutionToRecover_ `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// Additional metadata which will be used to overwrite any metadata in the reference execution when triggering a recovery execution.
	Metadata *AdminExecutionMetadata `json:"metadata,omitempty"`
}
