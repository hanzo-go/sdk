# SetEnvReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the application&#39;s slug, from the path. | [optional] 
**Env** | Pointer to [**[]EnvVarJSON**](EnvVarJSON.md) | Env is the app&#39;s whole environment set, REPLACING what it had. Keys must match &#x60;^[A-Za-z_][A-Za-z0-9_]*$&#x60;; a variable marked &#x60;secret: true&#x60; is sealed into KMS and blanked in the database. | [optional] 
**Project** | Pointer to **string** | Project is the project the application lives under, from the path. | [optional] 

## Methods

### NewSetEnvReq

`func NewSetEnvReq() *SetEnvReq`

NewSetEnvReq instantiates a new SetEnvReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetEnvReqWithDefaults

`func NewSetEnvReqWithDefaults() *SetEnvReq`

NewSetEnvReqWithDefaults instantiates a new SetEnvReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *SetEnvReq) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *SetEnvReq) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *SetEnvReq) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *SetEnvReq) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetEnv

`func (o *SetEnvReq) GetEnv() []EnvVarJSON`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *SetEnvReq) GetEnvOk() (*[]EnvVarJSON, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *SetEnvReq) SetEnv(v []EnvVarJSON)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *SetEnvReq) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetProject

`func (o *SetEnvReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SetEnvReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SetEnvReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SetEnvReq) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


