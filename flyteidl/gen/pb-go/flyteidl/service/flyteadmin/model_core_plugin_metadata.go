/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type CorePluginMetadata struct {
	// +optional It can be used to decide use sync plugin or async plugin during runtime.
	IsSyncPlugin bool `json:"is_sync_plugin,omitempty"`
}
