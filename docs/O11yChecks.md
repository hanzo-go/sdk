# O11yChecks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MissingDefaultEnabledMetrics** | Pointer to [**[]O11yMissingMetricsComponentEntry**](O11yMissingMetricsComponentEntry.md) |  | [optional] 
**MissingOptionalMetrics** | Pointer to [**[]O11yMissingMetricsComponentEntry**](O11yMissingMetricsComponentEntry.md) |  | [optional] 
**MissingRequiredAttributes** | Pointer to [**[]O11yMissingAttributesComponentEntry**](O11yMissingAttributesComponentEntry.md) |  | [optional] 
**PresentDefaultEnabledMetrics** | Pointer to [**[]O11yMetricsComponentEntry**](O11yMetricsComponentEntry.md) |  | [optional] 
**PresentOptionalMetrics** | Pointer to [**[]O11yMetricsComponentEntry**](O11yMetricsComponentEntry.md) |  | [optional] 
**PresentRequiredAttributes** | Pointer to [**[]O11yAttributesComponentEntry**](O11yAttributesComponentEntry.md) |  | [optional] 
**Ready** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewO11yChecks

`func NewO11yChecks() *O11yChecks`

NewO11yChecks instantiates a new O11yChecks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yChecksWithDefaults

`func NewO11yChecksWithDefaults() *O11yChecks`

NewO11yChecksWithDefaults instantiates a new O11yChecks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMissingDefaultEnabledMetrics

`func (o *O11yChecks) GetMissingDefaultEnabledMetrics() []O11yMissingMetricsComponentEntry`

GetMissingDefaultEnabledMetrics returns the MissingDefaultEnabledMetrics field if non-nil, zero value otherwise.

### GetMissingDefaultEnabledMetricsOk

`func (o *O11yChecks) GetMissingDefaultEnabledMetricsOk() (*[]O11yMissingMetricsComponentEntry, bool)`

GetMissingDefaultEnabledMetricsOk returns a tuple with the MissingDefaultEnabledMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingDefaultEnabledMetrics

`func (o *O11yChecks) SetMissingDefaultEnabledMetrics(v []O11yMissingMetricsComponentEntry)`

SetMissingDefaultEnabledMetrics sets MissingDefaultEnabledMetrics field to given value.

### HasMissingDefaultEnabledMetrics

`func (o *O11yChecks) HasMissingDefaultEnabledMetrics() bool`

HasMissingDefaultEnabledMetrics returns a boolean if a field has been set.

### GetMissingOptionalMetrics

`func (o *O11yChecks) GetMissingOptionalMetrics() []O11yMissingMetricsComponentEntry`

GetMissingOptionalMetrics returns the MissingOptionalMetrics field if non-nil, zero value otherwise.

### GetMissingOptionalMetricsOk

`func (o *O11yChecks) GetMissingOptionalMetricsOk() (*[]O11yMissingMetricsComponentEntry, bool)`

GetMissingOptionalMetricsOk returns a tuple with the MissingOptionalMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingOptionalMetrics

`func (o *O11yChecks) SetMissingOptionalMetrics(v []O11yMissingMetricsComponentEntry)`

SetMissingOptionalMetrics sets MissingOptionalMetrics field to given value.

### HasMissingOptionalMetrics

`func (o *O11yChecks) HasMissingOptionalMetrics() bool`

HasMissingOptionalMetrics returns a boolean if a field has been set.

### GetMissingRequiredAttributes

`func (o *O11yChecks) GetMissingRequiredAttributes() []O11yMissingAttributesComponentEntry`

GetMissingRequiredAttributes returns the MissingRequiredAttributes field if non-nil, zero value otherwise.

### GetMissingRequiredAttributesOk

`func (o *O11yChecks) GetMissingRequiredAttributesOk() (*[]O11yMissingAttributesComponentEntry, bool)`

GetMissingRequiredAttributesOk returns a tuple with the MissingRequiredAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingRequiredAttributes

`func (o *O11yChecks) SetMissingRequiredAttributes(v []O11yMissingAttributesComponentEntry)`

SetMissingRequiredAttributes sets MissingRequiredAttributes field to given value.

### HasMissingRequiredAttributes

`func (o *O11yChecks) HasMissingRequiredAttributes() bool`

HasMissingRequiredAttributes returns a boolean if a field has been set.

### GetPresentDefaultEnabledMetrics

`func (o *O11yChecks) GetPresentDefaultEnabledMetrics() []O11yMetricsComponentEntry`

GetPresentDefaultEnabledMetrics returns the PresentDefaultEnabledMetrics field if non-nil, zero value otherwise.

### GetPresentDefaultEnabledMetricsOk

`func (o *O11yChecks) GetPresentDefaultEnabledMetricsOk() (*[]O11yMetricsComponentEntry, bool)`

GetPresentDefaultEnabledMetricsOk returns a tuple with the PresentDefaultEnabledMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentDefaultEnabledMetrics

`func (o *O11yChecks) SetPresentDefaultEnabledMetrics(v []O11yMetricsComponentEntry)`

SetPresentDefaultEnabledMetrics sets PresentDefaultEnabledMetrics field to given value.

### HasPresentDefaultEnabledMetrics

`func (o *O11yChecks) HasPresentDefaultEnabledMetrics() bool`

HasPresentDefaultEnabledMetrics returns a boolean if a field has been set.

### GetPresentOptionalMetrics

`func (o *O11yChecks) GetPresentOptionalMetrics() []O11yMetricsComponentEntry`

GetPresentOptionalMetrics returns the PresentOptionalMetrics field if non-nil, zero value otherwise.

### GetPresentOptionalMetricsOk

`func (o *O11yChecks) GetPresentOptionalMetricsOk() (*[]O11yMetricsComponentEntry, bool)`

GetPresentOptionalMetricsOk returns a tuple with the PresentOptionalMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentOptionalMetrics

`func (o *O11yChecks) SetPresentOptionalMetrics(v []O11yMetricsComponentEntry)`

SetPresentOptionalMetrics sets PresentOptionalMetrics field to given value.

### HasPresentOptionalMetrics

`func (o *O11yChecks) HasPresentOptionalMetrics() bool`

HasPresentOptionalMetrics returns a boolean if a field has been set.

### GetPresentRequiredAttributes

`func (o *O11yChecks) GetPresentRequiredAttributes() []O11yAttributesComponentEntry`

GetPresentRequiredAttributes returns the PresentRequiredAttributes field if non-nil, zero value otherwise.

### GetPresentRequiredAttributesOk

`func (o *O11yChecks) GetPresentRequiredAttributesOk() (*[]O11yAttributesComponentEntry, bool)`

GetPresentRequiredAttributesOk returns a tuple with the PresentRequiredAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentRequiredAttributes

`func (o *O11yChecks) SetPresentRequiredAttributes(v []O11yAttributesComponentEntry)`

SetPresentRequiredAttributes sets PresentRequiredAttributes field to given value.

### HasPresentRequiredAttributes

`func (o *O11yChecks) HasPresentRequiredAttributes() bool`

HasPresentRequiredAttributes returns a boolean if a field has been set.

### GetReady

`func (o *O11yChecks) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *O11yChecks) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *O11yChecks) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *O11yChecks) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetType

`func (o *O11yChecks) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yChecks) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yChecks) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *O11yChecks) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *O11yChecks) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *O11yChecks) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


