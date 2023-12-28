package util

type DeviceCategory = string

const (
	Device_Category_Android_Phone  = "Android Phone"
	Device_Category_Android_Pad    = "Android Pad"
	Device_Category_Android_Chrome = "Android Chrome"
	Device_Category_Apple_Phone    = "Apple Phone"
	Device_Category_Apple_Pad      = "Apple Pad"
	Device_Category_Web            = "Web"
	Device_Category_PC             = "PC"
	Device_Category_Mobile         = "Mobile"
)

func GetDeviceCategory(deviceOS, deviceType string) DeviceCategory {
	return Device_Category_Mobile
}
