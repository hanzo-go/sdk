# CloudSequenceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the sequence id from the path. | [optional] 
**Status** | Pointer to **string** | Status is draft, active or archived. Required; there is no default here, unlike on create. Only an active sequence accepts enrollments. | [optional] 

## Methods

### NewCloudSequenceStatus

`func NewCloudSequenceStatus() *CloudSequenceStatus`

NewCloudSequenceStatus instantiates a new CloudSequenceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSequenceStatusWithDefaults

`func NewCloudSequenceStatusWithDefaults() *CloudSequenceStatus`

NewCloudSequenceStatusWithDefaults instantiates a new CloudSequenceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudSequenceStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSequenceStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSequenceStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSequenceStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *CloudSequenceStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudSequenceStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudSequenceStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudSequenceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


