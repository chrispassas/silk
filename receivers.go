package silk

type FlowReceiver interface {
	HandleHeader(h Header)
	HandleFlow(f Flow)
	Close()
}

type SliceFlowReceiver struct {
	File
}

func NewSliceFlowReceiver(initialSize int) *SliceFlowReceiver {
	return &SliceFlowReceiver{
		File{
			Flows: make([]Flow, 0, initialSize),
		},
	}
}

func (a *SliceFlowReceiver) HandleHeader(h Header) {
	a.Header = h
}

func (a *SliceFlowReceiver) HandleFlow(f Flow) {
	a.Flows = append(a.Flows, f)
}

func (a *SliceFlowReceiver) Close() {
	// Nothing to do in the slice case
}

type ChannelFlowReceiver struct {
	Header    Header
	rwChannel chan Flow
}

func NewChannelFlowReceiver(channelBufferSize int) *ChannelFlowReceiver {
	rwChannel := make(chan Flow, channelBufferSize)
	return &ChannelFlowReceiver{
		rwChannel: rwChannel,
	}
}

func (c ChannelFlowReceiver) Read() <-chan Flow {
	return c.rwChannel
}

func (c *ChannelFlowReceiver) HandleHeader(h Header) {
	c.Header = h
}

func (c *ChannelFlowReceiver) HandleFlow(f Flow) {
	c.rwChannel <- f
}

func (c *ChannelFlowReceiver) Close() {
	close(c.rwChannel)
}
