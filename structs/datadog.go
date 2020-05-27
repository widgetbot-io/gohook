package structs

type DataDogMain struct {
	Body        string     `json:"body"`
	LastUpdated string     `json:"last_updated"`
	Transition  string     `json:"transition"`
	EventType   string     `json:"eventType"`
	Metric      string     `json:"metric"`
	Status      string     `json:"status"`
	Title       string     `json:"title"`
	Tags        string     `json:"tags"`
	Msg         string     `json:"msg"`
	Link        string     `json:"link"`
	Date        string     `json:"date"`
	Org         DataDogOrg `json:"org"`
	Id          string     `json:"id"`
}

//{
//"body": "$EVENT_MSG",
//"last_updated": "$LAST_UPDATED",
//"event_type": "$EVENT_TYPE",
//"title": "$ALERT_TITLE",
//"msg": "$EVENT_MSG",
//"date": "$DATE",
//"org": {
//	"id": "$ORG_ID",
//	"name": "$ORG_NAME"
//},
//"id": "$ID"
//}

type DataDogOrg struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
