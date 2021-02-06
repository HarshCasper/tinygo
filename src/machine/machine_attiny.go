// +build avr,attiny

package machine

// UART on the AVR is a dummy implementation. UART has not been implemented for ATtiny
// devices.
type UART struct {
	Buffer *RingBuffer
}

// Configure is a dummy implementation. UART has not been implemented for ATtiny
// devices.
func (uart UART) Configure(config UARTConfig) {
}

// WriteByte is a dummy implementation. UART has not been implemented for ATtiny
// devices.
func (uart UART) WriteByte(c byte) error {
	return nil
}
