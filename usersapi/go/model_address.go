/*
 * Sample API
 *
 * A short description of API.
 *
 * API version: 1.0.0
 * Contact: support@example.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package usersapi

type Address struct {

	Country Country `json:"country,omitempty"`

	ZipCode int32 `json:"zipCode,omitempty"`

	State string `json:"state"`

	City *string `json:"city,omitempty"`
}
