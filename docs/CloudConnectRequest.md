# CloudConnectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GithubLogin** | Pointer to **string** | GithubLogin is the account to link. Used only when IAM holds no linked account for the provider — a linked account is stronger proof and always wins. | [optional] 
**Login** | Pointer to **string** | Login is the provider-neutral alias for GithubLogin, preferred when both are sent. | [optional] 
**Provider** | Pointer to **string** | Provider is the forge to enrol with: github (the default) or gitlab. | [optional] 

## Methods

### NewCloudConnectRequest

`func NewCloudConnectRequest() *CloudConnectRequest`

NewCloudConnectRequest instantiates a new CloudConnectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConnectRequestWithDefaults

`func NewCloudConnectRequestWithDefaults() *CloudConnectRequest`

NewCloudConnectRequestWithDefaults instantiates a new CloudConnectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGithubLogin

`func (o *CloudConnectRequest) GetGithubLogin() string`

GetGithubLogin returns the GithubLogin field if non-nil, zero value otherwise.

### GetGithubLoginOk

`func (o *CloudConnectRequest) GetGithubLoginOk() (*string, bool)`

GetGithubLoginOk returns a tuple with the GithubLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLogin

`func (o *CloudConnectRequest) SetGithubLogin(v string)`

SetGithubLogin sets GithubLogin field to given value.

### HasGithubLogin

`func (o *CloudConnectRequest) HasGithubLogin() bool`

HasGithubLogin returns a boolean if a field has been set.

### GetLogin

`func (o *CloudConnectRequest) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CloudConnectRequest) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *CloudConnectRequest) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *CloudConnectRequest) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetProvider

`func (o *CloudConnectRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudConnectRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudConnectRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudConnectRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


