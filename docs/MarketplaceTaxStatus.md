# MarketplaceTaxStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certified** | Pointer to **bool** | Certified is whether its signature covers the form as it stands. | [optional] 
**Expires** | Pointer to **int64** | Expires is when a W-8 lapses, unix seconds; a W-9 does not. | [optional] 
**Form** | Pointer to **string** | Form is w9, w8ben or w8bene. | [optional] 
**Valid** | Pointer to **bool** | Valid is whether it is complete and unexpired. | [optional] 

## Methods

### NewMarketplaceTaxStatus

`func NewMarketplaceTaxStatus() *MarketplaceTaxStatus`

NewMarketplaceTaxStatus instantiates a new MarketplaceTaxStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceTaxStatusWithDefaults

`func NewMarketplaceTaxStatusWithDefaults() *MarketplaceTaxStatus`

NewMarketplaceTaxStatusWithDefaults instantiates a new MarketplaceTaxStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertified

`func (o *MarketplaceTaxStatus) GetCertified() bool`

GetCertified returns the Certified field if non-nil, zero value otherwise.

### GetCertifiedOk

`func (o *MarketplaceTaxStatus) GetCertifiedOk() (*bool, bool)`

GetCertifiedOk returns a tuple with the Certified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertified

`func (o *MarketplaceTaxStatus) SetCertified(v bool)`

SetCertified sets Certified field to given value.

### HasCertified

`func (o *MarketplaceTaxStatus) HasCertified() bool`

HasCertified returns a boolean if a field has been set.

### GetExpires

`func (o *MarketplaceTaxStatus) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *MarketplaceTaxStatus) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *MarketplaceTaxStatus) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *MarketplaceTaxStatus) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetForm

`func (o *MarketplaceTaxStatus) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *MarketplaceTaxStatus) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *MarketplaceTaxStatus) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *MarketplaceTaxStatus) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetValid

`func (o *MarketplaceTaxStatus) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *MarketplaceTaxStatus) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *MarketplaceTaxStatus) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *MarketplaceTaxStatus) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


