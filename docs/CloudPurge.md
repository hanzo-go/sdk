# CloudPurge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | Filter purges only messages on this org-relative subject (wildcards supported). | [optional] 
**Keep** | Pointer to **int32** | Keep retains that many newest messages. | [optional] 
**Name** | Pointer to **string** | Name is the stream name, from the path. | [optional] 

## Methods

### NewCloudPurge

`func NewCloudPurge() *CloudPurge`

NewCloudPurge instantiates a new CloudPurge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPurgeWithDefaults

`func NewCloudPurgeWithDefaults() *CloudPurge`

NewCloudPurgeWithDefaults instantiates a new CloudPurge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *CloudPurge) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CloudPurge) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CloudPurge) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CloudPurge) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetKeep

`func (o *CloudPurge) GetKeep() int32`

GetKeep returns the Keep field if non-nil, zero value otherwise.

### GetKeepOk

`func (o *CloudPurge) GetKeepOk() (*int32, bool)`

GetKeepOk returns a tuple with the Keep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeep

`func (o *CloudPurge) SetKeep(v int32)`

SetKeep sets Keep field to given value.

### HasKeep

`func (o *CloudPurge) HasKeep() bool`

HasKeep returns a boolean if a field has been set.

### GetName

`func (o *CloudPurge) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudPurge) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudPurge) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudPurge) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


