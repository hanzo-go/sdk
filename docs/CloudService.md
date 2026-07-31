# CloudService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backends** | Pointer to [**[]CloudBackend**](CloudBackend.md) | Backends are the upstream servers to balance across: 1..32 of them. | [optional] 
**Id** | Pointer to **string** | ID identifies the pool within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. Routes reference it by this id. | [optional] 
**PassHostHeader** | Pointer to **bool** | PassHostHeader forwards the client&#39;s original Host header upstream instead of rewriting it to the backend&#39;s. | [optional] 

## Methods

### NewCloudService

`func NewCloudService() *CloudService`

NewCloudService instantiates a new CloudService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudServiceWithDefaults

`func NewCloudServiceWithDefaults() *CloudService`

NewCloudServiceWithDefaults instantiates a new CloudService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackends

`func (o *CloudService) GetBackends() []CloudBackend`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *CloudService) GetBackendsOk() (*[]CloudBackend, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *CloudService) SetBackends(v []CloudBackend)`

SetBackends sets Backends field to given value.

### HasBackends

`func (o *CloudService) HasBackends() bool`

HasBackends returns a boolean if a field has been set.

### GetId

`func (o *CloudService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassHostHeader

`func (o *CloudService) GetPassHostHeader() bool`

GetPassHostHeader returns the PassHostHeader field if non-nil, zero value otherwise.

### GetPassHostHeaderOk

`func (o *CloudService) GetPassHostHeaderOk() (*bool, bool)`

GetPassHostHeaderOk returns a tuple with the PassHostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassHostHeader

`func (o *CloudService) SetPassHostHeader(v bool)`

SetPassHostHeader sets PassHostHeader field to given value.

### HasPassHostHeader

`func (o *CloudService) HasPassHostHeader() bool`

HasPassHostHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


