# DbUpdateProjectRequestProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DefaultEndpointSettings** | Pointer to [**DbEndpointSettings**](DbEndpointSettings.md) |  | [optional] 

## Methods

### NewDbUpdateProjectRequestProject

`func NewDbUpdateProjectRequestProject() *DbUpdateProjectRequestProject`

NewDbUpdateProjectRequestProject instantiates a new DbUpdateProjectRequestProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbUpdateProjectRequestProjectWithDefaults

`func NewDbUpdateProjectRequestProjectWithDefaults() *DbUpdateProjectRequestProject`

NewDbUpdateProjectRequestProjectWithDefaults instantiates a new DbUpdateProjectRequestProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DbUpdateProjectRequestProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbUpdateProjectRequestProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbUpdateProjectRequestProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbUpdateProjectRequestProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultEndpointSettings

`func (o *DbUpdateProjectRequestProject) GetDefaultEndpointSettings() DbEndpointSettings`

GetDefaultEndpointSettings returns the DefaultEndpointSettings field if non-nil, zero value otherwise.

### GetDefaultEndpointSettingsOk

`func (o *DbUpdateProjectRequestProject) GetDefaultEndpointSettingsOk() (*DbEndpointSettings, bool)`

GetDefaultEndpointSettingsOk returns a tuple with the DefaultEndpointSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEndpointSettings

`func (o *DbUpdateProjectRequestProject) SetDefaultEndpointSettings(v DbEndpointSettings)`

SetDefaultEndpointSettings sets DefaultEndpointSettings field to given value.

### HasDefaultEndpointSettings

`func (o *DbUpdateProjectRequestProject) HasDefaultEndpointSettings() bool`

HasDefaultEndpointSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


