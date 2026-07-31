# CloudIngressRoutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | Pointer to [**[]CloudRoute**](CloudRoute.md) | Routes is the org&#39;s routes, ordered by id. | [optional] 

## Methods

### NewCloudIngressRoutes

`func NewCloudIngressRoutes() *CloudIngressRoutes`

NewCloudIngressRoutes instantiates a new CloudIngressRoutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudIngressRoutesWithDefaults

`func NewCloudIngressRoutesWithDefaults() *CloudIngressRoutes`

NewCloudIngressRoutesWithDefaults instantiates a new CloudIngressRoutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *CloudIngressRoutes) GetRoutes() []CloudRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *CloudIngressRoutes) GetRoutesOk() (*[]CloudRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *CloudIngressRoutes) SetRoutes(v []CloudRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *CloudIngressRoutes) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


