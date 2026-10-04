# DomainRenewResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaidCents** | Pointer to **int64** | what this renewal cost, in cents | [optional] 
**Record** | Pointer to [**DomainHolding**](DomainHolding.md) | the ownership row with its new expiry | [optional] 

## Methods

### NewDomainRenewResult

`func NewDomainRenewResult() *DomainRenewResult`

NewDomainRenewResult instantiates a new DomainRenewResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRenewResultWithDefaults

`func NewDomainRenewResultWithDefaults() *DomainRenewResult`

NewDomainRenewResultWithDefaults instantiates a new DomainRenewResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaidCents

`func (o *DomainRenewResult) GetPaidCents() int64`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *DomainRenewResult) GetPaidCentsOk() (*int64, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *DomainRenewResult) SetPaidCents(v int64)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *DomainRenewResult) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetRecord

`func (o *DomainRenewResult) GetRecord() DomainHolding`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *DomainRenewResult) GetRecordOk() (*DomainHolding, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *DomainRenewResult) SetRecord(v DomainHolding)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *DomainRenewResult) HasRecord() bool`

HasRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


