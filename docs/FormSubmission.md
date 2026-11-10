# FormSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | Pointer to [**[]FormAnswer**](FormAnswer.md) | Answers are the answers, one per field at most. | [optional] 
**Completed** | Pointer to **bool** | Completed is false to save the answers so far and keep the response open (a later submission under the same token completes it); true or absent submits it whole, every required field answered. A completed response is final. | [optional] 
**Hidden** | Pointer to **map[string]string** | Hidden are the values the page passes in, by the names the form declares; each one line of at most 500 characters. | [optional] 
**Id** | Pointer to **string** | ID is the form&#39;s public id, from the path. | [optional] 
**LandedAt** | Pointer to **string** | LandedAt is when the respondent opened the form, RFC 3339. | [optional] 
**Path** | Pointer to **string** | Path is the page the form was answered on, at most 500 characters. | [optional] 
**Token** | Pointer to **string** | Token is the client&#39;s idempotency token for this response: 8 to 64 letters, digits, dashes and underscores. A submission sent twice under one token is one response, so a double click or a retry files nothing twice. Absent, the response is new each time. | [optional] 
**Visitor** | Pointer to **string** | Visitor is the page&#39;s visitor id: letters, digits, dot, dash and underscore, at most 64. It is who the response is from when nobody is signed in and no email field is answered. | [optional] 
**Website** | Pointer to **string** | Website is a field no person sees and no person fills: a submission that fills it is refused. | [optional] 

## Methods

### NewFormSubmission

`func NewFormSubmission() *FormSubmission`

NewFormSubmission instantiates a new FormSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSubmissionWithDefaults

`func NewFormSubmissionWithDefaults() *FormSubmission`

NewFormSubmissionWithDefaults instantiates a new FormSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *FormSubmission) GetAnswers() []FormAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *FormSubmission) GetAnswersOk() (*[]FormAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *FormSubmission) SetAnswers(v []FormAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *FormSubmission) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetCompleted

`func (o *FormSubmission) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *FormSubmission) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *FormSubmission) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *FormSubmission) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetHidden

`func (o *FormSubmission) GetHidden() map[string]string`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *FormSubmission) GetHiddenOk() (*map[string]string, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *FormSubmission) SetHidden(v map[string]string)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *FormSubmission) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *FormSubmission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormSubmission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormSubmission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormSubmission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLandedAt

`func (o *FormSubmission) GetLandedAt() string`

GetLandedAt returns the LandedAt field if non-nil, zero value otherwise.

### GetLandedAtOk

`func (o *FormSubmission) GetLandedAtOk() (*string, bool)`

GetLandedAtOk returns a tuple with the LandedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandedAt

`func (o *FormSubmission) SetLandedAt(v string)`

SetLandedAt sets LandedAt field to given value.

### HasLandedAt

`func (o *FormSubmission) HasLandedAt() bool`

HasLandedAt returns a boolean if a field has been set.

### GetPath

`func (o *FormSubmission) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FormSubmission) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FormSubmission) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FormSubmission) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetToken

`func (o *FormSubmission) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FormSubmission) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FormSubmission) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FormSubmission) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetVisitor

`func (o *FormSubmission) GetVisitor() string`

GetVisitor returns the Visitor field if non-nil, zero value otherwise.

### GetVisitorOk

`func (o *FormSubmission) GetVisitorOk() (*string, bool)`

GetVisitorOk returns a tuple with the Visitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitor

`func (o *FormSubmission) SetVisitor(v string)`

SetVisitor sets Visitor field to given value.

### HasVisitor

`func (o *FormSubmission) HasVisitor() bool`

HasVisitor returns a boolean if a field has been set.

### GetWebsite

`func (o *FormSubmission) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *FormSubmission) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *FormSubmission) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *FormSubmission) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


