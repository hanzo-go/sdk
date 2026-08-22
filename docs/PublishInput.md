# PublishInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctype** | Pointer to **string** | DocType is the content type holding the item: Campaign, SocialPost or Asset. Any other name is refused as an unknown content type. | [optional] 
**Name** | Pointer to **string** | Name is the document within that type — the item to distribute. Its caption, media and channel list come off the stored document, so this names WHICH item and says nothing about what goes out. | [optional] 
**ScheduleAt** | Pointer to **string** | ScheduleAt hands a future go-live to the channel&#39;s own scheduler, as an ISO-8601 time. Empty posts now. | [optional] 

## Methods

### NewPublishInput

`func NewPublishInput() *PublishInput`

NewPublishInput instantiates a new PublishInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishInputWithDefaults

`func NewPublishInputWithDefaults() *PublishInput`

NewPublishInputWithDefaults instantiates a new PublishInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctype

`func (o *PublishInput) GetDoctype() string`

GetDoctype returns the Doctype field if non-nil, zero value otherwise.

### GetDoctypeOk

`func (o *PublishInput) GetDoctypeOk() (*string, bool)`

GetDoctypeOk returns a tuple with the Doctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctype

`func (o *PublishInput) SetDoctype(v string)`

SetDoctype sets Doctype field to given value.

### HasDoctype

`func (o *PublishInput) HasDoctype() bool`

HasDoctype returns a boolean if a field has been set.

### GetName

`func (o *PublishInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublishInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublishInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublishInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleAt

`func (o *PublishInput) GetScheduleAt() string`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *PublishInput) GetScheduleAtOk() (*string, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *PublishInput) SetScheduleAt(v string)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *PublishInput) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


