# Middleware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]string** | Config is the transform&#39;s parameters: redirectScheme takes scheme (default https) and permanent (\&quot;true\&quot; ⇒ 301, else 302); stripPrefix REQUIRES prefixes (comma-separated, first match wins); addPrefix REQUIRES prefix; headers is a header→value map set on the response. | [optional] 
**Id** | Pointer to **string** | ID identifies the transform within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. Routes reference it by this id. | [optional] 
**Type** | Pointer to **string** | Type is the transform: redirectScheme, stripPrefix, addPrefix or headers. | [optional] 

## Methods

### NewMiddleware

`func NewMiddleware() *Middleware`

NewMiddleware instantiates a new Middleware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiddlewareWithDefaults

`func NewMiddlewareWithDefaults() *Middleware`

NewMiddlewareWithDefaults instantiates a new Middleware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *Middleware) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Middleware) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Middleware) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Middleware) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetId

`func (o *Middleware) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Middleware) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Middleware) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Middleware) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Middleware) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Middleware) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Middleware) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Middleware) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


