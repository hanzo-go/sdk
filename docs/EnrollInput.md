# EnrollInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is a single recipient, normalized (lower-cased, trimmed) before use. Give this OR audienceId, never both and never neither. | [optional] 
**AudienceId** | Pointer to **string** | AudienceID fans the sequence out over a saved audience, resolved live to the org&#39;s mailable customers. Email only. | [optional] 
**Channel** | Pointer to **string** | Channel is the delivery surface; empty means email. An audience resolves mailboxes, so an audience enroll must be email. | [optional] 
**Id** | Pointer to **string** | ID is the sequence id from the path. | [optional] 

## Methods

### NewEnrollInput

`func NewEnrollInput() *EnrollInput`

NewEnrollInput instantiates a new EnrollInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollInputWithDefaults

`func NewEnrollInputWithDefaults() *EnrollInput`

NewEnrollInputWithDefaults instantiates a new EnrollInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *EnrollInput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EnrollInput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EnrollInput) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *EnrollInput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAudienceId

`func (o *EnrollInput) GetAudienceId() string`

GetAudienceId returns the AudienceId field if non-nil, zero value otherwise.

### GetAudienceIdOk

`func (o *EnrollInput) GetAudienceIdOk() (*string, bool)`

GetAudienceIdOk returns a tuple with the AudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceId

`func (o *EnrollInput) SetAudienceId(v string)`

SetAudienceId sets AudienceId field to given value.

### HasAudienceId

`func (o *EnrollInput) HasAudienceId() bool`

HasAudienceId returns a boolean if a field has been set.

### GetChannel

`func (o *EnrollInput) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *EnrollInput) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *EnrollInput) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *EnrollInput) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetId

`func (o *EnrollInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnrollInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnrollInput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnrollInput) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


