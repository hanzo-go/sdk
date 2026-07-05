# TrackerCreateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Optional; derived from name when omitted. ^[A-Z][A-Z0-9]{1,7}$ | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewTrackerCreateProjectRequest

`func NewTrackerCreateProjectRequest(name string, ) *TrackerCreateProjectRequest`

NewTrackerCreateProjectRequest instantiates a new TrackerCreateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackerCreateProjectRequestWithDefaults

`func NewTrackerCreateProjectRequestWithDefaults() *TrackerCreateProjectRequest`

NewTrackerCreateProjectRequestWithDefaults instantiates a new TrackerCreateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TrackerCreateProjectRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TrackerCreateProjectRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TrackerCreateProjectRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TrackerCreateProjectRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *TrackerCreateProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrackerCreateProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrackerCreateProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TrackerCreateProjectRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrackerCreateProjectRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrackerCreateProjectRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrackerCreateProjectRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


