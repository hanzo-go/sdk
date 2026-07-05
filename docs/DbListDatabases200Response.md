# DbListDatabases200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Databases** | Pointer to [**[]DbDatabase**](DbDatabase.md) |  | [optional] 

## Methods

### NewDbListDatabases200Response

`func NewDbListDatabases200Response() *DbListDatabases200Response`

NewDbListDatabases200Response instantiates a new DbListDatabases200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbListDatabases200ResponseWithDefaults

`func NewDbListDatabases200ResponseWithDefaults() *DbListDatabases200Response`

NewDbListDatabases200ResponseWithDefaults instantiates a new DbListDatabases200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabases

`func (o *DbListDatabases200Response) GetDatabases() []DbDatabase`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *DbListDatabases200Response) GetDatabasesOk() (*[]DbDatabase, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *DbListDatabases200Response) SetDatabases(v []DbDatabase)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *DbListDatabases200Response) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


