# CloudQuestionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Questions** | Pointer to [**[]CloudQuestion**](CloudQuestion.md) |  | [optional] 

## Methods

### NewCloudQuestionsResponse

`func NewCloudQuestionsResponse() *CloudQuestionsResponse`

NewCloudQuestionsResponse instantiates a new CloudQuestionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudQuestionsResponseWithDefaults

`func NewCloudQuestionsResponseWithDefaults() *CloudQuestionsResponse`

NewCloudQuestionsResponseWithDefaults instantiates a new CloudQuestionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestions

`func (o *CloudQuestionsResponse) GetQuestions() []CloudQuestion`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *CloudQuestionsResponse) GetQuestionsOk() (*[]CloudQuestion, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *CloudQuestionsResponse) SetQuestions(v []CloudQuestion)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *CloudQuestionsResponse) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


