# AutoReviewIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the automation, and Run its run in review, from the path. | [optional] 
**Post** | Pointer to **bool** | Post posts it when true and discards it when false. | [optional] 
**Run** | Pointer to **string** |  | [optional] 

## Methods

### NewAutoReviewIn

`func NewAutoReviewIn() *AutoReviewIn`

NewAutoReviewIn instantiates a new AutoReviewIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoReviewInWithDefaults

`func NewAutoReviewInWithDefaults() *AutoReviewIn`

NewAutoReviewInWithDefaults instantiates a new AutoReviewIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoReviewIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoReviewIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoReviewIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoReviewIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPost

`func (o *AutoReviewIn) GetPost() bool`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *AutoReviewIn) GetPostOk() (*bool, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *AutoReviewIn) SetPost(v bool)`

SetPost sets Post field to given value.

### HasPost

`func (o *AutoReviewIn) HasPost() bool`

HasPost returns a boolean if a field has been set.

### GetRun

`func (o *AutoReviewIn) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *AutoReviewIn) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *AutoReviewIn) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *AutoReviewIn) HasRun() bool`

HasRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


