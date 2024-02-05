/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AttributeRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeRule{}

// AttributeRule Authentication policy rules using attributes from the previous authentication source. Each rule is evaluated to determine the next action in the policy.
type AttributeRule struct {
	AttributeSource *SourceTypeIdKey `json:"attributeSource,omitempty" tfsdk:"attribute_source"`
	// The name of the attribute to use in this attribute rule. This field is required if the Attribute Source type is not 'EXPRESSION'.
	AttributeName *string `json:"attributeName,omitempty" tfsdk:"attribute_name"`
	// The condition that will be applied to the attribute's expected value. This field is required if the Attribute Source type is not 'EXPRESSION'.
	Condition *string `json:"condition,omitempty" tfsdk:"condition"`
	// The expected value of this attribute rule. This field is required if the Attribute Source type is not 'EXPRESSION'.
	ExpectedValue *string `json:"expectedValue,omitempty" tfsdk:"expected_value"`
	// The expression of this attribute rule. This field is required if the Attribute Source type is 'EXPRESSION'.
	Expression *string `json:"expression,omitempty" tfsdk:"expression"`
	// The result of this attribute rule.
	Result string `json:"result" tfsdk:"result"`
}

// NewAttributeRule instantiates a new AttributeRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeRule(result string) *AttributeRule {
	this := AttributeRule{}
	this.Result = result
	return &this
}

// NewAttributeRuleWithDefaults instantiates a new AttributeRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeRuleWithDefaults() *AttributeRule {
	this := AttributeRule{}
	return &this
}

// GetAttributeSource returns the AttributeSource field value if set, zero value otherwise.
func (o *AttributeRule) GetAttributeSource() SourceTypeIdKey {
	if o == nil || IsNil(o.AttributeSource) {
		var ret SourceTypeIdKey
		return ret
	}
	return *o.AttributeSource
}

// GetAttributeSourceOk returns a tuple with the AttributeSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeRule) GetAttributeSourceOk() (*SourceTypeIdKey, bool) {
	if o == nil || IsNil(o.AttributeSource) {
		return nil, false
	}
	return o.AttributeSource, true
}

// HasAttributeSource returns a boolean if a field has been set.
func (o *AttributeRule) HasAttributeSource() bool {
	if o != nil && !IsNil(o.AttributeSource) {
		return true
	}

	return false
}

// SetAttributeSource gets a reference to the given SourceTypeIdKey and assigns it to the AttributeSource field.
func (o *AttributeRule) SetAttributeSource(v SourceTypeIdKey) {
	o.AttributeSource = &v
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *AttributeRule) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeRule) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *AttributeRule) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *AttributeRule) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *AttributeRule) GetCondition() string {
	if o == nil || IsNil(o.Condition) {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeRule) GetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *AttributeRule) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *AttributeRule) SetCondition(v string) {
	o.Condition = &v
}

// GetExpectedValue returns the ExpectedValue field value if set, zero value otherwise.
func (o *AttributeRule) GetExpectedValue() string {
	if o == nil || IsNil(o.ExpectedValue) {
		var ret string
		return ret
	}
	return *o.ExpectedValue
}

// GetExpectedValueOk returns a tuple with the ExpectedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeRule) GetExpectedValueOk() (*string, bool) {
	if o == nil || IsNil(o.ExpectedValue) {
		return nil, false
	}
	return o.ExpectedValue, true
}

// HasExpectedValue returns a boolean if a field has been set.
func (o *AttributeRule) HasExpectedValue() bool {
	if o != nil && !IsNil(o.ExpectedValue) {
		return true
	}

	return false
}

// SetExpectedValue gets a reference to the given string and assigns it to the ExpectedValue field.
func (o *AttributeRule) SetExpectedValue(v string) {
	o.ExpectedValue = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *AttributeRule) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeRule) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *AttributeRule) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *AttributeRule) SetExpression(v string) {
	o.Expression = &v
}

// GetResult returns the Result field value
func (o *AttributeRule) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *AttributeRule) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *AttributeRule) SetResult(v string) {
	o.Result = v
}

func (o AttributeRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeSource) {
		toSerialize["attributeSource"] = o.AttributeSource
	}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.ExpectedValue) {
		toSerialize["expectedValue"] = o.ExpectedValue
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableAttributeRule struct {
	value *AttributeRule
	isSet bool
}

func (v NullableAttributeRule) Get() *AttributeRule {
	return v.value
}

func (v *NullableAttributeRule) Set(val *AttributeRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeRule(val *AttributeRule) *NullableAttributeRule {
	return &NullableAttributeRule{value: val, isSet: true}
}

func (v NullableAttributeRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
