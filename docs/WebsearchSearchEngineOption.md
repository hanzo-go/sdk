# WebsearchSearchEngineOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **int64** | Default is the count used when none is set. Absent for a text option, whose default is unset. | [optional] 
**Max** | Pointer to **int64** | Max is the most count the engine accepts. Absent for a text option. | [optional] 
**Min** | Pointer to **int64** | Min is the least count the engine accepts. Absent for a text option. | [optional] 
**Name** | Pointer to **string** | Name is the option: &#x60;count&#x60;, &#x60;language&#x60; or &#x60;region&#x60;. | [optional] 
**Param** | Pointer to **string** | Param is the engine&#39;s own query parameter the option sets, such as &#x60;t&#x60; or &#x60;rb&#x60; for Mojeek. | [optional] 

## Methods

### NewWebsearchSearchEngineOption

`func NewWebsearchSearchEngineOption() *WebsearchSearchEngineOption`

NewWebsearchSearchEngineOption instantiates a new WebsearchSearchEngineOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsearchSearchEngineOptionWithDefaults

`func NewWebsearchSearchEngineOptionWithDefaults() *WebsearchSearchEngineOption`

NewWebsearchSearchEngineOptionWithDefaults instantiates a new WebsearchSearchEngineOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *WebsearchSearchEngineOption) GetDefault() int64`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WebsearchSearchEngineOption) GetDefaultOk() (*int64, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WebsearchSearchEngineOption) SetDefault(v int64)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WebsearchSearchEngineOption) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetMax

`func (o *WebsearchSearchEngineOption) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *WebsearchSearchEngineOption) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *WebsearchSearchEngineOption) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *WebsearchSearchEngineOption) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *WebsearchSearchEngineOption) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *WebsearchSearchEngineOption) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *WebsearchSearchEngineOption) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *WebsearchSearchEngineOption) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetName

`func (o *WebsearchSearchEngineOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebsearchSearchEngineOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebsearchSearchEngineOption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebsearchSearchEngineOption) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParam

`func (o *WebsearchSearchEngineOption) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *WebsearchSearchEngineOption) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *WebsearchSearchEngineOption) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *WebsearchSearchEngineOption) HasParam() bool`

HasParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


