package jsontypes

func (r *CurrentData) GetInfluxFields() map[string]interface{} {
	return map[string]interface{}{
		"DeviceStatus":    r.DeviceStatus,
		"DPowerFlow":      r.PowerFlow,
		"Export":          r.Export * 1000,
		"Import":          r.Import * 1000,
		"Load":            r.Load * 1000,
		"Solar":           (r.Ppv1 + r.Ppv2) * 1000,
		"Soc":             r.Soc,
		"Charge":          r.Charge * 1000,
		"Discharge":       r.Discharge * 1000,
		"GridPowerOutage": r.GridPowerOutage,
		"Temperature":     r.Temperature,
		"BatteryVoltage":  r.BatteryVoltage,
		"BatteryCurrent":  r.BatteryCurrent,
	}
}

func (r *DayDatum) GetInfluxFields() *map[string]interface{} {
	if r.ReceivedDate != "" {
		return &map[string]interface{}{
			"Hour":     r.ReceivedDate,
			"Export":   r.Export,
			"Grid":     r.Grid,
			"SolarUse": r.SolarUse,
			"SoC":      r.StateOfCharge,
		}
	} else {
		return nil
	}
}
