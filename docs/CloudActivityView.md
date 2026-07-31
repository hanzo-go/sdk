# CloudActivityView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | agent name | [optional] 
**At** | Pointer to **string** | RFC3339 UTC | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** | invoked|failed|created|updated (from real events) | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudActivityView

`func NewCloudActivityView() *CloudActivityView`

NewCloudActivityView instantiates a new CloudActivityView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudActivityViewWithDefaults

`func NewCloudActivityViewWithDefaults() *CloudActivityView`

NewCloudActivityViewWithDefaults instantiates a new CloudActivityView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *CloudActivityView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CloudActivityView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CloudActivityView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CloudActivityView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAt

`func (o *CloudActivityView) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *CloudActivityView) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *CloudActivityView) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *CloudActivityView) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetId

`func (o *CloudActivityView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudActivityView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudActivityView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudActivityView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudActivityView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudActivityView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudActivityView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudActivityView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMessage

`func (o *CloudActivityView) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudActivityView) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudActivityView) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudActivityView) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


