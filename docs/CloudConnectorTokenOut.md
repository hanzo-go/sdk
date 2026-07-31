# CloudConnectorTokenOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **string** | ExpiresAt is when this token expires, RFC 3339 UTC; empty if non-expiring. | [optional] 
**Label** | Pointer to **string** | Label is the connector&#39;s label. | [optional] 
**Provider** | Pointer to **string** | Provider is the connector&#39;s provider id. | [optional] 
**Token** | Pointer to **string** | Token is the access token, rotated first if it was within the refresh window. | [optional] 

## Methods

### NewCloudConnectorTokenOut

`func NewCloudConnectorTokenOut() *CloudConnectorTokenOut`

NewCloudConnectorTokenOut instantiates a new CloudConnectorTokenOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConnectorTokenOutWithDefaults

`func NewCloudConnectorTokenOutWithDefaults() *CloudConnectorTokenOut`

NewCloudConnectorTokenOutWithDefaults instantiates a new CloudConnectorTokenOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *CloudConnectorTokenOut) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CloudConnectorTokenOut) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CloudConnectorTokenOut) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CloudConnectorTokenOut) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetLabel

`func (o *CloudConnectorTokenOut) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudConnectorTokenOut) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudConnectorTokenOut) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudConnectorTokenOut) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetProvider

`func (o *CloudConnectorTokenOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudConnectorTokenOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudConnectorTokenOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudConnectorTokenOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetToken

`func (o *CloudConnectorTokenOut) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CloudConnectorTokenOut) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CloudConnectorTokenOut) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CloudConnectorTokenOut) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


