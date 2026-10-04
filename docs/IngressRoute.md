# IngressRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the exact hostname this route matches, lowercased with any trailing dot stripped. It is a GLOBALLY unique claim — one route across the whole edge may hold a host, so no tenant can hijack another&#39;s. | [optional] 
**Id** | Pointer to **string** | ID identifies the route within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. | [optional] 
**Middlewares** | Pointer to **[]string** | Middlewares are the ids of the edge transforms to apply, in this order, before the request reaches the service. At most 16. | [optional] 
**PathPrefix** | Pointer to **string** | PathPrefix narrows the match to requests under this path; it must start with \&quot;/\&quot;. Empty matches every path on the host. | [optional] 
**Priority** | Pointer to **int64** | Priority orders routes that share a host: higher wins, and equal priorities fall back to the longer PathPrefix. | [optional] 
**Service** | Pointer to **string** | Service is the id of the backend pool this route dispatches to. A route naming a service that does not exist is skipped at compile, not served. | [optional] 
**Tls** | Pointer to **bool** | TLS asks the edge to terminate TLS for Host with an ACME-managed certificate. | [optional] 

## Methods

### NewIngressRoute

`func NewIngressRoute() *IngressRoute`

NewIngressRoute instantiates a new IngressRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressRouteWithDefaults

`func NewIngressRouteWithDefaults() *IngressRoute`

NewIngressRouteWithDefaults instantiates a new IngressRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *IngressRoute) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IngressRoute) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IngressRoute) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IngressRoute) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *IngressRoute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngressRoute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngressRoute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IngressRoute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMiddlewares

`func (o *IngressRoute) GetMiddlewares() []string`

GetMiddlewares returns the Middlewares field if non-nil, zero value otherwise.

### GetMiddlewaresOk

`func (o *IngressRoute) GetMiddlewaresOk() (*[]string, bool)`

GetMiddlewaresOk returns a tuple with the Middlewares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddlewares

`func (o *IngressRoute) SetMiddlewares(v []string)`

SetMiddlewares sets Middlewares field to given value.

### HasMiddlewares

`func (o *IngressRoute) HasMiddlewares() bool`

HasMiddlewares returns a boolean if a field has been set.

### GetPathPrefix

`func (o *IngressRoute) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *IngressRoute) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *IngressRoute) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *IngressRoute) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetPriority

`func (o *IngressRoute) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IngressRoute) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IngressRoute) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IngressRoute) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetService

`func (o *IngressRoute) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IngressRoute) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IngressRoute) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *IngressRoute) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTls

`func (o *IngressRoute) GetTls() bool`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IngressRoute) GetTlsOk() (*bool, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IngressRoute) SetTls(v bool)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IngressRoute) HasTls() bool`

HasTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


