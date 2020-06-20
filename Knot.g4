grammar Knot;

// Tokens
THING: 'thing';
SENSOR: 'sensor';
ACTUATOR: 'actuator';
VALUE: 'value';
UNIT: 'unit';
TYPE: 'type';
END_CHAR: ';';
SEPARATOR: ':';
WS  :   [ \t\r\n\u000C]+ -> skip;
COMMENT :   '/*' .*? '*/' -> skip;
LINE_COMMENT    :   '//' ~[\r\n]* -> skip;
// Values
BOOL: 'bool';
INT: 'int';
FLOAT: 'float';
BYTES: 'bytes';

// TYPES
CURRENT: 'current';
RESISTANCE: 'resistance';
POWER: 'power';
TEMPERATURE: 'temperature';
RELATIVEHUMIDITY: 'relativehumidity';
LUMINOSITY: 'luminosity';
TIME: 'time';
MASS: 'mass';
PRESSURE: 'pressure';
DISTANCE: 'distance';
ANGLE: 'angle';
VOLUME: 'volume';
AREA: 'area';
RAIN: 'rain';
DENSITY: 'density';
LATITUDE: 'latitude';
LONGITUDE: 'longitude';
SPEED: 'speed';
VOLUMEFLOW: 'volumeflow';
ENERGY: 'energy';
SWITCH: 'switch';
PRESENCE: 'presence';
COMMAND: 'command';

// Lexical elements
fragment Letter :[A-Za-z_];
fragment DecimalDigit   :[0-9]+;
fragment DecimalLit:   [1-9] DecimalDigit*;
fragment NUMBER: ('0' .. '9') + ('.' ('0' .. '9') +)?;
fragment UNSIGNED_INTEGER: ('0' .. '9')+;
fragment Ident   : Letter (Letter | DecimalDigit)*;
IDENTIFIER     : Ident ;
ID: DecimalLit;

// TODO: Add rule to data config event flag

// Rules
start: definition+ EOF;
definition: THING IDENTIFIER '{' thingContent+ '}';
thingContent: op=(SENSOR|ACTUATOR) valueOptions IDENTIFIER '(' unitTypeOptions ')' END_CHAR;
valueOptions: op=(BOOL | INT | FLOAT | BYTES);
unitTypeOptions: ( voltage | CURRENT | RESISTANCE | POWER | temperature | RELATIVEHUMIDITY | LUMINOSITY | 
TIME| MASS | PRESSURE | DISTANCE | ANGLE | VOLUME | AREA | RAIN | DENSITY | LATITUDE | LONGITUDE | 
SPEED | VOLUMEFLOW | ENERGY | SWITCH | PRESENCE | COMMAND);
voltage: 'voltage in ' voltagesUnits;
voltagesUnits: op=('V' | 'mV' | 'kV');
temperature: 'temperature in ' temperatureUnits;
temperatureUnits: op=('F' | 'C' | 'K');
