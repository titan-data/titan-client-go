/*
 * Titan API
 *
 * API used by the Titan CLI
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package titan-client
// Volume struct for Volume
type Volume struct {
	// Volume name
	Name string `json:"name"`
	// Client-specific properties
	Properties map[string]interface{} `json:"properties"`
	// Server-generated configuration
	Config map[string]interface{} `json:"config,omitempty"`
}