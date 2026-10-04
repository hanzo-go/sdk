# TaxAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | City is the city or town. | [optional] 
**Country** | Pointer to **string** | Country is ISO 3166-1 alpha-2; \&quot;US\&quot; when absent. | [optional] 
**Line1** | Pointer to **string** | Line1 is the number, street, and apartment or suite. No periods or \&quot;#\&quot;: IRIS rejects them in an address line. | [optional] 
**Line2** | Pointer to **string** | Line2 continues the street address, when there is more of it. | [optional] 
**State** | Pointer to **string** | State is the two-letter state or territory code. | [optional] 
**Zip** | Pointer to **string** | ZIP is five or nine digits. | [optional] 

## Methods

### NewTaxAddress

`func NewTaxAddress() *TaxAddress`

NewTaxAddress instantiates a new TaxAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxAddressWithDefaults

`func NewTaxAddressWithDefaults() *TaxAddress`

NewTaxAddressWithDefaults instantiates a new TaxAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *TaxAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *TaxAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *TaxAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *TaxAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *TaxAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TaxAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TaxAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TaxAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLine1

`func (o *TaxAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *TaxAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *TaxAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *TaxAddress) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *TaxAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *TaxAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *TaxAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *TaxAddress) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetState

`func (o *TaxAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaxAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaxAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TaxAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *TaxAddress) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *TaxAddress) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *TaxAddress) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *TaxAddress) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


