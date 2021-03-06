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

import (
	"time"
)

type Birth struct {

	Datetime time.Time `json:"datetime"`

	Weight *int32 `json:"weight,omitempty"`

	Hospital *string `json:"hospital,omitempty"`
}
