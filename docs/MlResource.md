# MlResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when Kubernetes admitted the object, RFC 3339 in UTC. | [optional] 
**Name** | Pointer to **string** | Name is the object&#39;s metadata.name, unique within the caller&#39;s namespace. | [optional] 
**Spec** | Pointer to **map[string]map[string]interface{}** | Spec is the resource spec, verbatim as Kubernetes stores it. Present on a single-object read, absent from a list. | [optional] 
**Status** | Pointer to **map[string]map[string]interface{}** | Status is the live status kserve owns, verbatim. Absent until kserve has written one. | [optional] 

## Methods

### NewMlResource

`func NewMlResource() *MlResource`

NewMlResource instantiates a new MlResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlResourceWithDefaults

`func NewMlResourceWithDefaults() *MlResource`

NewMlResourceWithDefaults instantiates a new MlResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MlResource) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MlResource) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MlResource) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MlResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *MlResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MlResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpec

`func (o *MlResource) GetSpec() map[string]map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MlResource) GetSpecOk() (*map[string]map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MlResource) SetSpec(v map[string]map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MlResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *MlResource) GetStatus() map[string]map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MlResource) GetStatusOk() (*map[string]map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MlResource) SetStatus(v map[string]map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MlResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


