/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.21
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SetActiveProjectInConsoleBody Set active project in the Ory Network Console Request Body
type SetActiveProjectInConsoleBody struct {
	// Project ID  The Project ID you want to set active.  format: uuid
	ProjectId string `json:"project_id"`
}

// NewSetActiveProjectInConsoleBody instantiates a new SetActiveProjectInConsoleBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetActiveProjectInConsoleBody(projectId string) *SetActiveProjectInConsoleBody {
	this := SetActiveProjectInConsoleBody{}
	this.ProjectId = projectId
	return &this
}

// NewSetActiveProjectInConsoleBodyWithDefaults instantiates a new SetActiveProjectInConsoleBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetActiveProjectInConsoleBodyWithDefaults() *SetActiveProjectInConsoleBody {
	this := SetActiveProjectInConsoleBody{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *SetActiveProjectInConsoleBody) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SetActiveProjectInConsoleBody) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SetActiveProjectInConsoleBody) SetProjectId(v string) {
	o.ProjectId = v
}

func (o SetActiveProjectInConsoleBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project_id"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableSetActiveProjectInConsoleBody struct {
	value *SetActiveProjectInConsoleBody
	isSet bool
}

func (v NullableSetActiveProjectInConsoleBody) Get() *SetActiveProjectInConsoleBody {
	return v.value
}

func (v *NullableSetActiveProjectInConsoleBody) Set(val *SetActiveProjectInConsoleBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetActiveProjectInConsoleBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetActiveProjectInConsoleBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetActiveProjectInConsoleBody(val *SetActiveProjectInConsoleBody) *NullableSetActiveProjectInConsoleBody {
	return &NullableSetActiveProjectInConsoleBody{value: val, isSet: true}
}

func (v NullableSetActiveProjectInConsoleBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetActiveProjectInConsoleBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

