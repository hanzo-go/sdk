# EmailRefused

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is the recipient. | [optional] 
**Reason** | Pointer to **string** | Reason is why: unsubscribed, company unsubscribed (the address&#39;s domain opted out of the org&#39;s mail) or blocklisted (the platform&#39;s list). | [optional] 
**Status** | Pointer to **int64** | Status is 422: the address may not be mailed. | [optional] 

## Methods

### NewEmailRefused

`func NewEmailRefused() *EmailRefused`

NewEmailRefused instantiates a new EmailRefused object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRefusedWithDefaults

`func NewEmailRefusedWithDefaults() *EmailRefused`

NewEmailRefusedWithDefaults instantiates a new EmailRefused object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *EmailRefused) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EmailRefused) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EmailRefused) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *EmailRefused) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetReason

`func (o *EmailRefused) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EmailRefused) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EmailRefused) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EmailRefused) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *EmailRefused) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailRefused) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailRefused) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailRefused) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


