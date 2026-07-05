# FlowAppConnection

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

### NewFlowAppConnection

`func NewFlowAppConnection() *FlowAppConnection`

NewFlowAppConnection instantiates a new FlowAppConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowAppConnectionWithDefaults

`func NewFlowAppConnectionWithDefaults() *FlowAppConnection`

NewFlowAppConnectionWithDefaults instantiates a new FlowAppConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowAppConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowAppConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowAppConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowAppConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FlowAppConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowAppConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowAppConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowAppConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *FlowAppConnection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FlowAppConnection) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FlowAppConnection) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FlowAppConnection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPieceName

`func (o *FlowAppConnection) GetPieceName() string`

GetPieceName returns the PieceName field if non-nil, zero value otherwise.

### GetPieceNameOk

`func (o *FlowAppConnection) GetPieceNameOk() (*string, bool)`

GetPieceNameOk returns a tuple with the PieceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceName

`func (o *FlowAppConnection) SetPieceName(v string)`

SetPieceName sets PieceName field to given value.

### HasPieceName

`func (o *FlowAppConnection) HasPieceName() bool`

HasPieceName returns a boolean if a field has been set.

### GetProjectId

`func (o *FlowAppConnection) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FlowAppConnection) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FlowAppConnection) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *FlowAppConnection) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetType

`func (o *FlowAppConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowAppConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowAppConnection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FlowAppConnection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *FlowAppConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowAppConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowAppConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlowAppConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExternalId

`func (o *FlowAppConnection) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *FlowAppConnection) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *FlowAppConnection) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *FlowAppConnection) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCreated

`func (o *FlowAppConnection) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FlowAppConnection) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FlowAppConnection) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FlowAppConnection) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *FlowAppConnection) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FlowAppConnection) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FlowAppConnection) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FlowAppConnection) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


