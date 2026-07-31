# AuthzEnforceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sub** | **string** | Subject — the user or role requesting access. | 
**Obj** | **string** | Object — the resource being accessed. | 
**Act** | **string** | Action — the operation on the object. | 

## Methods

### NewAuthzEnforceRequest

`func NewAuthzEnforceRequest(sub string, obj string, act string, ) *AuthzEnforceRequest`

NewAuthzEnforceRequest instantiates a new AuthzEnforceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthzEnforceRequestWithDefaults

`func NewAuthzEnforceRequestWithDefaults() *AuthzEnforceRequest`

NewAuthzEnforceRequestWithDefaults instantiates a new AuthzEnforceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSub

`func (o *AuthzEnforceRequest) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *AuthzEnforceRequest) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *AuthzEnforceRequest) SetSub(v string)`

SetSub sets Sub field to given value.


### GetObj

`func (o *AuthzEnforceRequest) GetObj() string`

GetObj returns the Obj field if non-nil, zero value otherwise.

### GetObjOk

`func (o *AuthzEnforceRequest) GetObjOk() (*string, bool)`

GetObjOk returns a tuple with the Obj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObj

`func (o *AuthzEnforceRequest) SetObj(v string)`

SetObj sets Obj field to given value.


### GetAct

`func (o *AuthzEnforceRequest) GetAct() string`

GetAct returns the Act field if non-nil, zero value otherwise.

### GetActOk

`func (o *AuthzEnforceRequest) GetActOk() (*string, bool)`

GetActOk returns a tuple with the Act field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAct

`func (o *AuthzEnforceRequest) SetAct(v string)`

SetAct sets Act field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


