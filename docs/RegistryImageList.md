# RegistryImageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RegistryImage**](RegistryImage.md) | Data is the org&#39;s repositories. | [optional] 
**Truncated** | Pointer to **bool** | Truncated is true when the catalog walk hit its page bound before the registry was exhausted — the list is a prefix, not the whole. | [optional] 

## Methods

### NewRegistryImageList

`func NewRegistryImageList() *RegistryImageList`

NewRegistryImageList instantiates a new RegistryImageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryImageListWithDefaults

`func NewRegistryImageListWithDefaults() *RegistryImageList`

NewRegistryImageListWithDefaults instantiates a new RegistryImageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RegistryImageList) GetData() []RegistryImage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RegistryImageList) GetDataOk() (*[]RegistryImage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RegistryImageList) SetData(v []RegistryImage)`

SetData sets Data field to given value.

### HasData

`func (o *RegistryImageList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTruncated

`func (o *RegistryImageList) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *RegistryImageList) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *RegistryImageList) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *RegistryImageList) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


