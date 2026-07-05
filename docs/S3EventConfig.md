# S3EventConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to [**S3EventConfigFilter**](S3EventConfigFilter.md) |  | [optional] 
**Destination** | Pointer to [**S3EventConfigDestination**](S3EventConfigDestination.md) |  | [optional] 

## Methods

### NewS3EventConfig

`func NewS3EventConfig() *S3EventConfig`

NewS3EventConfig instantiates a new S3EventConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3EventConfigWithDefaults

`func NewS3EventConfigWithDefaults() *S3EventConfig`

NewS3EventConfigWithDefaults instantiates a new S3EventConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3EventConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3EventConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3EventConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *S3EventConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEvents

`func (o *S3EventConfig) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *S3EventConfig) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *S3EventConfig) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *S3EventConfig) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFilter

`func (o *S3EventConfig) GetFilter() S3EventConfigFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *S3EventConfig) GetFilterOk() (*S3EventConfigFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *S3EventConfig) SetFilter(v S3EventConfigFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *S3EventConfig) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetDestination

`func (o *S3EventConfig) GetDestination() S3EventConfigDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *S3EventConfig) GetDestinationOk() (*S3EventConfigDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *S3EventConfig) SetDestination(v S3EventConfigDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *S3EventConfig) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


