/*
Vearch Database API

API for sending dynamic records to the Vearch database.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vearch_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SearchDocumentsRequestRetrievalParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchDocumentsRequestRetrievalParam{}

// SearchDocumentsRequestRetrievalParam struct for SearchDocumentsRequestRetrievalParam
type SearchDocumentsRequestRetrievalParam struct {
	MetricType string `json:"metric_type"`
}

type _SearchDocumentsRequestRetrievalParam SearchDocumentsRequestRetrievalParam

// NewSearchDocumentsRequestRetrievalParam instantiates a new SearchDocumentsRequestRetrievalParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchDocumentsRequestRetrievalParam(metricType string) *SearchDocumentsRequestRetrievalParam {
	this := SearchDocumentsRequestRetrievalParam{}
	this.MetricType = metricType
	return &this
}

// NewSearchDocumentsRequestRetrievalParamWithDefaults instantiates a new SearchDocumentsRequestRetrievalParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchDocumentsRequestRetrievalParamWithDefaults() *SearchDocumentsRequestRetrievalParam {
	this := SearchDocumentsRequestRetrievalParam{}
	return &this
}

// GetMetricType returns the MetricType field value
func (o *SearchDocumentsRequestRetrievalParam) GetMetricType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value
// and a boolean to check if the value has been set.
func (o *SearchDocumentsRequestRetrievalParam) GetMetricTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricType, true
}

// SetMetricType sets field value
func (o *SearchDocumentsRequestRetrievalParam) SetMetricType(v string) {
	o.MetricType = v
}

func (o SearchDocumentsRequestRetrievalParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchDocumentsRequestRetrievalParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metric_type"] = o.MetricType
	return toSerialize, nil
}

func (o *SearchDocumentsRequestRetrievalParam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metric_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSearchDocumentsRequestRetrievalParam := _SearchDocumentsRequestRetrievalParam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchDocumentsRequestRetrievalParam)

	if err != nil {
		return err
	}

	*o = SearchDocumentsRequestRetrievalParam(varSearchDocumentsRequestRetrievalParam)

	return err
}

type NullableSearchDocumentsRequestRetrievalParam struct {
	value *SearchDocumentsRequestRetrievalParam
	isSet bool
}

func (v NullableSearchDocumentsRequestRetrievalParam) Get() *SearchDocumentsRequestRetrievalParam {
	return v.value
}

func (v *NullableSearchDocumentsRequestRetrievalParam) Set(val *SearchDocumentsRequestRetrievalParam) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchDocumentsRequestRetrievalParam) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchDocumentsRequestRetrievalParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchDocumentsRequestRetrievalParam(val *SearchDocumentsRequestRetrievalParam) *NullableSearchDocumentsRequestRetrievalParam {
	return &NullableSearchDocumentsRequestRetrievalParam{value: val, isSet: true}
}

func (v NullableSearchDocumentsRequestRetrievalParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchDocumentsRequestRetrievalParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

