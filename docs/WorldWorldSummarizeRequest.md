# WorldWorldSummarizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headlines** | **[]string** | Headlines to summarize (non-empty). | 
**Mode** | Pointer to **string** | Summary mode (default: brief). | [optional] 
**GeoContext** | Pointer to **string** | Optional geographic context hint. | [optional] 
**Variant** | Pointer to **string** | Output variant (default: full). | [optional] 
**Lang** | Pointer to **string** | Optional target language. | [optional] 

## Methods

### NewWorldWorldSummarizeRequest

`func NewWorldWorldSummarizeRequest(headlines []string, ) *WorldWorldSummarizeRequest`

NewWorldWorldSummarizeRequest instantiates a new WorldWorldSummarizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldWorldSummarizeRequestWithDefaults

`func NewWorldWorldSummarizeRequestWithDefaults() *WorldWorldSummarizeRequest`

NewWorldWorldSummarizeRequestWithDefaults instantiates a new WorldWorldSummarizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeadlines

`func (o *WorldWorldSummarizeRequest) GetHeadlines() []string`

GetHeadlines returns the Headlines field if non-nil, zero value otherwise.

### GetHeadlinesOk

`func (o *WorldWorldSummarizeRequest) GetHeadlinesOk() (*[]string, bool)`

GetHeadlinesOk returns a tuple with the Headlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadlines

`func (o *WorldWorldSummarizeRequest) SetHeadlines(v []string)`

SetHeadlines sets Headlines field to given value.


### GetMode

`func (o *WorldWorldSummarizeRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WorldWorldSummarizeRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WorldWorldSummarizeRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WorldWorldSummarizeRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetGeoContext

`func (o *WorldWorldSummarizeRequest) GetGeoContext() string`

GetGeoContext returns the GeoContext field if non-nil, zero value otherwise.

### GetGeoContextOk

`func (o *WorldWorldSummarizeRequest) GetGeoContextOk() (*string, bool)`

GetGeoContextOk returns a tuple with the GeoContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoContext

`func (o *WorldWorldSummarizeRequest) SetGeoContext(v string)`

SetGeoContext sets GeoContext field to given value.

### HasGeoContext

`func (o *WorldWorldSummarizeRequest) HasGeoContext() bool`

HasGeoContext returns a boolean if a field has been set.

### GetVariant

`func (o *WorldWorldSummarizeRequest) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *WorldWorldSummarizeRequest) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *WorldWorldSummarizeRequest) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *WorldWorldSummarizeRequest) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetLang

`func (o *WorldWorldSummarizeRequest) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *WorldWorldSummarizeRequest) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *WorldWorldSummarizeRequest) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *WorldWorldSummarizeRequest) HasLang() bool`

HasLang returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


