package easee

// ChargeStateHasChanged check if the charge state has changed.
func (p *Product) ChargeStateHasChanged() bool {
	if p.LastState.ChargerOpMode == 0 {
		return false
	}
	if p.ChargerState.ChargerOpMode != p.LastState.ChargerOpMode {
		return true
	}
	return false
}

// WattHasChanged check if the charger totalt power is changed.
func (p *Product) WattHasChanged() bool {
	if p.ChargerState.TotalPower != p.LastState.TotalPower {
		return true
	}
	return false
}
