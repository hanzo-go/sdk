# MarketingAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | City is the city. Required. | [optional] 
**Country** | Pointer to **string** | Country is the country. Required. | [optional] 
**Postcode** | Pointer to **string** | Postcode is the postal code. Optional. | [optional] 
**Region** | Pointer to **string** | Region is the state, province or region. Optional. | [optional] 
**Street** | Pointer to **string** | Street is the street address, on one line. Required. | [optional] 

## Methods

### NewMarketingAddress

`func NewMarketingAddress() *MarketingAddress`

NewMarketingAddress instantiates a new MarketingAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingAddressWithDefaults

`func NewMarketingAddressWithDefaults() *MarketingAddress`

NewMarketingAddressWithDefaults instantiates a new MarketingAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *MarketingAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MarketingAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MarketingAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MarketingAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *MarketingAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MarketingAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MarketingAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MarketingAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPostcode

`func (o *MarketingAddress) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *MarketingAddress) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *MarketingAddress) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *MarketingAddress) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetRegion

`func (o *MarketingAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MarketingAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MarketingAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MarketingAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStreet

`func (o *MarketingAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *MarketingAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *MarketingAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *MarketingAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


