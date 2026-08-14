# O11yO11ySessionContextOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yO11ySessionContext**](O11yO11ySessionContext.md) | Data is the context. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11ySessionContextOut

`func NewO11yO11ySessionContextOut() *O11yO11ySessionContextOut`

NewO11yO11ySessionContextOut instantiates a new O11yO11ySessionContextOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11ySessionContextOutWithDefaults

`func NewO11yO11ySessionContextOutWithDefaults() *O11yO11ySessionContextOut`

NewO11yO11ySessionContextOutWithDefaults instantiates a new O11yO11ySessionContextOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11ySessionContextOut) GetData() O11yO11ySessionContext`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11ySessionContextOut) GetDataOk() (*O11yO11ySessionContext, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11ySessionContextOut) SetData(v O11yO11ySessionContext)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11ySessionContextOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11ySessionContextOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11ySessionContextOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11ySessionContextOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11ySessionContextOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


