# BoardView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditUrl** | Pointer to **string** |  | [optional] 
**Configured** | Pointer to **bool** |  | [optional] 
**Engine** | Pointer to **string** |  | [optional] 
**ManageUrl** | Pointer to **string** |  | [optional] 
**Switches** | Pointer to [**[]SwitchView**](SwitchView.md) |  | [optional] 

## Methods

### NewBoardView

`func NewBoardView() *BoardView`

NewBoardView instantiates a new BoardView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardViewWithDefaults

`func NewBoardViewWithDefaults() *BoardView`

NewBoardViewWithDefaults instantiates a new BoardView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditUrl

`func (o *BoardView) GetAuditUrl() string`

GetAuditUrl returns the AuditUrl field if non-nil, zero value otherwise.

### GetAuditUrlOk

`func (o *BoardView) GetAuditUrlOk() (*string, bool)`

GetAuditUrlOk returns a tuple with the AuditUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditUrl

`func (o *BoardView) SetAuditUrl(v string)`

SetAuditUrl sets AuditUrl field to given value.

### HasAuditUrl

`func (o *BoardView) HasAuditUrl() bool`

HasAuditUrl returns a boolean if a field has been set.

### GetConfigured

`func (o *BoardView) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *BoardView) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *BoardView) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *BoardView) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetEngine

`func (o *BoardView) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *BoardView) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *BoardView) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *BoardView) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetManageUrl

`func (o *BoardView) GetManageUrl() string`

GetManageUrl returns the ManageUrl field if non-nil, zero value otherwise.

### GetManageUrlOk

`func (o *BoardView) GetManageUrlOk() (*string, bool)`

GetManageUrlOk returns a tuple with the ManageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageUrl

`func (o *BoardView) SetManageUrl(v string)`

SetManageUrl sets ManageUrl field to given value.

### HasManageUrl

`func (o *BoardView) HasManageUrl() bool`

HasManageUrl returns a boolean if a field has been set.

### GetSwitches

`func (o *BoardView) GetSwitches() []SwitchView`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *BoardView) GetSwitchesOk() (*[]SwitchView, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *BoardView) SetSwitches(v []SwitchView)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *BoardView) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


