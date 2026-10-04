# AutoRunIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is the name of the provider action to invoke. | [optional] 
**Auth** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID is the provider to run, from the path. | [optional] 
**Props** | Pointer to **map[string]interface{}** | Props are the action&#39;s input properties, keyed by property name. | [optional] 

## Methods

### NewAutoRunIn

`func NewAutoRunIn() *AutoRunIn`

NewAutoRunIn instantiates a new AutoRunIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRunInWithDefaults

`func NewAutoRunInWithDefaults() *AutoRunIn`

NewAutoRunInWithDefaults instantiates a new AutoRunIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AutoRunIn) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AutoRunIn) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AutoRunIn) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AutoRunIn) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAuth

`func (o *AutoRunIn) GetAuth() interface{}`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *AutoRunIn) GetAuthOk() (*interface{}, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *AutoRunIn) SetAuth(v interface{})`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *AutoRunIn) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### SetAuthNil

`func (o *AutoRunIn) SetAuthNil(b bool)`

 SetAuthNil sets the value for Auth to be an explicit nil

### UnsetAuth
`func (o *AutoRunIn) UnsetAuth()`

UnsetAuth ensures that no value is present for Auth, not even an explicit nil
### GetId

`func (o *AutoRunIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoRunIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoRunIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoRunIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProps

`func (o *AutoRunIn) GetProps() map[string]interface{}`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *AutoRunIn) GetPropsOk() (*map[string]interface{}, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *AutoRunIn) SetProps(v map[string]interface{})`

SetProps sets Props field to given value.

### HasProps

`func (o *AutoRunIn) HasProps() bool`

HasProps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


