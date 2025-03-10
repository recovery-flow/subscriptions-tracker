# SubscriptionDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | [**Relationships**](Relationships.md) |  | 
**Type** | [**Relationships**](Relationships.md) |  | 
**PaymentMethod** | [**Relationships**](Relationships.md) |  | 

## Methods

### NewSubscriptionDataRelationships

`func NewSubscriptionDataRelationships(plan Relationships, type_ Relationships, paymentMethod Relationships, ) *SubscriptionDataRelationships`

NewSubscriptionDataRelationships instantiates a new SubscriptionDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDataRelationshipsWithDefaults

`func NewSubscriptionDataRelationshipsWithDefaults() *SubscriptionDataRelationships`

NewSubscriptionDataRelationshipsWithDefaults instantiates a new SubscriptionDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *SubscriptionDataRelationships) GetPlan() Relationships`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SubscriptionDataRelationships) GetPlanOk() (*Relationships, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SubscriptionDataRelationships) SetPlan(v Relationships)`

SetPlan sets Plan field to given value.


### GetType

`func (o *SubscriptionDataRelationships) GetType() Relationships`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionDataRelationships) GetTypeOk() (*Relationships, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionDataRelationships) SetType(v Relationships)`

SetType sets Type field to given value.


### GetPaymentMethod

`func (o *SubscriptionDataRelationships) GetPaymentMethod() Relationships`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *SubscriptionDataRelationships) GetPaymentMethodOk() (*Relationships, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *SubscriptionDataRelationships) SetPaymentMethod(v Relationships)`

SetPaymentMethod sets PaymentMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


