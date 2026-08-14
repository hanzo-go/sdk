# RenewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain is the name to extend. It is required, and the caller&#39;s org must hold it. | 
**Years** | Pointer to **int32** | Years is how much longer to hold it, defaulting to 1. | [optional] 

## Methods

### NewRenewReq

`func NewRenewReq(domain string, ) *RenewReq`

NewRenewReq instantiates a new RenewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewReqWithDefaults

`func NewRenewReqWithDefaults() *RenewReq`

NewRenewReqWithDefaults instantiates a new RenewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RenewReq) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RenewReq) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RenewReq) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetYears

`func (o *RenewReq) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *RenewReq) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *RenewReq) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *RenewReq) HasYears() bool`

HasYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


