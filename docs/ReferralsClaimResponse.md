# ReferralsClaimResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ReferralsReferralStatus**](ReferralsReferralStatus.md) |  | [optional] 
**Created** | Pointer to **bool** | True when a new referral was created (201); false on idempotent replay (200). | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp. | [optional] 

## Methods

### NewReferralsClaimResponse

`func NewReferralsClaimResponse() *ReferralsClaimResponse`

NewReferralsClaimResponse instantiates a new ReferralsClaimResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralsClaimResponseWithDefaults

`func NewReferralsClaimResponseWithDefaults() *ReferralsClaimResponse`

NewReferralsClaimResponseWithDefaults instantiates a new ReferralsClaimResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReferralsClaimResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReferralsClaimResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReferralsClaimResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReferralsClaimResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ReferralsClaimResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReferralsClaimResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReferralsClaimResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ReferralsClaimResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *ReferralsClaimResponse) GetStatus() ReferralsReferralStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReferralsClaimResponse) GetStatusOk() (*ReferralsReferralStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReferralsClaimResponse) SetStatus(v ReferralsReferralStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReferralsClaimResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *ReferralsClaimResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ReferralsClaimResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ReferralsClaimResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ReferralsClaimResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReferralsClaimResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReferralsClaimResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReferralsClaimResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReferralsClaimResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


