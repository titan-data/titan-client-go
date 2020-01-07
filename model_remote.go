/*
 * Titan API
 *
 * API used by the Titan CLI
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package titanclient
// Remote struct for Remote
type Remote struct {
	// Remote type
	Provider string `json:"provider"`
	// Name of remote
	Name string `json:"name"`
	// Provider-specific remote properties
	Properties map[string]interface{} `json:"properties"`
}
