# AutoUpsertAppConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**PieceName** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Value** | **map[string]interface{}** |  | 

## Methods

### NewAutoUpsertAppConnectionRequest

`func NewAutoUpsertAppConnectionRequest(displayName string, pieceName string, type_ string, value map[string]interface{}, ) *AutoUpsertAppConnectionRequest`

NewAutoUpsertAppConnectionRequest instantiates a new AutoUpsertAppConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoUpsertAppConnectionRequestWithDefaults

`func NewAutoUpsertAppConnectionRequestWithDefaults() *AutoUpsertAppConnectionRequest`

NewAutoUpsertAppConnectionRequestWithDefaults instantiates a new AutoUpsertAppConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AutoUpsertAppConnectionRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutoUpsertAppConnectionRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutoUpsertAppConnectionRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPieceName

`func (o *AutoUpsertAppConnectionRequest) GetPieceName() string`

GetPieceName returns the PieceName field if non-nil, zero value otherwise.

### GetPieceNameOk

`func (o *AutoUpsertAppConnectionRequest) GetPieceNameOk() (*string, bool)`

GetPieceNameOk returns a tuple with the PieceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceName

`func (o *AutoUpsertAppConnectionRequest) SetPieceName(v string)`

SetPieceName sets PieceName field to given value.


### GetExternalId

`func (o *AutoUpsertAppConnectionRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutoUpsertAppConnectionRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutoUpsertAppConnectionRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AutoUpsertAppConnectionRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetType

`func (o *AutoUpsertAppConnectionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoUpsertAppConnectionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoUpsertAppConnectionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *AutoUpsertAppConnectionRequest) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AutoUpsertAppConnectionRequest) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AutoUpsertAppConnectionRequest) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


