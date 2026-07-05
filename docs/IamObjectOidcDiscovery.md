# IamObjectOidcDiscovery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationEndpoint** | Pointer to **string** |  | [optional] 
**ClaimsSupported** | Pointer to **[]string** |  | [optional] 
**DeviceAuthorizationEndpoint** | Pointer to **string** |  | [optional] 
**EndSessionEndpoint** | Pointer to **string** |  | [optional] 
**GrantTypesSupported** | Pointer to **[]string** |  | [optional] 
**IdTokenSigningAlgValuesSupported** | Pointer to **[]string** |  | [optional] 
**IntrospectionEndpoint** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**JwksUri** | Pointer to **string** |  | [optional] 
**RequestObjectSigningAlgValuesSupported** | Pointer to **[]string** |  | [optional] 
**RequestParameterSupported** | Pointer to **bool** |  | [optional] 
**ResponseModesSupported** | Pointer to **[]string** |  | [optional] 
**ResponseTypesSupported** | Pointer to **[]string** |  | [optional] 
**ScopesSupported** | Pointer to **[]string** |  | [optional] 
**SubjectTypesSupported** | Pointer to **[]string** |  | [optional] 
**TokenEndpoint** | Pointer to **string** |  | [optional] 
**UserinfoEndpoint** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectOidcDiscovery

`func NewIamObjectOidcDiscovery() *IamObjectOidcDiscovery`

NewIamObjectOidcDiscovery instantiates a new IamObjectOidcDiscovery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectOidcDiscoveryWithDefaults

`func NewIamObjectOidcDiscoveryWithDefaults() *IamObjectOidcDiscovery`

NewIamObjectOidcDiscoveryWithDefaults instantiates a new IamObjectOidcDiscovery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationEndpoint

