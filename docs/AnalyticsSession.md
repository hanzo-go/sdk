# AnalyticsSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**WebsiteId** | Pointer to **string** |  | [optional] 
**Browser** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Device** | Pointer to **string** |  | [optional] 
**Screen** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**DistinctId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAnalyticsSession

`func NewAnalyticsSession() *AnalyticsSession`

NewAnalyticsSession instantiates a new AnalyticsSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsSessionWithDefaults

`func NewAnalyticsSessionWithDefaults() *AnalyticsSession`

NewAnalyticsSessionWithDefaults instantiates a new AnalyticsSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnalyticsSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnalyticsSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWebsiteId

`func (o *AnalyticsSession) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *AnalyticsSession) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *AnalyticsSession) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.

### HasWebsiteId

`func (o *AnalyticsSession) HasWebsiteId() bool`

HasWebsiteId returns a boolean if a field has been set.

### GetBrowser

`func (o *AnalyticsSession) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *AnalyticsSession) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *AnalyticsSession) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *AnalyticsSession) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetOs

`func (o *AnalyticsSession) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *AnalyticsSession) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *AnalyticsSession) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *AnalyticsSession) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetDevice

`func (o *AnalyticsSession) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AnalyticsSession) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AnalyticsSession) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AnalyticsSession) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetScreen

`func (o *AnalyticsSession) GetScreen() string`

GetScreen returns the Screen field if non-nil, zero value otherwise.

### GetScreenOk

`func (o *AnalyticsSession) GetScreenOk() (*string, bool)`

GetScreenOk returns a tuple with the Screen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreen

`func (o *AnalyticsSession) SetScreen(v string)`

SetScreen sets Screen field to given value.

### HasScreen

`func (o *AnalyticsSession) HasScreen() bool`

HasScreen returns a boolean if a field has been set.

### GetLanguage

`func (o *AnalyticsSession) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AnalyticsSession) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AnalyticsSession) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AnalyticsSession) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCountry

`func (o *AnalyticsSession) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AnalyticsSession) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AnalyticsSession) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AnalyticsSession) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRegion

`func (o *AnalyticsSession) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AnalyticsSession) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AnalyticsSession) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AnalyticsSession) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCity

`func (o *AnalyticsSession) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AnalyticsSession) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AnalyticsSession) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AnalyticsSession) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetDistinctId

`func (o *AnalyticsSession) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *AnalyticsSession) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *AnalyticsSession) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *AnalyticsSession) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AnalyticsSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnalyticsSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnalyticsSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AnalyticsSession) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


