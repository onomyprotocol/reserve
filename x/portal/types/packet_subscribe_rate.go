package types

// ValidateBasic is used for validating the packet
func (p SubscribeRatePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p SubscribeRatePacketData) GetBytes() ([]byte, error) {
	var modulePacket PortalPacketData

	modulePacket.Packet = &PortalPacketData_SubscribeRatePacket{&p}

	return modulePacket.Marshal()
}
