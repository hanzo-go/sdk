# EnrollmentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Enrollment**](Enrollment.md) | Data is the page: every contact walking this ONE sequence, in any state — active, completed and canceled walks all appear, since the history of who was reached is the point. An empty array when nobody has been enrolled. | [optional] 

## Methods

### NewEnrollmentList

`func NewEnrollmentList() *EnrollmentList`

NewEnrollmentList instantiates a new EnrollmentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentListWithDefaults

`func NewEnrollmentListWithDefaults() *EnrollmentList`

NewEnrollmentListWithDefaults instantiates a new EnrollmentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnrollmentList) GetData() []Enrollment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnrollmentList) GetDataOk() (*[]Enrollment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnrollmentList) SetData(v []Enrollment)`

SetData sets Data field to given value.

### HasData

`func (o *EnrollmentList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


