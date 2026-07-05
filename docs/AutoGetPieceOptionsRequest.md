# AutoGetPieceOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **string** |  | 
**FlowVersionId** | **string** |  | 
**ActionOrTriggerName** | **string** |  | 
**PropertyName** | **string** |  | 
**PieceName** | **string** |  | 
**PieceVersion** | **string** |  | 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**SearchValue** | Pointer to **string** |  | [optional] 

## Methods

### NewAutoGetPieceOptionsRequest

`func NewAutoGetPieceOptionsRequest(flowId string, flowVersionId string, actionOrTriggerName string, propertyName string, pieceName string, pieceVersion string, ) *AutoGetPieceOptionsRequest`

NewAutoGetPieceOptionsRequest instantiates a new AutoGetPieceOptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoGetPieceOptionsRequestWithDefaults

`func NewAutoGetPieceOptionsRequestWithDefaults() *AutoGetPieceOptionsRequest`

NewAutoGetPieceOptionsRequestWithDefaults instantiates a new AutoGetPieceOptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *AutoGetPieceOptionsRequest) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AutoGetPieceOptionsRequest) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AutoGetPieceOptionsRequest) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetFlowVersionId

`func (o *AutoGetPieceOptionsRequest) GetFlowVersionId() string`

GetFlowVersionId returns the FlowVersionId field if non-nil, zero value otherwise.

### GetFlowVersionIdOk

`func (o *AutoGetPieceOptionsRequest) GetFlowVersionIdOk() (*string, bool)`

GetFlowVersionIdOk returns a tuple with the FlowVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowVersionId

`func (o *AutoGetPieceOptionsRequest) SetFlowVersionId(v string)`

SetFlowVersionId sets FlowVersionId field to given value.


### GetActionOrTriggerName

`func (o *AutoGetPieceOptionsRequest) GetActionOrTriggerName() string`

GetActionOrTriggerName returns the ActionOrTriggerName field if non-nil, zero value otherwise.

### GetActionOrTriggerNameOk

`func (o *AutoGetPieceOptionsRequest) GetActionOrTriggerNameOk() (*string, bool)`

GetActionOrTriggerNameOk returns a tuple with the ActionOrTriggerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOrTriggerName

`func (o *AutoGetPieceOptionsRequest) SetActionOrTriggerName(v string)`

SetActionOrTriggerName sets ActionOrTriggerName field to given value.


### GetPropertyName

`func (o *AutoGetPieceOptionsRequest) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *AutoGetPieceOptionsRequest) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *AutoGetPieceOptionsRequest) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.


### GetPieceName

`func (o *AutoGetPieceOptionsRequest) GetPieceName() string`

GetPieceName returns the PieceName field if non-nil, zero value otherwise.

### GetPieceNameOk

`func (o *AutoGetPieceOptionsRequest) GetPieceNameOk() (*string, bool)`

GetPieceNameOk returns a tuple with the PieceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceName

`func (o *AutoGetPieceOptionsRequest) SetPieceName(v string)`

SetPieceName sets PieceName field to given value.


### GetPieceVersion

`func (o *AutoGetPieceOptionsRequest) GetPieceVersion() string`

GetPieceVersion returns the PieceVersion field if non-nil, zero value otherwise.

### GetPieceVersionOk

`func (o *AutoGetPieceOptionsRequest) GetPieceVersionOk() (*string, bool)`

GetPieceVersionOk returns a tuple with the PieceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceVersion

`func (o *AutoGetPieceOptionsRequest) SetPieceVersion(v string)`

SetPieceVersion sets PieceVersion field to given value.


### GetInput

`func (o *AutoGetPieceOptionsRequest) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AutoGetPieceOptionsRequest) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AutoGetPieceOptionsRequest) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *AutoGetPieceOptionsRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetSearchValue

`func (o *AutoGetPieceOptionsRequest) GetSearchValue() string`

GetSearchValue returns the SearchValue field if non-nil, zero value otherwise.

### GetSearchValueOk

`func (o *AutoGetPieceOptionsRequest) GetSearchValueOk() (*string, bool)`

GetSearchValueOk returns a tuple with the SearchValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchValue

`func (o *AutoGetPieceOptionsRequest) SetSearchValue(v string)`

SetSearchValue sets SearchValue field to given value.

### HasSearchValue

`func (o *AutoGetPieceOptionsRequest) HasSearchValue() bool`

HasSearchValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


