# PrincipalClearIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Amount is the gross amount in U.S. dollars, e.g. \&quot;1250.00\&quot;. | [optional] 
**Category** | Pointer to **string** | Category is what is bought: services, attorney, rents, royalties, other or merchandise. Services when omitted. | [optional] 
**Payee** | Pointer to **string** | Payee is the org id of the org to be paid. | [optional] 
**Performed** | Pointer to **string** | Performed is where the service is performed — for rents, where the property is; for royalties, where it is used — as an ISO 3166-1 alpha-2 code. It decides the source of a foreign payee&#39;s income, and so its withholding. | [optional] 
**Platform** | Pointer to **bool** | Platform is whether the caller operates the platform through which the payee provides this to the platform&#39;s users, which decides platform-operator reporting. Omitted, it is asked back as a fact when it matters. | [optional] 
**Rail** | Pointer to **string** | Rail is how it will move: x402, ledger, chain, ach, wire, card, network or other — the rails an economic event names. Ledger when omitted. | [optional] 

## Methods

### NewPrincipalClearIn

`func NewPrincipalClearIn() *PrincipalClearIn`

NewPrincipalClearIn instantiates a new PrincipalClearIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalClearInWithDefaults

`func NewPrincipalClearInWithDefaults() *PrincipalClearIn`

NewPrincipalClearInWithDefaults instantiates a new PrincipalClearIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PrincipalClearIn) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PrincipalClearIn) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PrincipalClearIn) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PrincipalClearIn) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCategory

`func (o *PrincipalClearIn) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PrincipalClearIn) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PrincipalClearIn) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PrincipalClearIn) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPayee

`func (o *PrincipalClearIn) GetPayee() string`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *PrincipalClearIn) GetPayeeOk() (*string, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *PrincipalClearIn) SetPayee(v string)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *PrincipalClearIn) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### GetPerformed

`func (o *PrincipalClearIn) GetPerformed() string`

GetPerformed returns the Performed field if non-nil, zero value otherwise.

### GetPerformedOk

`func (o *PrincipalClearIn) GetPerformedOk() (*string, bool)`

GetPerformedOk returns a tuple with the Performed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformed

`func (o *PrincipalClearIn) SetPerformed(v string)`

SetPerformed sets Performed field to given value.

### HasPerformed

`func (o *PrincipalClearIn) HasPerformed() bool`

HasPerformed returns a boolean if a field has been set.

### GetPlatform

`func (o *PrincipalClearIn) GetPlatform() bool`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PrincipalClearIn) GetPlatformOk() (*bool, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PrincipalClearIn) SetPlatform(v bool)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PrincipalClearIn) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRail

`func (o *PrincipalClearIn) GetRail() string`

GetRail returns the Rail field if non-nil, zero value otherwise.

### GetRailOk

`func (o *PrincipalClearIn) GetRailOk() (*string, bool)`

GetRailOk returns a tuple with the Rail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRail

`func (o *PrincipalClearIn) SetRail(v string)`

SetRail sets Rail field to given value.

### HasRail

`func (o *PrincipalClearIn) HasRail() bool`

HasRail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


