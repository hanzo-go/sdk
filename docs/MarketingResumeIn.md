# MarketingResumeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the file&#39;s name; its extension must be .pdf, .doc or .docx. | [optional] 
**Size** | Pointer to **int64** | Size is the file&#39;s size in bytes, at most 10 MB. | [optional] 
**Type** | Pointer to **string** | Type is the file&#39;s content type, which must be that extension&#39;s. | [optional] 

## Methods

### NewMarketingResumeIn

`func NewMarketingResumeIn() *MarketingResumeIn`

NewMarketingResumeIn instantiates a new MarketingResumeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingResumeInWithDefaults

`func NewMarketingResumeInWithDefaults() *MarketingResumeIn`

NewMarketingResumeInWithDefaults instantiates a new MarketingResumeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MarketingResumeIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingResumeIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingResumeIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingResumeIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *MarketingResumeIn) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MarketingResumeIn) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MarketingResumeIn) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MarketingResumeIn) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *MarketingResumeIn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketingResumeIn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketingResumeIn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MarketingResumeIn) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


