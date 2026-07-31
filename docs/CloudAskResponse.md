# CloudAskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** | Answer is one or two sentences answering the question, every number in it taken from Figures. | [optional] 
**Figures** | Pointer to [**[]CloudFigure**](CloudFigure.md) | Figures are the grounded numbers the answer states, each already formatted. | [optional] 
**Followups** | Pointer to **[]string** | Followups are sharper questions to ask next, chosen from the same intent. | [optional] 
**Sources** | Pointer to **[]string** | Sources name the books reports the figures were computed from — \&quot;pnl\&quot;, \&quot;balance-sheet\&quot;, \&quot;trial-balance\&quot;. | [optional] 

## Methods

### NewCloudAskResponse

`func NewCloudAskResponse() *CloudAskResponse`

NewCloudAskResponse instantiates a new CloudAskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAskResponseWithDefaults

`func NewCloudAskResponseWithDefaults() *CloudAskResponse`

NewCloudAskResponseWithDefaults instantiates a new CloudAskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *CloudAskResponse) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *CloudAskResponse) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *CloudAskResponse) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *CloudAskResponse) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetFigures

`func (o *CloudAskResponse) GetFigures() []CloudFigure`

GetFigures returns the Figures field if non-nil, zero value otherwise.

### GetFiguresOk

`func (o *CloudAskResponse) GetFiguresOk() (*[]CloudFigure, bool)`

GetFiguresOk returns a tuple with the Figures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigures

`func (o *CloudAskResponse) SetFigures(v []CloudFigure)`

SetFigures sets Figures field to given value.

### HasFigures

`func (o *CloudAskResponse) HasFigures() bool`

HasFigures returns a boolean if a field has been set.

### GetFollowups

`func (o *CloudAskResponse) GetFollowups() []string`

GetFollowups returns the Followups field if non-nil, zero value otherwise.

### GetFollowupsOk

`func (o *CloudAskResponse) GetFollowupsOk() (*[]string, bool)`

GetFollowupsOk returns a tuple with the Followups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowups

`func (o *CloudAskResponse) SetFollowups(v []string)`

SetFollowups sets Followups field to given value.

### HasFollowups

`func (o *CloudAskResponse) HasFollowups() bool`

HasFollowups returns a boolean if a field has been set.

### GetSources

`func (o *CloudAskResponse) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudAskResponse) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudAskResponse) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudAskResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


