# PlatformProjectDuplicateRequestJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceEnvironmentId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IncludeServices** | Pointer to **bool** |  | [optional] [default to true]
**SelectedServices** | Pointer to [**[]PlatformProjectDuplicateRequestJsonSelectedServicesInner**](PlatformProjectDuplicateRequestJsonSelectedServicesInner.md) |  | [optional] 
**DuplicateInSameProject** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPlatformProjectDuplicateRequestJson

`func NewPlatformProjectDuplicateRequestJson(sourceEnvironmentId string, name string, ) *PlatformProjectDuplicateRequestJson`

NewPlatformProjectDuplicateRequestJson instantiates a new PlatformProjectDuplicateRequestJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformProjectDuplicateRequestJsonWithDefaults

`func NewPlatformProjectDuplicateRequestJsonWithDefaults() *PlatformProjectDuplicateRequestJson`

NewPlatformProjectDuplicateRequestJsonWithDefaults instantiates a new PlatformProjectDuplicateRequestJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceEnvironmentId

`func (o *PlatformProjectDuplicateRequestJson) GetSourceEnvironmentId() string`

GetSourceEnvironmentId returns the SourceEnvironmentId field if non-nil, zero value otherwise.

### GetSourceEnvironmentIdOk

`func (o *PlatformProjectDuplicateRequestJson) GetSourceEnvironmentIdOk() (*string, bool)`

GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironmentId

`func (o *PlatformProjectDuplicateRequestJson) SetSourceEnvironmentId(v string)`

SetSourceEnvironmentId sets SourceEnvironmentId field to given value.


### GetName

`func (o *PlatformProjectDuplicateRequestJson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformProjectDuplicateRequestJson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformProjectDuplicateRequestJson) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PlatformProjectDuplicateRequestJson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformProjectDuplicateRequestJson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformProjectDuplicateRequestJson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformProjectDuplicateRequestJson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludeServices

`func (o *PlatformProjectDuplicateRequestJson) GetIncludeServices() bool`

GetIncludeServices returns the IncludeServices field if non-nil, zero value otherwise.

### GetIncludeServicesOk

`func (o *PlatformProjectDuplicateRequestJson) GetIncludeServicesOk() (*bool, bool)`

GetIncludeServicesOk returns a tuple with the IncludeServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeServices

`func (o *PlatformProjectDuplicateRequestJson) SetIncludeServices(v bool)`

SetIncludeServices sets IncludeServices field to given value.

### HasIncludeServices

`func (o *PlatformProjectDuplicateRequestJson) HasIncludeServices() bool`

HasIncludeServices returns a boolean if a field has been set.

### GetSelectedServices

`func (o *PlatformProjectDuplicateRequestJson) GetSelectedServices() []PlatformProjectDuplicateRequestJsonSelectedServicesInner`

GetSelectedServices returns the SelectedServices field if non-nil, zero value otherwise.

### GetSelectedServicesOk

`func (o *PlatformProjectDuplicateRequestJson) GetSelectedServicesOk() (*[]PlatformProjectDuplicateRequestJsonSelectedServicesInner, bool)`

GetSelectedServicesOk returns a tuple with the SelectedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedServices

`func (o *PlatformProjectDuplicateRequestJson) SetSelectedServices(v []PlatformProjectDuplicateRequestJsonSelectedServicesInner)`

SetSelectedServices sets SelectedServices field to given value.

### HasSelectedServices

`func (o *PlatformProjectDuplicateRequestJson) HasSelectedServices() bool`

HasSelectedServices returns a boolean if a field has been set.

### GetDuplicateInSameProject

`func (o *PlatformProjectDuplicateRequestJson) GetDuplicateInSameProject() bool`

GetDuplicateInSameProject returns the DuplicateInSameProject field if non-nil, zero value otherwise.

### GetDuplicateInSameProjectOk

`func (o *PlatformProjectDuplicateRequestJson) GetDuplicateInSameProjectOk() (*bool, bool)`

GetDuplicateInSameProjectOk returns a tuple with the DuplicateInSameProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateInSameProject

`func (o *PlatformProjectDuplicateRequestJson) SetDuplicateInSameProject(v bool)`

SetDuplicateInSameProject sets DuplicateInSameProject field to given value.

### HasDuplicateInSameProject

`func (o *PlatformProjectDuplicateRequestJson) HasDuplicateInSameProject() bool`

HasDuplicateInSameProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


