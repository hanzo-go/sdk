# ProviderRunIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is the action&#39;s name, as GET /v1/provider/{provider} lists it. | [optional] 
**Input** | Pointer to **map[string]interface{}** | Input is the action&#39;s props, by name. | [optional] 
**Provider** | Pointer to **string** | Provider is the connector, from the path. | [optional] 

## Methods

### NewProviderRunIn

`func NewProviderRunIn() *ProviderRunIn`

NewProviderRunIn instantiates a new ProviderRunIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderRunInWithDefaults

`func NewProviderRunInWithDefaults() *ProviderRunIn`

NewProviderRunInWithDefaults instantiates a new ProviderRunIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ProviderRunIn) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProviderRunIn) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProviderRunIn) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ProviderRunIn) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetInput

`func (o *ProviderRunIn) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ProviderRunIn) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ProviderRunIn) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ProviderRunIn) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetProvider

`func (o *ProviderRunIn) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderRunIn) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderRunIn) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ProviderRunIn) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


