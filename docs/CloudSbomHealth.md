# CloudSbomHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **bool** | Datastore reports whether the shared datastore connection this subsystem reads and writes through is established. False means the data endpoints answer 503. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem answering: always \&quot;sbom\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is the liveness verdict: always \&quot;ok\&quot; here, because the process answering at all IS the liveness fact. | [optional] 
**Table** | Pointer to **string** | Table is the fully-qualified datastore table the components live in. | [optional] 

## Methods

### NewCloudSbomHealth

`func NewCloudSbomHealth() *CloudSbomHealth`

NewCloudSbomHealth instantiates a new CloudSbomHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSbomHealthWithDefaults

`func NewCloudSbomHealthWithDefaults() *CloudSbomHealth`

NewCloudSbomHealthWithDefaults instantiates a new CloudSbomHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *CloudSbomHealth) GetDatastore() bool`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *CloudSbomHealth) GetDatastoreOk() (*bool, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *CloudSbomHealth) SetDatastore(v bool)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *CloudSbomHealth) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetService

`func (o *CloudSbomHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudSbomHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudSbomHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CloudSbomHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *CloudSbomHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudSbomHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudSbomHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudSbomHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTable

`func (o *CloudSbomHealth) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *CloudSbomHealth) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *CloudSbomHealth) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *CloudSbomHealth) HasTable() bool`

HasTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


