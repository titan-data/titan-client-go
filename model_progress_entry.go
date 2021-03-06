/*
 * Titan API
 *
 * API used by the Titan CLI
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package titanclient
// ProgressEntry struct for ProgressEntry
type ProgressEntry struct {
	// Sequenced entry identifier
	Id int32 `json:"id"`
	// Entry type
	Type string `json:"type"`
	// Optional message for progress step
	Message string `json:"message,omitempty"`
	// Optional percent for step
	Percent int32 `json:"percent,omitempty"`
}
