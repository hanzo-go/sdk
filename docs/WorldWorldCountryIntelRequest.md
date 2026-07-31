# WorldWorldCountryIntelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** | Country name. | 
**Code** | **string** | Country code. | 
**Context** | Pointer to **map[string]interface{}** | Free-form sensor/context object marshaled into the prompt. | [optional] 

## Methods

### NewWorldWorldCountryIntelRequest

`func NewWorldWorldCountryIntelRequest(country string, code string, ) *WorldWorldCountryIntelRequest`

NewWorldWorldCountryIntelRequest instantiates a new WorldWorldCountryIntelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldWorldCountryIntelRequestWithDefaults

`func NewWorldWorldCountryIntelRequestWithDefaults() *WorldWorldCountryIntelRequest`

NewWorldWorldCountryIntelRequestWithDefaults instantiates a new WorldWorldCountryIntelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *WorldWorldCountryIntelRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *WorldWorldCountryIntelRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *WorldWorldCountryIntelRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCode

`func (o *WorldWorldCountryIntelRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WorldWorldCountryIntelRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WorldWorldCountryIntelRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetContext

`func (o *WorldWorldCountryIntelRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *WorldWorldCountryIntelRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *WorldWorldCountryIntelRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *WorldWorldCountryIntelRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