`func (o *IamObjectOidcDiscovery) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *IamObjectOidcDiscovery) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *IamObjectOidcDiscovery) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *IamObjectOidcDiscovery) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetClaimsSupported

`func (o *IamObjectOidcDiscovery) GetClaimsSupported() []string`

GetClaimsSupported returns the ClaimsSupported field if non-nil, zero value otherwise.

### GetClaimsSupportedOk

`func (o *IamObjectOidcDiscovery) GetClaimsSupportedOk() (*[]string, bool)`

GetClaimsSupportedOk returns a tuple with the ClaimsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsSupported

`func (o *IamObjectOidcDiscovery) SetClaimsSupported(v []string)`

SetClaimsSupported sets ClaimsSupported field to given value.

### HasClaimsSupported

`func (o *IamObjectOidcDiscovery) HasClaimsSupported() bool`

HasClaimsSupported returns a boolean if a field has been set.

### GetDeviceAuthorizationEndpoint

`func (o *IamObjectOidcDiscovery) GetDeviceAuthorizationEndpoint() string`

GetDeviceAuthorizationEndpoint returns the DeviceAuthorizationEndpoint field if non-nil, zero value otherwise.

### GetDeviceAuthorizationEndpointOk

`func (o *IamObjectOidcDiscovery) GetDeviceAuthorizationEndpointOk() (*string, bool)`

GetDeviceAuthorizationEndpointOk returns a tuple with the DeviceAuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthorizationEndpoint

`func (o *IamObjectOidcDiscovery) SetDeviceAuthorizationEndpoint(v string)`

SetDeviceAuthorizationEndpoint sets DeviceAuthorizationEndpoint field to given value.

### HasDeviceAuthorizationEndpoint

`func (o *IamObjectOidcDiscovery) HasDeviceAuthorizationEndpoint() bool`

HasDeviceAuthorizationEndpoint returns a boolean if a field has been set.

### GetEndSessionEndpoint

`func (o *IamObjectOidcDiscovery) GetEndSessionEndpoint() string`

GetEndSessionEndpoint returns the EndSessionEndpoint field if non-nil, zero value otherwise.

### GetEndSessionEndpointOk

`func (o *IamObjectOidcDiscovery) GetEndSessionEndpointOk() (*string, bool)`

GetEndSessionEndpointOk returns a tuple with the EndSessionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSessionEndpoint

`func (o *IamObjectOidcDiscovery) SetEndSessionEndpoint(v string)`

SetEndSessionEndpoint sets EndSessionEndpoint field to given value.

### HasEndSessionEndpoint

`func (o *IamObjectOidcDiscovery) HasEndSessionEndpoint() bool`

HasEndSessionEndpoint returns a boolean if a field has been set.

### GetGrantTypesSupported

`func (o *IamObjectOidcDiscovery) GetGrantTypesSupported() []string`

GetGrantTypesSupported returns the GrantTypesSupported field if non-nil, zero value otherwise.

### GetGrantTypesSupportedOk

`func (o *IamObjectOidcDiscovery) GetGrantTypesSupportedOk() (*[]string, bool)`

GetGrantTypesSupportedOk returns a tuple with the GrantTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypesSupported

`func (o *IamObjectOidcDiscovery) SetGrantTypesSupported(v []string)`

SetGrantTypesSupported sets GrantTypesSupported field to given value.

### HasGrantTypesSupported

`func (o *IamObjectOidcDiscovery) HasGrantTypesSupported() bool`

HasGrantTypesSupported returns a boolean if a field has been set.

### GetIdTokenSigningAlgValuesSupported

`func (o *IamObjectOidcDiscovery) GetIdTokenSigningAlgValuesSupported() []string`

GetIdTokenSigningAlgValuesSupported returns the IdTokenSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetIdTokenSigningAlgValuesSupportedOk

`func (o *IamObjectOidcDiscovery) GetIdTokenSigningAlgValuesSupportedOk() (*[]string, bool)`

GetIdTokenSigningAlgValuesSupportedOk returns a tuple with the IdTokenSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenSigningAlgValuesSupported

`func (o *IamObjectOidcDiscovery) SetIdTokenSigningAlgValuesSupported(v []string)`

SetIdTokenSigningAlgValuesSupported sets IdTokenSigningAlgValuesSupported field to given value.

### HasIdTokenSigningAlgValuesSupported

`func (o *IamObjectOidcDiscovery) HasIdTokenSigningAlgValuesSupported() bool`

HasIdTokenSigningAlgValuesSupported returns a boolean if a field has been set.

### GetIntrospectionEndpoint

`func (o *IamObjectOidcDiscovery) GetIntrospectionEndpoint() string`

GetIntrospectionEndpoint returns the IntrospectionEndpoint field if non-nil, zero value otherwise.

### GetIntrospectionEndpointOk

`func (o *IamObjectOidcDiscovery) GetIntrospectionEndpointOk() (*string, bool)`

GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpoint

`func (o *IamObjectOidcDiscovery) SetIntrospectionEndpoint(v string)`

SetIntrospectionEndpoint sets IntrospectionEndpoint field to given value.

### HasIntrospectionEndpoint

`func (o *IamObjectOidcDiscovery) HasIntrospectionEndpoint() bool`

HasIntrospectionEndpoint returns a boolean if a field has been set.

### GetIssuer

`func (o *IamObjectOidcDiscovery) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *IamObjectOidcDiscovery) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *IamObjectOidcDiscovery) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *IamObjectOidcDiscovery) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJwksUri

