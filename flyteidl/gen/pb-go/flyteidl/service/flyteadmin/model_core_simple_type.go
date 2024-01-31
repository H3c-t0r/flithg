/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin
// CoreSimpleType : Define a set of simple types.
type CoreSimpleType string

// List of coreSimpleType
const (
	NONE_CoreSimpleType CoreSimpleType = "NONE"
	INTEGER_CoreSimpleType CoreSimpleType = "INTEGER"
	FLOAT_CoreSimpleType CoreSimpleType = "FLOAT"
	STRING__CoreSimpleType CoreSimpleType = "STRING"
	BOOLEAN_CoreSimpleType CoreSimpleType = "BOOLEAN"
	DATETIME_CoreSimpleType CoreSimpleType = "DATETIME"
	DURATION_CoreSimpleType CoreSimpleType = "DURATION"
	BINARY_CoreSimpleType CoreSimpleType = "BINARY"
	ERROR__CoreSimpleType CoreSimpleType = "ERROR"
	STRUCT__CoreSimpleType CoreSimpleType = "STRUCT"
)
