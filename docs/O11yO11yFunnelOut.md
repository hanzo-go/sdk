# O11yO11yFunnelOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yGettableFunnel**](O11yGettableFunnel.md) | Data is the funnel with its steps. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yFunnelOut

`func NewO11yO11yFunnelOut() *O11yO11yFunnelOut`

NewO11yO11yFunnelOut instantiates a new O11yO11yFunnelOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFunnelOutWithDefaults

`func NewO11yO11yFunnelOutWithDefaults() *O11yO11yFunnelOut`

NewO11yO11yFunnelOutWithDefaults instantiates a new O11yO11yFunnelOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yFunnelOut) GetData() O11yGettableFunnel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yFunnelOut) GetDataOk() (*O11yGettableFunnel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yFunnelOut) SetData(v O11yGettableFunnel)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yFunnelOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yFunnelOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yFunnelOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yFunnelOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yFunnelOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


