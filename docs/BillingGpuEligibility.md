# BillingGpuEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eligible** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**PrepaidAvailable** | Pointer to **int64** |  | [optional] 
**CardOnFile** | Pointer to **bool** |  | [optional] 
**RequiredCents** | Pointer to **int64** |  | [optional] 

## Methods

### NewBillingGpuEligibility

`func NewBillingGpuEligibility() *BillingGpuEligibility`

NewBillingGpuEligibility instantiates a new BillingGpuEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingGpuEligibilityWithDefaults

`func NewBillingGpuEligibilityWithDefaults() *BillingGpuEligibility`

NewBillingGpuEligibilityWithDefaults instantiates a new BillingGpuEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligible

`func (o *BillingGpuEligibility) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *BillingGpuEligibility) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *BillingGpuEligibility) SetEligible(v bool)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *BillingGpuEligibility) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetReason

`func (o *BillingGpuEligibility) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BillingGpuEligibility) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BillingGpuEligibility) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BillingGpuEligibility) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPrepaidAvailable

`func (o *BillingGpuEligibility) GetPrepaidAvailable() int64`

GetPrepaidAvailable returns the PrepaidAvailable field if non-nil, zero value otherwise.

### GetPrepaidAvailableOk

`func (o *BillingGpuEligibility) GetPrepaidAvailableOk() (*int64, bool)`

GetPrepaidAvailableOk returns a tuple with the PrepaidAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidAvailable

`func (o *BillingGpuEligibility) SetPrepaidAvailable(v int64)`

SetPrepaidAvailable sets PrepaidAvailable field to given value.

### HasPrepaidAvailable

`func (o *BillingGpuEligibility) HasPrepaidAvailable() bool`

HasPrepaidAvailable returns a boolean if a field has been set.

### GetCardOnFile

`func (o *BillingGpuEligibility) GetCardOnFile() bool`

GetCardOnFile returns the CardOnFile field if non-nil, zero value otherwise.

### GetCardOnFileOk

`func (o *BillingGpuEligibility) GetCardOnFileOk() (*bool, bool)`

GetCardOnFileOk returns a tuple with the CardOnFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardOnFile

`func (o *BillingGpuEligibility) SetCardOnFile(v bool)`

SetCardOnFile sets CardOnFile field to given value.

### HasCardOnFile

`func (o *BillingGpuEligibility) HasCardOnFile() bool`

HasCardOnFile returns a boolean if a field has been set.

### GetRequiredCents

`func (o *BillingGpuEligibility) GetRequiredCents() int64`

GetRequiredCents returns the RequiredCents field if non-nil, zero value otherwise.

### GetRequiredCentsOk

`func (o *BillingGpuEligibility) GetRequiredCentsOk() (*int64, bool)`

GetRequiredCentsOk returns a tuple with the RequiredCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCents

`func (o *BillingGpuEligibility) SetRequiredCents(v int64)`

SetRequiredCents sets RequiredCents field to given value.

### HasRequiredCents

`func (o *BillingGpuEligibility) HasRequiredCents() bool`

HasRequiredCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


