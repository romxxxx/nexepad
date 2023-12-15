package flowcontext

import "github.com/romxxxx/nexepad/infrastructure/config"

// Config returns an instance of *config.Config associated to the flow context.
func (f *FlowContext) Config() *config.Config {
	return f.cfg
}
