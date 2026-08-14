# KycSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the founder the session belongs to. | [optional] 
**Ref** | Pointer to **string** | Ref is the provider&#39;s reference for the session. | [optional] 
**Status** | Pointer to **string** | Status is the session&#39;s status at start, which is always pending. | [optional] 
**VerifyUrl** | Pointer to **string** | VerifyURL is the hosted flow the founder visits; empty for the manual provider. | [optional] 

## Methods

### NewKycSession

`func NewKycSession() *KycSession`

NewKycSession instantiates a new KycSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKycSessionWithDefaults

`func NewKycSessionWithDefaults() *KycSession`

NewKycSessionWithDefaults instantiates a new KycSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *KycSession) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *KycSession) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *KycSession) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *KycSession) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRef

`func (o *KycSession) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *KycSession) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *KycSession) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *KycSession) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetStatus

`func (o *KycSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KycSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KycSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KycSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerifyUrl

`func (o *KycSession) GetVerifyUrl() string`

GetVerifyUrl returns the VerifyUrl field if non-nil, zero value otherwise.

### GetVerifyUrlOk

`func (o *KycSession) GetVerifyUrlOk() (*string, bool)`

GetVerifyUrlOk returns a tuple with the VerifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyUrl

`func (o *KycSession) SetVerifyUrl(v string)`

SetVerifyUrl sets VerifyUrl field to given value.

### HasVerifyUrl

`func (o *KycSession) HasVerifyUrl() bool`

HasVerifyUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


