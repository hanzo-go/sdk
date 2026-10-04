# PrincipalFounder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquityBps** | Pointer to **int64** | EquityBps is ownership in basis points, 0–10000. | [optional] 
**Kyc** | Pointer to **string** | KYC is pending, verified, reviewer_confirmed or failed, verbatim. | [optional] 
**Name** | Pointer to **string** | Name is the full legal name. | [optional] 

## Methods

### NewPrincipalFounder

`func NewPrincipalFounder() *PrincipalFounder`

NewPrincipalFounder instantiates a new PrincipalFounder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalFounderWithDefaults

`func NewPrincipalFounderWithDefaults() *PrincipalFounder`

NewPrincipalFounderWithDefaults instantiates a new PrincipalFounder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquityBps

`func (o *PrincipalFounder) GetEquityBps() int64`

GetEquityBps returns the EquityBps field if non-nil, zero value otherwise.

### GetEquityBpsOk

`func (o *PrincipalFounder) GetEquityBpsOk() (*int64, bool)`

GetEquityBpsOk returns a tuple with the EquityBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityBps

`func (o *PrincipalFounder) SetEquityBps(v int64)`

SetEquityBps sets EquityBps field to given value.

### HasEquityBps

`func (o *PrincipalFounder) HasEquityBps() bool`

HasEquityBps returns a boolean if a field has been set.

### GetKyc

`func (o *PrincipalFounder) GetKyc() string`

GetKyc returns the Kyc field if non-nil, zero value otherwise.

### GetKycOk

`func (o *PrincipalFounder) GetKycOk() (*string, bool)`

GetKycOk returns a tuple with the Kyc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKyc

`func (o *PrincipalFounder) SetKyc(v string)`

SetKyc sets Kyc field to given value.

### HasKyc

`func (o *PrincipalFounder) HasKyc() bool`

HasKyc returns a boolean if a field has been set.

### GetName

`func (o *PrincipalFounder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalFounder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalFounder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrincipalFounder) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


