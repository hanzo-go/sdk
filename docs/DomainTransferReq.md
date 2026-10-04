# DomainTransferReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | **string** | AuthCode is the transfer authorization the losing registrar issued. It is required. | 
**Domain** | **string** | Domain is the name to move in. It is required. | 
**Years** | Pointer to **int64** | Years is the term to buy on transfer, defaulting to 1. | [optional] 

## Methods

### NewDomainTransferReq

`func NewDomainTransferReq(authCode string, domain string, ) *DomainTransferReq`

NewDomainTransferReq instantiates a new DomainTransferReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainTransferReqWithDefaults

`func NewDomainTransferReqWithDefaults() *DomainTransferReq`

NewDomainTransferReqWithDefaults instantiates a new DomainTransferReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCode

`func (o *DomainTransferReq) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *DomainTransferReq) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *DomainTransferReq) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetDomain

`func (o *DomainTransferReq) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainTransferReq) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainTransferReq) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetYears

`func (o *DomainTransferReq) GetYears() int64`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *DomainTransferReq) GetYearsOk() (*int64, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *DomainTransferReq) SetYears(v int64)`

SetYears sets Years field to given value.

### HasYears

`func (o *DomainTransferReq) HasYears() bool`

HasYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


