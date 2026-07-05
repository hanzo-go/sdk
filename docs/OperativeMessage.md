# OperativeMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] 
**Content** | Pointer to [**OperativeMessageContent**](OperativeMessageContent.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOperativeMessage

`func NewOperativeMessage() *OperativeMessage`

NewOperativeMessage instantiates a new OperativeMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperativeMessageWithDefaults

`func NewOperativeMessageWithDefaults() *OperativeMessage`

NewOperativeMessageWithDefaults instantiates a new OperativeMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *OperativeMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OperativeMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OperativeMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OperativeMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *OperativeMessage) GetContent() OperativeMessageContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OperativeMessage) GetContentOk() (*OperativeMessageContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OperativeMessage) SetContent(v OperativeMessageContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *OperativeMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTimestamp

`func (o *OperativeMessage) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OperativeMessage) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OperativeMessage) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OperativeMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


