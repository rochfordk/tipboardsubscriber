package tipboardsubscriber

//TipboardDash - A struct to represent a TipBoard dashboard host.
type TipboardDash struct {
	DashHost   string //host running the dashboard
	DashPort   int    //port the dashboard is listening on
	DashAPIKey string //API key for the dashboard
}
