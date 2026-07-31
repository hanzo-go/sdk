# CloudTargetDeleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** | Deleted is true when the target was removed. | [optional] 
**Id** | Pointer to **string** | ID is the target that was removed. | [optional] 

## Methods

### NewCloudTargetDeleted

`func NewCloudTargetDeleted() *CloudTargetDeleted`

NewCloudTargetDeleted instantiates a new CloudTargetDeleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTargetDeletedWithDefaults

`func NewCloudTargetDeletedWithDefaults() *CloudTargetDeleted`

NewCloudTargetDeletedWithDefaults instantiates a new CloudTargetDeleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *CloudTargetDeleted) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CloudTargetDeleted) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CloudTargetDeleted) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CloudTargetDeleted) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *CloudTargetDeleted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudTargetDeleted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudTargetDeleted) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudTargetDeleted) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


