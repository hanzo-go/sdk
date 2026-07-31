# CloudClaimView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the referral code the referral was recorded against. | [optional] 
**Created** | Pointer to **bool** | Created is true when this call recorded the referral and false when it found one already recorded for this referee — the idempotent replay. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the referral was first recorded, as a Unix timestamp. | [optional] 
**Id** | Pointer to **string** | ID is the referral&#39;s handle. | [optional] 
**Status** | Pointer to **string** | Status is the referral&#39;s lifecycle state: \&quot;signup\&quot; until the referee makes metered spend, then \&quot;qualified\&quot;, then \&quot;credited\&quot;. | [optional] 

## Methods

### NewCloudClaimView

`func NewCloudClaimView() *CloudClaimView`

NewCloudClaimView instantiates a new CloudClaimView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudClaimViewWithDefaults

`func NewCloudClaimViewWithDefaults() *CloudClaimView`

NewCloudClaimViewWithDefaults instantiates a new CloudClaimView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CloudClaimView) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudClaimView) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudClaimView) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CloudClaimView) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreated

`func (o *CloudClaimView) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudClaimView) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudClaimView) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudClaimView) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudClaimView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudClaimView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudClaimView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudClaimView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudClaimView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudClaimView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudClaimView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudClaimView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *CloudClaimView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudClaimView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudClaimView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudClaimView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


