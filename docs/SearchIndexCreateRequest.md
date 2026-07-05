# SearchIndexCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**PrimaryKey** | Pointer to **string** |  | [optional] 

## Methods

### NewSearchIndexCreateRequest

`func NewSearchIndexCreateRequest(uid string, ) *SearchIndexCreateRequest`

NewSearchIndexCreateRequest instantiates a new SearchIndexCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexCreateRequestWithDefaults

`func NewSearchIndexCreateRequestWithDefaults() *SearchIndexCreateRequest`

NewSearchIndexCreateRequestWithDefaults instantiates a new SearchIndexCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *SearchIndexCreateRequest) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SearchIndexCreateRequest) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SearchIndexCreateRequest) SetUid(v string)`

SetUid sets Uid field to given value.


### GetPrimaryKey

`func (o *SearchIndexCreateRequest) GetPrimaryKey() string`

GetPrimaryKey returns the PrimaryKey field if non-nil, zero value otherwise.

### GetPrimaryKeyOk

`func (o *SearchIndexCreateRequest) GetPrimaryKeyOk() (*string, bool)`

GetPrimaryKeyOk returns a tuple with the PrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKey

`func (o *SearchIndexCreateRequest) SetPrimaryKey(v string)`

SetPrimaryKey sets PrimaryKey field to given value.

### HasPrimaryKey

`func (o *SearchIndexCreateRequest) HasPrimaryKey() bool`

HasPrimaryKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


