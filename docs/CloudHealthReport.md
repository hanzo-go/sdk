# CloudHealthReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **bool** |  | [optional] 
**Lenses** | Pointer to [**CloudHealthLenses**](CloudHealthLenses.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Warehouse** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudHealthReport

`func NewCloudHealthReport() *CloudHealthReport`

NewCloudHealthReport instantiates a new CloudHealthReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudHealthReportWithDefaults

`func NewCloudHealthReportWithDefaults() *CloudHealthReport`

NewCloudHealthReportWithDefaults instantiates a new CloudHealthReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *CloudHealthReport) GetDatastore() bool`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *CloudHealthReport) GetDatastoreOk() (*bool, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *CloudHealthReport) SetDatastore(v bool)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *CloudHealthReport) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetLenses

`func (o *CloudHealthReport) GetLenses() CloudHealthLenses`

GetLenses returns the Lenses field if non-nil, zero value otherwise.

### GetLensesOk

`func (o *CloudHealthReport) GetLensesOk() (*CloudHealthLenses, bool)`

GetLensesOk returns a tuple with the Lenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLenses

`func (o *CloudHealthReport) SetLenses(v CloudHealthLenses)`

SetLenses sets Lenses field to given value.

### HasLenses

`func (o *CloudHealthReport) HasLenses() bool`

HasLenses returns a boolean if a field has been set.

### GetReason

`func (o *CloudHealthReport) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudHealthReport) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudHealthReport) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudHealthReport) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetService

`func (o *CloudHealthReport) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudHealthReport) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudHealthReport) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CloudHealthReport) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *CloudHealthReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudHealthReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudHealthReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudHealthReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWarehouse

`func (o *CloudHealthReport) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *CloudHealthReport) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *CloudHealthReport) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *CloudHealthReport) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


