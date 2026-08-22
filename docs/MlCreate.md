# MlCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **map[string]string** | Labels are extra labels to set on the object, merged UNDER the tenancy labels this plane derives from the validated principal — so a label naming another org&#39;s scope cannot displace the real one. | [optional] 
**Name** | Pointer to **string** | Name is the resource&#39;s name: a DNS-1123 label (^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$), lowercased and trimmed. It is the name the resource answers to for the life of the caller&#39;s org. | [optional] 
**Spec** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewMlCreate

`func NewMlCreate() *MlCreate`

NewMlCreate instantiates a new MlCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlCreateWithDefaults

`func NewMlCreateWithDefaults() *MlCreate`

NewMlCreateWithDefaults instantiates a new MlCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *MlCreate) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MlCreate) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MlCreate) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MlCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *MlCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MlCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpec

`func (o *MlCreate) GetSpec() interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MlCreate) GetSpecOk() (*interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MlCreate) SetSpec(v interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MlCreate) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *MlCreate) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *MlCreate) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


