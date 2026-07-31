# O11yIngestEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | **string** | Event type (trace, observation, score). | 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Body** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewO11yIngestEvent

`func NewO11yIngestEvent(type_ string, ) *O11yIngestEvent`

NewO11yIngestEvent instantiates a new O11yIngestEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yIngestEventWithDefaults

`func NewO11yIngestEventWithDefaults() *O11yIngestEvent`

NewO11yIngestEventWithDefaults instantiates a new O11yIngestEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *O11yIngestEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yIngestEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yIngestEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yIngestEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *O11yIngestEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yIngestEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yIngestEvent) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *O11yIngestEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yIngestEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yIngestEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yIngestEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBody

`func (o *O11yIngestEvent) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *O11yIngestEvent) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *O11yIngestEvent) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *O11yIngestEvent) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


