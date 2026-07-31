# CloudGrantOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | Pointer to **int32** | Updated is the number of records the grant applied to. Zero is a 404, never a silent no-op. | [optional] 

## Methods

### NewCloudGrantOut

`func NewCloudGrantOut() *CloudGrantOut`

NewCloudGrantOut instantiates a new CloudGrantOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGrantOutWithDefaults

`func NewCloudGrantOutWithDefaults() *CloudGrantOut`

NewCloudGrantOutWithDefaults instantiates a new CloudGrantOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *CloudGrantOut) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudGrantOut) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudGrantOut) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudGrantOut) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


