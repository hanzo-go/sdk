# RenewResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaidCents** | Pointer to **int32** | what this renewal cost, in cents | [optional] 
**Record** | Pointer to [**Holding**](Holding.md) | the ownership row with its new expiry | [optional] 

## Methods

### NewRenewResult

`func NewRenewResult() *RenewResult`

NewRenewResult instantiates a new RenewResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewResultWithDefaults

`func NewRenewResultWithDefaults() *RenewResult`

NewRenewResultWithDefaults instantiates a new RenewResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaidCents

`func (o *RenewResult) GetPaidCents() int32`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *RenewResult) GetPaidCentsOk() (*int32, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *RenewResult) SetPaidCents(v int32)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *RenewResult) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetRecord

`func (o *RenewResult) GetRecord() Holding`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *RenewResult) GetRecordOk() (*Holding, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *RenewResult) SetRecord(v Holding)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *RenewResult) HasRecord() bool`

HasRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


