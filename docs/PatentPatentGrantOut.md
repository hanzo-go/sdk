# PatentPatentGrantOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | Pointer to **int64** | Updated is how many records the grant moved to the visibility asked for. It counts every RETAINED version of the target, so it is at least one. | [optional] 

## Methods

### NewPatentPatentGrantOut

`func NewPatentPatentGrantOut() *PatentPatentGrantOut`

NewPatentPatentGrantOut instantiates a new PatentPatentGrantOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentGrantOutWithDefaults

`func NewPatentPatentGrantOutWithDefaults() *PatentPatentGrantOut`

NewPatentPatentGrantOutWithDefaults instantiates a new PatentPatentGrantOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *PatentPatentGrantOut) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PatentPatentGrantOut) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PatentPatentGrantOut) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PatentPatentGrantOut) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


