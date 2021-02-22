# InvoicePaymentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | The Square-generated ID of the payment request in an &#x60;invoice&#x60;. | [optional] [default to null]
**RequestMethod** | **string** | Indicates how Square processes the payment request. DEPRECATED at version 2021-01-21. Replaced by the &#x60;Invoice.delivery_method&#x60; and &#x60;InvoicePaymentRequest.automatic_payment_source&#x60; fields.  One of the following is required when creating an invoice: - (Recommended) The &#x60;delivery_method&#x60; field of the invoice. To configure an automatic payment, the &#x60;automatic_payment_source&#x60; field of the payment request is also required. - This &#x60;request_method&#x60; field. Note that &#x60;invoice&#x60; objects returned in responses do not include &#x60;request_method&#x60;. See [InvoiceRequestMethod](#type-invoicerequestmethod) for possible values | [optional] [default to null]
**RequestType** | **string** | Identifies the payment request type. This type defines how the payment request amount is determined. This field is required to create a payment request. See [InvoiceRequestType](#type-invoicerequesttype) for possible values | [optional] [default to null]
**DueDate** | **string** | The due date (in the invoice location&#x27;s time zone) for the payment request, in &#x60;YYYY-MM-DD&#x60; format.  After this date, the invoice becomes overdue. This field is required to create a payment request. | [optional] [default to null]
**FixedAmountRequestedMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**PercentageRequested** | **string** | Specifies the amount for the payment request in percentage:  - When the payment &#x60;request_type&#x60; is &#x60;DEPOSIT&#x60;, it is the percentage of the order total amount. - When the payment &#x60;request_type&#x60; is &#x60;INSTALLMENT&#x60;, it is the percentage of the order total less  the deposit, if requested. The sum of the &#x60;percentage_requested&#x60; in all installment  payment requests must be equal to 100.  You cannot specify this when the payment &#x60;request_type&#x60; is &#x60;BALANCE&#x60; or when the  payment request specifies the &#x60;fixed_amount_requested_money&#x60; field. | [optional] [default to null]
**TippingEnabled** | **bool** | If set to true, the Square-hosted invoice page (the &#x60;public_url&#x60; field of the invoice)  provides a place for the customer to pay a tip.   This field is allowed only on the final payment request   and the payment &#x60;request_type&#x60; must be &#x60;BALANCE&#x60; or &#x60;INSTALLMENT&#x60;. | [optional] [default to null]
**AutomaticPaymentSource** | **string** | The payment method for an automatic payment.  The default value is &#x60;NONE&#x60;. See [InvoiceAutomaticPaymentSource](#type-invoiceautomaticpaymentsource) for possible values | [optional] [default to null]
**CardId** | **string** | The ID of the card on file to charge for the payment request. To get the customer’s card on file, use the &#x60;customer_id&#x60; of the invoice recipient to call &#x60;RetrieveCustomer&#x60; in the Customers API. Then, get the ID of the target card from the &#x60;cards&#x60; field in the response. | [optional] [default to null]
**Reminders** | [**[]InvoicePaymentReminder**](InvoicePaymentReminder.md) | A list of one or more reminders to send for the payment request. | [optional] [default to null]
**ComputedAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalCompletedAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**RoundingAdjustmentIncludedMoney** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

