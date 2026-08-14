# RunIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is the name of the connector action to invoke. | [optional] 
**Auth** | Pointer to **map[string]interface{}** | Auth is the caller&#39;s resolved credential for the connector, handed to the action verbatim. Its shape is whatever the connector&#39;s auth descriptor declares (a token string, an object), so it is opaque here. | [optional] 
**Id** | Pointer to **string** | ID is the connector to run, from the path. | [optional] 
**Props** | Pointer to **map[string]map[string]interface{}** | Props are the action&#39;s input properties, keyed by property name. | [optional] 

## Methods

### NewRunIn

`func NewRunIn() *RunIn`

NewRunIn instantiates a new RunIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunInWithDefaults

`func NewRunInWithDefaults() *RunIn`

NewRunInWithDefaults instantiates a new RunIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RunIn) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RunIn) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RunIn) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RunIn) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAuth

`func (o *RunIn) GetAuth() map[string]interface{}`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *RunIn) GetAuthOk() (*map[string]interface{}, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *RunIn) SetAuth(v map[string]interface{})`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *RunIn) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetId

`func (o *RunIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RunIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProps

`func (o *RunIn) GetProps() map[string]map[string]interface{}`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *RunIn) GetPropsOk() (*map[string]map[string]interface{}, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *RunIn) SetProps(v map[string]map[string]interface{})`

SetProps sets Props field to given value.

### HasProps

`func (o *RunIn) HasProps() bool`

HasProps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


