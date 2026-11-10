# MarketplaceExpertIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Areas** | Pointer to **[]string** | Areas is what the expert works on, one or more of configure, base, commerce, bot, agents, cloud, models, creative and freelance. | [optional] 
**Availability** | Pointer to **string** | Availability is the weekly hours in free form, at most 512 characters. | [optional] 
**Languages** | Pointer to **[]string** | Languages are 1 to 8 lowercase BCP 47 tags: \&quot;en\&quot;, \&quot;es\&quot;, \&quot;zh-cn\&quot;. | [optional] 
**Rate** | Pointer to **string** | Rate is the hourly rate as a decimal USD string, exact to 18 places; empty or \&quot;0\&quot; sells no hours. | [optional] 
**Timezone** | Pointer to **string** | Timezone is an IANA time zone: \&quot;America/Los_Angeles\&quot;. | [optional] 

## Methods

### NewMarketplaceExpertIn

`func NewMarketplaceExpertIn() *MarketplaceExpertIn`

NewMarketplaceExpertIn instantiates a new MarketplaceExpertIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceExpertInWithDefaults

`func NewMarketplaceExpertInWithDefaults() *MarketplaceExpertIn`

NewMarketplaceExpertInWithDefaults instantiates a new MarketplaceExpertIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreas

`func (o *MarketplaceExpertIn) GetAreas() []string`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *MarketplaceExpertIn) GetAreasOk() (*[]string, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *MarketplaceExpertIn) SetAreas(v []string)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *MarketplaceExpertIn) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetAvailability

`func (o *MarketplaceExpertIn) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *MarketplaceExpertIn) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *MarketplaceExpertIn) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *MarketplaceExpertIn) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetLanguages

`func (o *MarketplaceExpertIn) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *MarketplaceExpertIn) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *MarketplaceExpertIn) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *MarketplaceExpertIn) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetRate

`func (o *MarketplaceExpertIn) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *MarketplaceExpertIn) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *MarketplaceExpertIn) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *MarketplaceExpertIn) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetTimezone

`func (o *MarketplaceExpertIn) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MarketplaceExpertIn) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MarketplaceExpertIn) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *MarketplaceExpertIn) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


