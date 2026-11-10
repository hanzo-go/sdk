# PatentPatentTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | Pointer to **bool** | Found is false when no version of that id exists in this org. | [optional] 
**Id** | Pointer to **string** | ID is the entry named. | [optional] 
**Kind** | Pointer to **string** | Kind is which log was read. | [optional] 
**Project** | Pointer to **string** | Project is the scope it is filed under. | [optional] 
**Retained** | Pointer to **int64** | Retained is len(versions). | [optional] 
**Versions** | Pointer to [**[]PatentPatentTimelineEntry**](PatentPatentTimelineEntry.md) | Versions are every retained version, oldest append first. | [optional] 

## Methods

### NewPatentPatentTimeline

`func NewPatentPatentTimeline() *PatentPatentTimeline`

NewPatentPatentTimeline instantiates a new PatentPatentTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentTimelineWithDefaults

`func NewPatentPatentTimelineWithDefaults() *PatentPatentTimeline`

NewPatentPatentTimelineWithDefaults instantiates a new PatentPatentTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *PatentPatentTimeline) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *PatentPatentTimeline) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *PatentPatentTimeline) SetFound(v bool)`

SetFound sets Found field to given value.

### HasFound

`func (o *PatentPatentTimeline) HasFound() bool`

HasFound returns a boolean if a field has been set.

### GetId

`func (o *PatentPatentTimeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatentPatentTimeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatentPatentTimeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatentPatentTimeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *PatentPatentTimeline) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatentPatentTimeline) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatentPatentTimeline) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PatentPatentTimeline) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProject

`func (o *PatentPatentTimeline) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatentPatentTimeline) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatentPatentTimeline) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatentPatentTimeline) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRetained

`func (o *PatentPatentTimeline) GetRetained() int64`

GetRetained returns the Retained field if non-nil, zero value otherwise.

### GetRetainedOk

`func (o *PatentPatentTimeline) GetRetainedOk() (*int64, bool)`

GetRetainedOk returns a tuple with the Retained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetained

`func (o *PatentPatentTimeline) SetRetained(v int64)`

SetRetained sets Retained field to given value.

### HasRetained

`func (o *PatentPatentTimeline) HasRetained() bool`

HasRetained returns a boolean if a field has been set.

### GetVersions

`func (o *PatentPatentTimeline) GetVersions() []PatentPatentTimelineEntry`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PatentPatentTimeline) GetVersionsOk() (*[]PatentPatentTimelineEntry, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PatentPatentTimeline) SetVersions(v []PatentPatentTimelineEntry)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PatentPatentTimeline) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


