package jsontypes

// func (r *CurrentData) GetInfluxFields() map[string]interface{} {
// 	return map[string]interface{}{
// 		"DeviceStatus":    r.DeviceStatus,
// 		"DPowerFlow":      r.PowerFlow,
// 		"Export":          r.Export * 1000,
// 		"Import":          r.Import * 1000,
// 		"Load":            r.Load * 1000,
// 		"Solar":           (r.Ppv1 + r.Ppv2) * 1000,
// 		"Soc":             r.Soc,
// 		"Charge":          r.Charge * 1000,
// 		"Discharge":       r.Discharge * 1000,
// 		"GridPowerOutage": r.GridPowerOutage,
// 		"Temperature":     r.Temperature,
// 		"BatteryVoltage":  r.BatteryVoltage,
// 		"BatteryCurrent":  r.BatteryCurrent,
// 	}
// }

// func (r *CurrentData) GetMQTTFields() map[string]string {
// 	return map[string]string{
// 		"devicestatus":    strconv.FormatInt(r.DeviceStatus, 10),
// 		"dpowerflow":      strconv.FormatInt(r.PowerFlow, 10),
// 		"import":          strconv.FormatInt(int64(math.Abs(r.Import*1000)), 10),
// 		"export":          strconv.FormatInt(int64(math.Abs(r.Export*1000)), 10),
// 		"load":            strconv.FormatInt(int64(r.Load*1000), 10),
// 		"solar":           strconv.FormatInt(int64((r.Ppv1+r.Ppv2)*1000), 10),
// 		"soc":             strconv.FormatInt(r.Soc, 10),
// 		"charge":          strconv.FormatInt(int64(math.Abs(r.Charge*1000)), 10),
// 		"discharge":       strconv.FormatInt(int64(math.Abs(r.Discharge*1000)), 10),
// 		"gridpoweroutage": strconv.FormatInt(r.GridPowerOutage, 10),
// 		"temperature":     strconv.FormatFloat(r.Temperature, 'f', 2, 32),
// 		"batteryvoltage":  strconv.FormatFloat(r.BatteryVoltage, 'f', 2, 32),
// 		"batterycurrent":  strconv.FormatFloat(r.BatteryCurrent, 'f', 2, 32),
// 	}

// }

// func (r *DayDatum) GetInfluxFields() *map[string]interface{} {
// 	if r.ReceivedDate != "" {
// 		return &map[string]interface{}{
// 			"Hour":     r.ReceivedDate,
// 			"Export":   r.Export,
// 			"Grid":     r.Grid,
// 			"SolarUse": r.SolarUse,
// 			"SoC":      r.StateOfCharge,
// 		}
// 	} else {
// 		return nil
// 	}
// }

// func (r *DayDatum) GetMQTTFields() *map[string]string {
// 	if r.ReceivedDate != "" {
// 		return &map[string]string{
// 			"Hour":     r.ReceivedDate,
// 			"Export":   strconv.FormatInt(int64(r.Export*1000.0), 10),
// 			"Grid":     strconv.FormatInt(int64(r.Grid*1000.0), 10),
// 			"SolarUse": strconv.FormatInt(int64(r.SolarUse*1000.0), 10),
// 			"SoC":      strconv.FormatInt(r.StateOfCharge, 10),
// 		}
// 	} else {
// 		return nil
// 	}
// }
