# IngressRoutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | Pointer to [**[]Route**](Route.md) | Routes is the org&#39;s routes, ordered by id. | [optional] 

## Methods

### NewIngressRoutes

`func NewIngressRoutes() *IngressRoutes`

NewIngressRoutes instantiates a new IngressRoutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressRoutesWithDefaults

`func NewIngressRoutesWithDefaults() *IngressRoutes`

NewIngressRoutesWithDefaults instantiates a new IngressRoutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *IngressRoutes) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *IngressRoutes) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *IngressRoutes) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *IngressRoutes) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


