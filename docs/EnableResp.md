# EnableResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | Pointer to **string** | AccountToken is the org&#39;s own tunnel-account credential. Treat it as a secret: it is what the CLI enables an environment with. | [optional] 
**Controller** | Pointer to **string** | Controller is the public controller endpoint the CLI enables against. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the public frontend a share is published into, when the deployment names one. | [optional] 
**UrlTemplate** | Pointer to **string** | URLTemplate is the shape a share token expands to, so the CLI can print the resulting URL without asking again. | [optional] 

## Methods

### NewEnableResp

`func NewEnableResp() *EnableResp`

NewEnableResp instantiates a new EnableResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableRespWithDefaults

`func NewEnableRespWithDefaults() *EnableResp`

NewEnableRespWithDefaults instantiates a new EnableResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *EnableResp) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *EnableResp) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *EnableResp) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.

### HasAccountToken

`func (o *EnableResp) HasAccountToken() bool`

HasAccountToken returns a boolean if a field has been set.

### GetController

`func (o *EnableResp) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *EnableResp) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *EnableResp) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *EnableResp) HasController() bool`

HasController returns a boolean if a field has been set.

### GetNamespace

`func (o *EnableResp) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EnableResp) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EnableResp) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EnableResp) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetUrlTemplate

`func (o *EnableResp) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *EnableResp) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *EnableResp) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.

### HasUrlTemplate

`func (o *EnableResp) HasUrlTemplate() bool`

HasUrlTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


