package rabbitmq

const Exchange = "mbs_exchange"

// Telegram
const (
	TelegramMessageQueue    = "telegram_message_queue"
	TelegramMessageRouteKey = "telegram_message_route_key"
)

// Dead Letter Queue
const (
	DeadLetterExchange = "dead_letter_exchange"
	DeadLetterQueue    = "dead_letter_queue"
)
