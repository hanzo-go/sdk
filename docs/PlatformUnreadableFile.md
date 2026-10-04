# PlatformUnreadableFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path is the values file, relative to the repository root. | [optional] 
**Reason** | Pointer to **string** | Reason is why it could not be read. A SuperAdmin reads the parser&#39;s own words; anyone else reads that it does not parse, because those words can name the platform&#39;s cluster overlay. | [optional] 

## Methods

### NewPlatformUnreadableFile

`func NewPlatformUnreadableFile() *PlatformUnreadableFile`

NewPlatformUnreadableFile instantiates a new PlatformUnreadableFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformUnreadableFileWithDefaults

`func NewPlatformUnreadableFileWithDefaults() *PlatformUnreadableFile`

NewPlatformUnreadableFileWithDefaults instantiates a new PlatformUnreadableFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PlatformUnreadableFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PlatformUnreadableFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PlatformUnreadableFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PlatformUnreadableFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReason

`func (o *PlatformUnreadableFile) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PlatformUnreadableFile) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PlatformUnreadableFile) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PlatformUnreadableFile) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


