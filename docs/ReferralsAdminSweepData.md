# ReferralsAdminSweepData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Swept** | Pointer to **int32** | Number of pending referrals examined this sweep. | [optional] 
**Credited** | Pointer to **int32** | Number newly advanced to credited this sweep. | [optional] 

## Methods

### NewReferralsAdminSweepData

`func NewReferralsAdminSweepData() *ReferralsAdminSweepData`

NewReferralsAdminSweepData instantiates a new ReferralsAdminSweepData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralsAdminSweepDataWithDefaults

`func NewReferralsAdminSweepDataWithDefaults() *ReferralsAdminSweepData`

NewReferralsAdminSweepDataWithDefaults instantiates a new ReferralsAdminSweepData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwept

`func (o *ReferralsAdminSweepData) GetSwept() int32`

GetSwept returns the Swept field if non-nil, zero value otherwise.

### GetSweptOk

`func (o *ReferralsAdminSweepData) GetSweptOk() (*int32, bool)`

GetSweptOk returns a tuple with the Swept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwept

`func (o *ReferralsAdminSweepData) SetSwept(v int32)`

SetSwept sets Swept field to given value.

### HasSwept

`func (o *ReferralsAdminSweepData) HasSwept() bool`

HasSwept returns a boolean if a field has been set.

### GetCredited

`func (o *ReferralsAdminSweepData) GetCredited() int32`

GetCredited returns the Credited field if non-nil, zero value otherwise.

### GetCreditedOk

`func (o *ReferralsAdminSweepData) GetCreditedOk() (*int32, bool)`

GetCreditedOk returns a tuple with the Credited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredited

`func (o *ReferralsAdminSweepData) SetCredited(v int32)`

SetCredited sets Credited field to given value.

### HasCredited

`func (o *ReferralsAdminSweepData) HasCredited() bool`

HasCredited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


