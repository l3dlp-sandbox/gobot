package beaglebone

import "gobot.io/x/gobot/v2/platforms/adaptors"

var bbbPinMap = map[string]int{
	// P8_01 - P8_2 GND
	// P8_03 - P8_6 EMCC
	"P8_07": 66,
	"P8_08": 67,
	"P8_09": 69,
	"P8_10": 68,
	"P8_11": 45,
	"P8_12": 44,
	"P8_13": 23,
	"P8_14": 26,
	"P8_15": 47,
	"P8_16": 46,
	"P8_17": 27,
	"P8_18": 65,
	"P8_19": 22,
	// P8_20 - P8_25 EMCC
	"P8_26": 61,
	// P8_27 - P8_46 HDMI

	// P9_1 - P9_2 GND
	// P9_3 - P9_4 3V3
	// P9_5 - P9_6 5V
	// P9_7 - P9_8 5V SYS
	// P9_9 PWR_BUT
	// P9_10 SYS_RESET
	"P9_11": 30,
	"P9_12": 60,
	"P9_13": 31,
	"P9_14": 50,
	"P9_15": 48,
	"P9_16": 51,
	"P9_17": 5,
	"P9_18": 4,
	// P9_19 I2C2 SCL
	// P9_20 I2C2 SDA
	"P9_21": 3,
	"P9_22": 2,
	"P9_23": 49,
	"P9_24": 15,
	"P9_25": 117,
	"P9_26": 14,
	"P9_27": 115,
	"P9_28": 113,
	"P9_29": 111,
	"P9_30": 112,
	"P9_31": 110,
}

var bbbPwmPinMap = adaptors.PWMPinDefinitions{
	"P8_13": {
		Dir: "/sys/devices/platform/ocp/48304000.epwmss/48304200.pwm/pwm/", DirRegexp: "pwmchip[0-9]+$", Channel: 1,
	},
	"P8_19": {
		Dir: "/sys/devices/platform/ocp/48304000.epwmss/48304200.pwm/pwm/", DirRegexp: "pwmchip[0-9]+$", Channel: 0,
	},
	"P9_14": {
		Dir: "/sys/devices/platform/ocp/48302000.epwmss/48302200.pwm/pwm/", DirRegexp: "pwmchip[0-9]+$", Channel: 0,
	},
	"P9_16": {
		Dir: "/sys/devices/platform/ocp/48302000.epwmss/48302200.pwm/pwm/", DirRegexp: "pwmchip[0-9]+$", Channel: 1,
	},
	"P9_21": {
		Dir: "/sys/devices/platform/ocp/48300000.epwmss/48300200.pwm/pwm/", DirRegexp: "pwmchip[0-9]+$", Channel: 1,
	},
	"P9_22": {
		Dir: "/sys/devices/platform/ocp/48300000.epwmss/48300200.pwm/pwm/", DirRegexp: "pwmchip[0-9]+$", Channel: 0,
	},
	"P9_42": {
		Dir: "/sys/devices/platform/ocp/48300000.epwmss/48300100.ecap/pwm/", DirRegexp: "pwmchip[0-9]+$", Channel: 0,
	},
}

var bbbAnalogPinMap = adaptors.AnalogPinDefinitions{
	"P9_39": {Path: "/sys/bus/iio/devices/iio:device0/in_voltage0_raw", W: false, ReadBufLen: 1024},
	"P9_40": {Path: "/sys/bus/iio/devices/iio:device0/in_voltage1_raw", W: false, ReadBufLen: 1024},
	"P9_37": {Path: "/sys/bus/iio/devices/iio:device0/in_voltage2_raw", W: false, ReadBufLen: 1024},
	"P9_38": {Path: "/sys/bus/iio/devices/iio:device0/in_voltage3_raw", W: false, ReadBufLen: 1024},
	"P9_33": {Path: "/sys/bus/iio/devices/iio:device0/in_voltage4_raw", W: false, ReadBufLen: 1024},
	"P9_36": {Path: "/sys/bus/iio/devices/iio:device0/in_voltage5_raw", W: false, ReadBufLen: 1024},
	"P9_35": {Path: "/sys/bus/iio/devices/iio:device0/in_voltage6_raw", W: false, ReadBufLen: 1024},
}
