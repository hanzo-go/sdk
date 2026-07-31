# KvStreamInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | Pointer to **int32** |  | [optional] 
**FirstEntry** | Pointer to [**KvStreamEntry**](KvStreamEntry.md) |  | [optional] 
**LastEntry** | Pointer to [**KvStreamEntry**](KvStreamEntry.md) |  | [optional] 
**Groups** | Pointer to **int32** |  | [optional] 

## Methods

### NewKvStreamInfo200Response

`func NewKvStreamInfo200Response() *KvStreamInfo200Response`

NewKvStreamInfo200Response instantiates a new KvStreamInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvStreamInfo200ResponseWithDefaults

`func NewKvStreamInfo200ResponseWithDefaults() *KvStreamInfo200Response`

NewKvStreamInfo200ResponseWithDefaults instantiates a new KvStreamInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *KvStreamInfo200Response) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *KvStreamInfo200Response) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *KvStreamInfo200Response) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *KvStreamInfo200Response) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetFirstEntry

`func (o *KvStreamInfo200Response) GetFirstEntry() KvStreamEntry`

GetFirstEntry returns the FirstEntry field if non-nil, zero value otherwise.

### GetFirstEntryOk

`func (o *KvStreamInfo200Response) GetFirstEntryOk() (*KvStreamEntry, bool)`

GetFirstEntryOk returns a tuple with the FirstEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstEntry

`func (o *KvStreamInfo200Response) SetFirstEntry(v KvStreamEntry)`

SetFirstEntry sets FirstEntry field to given value.

### HasFirstEntry

`func (o *KvStreamInfo200Response) HasFirstEntry() bool`

HasFirstEntry returns a boolean if a field has been set.

### GetLastEntry

`func (o *KvStreamInfo200Response) GetLastEntry() KvStreamEntry`

GetLastEntry returns the LastEntry field if non-nil, zero value otherwise.

### GetLastEntryOk

`func (o *KvStreamInfo200Response) GetLastEntryOk() (*KvStreamEntry, bool)`

GetLastEntryOk returns a tuple with the LastEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEntry

`func (o *KvStreamInfo200Response) SetLastEntry(v KvStreamEntry)`

SetLastEntry sets LastEntry field to given value.

### HasLastEntry

`func (o *KvStreamInfo200Response) HasLastEntry() bool`

HasLastEntry returns a boolean if a field has been set.

### GetGroups

`func (o *KvStreamInfo200Response) GetGroups() int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *KvStreamInfo200Response) GetGroupsOk() (*int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *KvStreamInfo200Response) SetGroups(v int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *KvStreamInfo200Response) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


