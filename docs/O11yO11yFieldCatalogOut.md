# O11yO11yFieldCatalogOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interesting** | Pointer to [**[]O11yO11yTelemetryField**](O11yO11yTelemetryField.md) | Interesting are fields seen in the data that could be selected. | [optional] 
**Selected** | Pointer to [**[]O11yO11yTelemetryField**](O11yO11yTelemetryField.md) | Selected are the fields materialized as their own columns. | [optional] 

## Methods

### NewO11yO11yFieldCatalogOut

`func NewO11yO11yFieldCatalogOut() *O11yO11yFieldCatalogOut`

NewO11yO11yFieldCatalogOut instantiates a new O11yO11yFieldCatalogOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFieldCatalogOutWithDefaults

`func NewO11yO11yFieldCatalogOutWithDefaults() *O11yO11yFieldCatalogOut`

NewO11yO11yFieldCatalogOutWithDefaults instantiates a new O11yO11yFieldCatalogOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteresting

`func (o *O11yO11yFieldCatalogOut) GetInteresting() []O11yO11yTelemetryField`

GetInteresting returns the Interesting field if non-nil, zero value otherwise.

### GetInterestingOk

`func (o *O11yO11yFieldCatalogOut) GetInterestingOk() (*[]O11yO11yTelemetryField, bool)`

GetInterestingOk returns a tuple with the Interesting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteresting

`func (o *O11yO11yFieldCatalogOut) SetInteresting(v []O11yO11yTelemetryField)`

SetInteresting sets Interesting field to given value.

### HasInteresting

`func (o *O11yO11yFieldCatalogOut) HasInteresting() bool`

HasInteresting returns a boolean if a field has been set.

### GetSelected

`func (o *O11yO11yFieldCatalogOut) GetSelected() []O11yO11yTelemetryField`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *O11yO11yFieldCatalogOut) GetSelectedOk() (*[]O11yO11yTelemetryField, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *O11yO11yFieldCatalogOut) SetSelected(v []O11yO11yTelemetryField)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *O11yO11yFieldCatalogOut) HasSelected() bool`

HasSelected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


