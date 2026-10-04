# DomainContacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to [**DomainRegistrant**](DomainRegistrant.md) | who administers it | [optional] 
**Billing** | Pointer to [**DomainRegistrant**](DomainRegistrant.md) | who is reached about payment | [optional] 
**Registrant** | Pointer to [**DomainRegistrant**](DomainRegistrant.md) | who owns the domain | [optional] 
**Tech** | Pointer to [**DomainRegistrant**](DomainRegistrant.md) | who is reached about technical matters | [optional] 

## Methods

### NewDomainContacts

`func NewDomainContacts() *DomainContacts`

NewDomainContacts instantiates a new DomainContacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainContactsWithDefaults

`func NewDomainContactsWithDefaults() *DomainContacts`

NewDomainContactsWithDefaults instantiates a new DomainContacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *DomainContacts) GetAdmin() DomainRegistrant`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *DomainContacts) GetAdminOk() (*DomainRegistrant, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *DomainContacts) SetAdmin(v DomainRegistrant)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *DomainContacts) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetBilling

`func (o *DomainContacts) GetBilling() DomainRegistrant`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *DomainContacts) GetBillingOk() (*DomainRegistrant, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *DomainContacts) SetBilling(v DomainRegistrant)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *DomainContacts) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetRegistrant

`func (o *DomainContacts) GetRegistrant() DomainRegistrant`

GetRegistrant returns the Registrant field if non-nil, zero value otherwise.

### GetRegistrantOk

`func (o *DomainContacts) GetRegistrantOk() (*DomainRegistrant, bool)`

GetRegistrantOk returns a tuple with the Registrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrant

`func (o *DomainContacts) SetRegistrant(v DomainRegistrant)`

SetRegistrant sets Registrant field to given value.

### HasRegistrant

`func (o *DomainContacts) HasRegistrant() bool`

HasRegistrant returns a boolean if a field has been set.

### GetTech

`func (o *DomainContacts) GetTech() DomainRegistrant`

GetTech returns the Tech field if non-nil, zero value otherwise.

### GetTechOk

`func (o *DomainContacts) GetTechOk() (*DomainRegistrant, bool)`

GetTechOk returns a tuple with the Tech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTech

`func (o *DomainContacts) SetTech(v DomainRegistrant)`

SetTech sets Tech field to given value.

### HasTech

`func (o *DomainContacts) HasTech() bool`

HasTech returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


