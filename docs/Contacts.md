# Contacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to [**Registrant**](Registrant.md) | who administers it | [optional] 
**Billing** | Pointer to [**Registrant**](Registrant.md) | who is reached about payment | [optional] 
**Registrant** | Pointer to [**Registrant**](Registrant.md) | who owns the domain | [optional] 
**Tech** | Pointer to [**Registrant**](Registrant.md) | who is reached about technical matters | [optional] 

## Methods

### NewContacts

`func NewContacts() *Contacts`

NewContacts instantiates a new Contacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactsWithDefaults

`func NewContactsWithDefaults() *Contacts`

NewContactsWithDefaults instantiates a new Contacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *Contacts) GetAdmin() Registrant`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *Contacts) GetAdminOk() (*Registrant, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *Contacts) SetAdmin(v Registrant)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *Contacts) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetBilling

`func (o *Contacts) GetBilling() Registrant`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Contacts) GetBillingOk() (*Registrant, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *Contacts) SetBilling(v Registrant)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *Contacts) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetRegistrant

`func (o *Contacts) GetRegistrant() Registrant`

GetRegistrant returns the Registrant field if non-nil, zero value otherwise.

### GetRegistrantOk

`func (o *Contacts) GetRegistrantOk() (*Registrant, bool)`

GetRegistrantOk returns a tuple with the Registrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrant

`func (o *Contacts) SetRegistrant(v Registrant)`

SetRegistrant sets Registrant field to given value.

### HasRegistrant

`func (o *Contacts) HasRegistrant() bool`

HasRegistrant returns a boolean if a field has been set.

### GetTech

`func (o *Contacts) GetTech() Registrant`

GetTech returns the Tech field if non-nil, zero value otherwise.

### GetTechOk

`func (o *Contacts) GetTechOk() (*Registrant, bool)`

GetTechOk returns a tuple with the Tech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTech

`func (o *Contacts) SetTech(v Registrant)`

SetTech sets Tech field to given value.

### HasTech

`func (o *Contacts) HasTech() bool`

HasTech returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


