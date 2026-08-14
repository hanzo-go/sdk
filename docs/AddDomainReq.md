# AddDomainReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the application&#39;s slug, from the path. | [optional] 
**Host** | Pointer to **string** | Host is the hostname to attach. Required, and must be a valid DNS hostname. | [optional] 
**Project** | Pointer to **string** | Project is the project the application lives under, from the path. | [optional] 

## Methods

### NewAddDomainReq

`func NewAddDomainReq() *AddDomainReq`

NewAddDomainReq instantiates a new AddDomainReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDomainReqWithDefaults

`func NewAddDomainReqWithDefaults() *AddDomainReq`

NewAddDomainReqWithDefaults instantiates a new AddDomainReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *AddDomainReq) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AddDomainReq) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AddDomainReq) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *AddDomainReq) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetHost

`func (o *AddDomainReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AddDomainReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AddDomainReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AddDomainReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetProject

`func (o *AddDomainReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AddDomainReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AddDomainReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *AddDomainReq) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


