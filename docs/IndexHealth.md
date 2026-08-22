# IndexHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status is &#x60;available&#x60; when the store is readable and &#x60;unavailable&#x60; when it is not — the second answer rides a 503, so a pod with a broken volume is taken out of rotation rather than serving empty searches. | [optional] 

## Methods

### NewIndexHealth

`func NewIndexHealth() *IndexHealth`

NewIndexHealth instantiates a new IndexHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexHealthWithDefaults

`func NewIndexHealthWithDefaults() *IndexHealth`

NewIndexHealthWithDefaults instantiates a new IndexHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IndexHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IndexHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IndexHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IndexHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


