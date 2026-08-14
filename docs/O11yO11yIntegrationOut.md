# O11yO11yIntegrationOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yIntegration**](O11yIntegration.md) | Data holds the integration and its installation record. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yIntegrationOut

`func NewO11yO11yIntegrationOut() *O11yO11yIntegrationOut`

NewO11yO11yIntegrationOut instantiates a new O11yO11yIntegrationOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yIntegrationOutWithDefaults

`func NewO11yO11yIntegrationOutWithDefaults() *O11yO11yIntegrationOut`

NewO11yO11yIntegrationOutWithDefaults instantiates a new O11yO11yIntegrationOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yIntegrationOut) GetData() O11yIntegration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yIntegrationOut) GetDataOk() (*O11yIntegration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yIntegrationOut) SetData(v O11yIntegration)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yIntegrationOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yIntegrationOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yIntegrationOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yIntegrationOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yIntegrationOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


