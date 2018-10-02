package tipboardsubscriber

//TipboardDash - A struct to represent a TipBoard dashboard host.
type TipboardDash struct {
	dashHost   string //host running the dashboard
	dashPort   int    //port the dashboard is listening on
	dashAPIKey string //API key for the dashboard
}
