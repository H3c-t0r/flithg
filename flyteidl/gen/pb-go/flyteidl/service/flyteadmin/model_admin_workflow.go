/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Represents the workflow structure stored in the Admin A workflow is created by ordering tasks and associating outputs to inputs in order to produce a directed-acyclic execution graph.
type AdminWorkflow struct {
	// id represents the unique identifier of the workflow.
	Id *CoreIdentifier `json:"id,omitempty"`
	// closure encapsulates all the fields that maps to a compiled version of the workflow.
	Closure *FlyteidladminWorkflowClosure `json:"closure,omitempty"`
	// One-liner overview of the entity.
	ShortDescription string `json:"short_description,omitempty"`
}
