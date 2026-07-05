# VisorScaleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewVisorScaleRequest

`func NewVisorScaleRequest(provider string, ) *VisorScaleRequest`

NewVisorScaleRequest instantiates a new VisorScaleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisorScaleRequestWithDefaults

`func NewVisorScaleRequestWithDefaults() *VisorScaleRequest`

NewVisorScaleRequestWithDefaults instantiates a new VisorScaleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *VisorScaleRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VisorScaleRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VisorScaleRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetCount

`func (o *VisorScaleRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VisorScaleRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VisorScaleRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *VisorScaleRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


