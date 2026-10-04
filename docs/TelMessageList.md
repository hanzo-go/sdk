# TelMessageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TelSMS**](TelSMS.md) | Data is this org&#39;s own messages, newest first — from our store rather than the carrier&#39;s, so it is the set an audit or a bill has to agree with. | [optional] 

## Methods

### NewTelMessageList

`func NewTelMessageList() *TelMessageList`

NewTelMessageList instantiates a new TelMessageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelMessageListWithDefaults

`func NewTelMessageListWithDefaults() *TelMessageList`

NewTelMessageListWithDefaults instantiates a new TelMessageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TelMessageList) GetData() []TelSMS`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TelMessageList) GetDataOk() (*[]TelSMS, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TelMessageList) SetData(v []TelSMS)`

SetData sets Data field to given value.

### HasData

`func (o *TelMessageList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


