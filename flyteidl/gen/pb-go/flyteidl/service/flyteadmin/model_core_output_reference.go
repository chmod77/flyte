/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// A reference to an output produced by a node. The type can be retrieved -and validated- from the underlying interface of the node.
type CoreOutputReference struct {
	// Node id must exist at the graph layer.
	NodeId string `json:"node_id,omitempty"`
	// Variable name must refer to an output variable for the node.
	Var_     string                 `json:"var,omitempty"`
	AttrPath []CorePromiseAttribute `json:"attr_path,omitempty"`
}