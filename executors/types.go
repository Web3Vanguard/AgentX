package executors



type FlowContext struct {
    Text      string                 // Main text content
    Memory    interface{}           // Structured data storage
    Variables map[string]any
    Flow      *interface{}
    ImageURLs []string             // Image URLs for multimodal content
    Think     string               // Captured thinking process               // Reference to current flow
}


type Executor interface {
    Execute(context *FlowContext) (*FlowContext, error)
}