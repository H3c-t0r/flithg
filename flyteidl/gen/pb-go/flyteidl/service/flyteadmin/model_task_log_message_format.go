/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type TaskLogMessageFormat string

// List of TaskLogMessageFormat
const (
	UNKNOWN_TaskLogMessageFormat TaskLogMessageFormat = "UNKNOWN"
	CSV_TaskLogMessageFormat TaskLogMessageFormat = "CSV"
	JSON_TaskLogMessageFormat TaskLogMessageFormat = "JSON"
)
