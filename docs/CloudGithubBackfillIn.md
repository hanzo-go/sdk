# CloudGithubBackfillIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | State is the GitHub issue state to walk: \&quot;open\&quot; (the default), \&quot;closed\&quot; or \&quot;all\&quot;. Anything else is a 400. | [optional] 

## Methods

### NewCloudGithubBackfillIn

`func NewCloudGithubBackfillIn() *CloudGithubBackfillIn`

NewCloudGithubBackfillIn instantiates a new CloudGithubBackfillIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGithubBackfillInWithDefaults

`func NewCloudGithubBackfillInWithDefaults() *CloudGithubBackfillIn`

NewCloudGithubBackfillInWithDefaults instantiates a new CloudGithubBackfillIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CloudGithubBackfillIn) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudGithubBackfillIn) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudGithubBackfillIn) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudGithubBackfillIn) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


