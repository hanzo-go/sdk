# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contacts** | Pointer to [**Contacts**](Contacts.md) | Contacts is the WHOIS contact set. Omit it and the registrar uses the reseller account&#39;s default contacts. | [optional] 
**Domain** | **string** | Domain is the name to buy. It is required. | 
**Years** | Pointer to **int32** | Years is the term to buy, defaulting to 1. | [optional] 

## Methods

### NewOrder

`func NewOrder(domain string, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContacts

`func (o *Order) GetContacts() Contacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Order) GetContactsOk() (*Contacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Order) SetContacts(v Contacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *Order) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetDomain

`func (o *Order) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Order) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Order) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetYears

`func (o *Order) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *Order) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *Order) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *Order) HasYears() bool`

HasYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


