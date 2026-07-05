# NexusActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldCount** | Pointer to **map[string]int64** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 

## Methods

### NewNexusActivity

`func NewNexusActivity() *NexusActivity`

NewNexusActivity instantiates a new NexusActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusActivityWithDefaults

`func NewNexusActivityWithDefaults() *NexusActivity`

NewNexusActivityWithDefaults instantiates a new NexusActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldCount

`func (o *NexusActivity) GetFieldCount() map[string]int64`

GetFieldCount returns the FieldCount field if non-nil, zero value otherwise.

### GetFieldCountOk

`func (o *NexusActivity) GetFieldCountOk() (*map[string]int64, bool)`

GetFieldCountOk returns a tuple with the FieldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCount

`func (o *NexusActivity) SetFieldCount(v map[string]int64)`

SetFieldCount sets FieldCount field to given value.

### HasFieldCount

`func (o *NexusActivity) HasFieldCount() bool`

HasFieldCount returns a boolean if a field has been set.

### GetDate

`func (o *NexusActivity) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NexusActivity) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NexusActivity) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *NexusActivity) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


