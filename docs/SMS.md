# SMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | From is the sending number in E.164, and must be one this org holds. | [optional] 
**Id** | Pointer to **string** | ID is the carrier&#39;s handle for the message. | [optional] 
**Org** | Pointer to **string** | Org is the tenant the message was sent for or received by. | [optional] 
**Status** | Pointer to **string** | Status is where the message is: \&quot;queued\&quot;, \&quot;sent\&quot;, \&quot;delivered\&quot; or \&quot;failed\&quot;. \&quot;sent\&quot; means the carrier took it; \&quot;delivered\&quot; means the handset got it, and not every carrier or destination reports that. | [optional] 
**Text** | Pointer to **string** | Text is the message body. Empty is legal when the message carried only media. | [optional] 
**To** | Pointer to **string** | To is the receiving number in E.164. | [optional] 

## Methods

### NewSMS

`func NewSMS() *SMS`

NewSMS instantiates a new SMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSWithDefaults

`func NewSMSWithDefaults() *SMS`

NewSMSWithDefaults instantiates a new SMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *SMS) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SMS) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SMS) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SMS) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *SMS) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SMS) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SMS) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SMS) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrg

`func (o *SMS) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *SMS) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *SMS) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *SMS) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetStatus

`func (o *SMS) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SMS) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SMS) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SMS) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetText

`func (o *SMS) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SMS) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SMS) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SMS) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTo

`func (o *SMS) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SMS) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SMS) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *SMS) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


