# ContainerDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**ResourceRequests**](ResourceRequests.md) |  | [optional] 

## Methods

### NewContainerDetail

`func NewContainerDetail() *ContainerDetail`

NewContainerDetail instantiates a new ContainerDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerDetailWithDefaults

`func NewContainerDetailWithDefaults() *ContainerDetail`

NewContainerDetailWithDefaults instantiates a new ContainerDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *ContainerDetail) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ContainerDetail) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ContainerDetail) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ContainerDetail) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *ContainerDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContainerDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *ContainerDetail) GetResources() ResourceRequests`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ContainerDetail) GetResourcesOk() (*ResourceRequests, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ContainerDetail) SetResources(v ResourceRequests)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ContainerDetail) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


