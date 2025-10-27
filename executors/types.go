package executors

type FlowContext struct {
    Text   string      // Current text content
    Memory interface{} // Structured data
    Think  string      // Thinking process
    Images []string    // Image URLs
}


type Executor interface {
    Execute(context *FlowContext) (*FlowContext, error)
}