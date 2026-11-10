# ComputeLaneModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the weight set&#39;s id, as the node&#39;s engine loads it (triposr, wan2.2, ...). | [optional] 
**Lane** | Pointer to **string** | Lane is the Studio render lane these weights serve first (mesh, video.wan, ...). | [optional] 
**Licence** | Pointer to **string** | Licence is the weights&#39; SPDX licence identifier (MIT, Apache-2.0, ...). | [optional] 

## Methods

### NewComputeLaneModel

`func NewComputeLaneModel() *ComputeLaneModel`

NewComputeLaneModel instantiates a new ComputeLaneModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeLaneModelWithDefaults

`func NewComputeLaneModelWithDefaults() *ComputeLaneModel`

NewComputeLaneModelWithDefaults instantiates a new ComputeLaneModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputeLaneModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputeLaneModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputeLaneModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputeLaneModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLane

`func (o *ComputeLaneModel) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *ComputeLaneModel) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *ComputeLaneModel) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *ComputeLaneModel) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetLicence

`func (o *ComputeLaneModel) GetLicence() string`

GetLicence returns the Licence field if non-nil, zero value otherwise.

### GetLicenceOk

`func (o *ComputeLaneModel) GetLicenceOk() (*string, bool)`

GetLicenceOk returns a tuple with the Licence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicence

`func (o *ComputeLaneModel) SetLicence(v string)`

SetLicence sets Licence field to given value.

### HasLicence

`func (o *ComputeLaneModel) HasLicence() bool`

HasLicence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


