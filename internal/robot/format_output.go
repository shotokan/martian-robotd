package robot

func formatRadiansOrientation(orientation int32) string {
	switch orientation % 360 {
	case 90, -270:
		return "N"
	case 0, -360:
		return "E"
	case 270, -90:
		return "S"
	case 180, -180:
		return "W"
	default:
		panic("Could not convert radians to a specific orientation")
	}
}

func formatLost(isLost bool) string {
	if isLost {
		return "LOST"
	}
	return ""
}
