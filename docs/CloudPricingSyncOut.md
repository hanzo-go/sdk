# CloudPricingSyncOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when the sync completed. | [optional] 
**Updated** | Pointer to **string** | Updated is the RFC 3339 time the refreshed catalog was stamped with. | [optional] 

## Methods

### NewCloudPricingSyncOut

`func NewCloudPricingSyncOut() *CloudPricingSyncOut`

NewCloudPricingSyncOut instantiates a new CloudPricingSyncOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPricingSyncOutWithDefaults

`func NewCloudPricingSyncOutWithDefaults() *CloudPricingSyncOut`

NewCloudPricingSyncOutWithDefaults instantiates a new CloudPricingSyncOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudPricingSyncOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudPricingSyncOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudPricingSyncOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudPricingSyncOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *CloudPricingSyncOut) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudPricingSyncOut) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudPricingSyncOut) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudPricingSyncOut) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


