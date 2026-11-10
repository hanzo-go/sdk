# MarketplaceExpert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Areas** | Pointer to **[]string** | Areas is what the expert works on: configure, base, commerce, bot, agents, cloud, models, creative or freelance. | [optional] 
**Availability** | Pointer to **string** | Availability is the expert&#39;s weekly hours, in their own words. | [optional] 
**Languages** | Pointer to **[]string** | Languages are the languages the expert works in, lowercase BCP 47 tags. | [optional] 
**Person** | Pointer to **string** | Person is the expert&#39;s IAM identity: the verified principal that published the listing, never taken off the wire. Only the expert&#39;s own org reads it. | [optional] 
**Rate** | Pointer to **interface{}** |  | [optional] 
**Self** | Pointer to **bool** | Self is whether the expert is the caller. It is set on the caller&#39;s own org&#39;s listings and nowhere else, so a client finds its own listing without knowing how an identity is spelled. | [optional] 
**Timezone** | Pointer to **string** | Timezone is the expert&#39;s IANA time zone, \&quot;America/Los_Angeles\&quot;. | [optional] 
**Verified** | Pointer to **bool** | Verified is whether staff vouch for the expert. Only a SuperAdmin sets it. | [optional] 

## Methods

### NewMarketplaceExpert

`func NewMarketplaceExpert() *MarketplaceExpert`

NewMarketplaceExpert instantiates a new MarketplaceExpert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceExpertWithDefaults

`func NewMarketplaceExpertWithDefaults() *MarketplaceExpert`

NewMarketplaceExpertWithDefaults instantiates a new MarketplaceExpert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreas

`func (o *MarketplaceExpert) GetAreas() []string`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *MarketplaceExpert) GetAreasOk() (*[]string, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *MarketplaceExpert) SetAreas(v []string)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *MarketplaceExpert) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetAvailability

`func (o *MarketplaceExpert) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *MarketplaceExpert) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *MarketplaceExpert) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *MarketplaceExpert) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetLanguages

`func (o *MarketplaceExpert) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *MarketplaceExpert) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *MarketplaceExpert) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *MarketplaceExpert) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetPerson

`func (o *MarketplaceExpert) GetPerson() string`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *MarketplaceExpert) GetPersonOk() (*string, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *MarketplaceExpert) SetPerson(v string)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *MarketplaceExpert) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetRate

`func (o *MarketplaceExpert) GetRate() interface{}`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *MarketplaceExpert) GetRateOk() (*interface{}, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *MarketplaceExpert) SetRate(v interface{})`

SetRate sets Rate field to given value.

### HasRate

`func (o *MarketplaceExpert) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRateNil

`func (o *MarketplaceExpert) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *MarketplaceExpert) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetSelf

`func (o *MarketplaceExpert) GetSelf() bool`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MarketplaceExpert) GetSelfOk() (*bool, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *MarketplaceExpert) SetSelf(v bool)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *MarketplaceExpert) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetTimezone

`func (o *MarketplaceExpert) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MarketplaceExpert) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MarketplaceExpert) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *MarketplaceExpert) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetVerified

`func (o *MarketplaceExpert) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *MarketplaceExpert) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *MarketplaceExpert) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *MarketplaceExpert) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


