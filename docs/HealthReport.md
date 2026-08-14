# HealthReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **bool** | Datastore reports whether the shared warehouse client has a live connection. It is load-bearing for the READ path: false is one of the two ways this answers 503. | [optional] 
**Lenses** | Pointer to [**HealthLenses**](HealthLenses.md) | Lenses is per-lens table availability, probed only when connected — so it is absent from a degraded report, which has nothing to say about tables it could not reach. | [optional] 
**Lost** | Pointer to [**Loss**](Loss.md) | Lost is the count of facts the sink irrecoverably dropped since boot (warehouse.go). It is reported on the DEGRADED report too, and deliberately: a warehouse that is unreachable is exactly when facts start failing their deliveries, so suppressing the number here would hide it precisely when it moves. ANY NON-ZERO VALUE IS AN ALARM — it counts data the door already answered 200 for. | [optional] 
**Plane** | Pointer to [**HealthPlane**](HealthPlane.md) | Plane reports the event plane — the bus and the stream every accepted event is published to BEFORE any of it reaches the warehouse. It is load-bearing for the WRITE path, and it is here because its absence was a real outage: this endpoint answered 200/ok on warehouse connectivity alone while every POST /v1/event 503&#39;d on a stream that could not bind, so 100% ingest loss was invisible to monitoring. A probe that cannot see the write path cannot report the write path. | [optional] 
**Reason** | Pointer to **string** | Reason is the human-readable cause, present only on a degraded report. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem answering, so a probe aggregating several health endpoints can attribute a degraded one. | [optional] 
**Status** | Pointer to **string** | Status is ok or degraded. Degraded is the 503 and means EITHER load-bearing dependency is down — the warehouse this subsystem reads, or the event plane it writes. It is not moved by a missing lens table, which is honest-empty. | [optional] 
**Warehouse** | Pointer to **string** | Warehouse names the datastore database every lens reads. | [optional] 

## Methods

### NewHealthReport

`func NewHealthReport() *HealthReport`

NewHealthReport instantiates a new HealthReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthReportWithDefaults

`func NewHealthReportWithDefaults() *HealthReport`

NewHealthReportWithDefaults instantiates a new HealthReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *HealthReport) GetDatastore() bool`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *HealthReport) GetDatastoreOk() (*bool, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *HealthReport) SetDatastore(v bool)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *HealthReport) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetLenses

`func (o *HealthReport) GetLenses() HealthLenses`

GetLenses returns the Lenses field if non-nil, zero value otherwise.

### GetLensesOk

`func (o *HealthReport) GetLensesOk() (*HealthLenses, bool)`

GetLensesOk returns a tuple with the Lenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLenses

`func (o *HealthReport) SetLenses(v HealthLenses)`

SetLenses sets Lenses field to given value.

### HasLenses

`func (o *HealthReport) HasLenses() bool`

HasLenses returns a boolean if a field has been set.

### GetLost

`func (o *HealthReport) GetLost() Loss`

GetLost returns the Lost field if non-nil, zero value otherwise.

### GetLostOk

`func (o *HealthReport) GetLostOk() (*Loss, bool)`

GetLostOk returns a tuple with the Lost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLost

`func (o *HealthReport) SetLost(v Loss)`

SetLost sets Lost field to given value.

### HasLost

`func (o *HealthReport) HasLost() bool`

HasLost returns a boolean if a field has been set.

### GetPlane

`func (o *HealthReport) GetPlane() HealthPlane`

GetPlane returns the Plane field if non-nil, zero value otherwise.

### GetPlaneOk

`func (o *HealthReport) GetPlaneOk() (*HealthPlane, bool)`

GetPlaneOk returns a tuple with the Plane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlane

`func (o *HealthReport) SetPlane(v HealthPlane)`

SetPlane sets Plane field to given value.

### HasPlane

`func (o *HealthReport) HasPlane() bool`

HasPlane returns a boolean if a field has been set.

### GetReason

`func (o *HealthReport) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HealthReport) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HealthReport) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HealthReport) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetService

`func (o *HealthReport) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *HealthReport) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *HealthReport) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *HealthReport) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *HealthReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWarehouse

`func (o *HealthReport) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *HealthReport) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *HealthReport) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *HealthReport) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


