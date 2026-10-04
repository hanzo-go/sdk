# TaxParty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**TaxAddress**](TaxAddress.md) | Address is where the form says they are. | [optional] 
**BusinessName** | Pointer to **string** | BusinessName is W-9 line 2, when there is one. | [optional] 
**Name** | Pointer to **string** | Name is W-9 line 1. | [optional] 
**Org** | Pointer to **string** | Org is the Hanzo org. | [optional] 
**Phone** | Pointer to **string** | Phone is the payer&#39;s telephone number; a recipient has none on the form. | [optional] 
**Tin** | Pointer to **string** | TIN is masked here always. The recipient&#39;s is truncated on Copy B as Pub 1099 allows; the payer&#39;s prints in full on the form itself, where the law requires it. | [optional] 
**TinType** | Pointer to **string** | TINType is ssn or ein; empty when no TIN is held. | [optional] 

## Methods

### NewTaxParty

`func NewTaxParty() *TaxParty`

NewTaxParty instantiates a new TaxParty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxPartyWithDefaults

`func NewTaxPartyWithDefaults() *TaxParty`

NewTaxPartyWithDefaults instantiates a new TaxParty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TaxParty) GetAddress() TaxAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TaxParty) GetAddressOk() (*TaxAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TaxParty) SetAddress(v TaxAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TaxParty) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBusinessName

`func (o *TaxParty) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *TaxParty) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *TaxParty) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *TaxParty) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetName

`func (o *TaxParty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxParty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxParty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxParty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *TaxParty) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *TaxParty) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *TaxParty) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *TaxParty) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPhone

`func (o *TaxParty) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TaxParty) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TaxParty) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TaxParty) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetTin

`func (o *TaxParty) GetTin() string`

GetTin returns the Tin field if non-nil, zero value otherwise.

### GetTinOk

`func (o *TaxParty) GetTinOk() (*string, bool)`

GetTinOk returns a tuple with the Tin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTin

`func (o *TaxParty) SetTin(v string)`

SetTin sets Tin field to given value.

### HasTin

`func (o *TaxParty) HasTin() bool`

HasTin returns a boolean if a field has been set.

### GetTinType

`func (o *TaxParty) GetTinType() string`

GetTinType returns the TinType field if non-nil, zero value otherwise.

### GetTinTypeOk

`func (o *TaxParty) GetTinTypeOk() (*string, bool)`

GetTinTypeOk returns a tuple with the TinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinType

`func (o *TaxParty) SetTinType(v string)`

SetTinType sets TinType field to given value.

### HasTinType

`func (o *TaxParty) HasTinType() bool`

HasTinType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


