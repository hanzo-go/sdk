# CloudStepInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the message text. Required. | [optional] 
**DelaySeconds** | Pointer to **int32** | DelaySeconds is how long after the previous step this one sends (after enrollment, for the first step). Must be &gt;&#x3D; 0. | [optional] 
**Id** | Pointer to **string** | SequenceID is the sequence id from the path (the route&#39;s :id). | [optional] 
**Subject** | Pointer to **string** | Subject is the email subject line, capped at 1024 bytes. | [optional] 

## Methods

### NewCloudStepInput

`func NewCloudStepInput() *CloudStepInput`

NewCloudStepInput instantiates a new CloudStepInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStepInputWithDefaults

`func NewCloudStepInputWithDefaults() *CloudStepInput`

NewCloudStepInputWithDefaults instantiates a new CloudStepInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CloudStepInput) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CloudStepInput) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CloudStepInput) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CloudStepInput) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDelaySeconds

`func (o *CloudStepInput) GetDelaySeconds() int32`

GetDelaySeconds returns the DelaySeconds field if non-nil, zero value otherwise.

### GetDelaySecondsOk

`func (o *CloudStepInput) GetDelaySecondsOk() (*int32, bool)`

GetDelaySecondsOk returns a tuple with the DelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaySeconds

`func (o *CloudStepInput) SetDelaySeconds(v int32)`

SetDelaySeconds sets DelaySeconds field to given value.

### HasDelaySeconds

`func (o *CloudStepInput) HasDelaySeconds() bool`

HasDelaySeconds returns a boolean if a field has been set.

### GetId

`func (o *CloudStepInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudStepInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudStepInput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudStepInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubject

`func (o *CloudStepInput) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudStepInput) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudStepInput) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudStepInput) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


