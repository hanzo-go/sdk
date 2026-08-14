# TransferReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | **string** | AuthCode is the transfer authorization the losing registrar issued. It is required. | 
**Domain** | **string** | Domain is the name to move in. It is required. | 
**Years** | Pointer to **int32** | Years is the term to buy on transfer, defaulting to 1. | [optional] 

## Methods

### NewTransferReq

`func NewTransferReq(authCode string, domain string, ) *TransferReq`

NewTransferReq instantiates a new TransferReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferReqWithDefaults

`func NewTransferReqWithDefaults() *TransferReq`

NewTransferReqWithDefaults instantiates a new TransferReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCode

`func (o *TransferReq) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *TransferReq) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *TransferReq) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetDomain

`func (o *TransferReq) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TransferReq) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TransferReq) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetYears

`func (o *TransferReq) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *TransferReq) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *TransferReq) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *TransferReq) HasYears() bool`

HasYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


