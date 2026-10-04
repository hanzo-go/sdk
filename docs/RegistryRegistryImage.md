# RegistryRegistryImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the repository name inside the org&#39;s namespace (e.g. \&quot;cloud\&quot;). | [optional] 
**Ref** | Pointer to **string** | Ref is the full pullable reference (host + org + name) the docker CLI takes verbatim. | [optional] 

## Methods

### NewRegistryRegistryImage

`func NewRegistryRegistryImage() *RegistryRegistryImage`

NewRegistryRegistryImage instantiates a new RegistryRegistryImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryRegistryImageWithDefaults

`func NewRegistryRegistryImageWithDefaults() *RegistryRegistryImage`

NewRegistryRegistryImageWithDefaults instantiates a new RegistryRegistryImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegistryRegistryImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryRegistryImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryRegistryImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegistryRegistryImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRef

`func (o *RegistryRegistryImage) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RegistryRegistryImage) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RegistryRegistryImage) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RegistryRegistryImage) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


