package template

const (
	ZephyrSourceCode = `/* {{ .Name }}.c - KNoT Application Client

This is a generated file

*/


#include <zephyr.h>
#include <net/net_core.h>
#include <logging/log.h>
#include <device.h>
#include <gpio.h>

#include "knot.h"
#include <knot/knot_types.h>
#include <knot/knot_protocol.h>
{{range .Sensors }}
#define {{ .Name | Up }}_PORT		{{ .Name | Up }}_GPIO_CONTROLLER
#define {{ .Name | Up }}_PIN		{{ .Name | Up }}_GPIO_PIN
{{end}}
LOG_MODULE_REGISTER({{ .Name }}, LOG_LEVEL_DBG);

{{range .Sensors }}
{{ .Value }} {{ .Name | Lower }} = {{ .DefaultValue }}; 			/* Tracked value */
struct device *gpio_{{ .Name | Lower }};
{{end}}
{{range .Sensors }}
{{if .IsSensor}}{{else}}int write_{{ .Name | Lower }}(int id)
{
	/* Add here specific code for your actuator */
	return KNOT_CALLBACK_SUCCESS;
}{{end}}
{{end}}
void setup(void)
{
	bool success;
	{{range .Sensors}}
	/* Peripherals control */
	gpio_{{ .Name | Lower }} = device_get_binding({{ .Name | Up }}_PORT);
	/* Uncomment and configure the GPIO {{.Name}} here */
	//gpio_pin_configure(gpio_{{ .Name | Lower }}, {{ .Name | Up }}_PIN, /* GPIO_*/);

	/* KNoT config */
	if (knot_data_register({{ .ID }}, "{{ .Name }}", KNOT_TYPE_ID_{{ .TypeUnit | Up }},
				KNOT_VALUE_TYPE_{{ .Value | Up }}, {{ if .Unit }}KNOT_UNIT_{{ .Unit | Up }},{{ else }}KNOT_UNIT_NOT_APPLICABLE,{{ end }}
				&{{ .Name | Lower }}, sizeof({{ .Name | Lower }}), write_{{ .Name | Lower }}, NULL) < 1)
		LOG_ERR("{{.Name}} failed to register");
	success = knot_data_config({{ .ID }}, {{ range .Configs }}KNOT_EVT_FLAG_{{ .Type }}, {{if .Value}}{{ .Value }}, {{end}}{{end}}NULL);
	if (!success)
		LOG_ERR("{{.Name}} failed to configure");
	{{end}}
}
void loop(void)
{
}
`
	ZephyrPrjConf = `# KNoT
CONFIG_KNOT_NAME="{{ .Name }}"
CONFIG_KNOT_THING_DATA_MAX={{ .Len }}

# Logging
CONFIG_LOG=y
CONFIG_LOG_IMMEDIATE=n
CONFIG_KNOT_LOG=y
CONFIG_KNOT_LOG_LEVEL_DEBUG=y`
	ZephyrSetupConf = `CONFIG_BT_DEVICE_NAME="{{ .Name }}"`
	ZephyrCMakeLists = `cmake_minimum_required(VERSION 3.8.2)

if (NOT DEFINED ENV{KNOT_BASE})
	message(FATAL_ERROR "Source the KNoT shell initialize script!")
endif()

include($ENV{KNOT_BASE}/core/CMakeLists.txt)
project(Blink)

FILE(GLOB app_sources src/*.c)
target_sources(app PRIVATE ${app_sources})

include($ENV{ZEPHYR_BASE}/samples/net/common/common.cmake)`
)