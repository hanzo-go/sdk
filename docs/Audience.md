# Audience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is unix seconds when the filter was saved, server-assigned. | [optional] 
**Event** | Pointer to **string** | Event is the analytics event a member must have fired. EMPTY MEANS NO FILTER: the audience is then every mailable customer in the org, and no warehouse is consulted. | [optional] 
**Id** | Pointer to **string** | ID is the server-assigned audience id (\&quot;aud_\&quot; + 128 random bits). | [optional] 
**Name** | Pointer to **string** | Name is the audience&#39;s label. Required, trimmed, capped at 1024 bytes. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is unix seconds of the last write, server-assigned, and the key the audience list is ordered by (newest first). A saved audience has no update route, so in practice it stays equal to CreatedAt: to change a filter you save another one. | [optional] 
**WindowDays** | Pointer to **int32** | WindowDays is how far back the event counts, ending now. 0 means 30 and nothing above 3650 is honoured. Ignored when Event is empty. | [optional] 

## Methods

### NewAudience

`func NewAudience() *Audience`

NewAudience instantiates a new Audience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceWithDefaults

`func NewAudienceWithDefaults() *Audience`

NewAudienceWithDefaults instantiates a new Audience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Audience) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Audience) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Audience) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Audience) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvent

`func (o *Audience) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Audience) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Audience) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Audience) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId

`func (o *Audience) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Audience) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Audience) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Audience) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Audience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Audience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Audience) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Audience) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Audience) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Audience) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Audience) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Audience) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWindowDays

`func (o *Audience) GetWindowDays() int32`

GetWindowDays returns the WindowDays field if non-nil, zero value otherwise.

### GetWindowDaysOk

`func (o *Audience) GetWindowDaysOk() (*int32, bool)`

GetWindowDaysOk returns a tuple with the WindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowDays

`func (o *Audience) SetWindowDays(v int32)`

SetWindowDays sets WindowDays field to given value.

### HasWindowDays

`func (o *Audience) HasWindowDays() bool`

HasWindowDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


