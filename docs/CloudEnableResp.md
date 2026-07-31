# CloudEnableResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | Pointer to **string** | AccountToken is the org&#39;s own tunnel-account credential. Treat it as a secret: it is what the CLI enables an environment with. | [optional] 
**Controller** | Pointer to **string** | Controller is the public controller endpoint the CLI enables against. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the public frontend a share is published into, when the deployment names one. | [optional] 
**UrlTemplate** | Pointer to **string** | URLTemplate is the shape a share token expands to, so the CLI can print the resulting URL without asking again. | [optional] 

## Methods

### NewCloudEnableResp

`func NewCloudEnableResp() *CloudEnableResp`

NewCloudEnableResp instantiates a new CloudEnableResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEnableRespWithDefaults

`func NewCloudEnableRespWithDefaults() *CloudEnableResp`

NewCloudEnableRespWithDefaults instantiates a new CloudEnableResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *CloudEnableResp) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *CloudEnableResp) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *CloudEnableResp) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.

### HasAccountToken

`func (o *CloudEnableResp) HasAccountToken() bool`

HasAccountToken returns a boolean if a field has been set.

### GetController

`func (o *CloudEnableResp) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *CloudEnableResp) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *CloudEnableResp) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *CloudEnableResp) HasController() bool`

HasController returns a boolean if a field has been set.

### GetNamespace

`func (o *CloudEnableResp) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CloudEnableResp) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CloudEnableResp) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CloudEnableResp) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetUrlTemplate

`func (o *CloudEnableResp) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *CloudEnableResp) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *CloudEnableResp) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.

### HasUrlTemplate

`func (o *CloudEnableResp) HasUrlTemplate() bool`

HasUrlTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


