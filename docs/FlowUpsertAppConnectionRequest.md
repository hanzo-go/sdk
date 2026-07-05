# FlowUpsertAppConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**PieceName** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Value** | **map[string]interface{}** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewFlowUpsertAppConnectionRequest

`func NewFlowUpsertAppConnectionRequest(displayName string, pieceName string, type_ string, value map[string]interface{}, ) *FlowUpsertAppConnectionRequest`

NewFlowUpsertAppConnectionRequest instantiates a new FlowUpsertAppConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowUpsertAppConnectionRequestWithDefaults

`func NewFlowUpsertAppConnectionRequestWithDefaults() *FlowUpsertAppConnectionRequest`

NewFlowUpsertAppConnectionRequestWithDefaults instantiates a new FlowUpsertAppConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *FlowUpsertAppConnectionRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FlowUpsertAppConnectionRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FlowUpsertAppConnectionRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPieceName

`func (o *FlowUpsertAppConnectionRequest) GetPieceName() string`

GetPieceName returns the PieceName field if non-nil, zero value otherwise.

### GetPieceNameOk

`func (o *FlowUpsertAppConnectionRequest) GetPieceNameOk() (*string, bool)`

GetPieceNameOk returns a tuple with the PieceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceName

`func (o *FlowUpsertAppConnectionRequest) SetPieceName(v string)`

SetPieceName sets PieceName field to given value.


### GetExternalId

`func (o *FlowUpsertAppConnectionRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *FlowUpsertAppConnectionRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *FlowUpsertAppConnectionRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *FlowUpsertAppConnectionRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetType

`func (o *FlowUpsertAppConnectionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowUpsertAppConnectionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowUpsertAppConnectionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *FlowUpsertAppConnectionRequest) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FlowUpsertAppConnectionRequest) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FlowUpsertAppConnectionRequest) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetMetadata

`func (o *FlowUpsertAppConnectionRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlowUpsertAppConnectionRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlowUpsertAppConnectionRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FlowUpsertAppConnectionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


