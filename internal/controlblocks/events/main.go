package events

type MessageEvent struct {
	Text string
	Next []interface{}
}

type CronEvent struct {
	Cron string
	Next []interface{}
}

type ButtonEvent struct {
}
