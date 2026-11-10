# AiRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **int64** | Default is the platform&#39;s requests a minute, the most the organization may set. | [optional] 
**Org** | Pointer to **string** | Org is the organization. | [optional] 
**Own** | Pointer to **int64** | Own is the organization&#39;s own setting, 0 when it holds its people to the default. | [optional] 
**Rpm** | Pointer to **int64** | Rpm is the model requests a minute one person may make acting in it now: its own setting where it set less than the default, else the default. | [optional] 

## Methods

### NewAiRate

`func NewAiRate() *AiRate`

NewAiRate instantiates a new AiRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRateWithDefaults

`func NewAiRateWithDefaults() *AiRate`

NewAiRateWithDefaults instantiates a new AiRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *AiRate) GetDefault() int64`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *AiRate) GetDefaultOk() (*int64, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *AiRate) SetDefault(v int64)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *AiRate) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetOrg

`func (o *AiRate) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AiRate) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AiRate) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AiRate) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOwn

`func (o *AiRate) GetOwn() int64`

GetOwn returns the Own field if non-nil, zero value otherwise.

### GetOwnOk

`func (o *AiRate) GetOwnOk() (*int64, bool)`

GetOwnOk returns a tuple with the Own field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwn

`func (o *AiRate) SetOwn(v int64)`

SetOwn sets Own field to given value.

### HasOwn

`func (o *AiRate) HasOwn() bool`

HasOwn returns a boolean if a field has been set.

### GetRpm

`func (o *AiRate) GetRpm() int64`

GetRpm returns the Rpm field if non-nil, zero value otherwise.

### GetRpmOk

`func (o *AiRate) GetRpmOk() (*int64, bool)`

GetRpmOk returns a tuple with the Rpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpm

`func (o *AiRate) SetRpm(v int64)`

SetRpm sets Rpm field to given value.

### HasRpm

`func (o *AiRate) HasRpm() bool`

HasRpm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


