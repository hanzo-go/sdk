# AutoAppConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PieceName** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAutoAppConnection

`func NewAutoAppConnection() *AutoAppConnection`

NewAutoAppConnection instantiates a new AutoAppConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoAppConnectionWithDefaults

`func NewAutoAppConnectionWithDefaults() *AutoAppConnection`

NewAutoAppConnectionWithDefaults instantiates a new AutoAppConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoAppConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoAppConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoAppConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoAppConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AutoAppConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoAppConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoAppConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoAppConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *AutoAppConnection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutoAppConnection) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutoAppConnection) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutoAppConnection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPieceName

`func (o *AutoAppConnection) GetPieceName() string`

GetPieceName returns the PieceName field if non-nil, zero value otherwise.

### GetPieceNameOk

`func (o *AutoAppConnection) GetPieceNameOk() (*string, bool)`

GetPieceNameOk returns a tuple with the PieceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceName

`func (o *AutoAppConnection) SetPieceName(v string)`

SetPieceName sets PieceName field to given value.

### HasPieceName

`func (o *AutoAppConnection) HasPieceName() bool`

HasPieceName returns a boolean if a field has been set.

### GetProjectId

`func (o *AutoAppConnection) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoAppConnection) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoAppConnection) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AutoAppConnection) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetType

`func (o *AutoAppConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoAppConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoAppConnection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutoAppConnection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *AutoAppConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutoAppConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutoAppConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutoAppConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExternalId

`func (o *AutoAppConnection) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutoAppConnection) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutoAppConnection) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AutoAppConnection) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCreated

`func (o *AutoAppConnection) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AutoAppConnection) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AutoAppConnection) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AutoAppConnection) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *AutoAppConnection) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AutoAppConnection) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AutoAppConnection) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AutoAppConnection) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


