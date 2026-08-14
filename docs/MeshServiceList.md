# MeshServiceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to [**[]MeshView**](MeshView.md) | Services is one row per ZT edge service tagged with the caller&#39;s org role. | [optional] 

## Methods

### NewMeshServiceList

`func NewMeshServiceList() *MeshServiceList`

NewMeshServiceList instantiates a new MeshServiceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeshServiceListWithDefaults

`func NewMeshServiceListWithDefaults() *MeshServiceList`

NewMeshServiceListWithDefaults instantiates a new MeshServiceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *MeshServiceList) GetServices() []MeshView`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *MeshServiceList) GetServicesOk() (*[]MeshView, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *MeshServiceList) SetServices(v []MeshView)`

SetServices sets Services field to given value.

### HasServices

`func (o *MeshServiceList) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


