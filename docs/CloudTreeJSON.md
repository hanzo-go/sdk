# CloudTreeJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]CloudTreeEntryJSON**](CloudTreeEntryJSON.md) | Entries are the immediate children, directories before files. | [optional] 

## Methods

### NewCloudTreeJSON

`func NewCloudTreeJSON() *CloudTreeJSON`

NewCloudTreeJSON instantiates a new CloudTreeJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTreeJSONWithDefaults

`func NewCloudTreeJSONWithDefaults() *CloudTreeJSON`

NewCloudTreeJSONWithDefaults instantiates a new CloudTreeJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *CloudTreeJSON) GetEntries() []CloudTreeEntryJSON`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *CloudTreeJSON) GetEntriesOk() (*[]CloudTreeEntryJSON, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *CloudTreeJSON) SetEntries(v []CloudTreeEntryJSON)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *CloudTreeJSON) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


