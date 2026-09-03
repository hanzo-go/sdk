# VerificationTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByStatus** | Pointer to **map[string]int64** | ByStatus tallies the org&#39;s verifications by provider-reported status. | [optional] 
**Total** | Pointer to **int64** | Total is the sum over every status. | [optional] 

## Methods

### NewVerificationTally

`func NewVerificationTally() *VerificationTally`

NewVerificationTally instantiates a new VerificationTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationTallyWithDefaults

`func NewVerificationTallyWithDefaults() *VerificationTally`

NewVerificationTallyWithDefaults instantiates a new VerificationTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByStatus

`func (o *VerificationTally) GetByStatus() map[string]int64`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *VerificationTally) GetByStatusOk() (*map[string]int64, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *VerificationTally) SetByStatus(v map[string]int64)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *VerificationTally) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.

### GetTotal

`func (o *VerificationTally) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VerificationTally) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VerificationTally) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *VerificationTally) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


