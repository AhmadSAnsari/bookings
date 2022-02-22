package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float64
	DataMap map[string]interface{}
	CSRFToken string
	Flash string   // messages
	Warning string
	Error string
}
