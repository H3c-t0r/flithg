/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Runtime information. This is loosely defined to allow for extensibility.
type CoreRuntimeMetadata struct {
	// Type of runtime.
	Type_ *RuntimeMetadataRuntimeType `json:"type,omitempty"`
	// Version of the runtime. All versions should be backward compatible. However, certain cases call for version checks to ensure tighter validation or setting expectations.
	Version string `json:"version,omitempty"`
	// +optional It can be used to provide extra information about the runtime (e.g. python, golang... etc.).
	Flavor string `json:"flavor,omitempty"`
	// +optional It can be used to decide use sync plugin or async plugin during runtime.
	IsSyncPlugin bool `json:"is_sync_plugin,omitempty"`
}
