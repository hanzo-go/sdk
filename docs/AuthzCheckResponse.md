# AuthzCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | Pointer to **bool** | Whether the tuple is permitted by the org&#39;s policy set. | [optional] 
**Sub** | Pointer to **string** |  | [optional] 
**Obj** | Pointer to **string** |  | [optional] 
**Act** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthzCheckResponse

`func NewAuthzCheckResponse() *AuthzCheckResponse`

NewAuthzCheckResponse instantiates a new AuthzCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthzCheckResponseWithDefaults

`func NewAuthzCheckResponseWithDefaults() *AuthzCheckResponse`

NewAuthzCheckResponseWithDefaults instantiates a new AuthzCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *AuthzCheckResponse) GetAllow() bool`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *AuthzCheckResponse) GetAllowOk() (*bool, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *AuthzCheckResponse) SetAllow(v bool)`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *AuthzCheckResponse) HasAllow() bool`

HasAllow returns a boolean if a field has been set.

### GetSub

`func (o *AuthzCheckResponse) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *AuthzCheckResponse) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *AuthzCheckResponse) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *AuthzCheckResponse) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetObj

`func (o *AuthzCheckResponse) GetObj() string`

GetObj returns the Obj field if non-nil, zero value otherwise.

### GetObjOk

`func (o *AuthzCheckResponse) GetObjOk() (*string, bool)`

GetObjOk returns a tuple with the Obj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObj

`func (o *AuthzCheckResponse) SetObj(v string)`

SetObj sets Obj field to given value.

### HasObj

`func (o *AuthzCheckResponse) HasObj() bool`

HasObj returns a boolean if a field has been set.

### GetAct

`func (o *AuthzCheckResponse) GetAct() string`

GetAct returns the Act field if non-nil, zero value otherwise.

### GetActOk

`func (o *AuthzCheckResponse) GetActOk() (*string, bool)`

GetActOk returns a tuple with the Act field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAct

`func (o *AuthzCheckResponse) SetAct(v string)`

SetAct sets Act field to given value.

### HasAct

`func (o *AuthzCheckResponse) HasAct() bool`

HasAct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


