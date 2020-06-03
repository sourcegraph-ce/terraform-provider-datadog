/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsBrowserError Error response object for a browser test.
type SyntheticsBrowserError struct {
	// Description of the error.
	Description string `json:"description"`
	// Name of the error.
	Name string `json:"name"`
	// Status Code of the error.
	StatusCode *int64                     `json:"statusCode,omitempty"`
	Type       SyntheticsBrowserErrorType `json:"type"`
}

// NewSyntheticsBrowserError instantiates a new SyntheticsBrowserError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsBrowserError(description string, name string, type_ SyntheticsBrowserErrorType) *SyntheticsBrowserError {
	this := SyntheticsBrowserError{}
	this.Description = description
	this.Name = name
	this.Type = type_
	return &this
}

// NewSyntheticsBrowserErrorWithDefaults instantiates a new SyntheticsBrowserError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsBrowserErrorWithDefaults() *SyntheticsBrowserError {
	this := SyntheticsBrowserError{}
	return &this
}

// GetDescription returns the Description field value
func (o *SyntheticsBrowserError) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserError) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SyntheticsBrowserError) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *SyntheticsBrowserError) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserError) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SyntheticsBrowserError) SetName(v string) {
	o.Name = v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *SyntheticsBrowserError) GetStatusCode() int64 {
	if o == nil || o.StatusCode == nil {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserError) GetStatusCodeOk() (*int64, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *SyntheticsBrowserError) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *SyntheticsBrowserError) SetStatusCode(v int64) {
	o.StatusCode = &v
}

// GetType returns the Type field value
func (o *SyntheticsBrowserError) GetType() SyntheticsBrowserErrorType {
	if o == nil {
		var ret SyntheticsBrowserErrorType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserError) GetTypeOk() (*SyntheticsBrowserErrorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SyntheticsBrowserError) SetType(v SyntheticsBrowserErrorType) {
	o.Type = v
}

func (o SyntheticsBrowserError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.StatusCode != nil {
		toSerialize["statusCode"] = o.StatusCode
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsBrowserError struct {
	value *SyntheticsBrowserError
	isSet bool
}

func (v NullableSyntheticsBrowserError) Get() *SyntheticsBrowserError {
	return v.value
}

func (v *NullableSyntheticsBrowserError) Set(val *SyntheticsBrowserError) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBrowserError) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBrowserError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBrowserError(val *SyntheticsBrowserError) *NullableSyntheticsBrowserError {
	return &NullableSyntheticsBrowserError{value: val, isSet: true}
}

func (v NullableSyntheticsBrowserError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBrowserError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}