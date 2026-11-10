# FormWelcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the mail&#39;s text: plain text, a blank line a new paragraph. At most 10000 characters. | [optional] 
**Subject** | Pointer to **string** | Subject is the mail&#39;s subject. At most 200 characters. | [optional] 

## Methods

### NewFormWelcome

`func NewFormWelcome() *FormWelcome`

NewFormWelcome instantiates a new FormWelcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormWelcomeWithDefaults

`func NewFormWelcomeWithDefaults() *FormWelcome`

NewFormWelcomeWithDefaults instantiates a new FormWelcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *FormWelcome) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *FormWelcome) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *FormWelcome) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *FormWelcome) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSubject

`func (o *FormWelcome) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *FormWelcome) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *FormWelcome) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *FormWelcome) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


