# DbUpdateDatabaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | [**DbUpdateDatabaseRequestDatabase**](DbUpdateDatabaseRequestDatabase.md) |  | 

## Methods

### NewDbUpdateDatabaseRequest

`func NewDbUpdateDatabaseRequest(database DbUpdateDatabaseRequestDatabase, ) *DbUpdateDatabaseRequest`

NewDbUpdateDatabaseRequest instantiates a new DbUpdateDatabaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbUpdateDatabaseRequestWithDefaults

`func NewDbUpdateDatabaseRequestWithDefaults() *DbUpdateDatabaseRequest`

NewDbUpdateDatabaseRequestWithDefaults instantiates a new DbUpdateDatabaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *DbUpdateDatabaseRequest) GetDatabase() DbUpdateDatabaseRequestDatabase`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DbUpdateDatabaseRequest) GetDatabaseOk() (*DbUpdateDatabaseRequestDatabase, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DbUpdateDatabaseRequest) SetDatabase(v DbUpdateDatabaseRequestDatabase)`

SetDatabase sets Database field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


