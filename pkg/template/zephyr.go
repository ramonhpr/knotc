package template

const (
	ZephyrSourceCode = `/* {{ .Name }}.c - KNoT Application Client

This is a generated file

*/

#include <zephyr.h>
#include <logging/log.h>
#include <device.h>

#include "knot.h"
#include "knot_generated.h"

LOG_MODULE_REGISTER({{ .Name }}, LOG_LEVEL_DBG);

{{range .Items }}
{{ .Value }} {{ .Name | Lower }} = {{ .DefaultValue }}; 			/* Tracked value */
struct device *gpio_{{ .Name | Lower }};
{{end}}
{{range .Items }}
{{if .IsSensor}}{{else}}int write_{{ .Name | Lower }}(int id)
{
	/* Add here specific code for your actuator */
	return KNOT_CALLBACK_SUCCESS;
}{{end}}
{{end}}
void setup(void)
{
	bool success = setup_knot();

	if (!success)
		LOG_ERR("KNoT failed to configure");

	/* Configure the GPIO/devices bellow */
}

void loop(void)
{
}
`
	ZephyrGeneratedHeader=`/* knot_generated.h - KNoT Application Client
This is a generated file. Don't Update it!!

*/
{{range .Items }}
{{if .IsSensor}}{{else}}int write_{{ .Name | Lower }}(int id);{{end}}
{{end}}

bool setup_knot();

/* Tracked values */
{{range .Items }}
extern {{ .Value }} {{ .Name | Lower }};
{{end}}
`
	ZephyrGeneratedSource=`/* knot_generated.c - KNoT Application Client
This is a generated file. Don't Update it!!

*/
#include <zephyr.h>
#include <logging/log.h>
#include <knot/knot_types.h>
#include <knot/knot_protocol.h>

#include "knot.h"
#include "knot_generated.h"

LOG_MODULE_REGISTER(knot_generated, LOG_LEVEL_ERR);

bool setup_knot()
{
	bool success;
	{{range .Items}}
	if (knot_data_register({{ .ID }}, "{{ .Name }}", KNOT_TYPE_ID_{{ .TypeUnit | Up }},
				KNOT_VALUE_TYPE_{{ .Value | UpRaw }}, {{ if .Unit }}KNOT_UNIT_{{ .TypeUnit | Up }}_{{ .Unit | Up }},{{ else }}KNOT_UNIT_NOT_APPLICABLE,{{ end }}
				&{{ .Name | Lower }}, sizeof({{ .Name | Lower }}), {{if .IsSensor}}NULL{{else}}write_{{ .Name | Lower }}{{end}}, NULL) < 1)
		LOG_ERR("{{.Name}} failed to register");

	success = knot_data_config({{ .ID }}, {{ range .Configs }}KNOT_EVT_FLAG_{{ .Type }}, {{if .Value}}{{ .Value }}, {{end}}{{end}}NULL);
	if (!success) {
		LOG_ERR("{{.Name}} failed to configure");
		return false;
	}
	{{end}}
	return true;
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
project({{ .Name }})

FILE(GLOB app_sources src/*.c)
target_sources(app PRIVATE ${app_sources})

include($ENV{ZEPHYR_BASE}/samples/net/common/common.cmake)`
)