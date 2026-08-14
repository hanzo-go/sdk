# O11yO11yFieldKeysOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yGettableFieldKeys**](O11yGettableFieldKeys.md) | Data holds the field keys grouped by name, and whether the catalog is complete. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yFieldKeysOut

`func NewO11yO11yFieldKeysOut() *O11yO11yFieldKeysOut`

NewO11yO11yFieldKeysOut instantiates a new O11yO11yFieldKeysOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFieldKeysOutWithDefaults

`func NewO11yO11yFieldKeysOutWithDefaults() *O11yO11yFieldKeysOut`

NewO11yO11yFieldKeysOutWithDefaults instantiates a new O11yO11yFieldKeysOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yFieldKeysOut) GetData() O11yGettableFieldKeys`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yFieldKeysOut) GetDataOk() (*O11yGettableFieldKeys, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yFieldKeysOut) SetData(v O11yGettableFieldKeys)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yFieldKeysOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yFieldKeysOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yFieldKeysOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yFieldKeysOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yFieldKeysOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


