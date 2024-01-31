/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

import (
	"time"
)

// Container for task execution details and results.
type AdminTaskExecutionClosure struct {
	// Path to remote data store where output blob is stored if the execution succeeded (and produced outputs). DEPRECATED. Use GetTaskExecutionData to fetch output data instead.
	OutputUri string `json:"output_uri,omitempty"`
	// Error information for the task execution. Populated if the execution failed.
	Error_ *CoreExecutionError `json:"error,omitempty"`
	// Raw output data produced by this task execution. DEPRECATED. Use GetTaskExecutionData to fetch output data instead.
	OutputData *CoreLiteralMap `json:"output_data,omitempty"`
	// The last recorded phase for this task execution.
	Phase *CoreTaskExecutionPhase `json:"phase,omitempty"`
	// Detailed log information output by the task execution.
	Logs []CoreTaskLog `json:"logs,omitempty"`
	// Time at which the task execution began running.
	StartedAt time.Time `json:"started_at,omitempty"`
	// The amount of time the task execution spent running.
	Duration string `json:"duration,omitempty"`
	// Time at which the task execution was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Time at which the task execution was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Custom data specific to the task plugin.
	CustomInfo interface{} `json:"custom_info,omitempty"`
	// If there is an explanation for the most recent phase transition, the reason will capture it.
	Reason string `json:"reason,omitempty"`
	// A predefined yet extensible Task type identifier.
	TaskType string `json:"task_type,omitempty"`
	// Metadata around how a task was executed.
	Metadata *FlyteidleventTaskExecutionMetadata `json:"metadata,omitempty"`
	// The event version is used to indicate versioned changes in how data is maintained using this proto message. For example, event_verison > 0 means that maps tasks logs use the TaskExecutionMetadata ExternalResourceInfo fields for each subtask rather than the TaskLog in this message.
	EventVersion int32 `json:"event_version,omitempty"`
	// A time-series of the phase transition or update explanations. This, when compared to storing a singular reason as previously done, is much more valuable in visualizing and understanding historical evaluations.
	Reasons []AdminReason `json:"reasons,omitempty"`
}
