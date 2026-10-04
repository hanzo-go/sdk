# TelMessageInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | From is the number to send FROM, in E.164. It must be one this org holds and it must be sms-capable. | [optional] 
**Media** | Pointer to **[]string** | Media are URLs to attach. A message with any is an MMS to the carrier — the distinction is the carrier&#39;s to make, not something the caller declares. | [optional] 
**Text** | Pointer to **string** | Text is the message body. It may be empty when Media carries the message. | [optional] 
**To** | Pointer to **string** | To is the number to send to, in E.164. | [optional] 

## Methods

### NewTelMessageInput

`func NewTelMessageInput() *TelMessageInput`

NewTelMessageInput instantiates a new TelMessageInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelMessageInputWithDefaults

`func NewTelMessageInputWithDefaults() *TelMessageInput`

NewTelMessageInputWithDefaults instantiates a new TelMessageInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *TelMessageInput) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TelMessageInput) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TelMessageInput) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TelMessageInput) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetMedia

`func (o *TelMessageInput) GetMedia() []string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *TelMessageInput) GetMediaOk() (*[]string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *TelMessageInput) SetMedia(v []string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *TelMessageInput) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetText

`func (o *TelMessageInput) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TelMessageInput) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TelMessageInput) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TelMessageInput) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTo

`func (o *TelMessageInput) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TelMessageInput) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TelMessageInput) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *TelMessageInput) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


