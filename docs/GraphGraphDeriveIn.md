# GraphGraphDeriveIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsKnown** | Pointer to **string** | AsKnown is how much this plane had heard, RFC 3339 — of the statements and of the rules — so a past answer is reproducible. Absent is now. | [optional] 
**AsOf** | Pointer to **string** | AsOf derives from the graph as it stood at an instant of the world, RFC 3339: the statements that held then. Absent is now. | [optional] 
**File** | Pointer to **bool** | File files each tuple Derived lists as an assertion: (subject, predicate, object) from the instant it was derived at, open, its source the rule, its evidence the rule&#39;s ID, its confidence the decay, and an edge when the object is an entity by its proof. It requires an admin of the organization, the graph as it stands now — no as_of, as_known or without — and predicates with two arguments. It is refused whole when a relation that holds one value at a time would keep only one of the values a subject is concluded or asserted to have, and when a conclusion is about a &#x60;relation:&#x60; or &#x60;rule:&#x60; key. Each is judged as any assertion is, and the write is on the audit trail before it commits. A filed conclusion is an assertion like any other: it does not fall when a support later does. | [optional] 
**Predicates** | Pointer to **[]string** | Predicates are what to derive: every tuple the rules in force conclude of each, with its proof. Each must be the head of a rule in force, asserted as (&#x60;rule:&lt;name&gt;&#x60;, &#x60;rule&#x60;, &#x60;&lt;text&gt;&#x60;) — for example &#x60;ancestor(X, Z) :- parent(X, Y), ancestor(Y, Z)&#x60;. At least one, at most 256. | [optional] 
**Without** | Pointer to **[]string** | Without is assertion IDs, of statements or of rules, to derive without as though they had never been filed. The derivation then runs twice, and Lost and Gained say what that changes. Each must be a row this derivation reads. At most 256. | [optional] 

## Methods

### NewGraphGraphDeriveIn

`func NewGraphGraphDeriveIn() *GraphGraphDeriveIn`

NewGraphGraphDeriveIn instantiates a new GraphGraphDeriveIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphDeriveInWithDefaults

`func NewGraphGraphDeriveInWithDefaults() *GraphGraphDeriveIn`

NewGraphGraphDeriveInWithDefaults instantiates a new GraphGraphDeriveIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsKnown

`func (o *GraphGraphDeriveIn) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphDeriveIn) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphDeriveIn) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphDeriveIn) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphDeriveIn) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphDeriveIn) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphDeriveIn) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphDeriveIn) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetFile

`func (o *GraphGraphDeriveIn) GetFile() bool`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *GraphGraphDeriveIn) GetFileOk() (*bool, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *GraphGraphDeriveIn) SetFile(v bool)`

SetFile sets File field to given value.

### HasFile

`func (o *GraphGraphDeriveIn) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetPredicates

`func (o *GraphGraphDeriveIn) GetPredicates() []string`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *GraphGraphDeriveIn) GetPredicatesOk() (*[]string, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *GraphGraphDeriveIn) SetPredicates(v []string)`

SetPredicates sets Predicates field to given value.

### HasPredicates

`func (o *GraphGraphDeriveIn) HasPredicates() bool`

HasPredicates returns a boolean if a field has been set.

### GetWithout

`func (o *GraphGraphDeriveIn) GetWithout() []string`

GetWithout returns the Without field if non-nil, zero value otherwise.

### GetWithoutOk

`func (o *GraphGraphDeriveIn) GetWithoutOk() (*[]string, bool)`

GetWithoutOk returns a tuple with the Without field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithout

`func (o *GraphGraphDeriveIn) SetWithout(v []string)`

SetWithout sets Without field to given value.

### HasWithout

`func (o *GraphGraphDeriveIn) HasWithout() bool`

HasWithout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


