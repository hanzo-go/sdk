# O11yO11yEventIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]map[string]interface{}** | Attributes are free-form event properties. | [optional] 
**EventName** | Pointer to **string** | EventName names the event; required for track events. | [optional] 
**EventType** | **string** | EventType is the kind of event — track, identify or group. Required. | 
**RateLimited** | Pointer to **bool** | RateLimited marks an event the reporting client rate-limited. | [optional] 

## Methods

### NewO11yO11yEventIn

`func NewO11yO11yEventIn(eventType string, ) *O11yO11yEventIn`

NewO11yO11yEventIn instantiates a new O11yO11yEventIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yEventInWithDefaults

`func NewO11yO11yEventInWithDefaults() *O11yO11yEventIn`

NewO11yO11yEventInWithDefaults instantiates a new O11yO11yEventIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *O11yO11yEventIn) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *O11yO11yEventIn) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *O11yO11yEventIn) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *O11yO11yEventIn) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEventName

`func (o *O11yO11yEventIn) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *O11yO11yEventIn) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *O11yO11yEventIn) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *O11yO11yEventIn) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetEventType

`func (o *O11yO11yEventIn) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *O11yO11yEventIn) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *O11yO11yEventIn) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetRateLimited

`func (o *O11yO11yEventIn) GetRateLimited() bool`

GetRateLimited returns the RateLimited field if non-nil, zero value otherwise.

### GetRateLimitedOk

`func (o *O11yO11yEventIn) GetRateLimitedOk() (*bool, bool)`

GetRateLimitedOk returns a tuple with the RateLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimited

`func (o *O11yO11yEventIn) SetRateLimited(v bool)`

SetRateLimited sets RateLimited field to given value.

### HasRateLimited

`func (o *O11yO11yEventIn) HasRateLimited() bool`

HasRateLimited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


