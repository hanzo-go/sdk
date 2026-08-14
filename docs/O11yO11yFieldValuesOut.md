# O11yO11yFieldValuesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yGettableFieldValues**](O11yGettableFieldValues.md) | Data holds the values by data type, and whether the value list is complete. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yFieldValuesOut

`func NewO11yO11yFieldValuesOut() *O11yO11yFieldValuesOut`

NewO11yO11yFieldValuesOut instantiates a new O11yO11yFieldValuesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFieldValuesOutWithDefaults

`func NewO11yO11yFieldValuesOutWithDefaults() *O11yO11yFieldValuesOut`

NewO11yO11yFieldValuesOutWithDefaults instantiates a new O11yO11yFieldValuesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yFieldValuesOut) GetData() O11yGettableFieldValues`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yFieldValuesOut) GetDataOk() (*O11yGettableFieldValues, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yFieldValuesOut) SetData(v O11yGettableFieldValues)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yFieldValuesOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yFieldValuesOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yFieldValuesOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yFieldValuesOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yFieldValuesOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


