package volume // import "github.com/sequix/moby/api/types/volume"

// ----------------------------------------------------------------------------
// Code generated by `swagger generate operation`. DO NOT EDIT.
//
// See hack/generate-swagger-api.sh
// ----------------------------------------------------------------------------

import "github.com/sequix/moby/api/types"

// VolumeListOKBody Volume list response
// swagger:model VolumeListOKBody
type VolumeListOKBody struct {

	// List of volumes
	// Required: true
	Volumes []*types.Volume `json:"Volumes"`

	// Warnings that occurred when fetching the list of volumes
	// Required: true
	Warnings []string `json:"Warnings"`
}