`func (o *IamObjectOidcDiscovery) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *IamObjectOidcDiscovery) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *IamObjectOidcDiscovery) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *IamObjectOidcDiscovery) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetRequestObjectSigningAlgValuesSupported

`func (o *IamObjectOidcDiscovery) GetRequestObjectSigningAlgValuesSupported() []string`

GetRequestObjectSigningAlgValuesSupported returns the RequestObjectSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetRequestObjectSigningAlgValuesSupportedOk

`func (o *IamObjectOidcDiscovery) GetRequestObjectSigningAlgValuesSupportedOk() (*[]string, bool)`

GetRequestObjectSigningAlgValuesSupportedOk returns a tuple with the RequestObjectSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectSigningAlgValuesSupported

`func (o *IamObjectOidcDiscovery) SetRequestObjectSigningAlgValuesSupported(v []string)`

SetRequestObjectSigningAlgValuesSupported sets RequestObjectSigningAlgValuesSupported field to given value.

### HasRequestObjectSigningAlgValuesSupported

`func (o *IamObjectOidcDiscovery) HasRequestObjectSigningAlgValuesSupported() bool`

HasRequestObjectSigningAlgValuesSupported returns a boolean if a field has been set.

### GetRequestParameterSupported

`func (o *IamObjectOidcDiscovery) GetRequestParameterSupported() bool`

GetRequestParameterSupported returns the RequestParameterSupported field if non-nil, zero value otherwise.

### GetRequestParameterSupportedOk

`func (o *IamObjectOidcDiscovery) GetRequestParameterSupportedOk() (*bool, bool)`

GetRequestParameterSupportedOk returns a tuple with the RequestParameterSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameterSupported

`func (o *IamObjectOidcDiscovery) SetRequestParameterSupported(v bool)`

SetRequestParameterSupported sets RequestParameterSupported field to given value.

### HasRequestParameterSupported

`func (o *IamObjectOidcDiscovery) HasRequestParameterSupported() bool`

HasRequestParameterSupported returns a boolean if a field has been set.

### GetResponseModesSupported

`func (o *IamObjectOidcDiscovery) GetResponseModesSupported() []string`

GetResponseModesSupported returns the ResponseModesSupported field if non-nil, zero value otherwise.

### GetResponseModesSupportedOk

`func (o *IamObjectOidcDiscovery) GetResponseModesSupportedOk() (*[]string, bool)`

GetResponseModesSupportedOk returns a tuple with the ResponseModesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseModesSupported

`func (o *IamObjectOidcDiscovery) SetResponseModesSupported(v []string)`

SetResponseModesSupported sets ResponseModesSupported field to given value.

### HasResponseModesSupported

`func (o *IamObjectOidcDiscovery) HasResponseModesSupported() bool`

HasResponseModesSupported returns a boolean if a field has been set.

### GetResponseTypesSupported

`func (o *IamObjectOidcDiscovery) GetResponseTypesSupported() []string`

GetResponseTypesSupported returns the ResponseTypesSupported field if non-nil, zero value otherwise.

### GetResponseTypesSupportedOk

`func (o *IamObjectOidcDiscovery) GetResponseTypesSupportedOk() (*[]string, bool)`

GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypesSupported

`func (o *IamObjectOidcDiscovery) SetResponseTypesSupported(v []string)`

SetResponseTypesSupported sets ResponseTypesSupported field to given value.

### HasResponseTypesSupported

`func (o *IamObjectOidcDiscovery) HasResponseTypesSupported() bool`

HasResponseTypesSupported returns a boolean if a field has been set.

### GetScopesSupported

`func (o *IamObjectOidcDiscovery) GetScopesSupported() []string`

GetScopesSupported returns the ScopesSupported field if non-nil, zero value otherwise.

### GetScopesSupportedOk

`func (o *IamObjectOidcDiscovery) GetScopesSupportedOk() (*[]string, bool)`

GetScopesSupportedOk returns a tuple with the ScopesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesSupported

`func (o *IamObjectOidcDiscovery) SetScopesSupported(v []string)`

SetScopesSupported sets ScopesSupported field to given value.

### HasScopesSupported

`func (o *IamObjectOidcDiscovery) HasScopesSupported() bool`

HasScopesSupported returns a boolean if a field has been set.

### GetSubjectTypesSupported

`func (o *IamObjectOidcDiscovery) GetSubjectTypesSupported() []string`

GetSubjectTypesSupported returns the SubjectTypesSupported field if non-nil, zero value otherwise.

### GetSubjectTypesSupportedOk

`func (o *IamObjectOidcDiscovery) GetSubjectTypesSupportedOk() (*[]string, bool)`

GetSubjectTypesSupportedOk returns a tuple with the SubjectTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTypesSupported

`func (o *IamObjectOidcDiscovery) SetSubjectTypesSupported(v []string)`

SetSubjectTypesSupported sets SubjectTypesSupported field to given value.

### HasSubjectTypesSupported

`func (o *IamObjectOidcDiscovery) HasSubjectTypesSupported() bool`

HasSubjectTypesSupported returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *IamObjectOidcDiscovery) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *IamObjectOidcDiscovery) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *IamObjectOidcDiscovery) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *IamObjectOidcDiscovery) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetUserinfoEndpoint

`func (o *IamObjectOidcDiscovery) GetUserinfoEndpoint() string`

GetUserinfoEndpoint returns the UserinfoEndpoint field if non-nil, zero value otherwise.

### GetUserinfoEndpointOk

`func (o *IamObjectOidcDiscovery) GetUserinfoEndpointOk() (*string, bool)`

GetUserinfoEndpointOk returns a tuple with the UserinfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserinfoEndpoint

`func (o *IamObjectOidcDiscovery) SetUserinfoEndpoint(v string)`

SetUserinfoEndpoint sets UserinfoEndpoint field to given value.

### HasUserinfoEndpoint

`func (o *IamObjectOidcDiscovery) HasUserinfoEndpoint() bool`

HasUserinfoEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


