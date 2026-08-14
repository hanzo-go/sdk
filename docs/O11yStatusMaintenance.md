# O11yStatusMaintenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedComponents** | Pointer to [**[]O11yStatusComponent**](O11yStatusComponent.md) |  | [optional] 
**EndsAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdateAt** | Pointer to **string** |  | [optional] 
**LastUpdateMessage** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StartsAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yStatusMaintenance

`func NewO11yStatusMaintenance() *O11yStatusMaintenance`

NewO11yStatusMaintenance instantiates a new O11yStatusMaintenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yStatusMaintenanceWithDefaults

`func NewO11yStatusMaintenanceWithDefaults() *O11yStatusMaintenance`

NewO11yStatusMaintenanceWithDefaults instantiates a new O11yStatusMaintenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedComponents

`func (o *O11yStatusMaintenance) GetAffectedComponents() []O11yStatusComponent`

GetAffectedComponents returns the AffectedComponents field if non-nil, zero value otherwise.

### GetAffectedComponentsOk

`func (o *O11yStatusMaintenance) GetAffectedComponentsOk() (*[]O11yStatusComponent, bool)`

GetAffectedComponentsOk returns a tuple with the AffectedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedComponents

`func (o *O11yStatusMaintenance) SetAffectedComponents(v []O11yStatusComponent)`

SetAffectedComponents sets AffectedComponents field to given value.

### HasAffectedComponents

`func (o *O11yStatusMaintenance) HasAffectedComponents() bool`

HasAffectedComponents returns a boolean if a field has been set.

### GetEndsAt

`func (o *O11yStatusMaintenance) GetEndsAt() string`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *O11yStatusMaintenance) GetEndsAtOk() (*string, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *O11yStatusMaintenance) SetEndsAt(v string)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *O11yStatusMaintenance) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetId

`func (o *O11yStatusMaintenance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yStatusMaintenance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yStatusMaintenance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yStatusMaintenance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateAt

`func (o *O11yStatusMaintenance) GetLastUpdateAt() string`

GetLastUpdateAt returns the LastUpdateAt field if non-nil, zero value otherwise.

### GetLastUpdateAtOk

`func (o *O11yStatusMaintenance) GetLastUpdateAtOk() (*string, bool)`

GetLastUpdateAtOk returns a tuple with the LastUpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateAt

`func (o *O11yStatusMaintenance) SetLastUpdateAt(v string)`

SetLastUpdateAt sets LastUpdateAt field to given value.

### HasLastUpdateAt

`func (o *O11yStatusMaintenance) HasLastUpdateAt() bool`

HasLastUpdateAt returns a boolean if a field has been set.

### GetLastUpdateMessage

`func (o *O11yStatusMaintenance) GetLastUpdateMessage() string`

GetLastUpdateMessage returns the LastUpdateMessage field if non-nil, zero value otherwise.

### GetLastUpdateMessageOk

`func (o *O11yStatusMaintenance) GetLastUpdateMessageOk() (*string, bool)`

GetLastUpdateMessageOk returns a tuple with the LastUpdateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateMessage

`func (o *O11yStatusMaintenance) SetLastUpdateMessage(v string)`

SetLastUpdateMessage sets LastUpdateMessage field to given value.

### HasLastUpdateMessage

`func (o *O11yStatusMaintenance) HasLastUpdateMessage() bool`

HasLastUpdateMessage returns a boolean if a field has been set.

### GetName

`func (o *O11yStatusMaintenance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yStatusMaintenance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yStatusMaintenance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yStatusMaintenance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartsAt

`func (o *O11yStatusMaintenance) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *O11yStatusMaintenance) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *O11yStatusMaintenance) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *O11yStatusMaintenance) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetStatus

`func (o *O11yStatusMaintenance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yStatusMaintenance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yStatusMaintenance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yStatusMaintenance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *O11yStatusMaintenance) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *O11yStatusMaintenance) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *O11yStatusMaintenance) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *O11yStatusMaintenance) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


