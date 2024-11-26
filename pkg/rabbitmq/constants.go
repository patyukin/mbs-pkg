package rabbitmq

const Exchange = "banking_system_exchange"

// Auth Service
const (
	AuthNotifyQueue                 = "auth_notify_queue"
	NotifyAuthQueue                 = "notify_auth_queue"
	NotifySignUpConfirmCodeRouteKey = "notify_sign_up_confirm_code_route_key"
	AuthSignUpResultMessageRouteKey = "auth_sign_up_result_message_route_key"
	AuthSignInConfirmCodeRouteKey   = "auth_sign_in_confirm_code_route_key"
)

// Logger Service
const (
	LoggerNotifyQueue    = "logger_notify_queue"
	LoggerReportRouteKey = "logger_report_route_key"
	NotifyReportRouteKey = "notify_report_route_key"
)

// Dead Letter Queue
const (
	DeadLetterExchange = "dead_letter_exchange"
	DeadLetterQueue    = "dead_letter_queue"
)

// Payment Service
const (
	PaymentNotifyQueue               = "payment_notify_queue"
	PaymentExecutionInitiateRouteKey = "payment_execution_initiate_route_key"
	AccountCreationRouteKey          = "account_creation_route_key"
	PaymentStatusChangedRouteKey     = "payment_status_change_route_key"
)

// Credit Service
const (
	CreditNotifyQueue                = "credit_notify_queue"
	CreditApplicationCreatedRouteKey = "credit_application_created_route_key"
	CreditCreatedRouteKey            = "credit_created_route_key"
)
