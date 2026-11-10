# EmailClassificationCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disposable** | Pointer to **bool** | Disposable reports whether the domain belongs to a temporary or throwaway provider. | [optional] 
**Free** | Pointer to **bool** | Free reports whether the domain belongs to a free consumer webmail provider. | [optional] 
**Reason** | Pointer to **string** | Reason describes the failure or skip reason when Status is not pass. | [optional] 
**Role** | Pointer to **bool** | Role reports whether the address is a generic role or distribution account. | [optional] 
**Status** | Pointer to **string** | Status reports the check outcome: pass, fail, unknown, or skipped. | [optional] 

## Methods

### NewEmailClassificationCheck

`func NewEmailClassificationCheck() *EmailClassificationCheck`

NewEmailClassificationCheck instantiates a new EmailClassificationCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailClassificationCheckWithDefaults

`func NewEmailClassificationCheckWithDefaults() *EmailClassificationCheck`

NewEmailClassificationCheckWithDefaults instantiates a new EmailClassificationCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisposable

`func (o *EmailClassificationCheck) GetDisposable() bool`

GetDisposable returns the Disposable field if non-nil, zero value otherwise.

### GetDisposableOk

`func (o *EmailClassificationCheck) GetDisposableOk() (*bool, bool)`

GetDisposableOk returns a tuple with the Disposable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposable

`func (o *EmailClassificationCheck) SetDisposable(v bool)`

SetDisposable sets Disposable field to given value.

### HasDisposable

`func (o *EmailClassificationCheck) HasDisposable() bool`

HasDisposable returns a boolean if a field has been set.

### GetFree

`func (o *EmailClassificationCheck) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *EmailClassificationCheck) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *EmailClassificationCheck) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *EmailClassificationCheck) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetReason

`func (o *EmailClassificationCheck) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EmailClassificationCheck) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EmailClassificationCheck) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EmailClassificationCheck) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRole

`func (o *EmailClassificationCheck) GetRole() bool`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EmailClassificationCheck) GetRoleOk() (*bool, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EmailClassificationCheck) SetRole(v bool)`

SetRole sets Role field to given value.

### HasRole

`func (o *EmailClassificationCheck) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *EmailClassificationCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailClassificationCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailClassificationCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailClassificationCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


