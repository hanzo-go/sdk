# RegistryTagList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]string** | Data is the tag names, as the registry reports them. | [optional] 
**Image** | Pointer to **string** | Image is the repository name inside the org&#39;s namespace. | [optional] 
**Ref** | Pointer to **string** | Ref is the full repository reference the tags belong to. | [optional] 

## Methods

### NewRegistryTagList

`func NewRegistryTagList() *RegistryTagList`

NewRegistryTagList instantiates a new RegistryTagList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryTagListWithDefaults

`func NewRegistryTagListWithDefaults() *RegistryTagList`

NewRegistryTagListWithDefaults instantiates a new RegistryTagList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RegistryTagList) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RegistryTagList) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RegistryTagList) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *RegistryTagList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetImage

`func (o *RegistryTagList) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *RegistryTagList) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *RegistryTagList) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *RegistryTagList) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetRef

`func (o *RegistryTagList) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RegistryTagList) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RegistryTagList) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RegistryTagList) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


