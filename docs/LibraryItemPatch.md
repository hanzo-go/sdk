# LibraryItemPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to **bool** | Commit finishes a reserved upload once its bytes have been PUT. | [optional] 
**Id** | Pointer to **string** | ID is the item, from the path. | [optional] 
**Name** | Pointer to **string** | Name renames it. | [optional] 
**Project** | Pointer to **string** | Project moves it to one of the organization&#39;s projects; \&quot;\&quot; moves it out. | [optional] 
**Shared** | Pointer to **bool** | Shared shows it to every member of your organization, or stops. | [optional] 

## Methods

### NewLibraryItemPatch

`func NewLibraryItemPatch() *LibraryItemPatch`

NewLibraryItemPatch instantiates a new LibraryItemPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryItemPatchWithDefaults

`func NewLibraryItemPatchWithDefaults() *LibraryItemPatch`

NewLibraryItemPatchWithDefaults instantiates a new LibraryItemPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *LibraryItemPatch) GetCommit() bool`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *LibraryItemPatch) GetCommitOk() (*bool, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *LibraryItemPatch) SetCommit(v bool)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *LibraryItemPatch) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetId

`func (o *LibraryItemPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LibraryItemPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LibraryItemPatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LibraryItemPatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LibraryItemPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LibraryItemPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LibraryItemPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LibraryItemPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *LibraryItemPatch) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *LibraryItemPatch) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *LibraryItemPatch) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *LibraryItemPatch) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetShared

`func (o *LibraryItemPatch) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *LibraryItemPatch) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *LibraryItemPatch) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *LibraryItemPatch) HasShared() bool`

HasShared returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


