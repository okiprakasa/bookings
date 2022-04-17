package models

// TemplateData holds data set from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	DataMap   map[string]interface{}
	CSFRToken string //CSFRToken Cross Site Request Forgery Token (Security token for forms)
	Flash     string //Flash Message data holder
	Warning   string
	Error     string
}
