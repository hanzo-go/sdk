# WorldWorldClassifyBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Titles** | **[]string** | Titles to classify (non-empty). | 
**Variant** | Pointer to **string** | Classification variant (only &#39;tech&#39; is special-cased). | [optional] 

## Methods

### NewWorldWorldClassifyBatchRequest

`func NewWorldWorldClassifyBatchRequest(titles []string, ) *WorldWorldClassifyBatchRequest`

NewWorldWorldClassifyBatchRequest instantiates a new WorldWorldClassifyBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldWorldClassifyBatchRequestWithDefaults

`func NewWorldWorldClassifyBatchRequestWithDefaults() *WorldWorldClassifyBatchRequest`

NewWorldWorldClassifyBatchRequestWithDefaults instantiates a new WorldWorldClassifyBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitles

`func (o *WorldWorldClassifyBatchRequest) GetTitles() []string`

GetTitles returns the Titles field if non-nil, zero value otherwise.

### GetTitlesOk

`func (o *WorldWorldClassifyBatchRequest) GetTitlesOk() (*[]string, bool)`

GetTitlesOk returns a tuple with the Titles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitles

`func (o *WorldWorldClassifyBatchRequest) SetTitles(v []string)`

SetTitles sets Titles field to given value.


### GetVariant

`func (o *WorldWorldClassifyBatchRequest) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *WorldWorldClassifyBatchRequest) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *WorldWorldClassifyBatchRequest) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *WorldWorldClassifyBatchRequest) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


