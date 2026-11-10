# ShareRotated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ended** | Pointer to **int64** | Ended is how many environments the rotated accounts had enabled. Each is ended with its shares, so a machine that ran a tunnel on an account enables again with the new credential. | [optional] 
**Pending** | Pointer to **string** | Pending says what this call could not rotate, and is empty when it rotated everything. In a held org an admin rotates their own account first and the org&#39;s own after it, so the org&#39;s failing to rotate is said here rather than keeping them from their own; calling again tries it again. | [optional] 

## Methods

### NewShareRotated

`func NewShareRotated() *ShareRotated`

NewShareRotated instantiates a new ShareRotated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareRotatedWithDefaults

`func NewShareRotatedWithDefaults() *ShareRotated`

NewShareRotatedWithDefaults instantiates a new ShareRotated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnded

`func (o *ShareRotated) GetEnded() int64`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *ShareRotated) GetEndedOk() (*int64, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *ShareRotated) SetEnded(v int64)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *ShareRotated) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### GetPending

`func (o *ShareRotated) GetPending() string`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ShareRotated) GetPendingOk() (*string, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ShareRotated) SetPending(v string)`

SetPending sets Pending field to given value.

### HasPending

`func (o *ShareRotated) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


