# ConsoleIngestionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Body** | **map[string]interface{}** |  | 

## Methods

### NewConsoleIngestionEvent

`func NewConsoleIngestionEvent(id string, type_ string, timestamp time.Time, body map[string]interface{}, ) *ConsoleIngestionEvent`

NewConsoleIngestionEvent instantiates a new ConsoleIngestionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleIngestionEventWithDefaults

`func NewConsoleIngestionEventWithDefaults() *ConsoleIngestionEvent`

NewConsoleIngestionEventWithDefaults instantiates a new ConsoleIngestionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleIngestionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleIngestionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleIngestionEvent) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ConsoleIngestionEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsoleIngestionEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsoleIngestionEvent) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *ConsoleIngestionEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ConsoleIngestionEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ConsoleIngestionEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetMetadata

`func (o *ConsoleIngestionEvent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsoleIngestionEvent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsoleIngestionEvent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsoleIngestionEvent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetBody

`func (o *ConsoleIngestionEvent) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ConsoleIngestionEvent) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ConsoleIngestionEvent) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


