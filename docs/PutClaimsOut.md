# PutClaimsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recorded** | Pointer to **int64** | Recorded is how many rows were written. | [optional] 
**Rejected** | Pointer to **[]string** | Rejected names the rows that were not, and why. | [optional] 

## Methods

### NewPutClaimsOut

`func NewPutClaimsOut() *PutClaimsOut`

NewPutClaimsOut instantiates a new PutClaimsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutClaimsOutWithDefaults

`func NewPutClaimsOutWithDefaults() *PutClaimsOut`

NewPutClaimsOutWithDefaults instantiates a new PutClaimsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecorded

`func (o *PutClaimsOut) GetRecorded() int64`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *PutClaimsOut) GetRecordedOk() (*int64, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *PutClaimsOut) SetRecorded(v int64)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *PutClaimsOut) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### GetRejected

`func (o *PutClaimsOut) GetRejected() []string`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *PutClaimsOut) GetRejectedOk() (*[]string, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *PutClaimsOut) SetRejected(v []string)`

SetRejected sets Rejected field to given value.

### HasRejected

`func (o *PutClaimsOut) HasRejected() bool`

HasRejected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


