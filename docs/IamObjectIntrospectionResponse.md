# IamObjectIntrospectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Aud** | Pointer to **[]string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Exp** | Pointer to **int64** |  | [optional] 
**Iat** | Pointer to **int64** |  | [optional] 
**Iss** | Pointer to **string** |  | [optional] 
**Jti** | Pointer to **string** |  | [optional] 
**Nbf** | Pointer to **int64** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Sub** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectIntrospectionResponse

`func NewIamObjectIntrospectionResponse() *IamObjectIntrospectionResponse`

NewIamObjectIntrospectionResponse instantiates a new IamObjectIntrospectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectIntrospectionResponseWithDefaults

`func NewIamObjectIntrospectionResponseWithDefaults() *IamObjectIntrospectionResponse`

NewIamObjectIntrospectionResponseWithDefaults instantiates a new IamObjectIntrospectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *IamObjectIntrospectionResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *IamObjectIntrospectionResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *IamObjectIntrospectionResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *IamObjectIntrospectionResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAud

`func (o *IamObjectIntrospectionResponse) GetAud() []string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *IamObjectIntrospectionResponse) GetAudOk() (*[]string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *IamObjectIntrospectionResponse) SetAud(v []string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *IamObjectIntrospectionResponse) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetClientId

`func (o *IamObjectIntrospectionResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamObjectIntrospectionResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamObjectIntrospectionResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamObjectIntrospectionResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetExp

`func (o *IamObjectIntrospectionResponse) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *IamObjectIntrospectionResponse) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *IamObjectIntrospectionResponse) SetExp(v int64)`

SetExp sets Exp field to given value.

### HasExp

`func (o *IamObjectIntrospectionResponse) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetIat

`func (o *IamObjectIntrospectionResponse) GetIat() int64`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *IamObjectIntrospectionResponse) GetIatOk() (*int64, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *IamObjectIntrospectionResponse) SetIat(v int64)`

SetIat sets Iat field to given value.

### HasIat

`func (o *IamObjectIntrospectionResponse) HasIat() bool`

HasIat returns a boolean if a field has been set.

### GetIss

`func (o *IamObjectIntrospectionResponse) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *IamObjectIntrospectionResponse) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *IamObjectIntrospectionResponse) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *IamObjectIntrospectionResponse) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetJti

`func (o *IamObjectIntrospectionResponse) GetJti() string`

GetJti returns the Jti field if non-nil, zero value otherwise.

### GetJtiOk

`func (o *IamObjectIntrospectionResponse) GetJtiOk() (*string, bool)`

GetJtiOk returns a tuple with the Jti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJti

`func (o *IamObjectIntrospectionResponse) SetJti(v string)`

SetJti sets Jti field to given value.

### HasJti

`func (o *IamObjectIntrospectionResponse) HasJti() bool`

HasJti returns a boolean if a field has been set.

### GetNbf

`func (o *IamObjectIntrospectionResponse) GetNbf() int64`

GetNbf returns the Nbf field if non-nil, zero value otherwise.

### GetNbfOk

`func (o *IamObjectIntrospectionResponse) GetNbfOk() (*int64, bool)`

GetNbfOk returns a tuple with the Nbf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbf

`func (o *IamObjectIntrospectionResponse) SetNbf(v int64)`

SetNbf sets Nbf field to given value.

### HasNbf

`func (o *IamObjectIntrospectionResponse) HasNbf() bool`

HasNbf returns a boolean if a field has been set.

### GetScope

`func (o *IamObjectIntrospectionResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IamObjectIntrospectionResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IamObjectIntrospectionResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IamObjectIntrospectionResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSub

`func (o *IamObjectIntrospectionResponse) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *IamObjectIntrospectionResponse) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *IamObjectIntrospectionResponse) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *IamObjectIntrospectionResponse) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetTokenType

`func (o *IamObjectIntrospectionResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *IamObjectIntrospectionResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *IamObjectIntrospectionResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *IamObjectIntrospectionResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetUsername

`func (o *IamObjectIntrospectionResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IamObjectIntrospectionResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IamObjectIntrospectionResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IamObjectIntrospectionResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


