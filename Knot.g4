grammar Knot;

// Tokens
THING: 'thing';
SENSOR: 'sensor';
ACTUATOR: 'actuator';
END_CHAR: ';';
WS  :   [ \t\r\n\u000C]+ -> skip;
COMMENT :   '/*' .*? '*/' -> skip;
LINE_COMMENT    :   '//' ~[\r\n]* -> skip;

// TYPES
RELATIVEHUMIDITY: 'relativehumidity';
SWITCH: 'switch';
PRESENCE: 'presence';
COMMAND: 'command';

// Lexical elements
fragment Letter :[A-Za-z_];
fragment DecimalDigit   :[0-9]+;
fragment DecimalLit:   [1-9] DecimalDigit*;
fragment Ident   : Letter (Letter | DecimalDigit)*;
IDENTIFIER     : Ident ;
UNSIGNED_INTEGER: DecimalLit;
NUMBER: DecimalLit ('.' DecimalLit)*;

// Rules
start: things+=definition+ EOF;
definition: THING IDENTIFIER '{' sensors+=thingContent+ '}';
thingContent: op=(SENSOR|ACTUATOR) valueOptions;
config: configChanges? configTime? configUpper? configLower?;
configChanges: 'at changes';
configTime: 'each ' number=UNSIGNED_INTEGER 's';
configUpper: 'greater than ' number=UNSIGNED_INTEGER;
configLower: 'lower than ' number=UNSIGNED_INTEGER;
valueOptions: boolOpt | numberOpt | bytesOpt;
boolOpt: 'bool' IDENTIFIER '(' op=(SWITCH | PRESENCE) ')' 'sends' configChanges? configTime? END_CHAR;
numberOpt: op=('int' | 'float') IDENTIFIER '(' unitTypeOptions ')' 'sends' config END_CHAR;
bytesOpt: 'bytes' IDENTIFIER '(' COMMAND ')' 'sends' configChanges? configTime? END_CHAR;
unitTypeOptions: ( voltage | current | resistance | power | temperature | luminosity | 
time| mass | pressure | distance | angle | volume | area | rain | density | latitude | longitude | 
speed | volumeflow | energy | RELATIVEHUMIDITY);
logicUnits: SWITCH | PRESENCE | COMMAND;
voltage: 'voltage in ' voltagesUnits;
voltagesUnits: op=('V' | 'mV' | 'kV');
current: 'current in ' currentUnits;
currentUnits: op=('A' | 'mA');
resistance: 'resistance in ' resistanceUnits;
resistanceUnits: op='ohm';
power: 'power in ' powerUnits;
powerUnits: op=('W' | 'kW' | 'MW');
temperature: 'temperature in ' temperatureUnits;
temperatureUnits: op=('F' | 'C' | 'K');
luminosity: 'luminosity in ' luminosityUnits;
luminosityUnits: op=('lm' | 'cd' | 'lx');
time: 'time in ' timeUnits;
timeUnits: op=('s' | 'ms' | 'us' | 'min' | 'h');
mass: 'mass in ' massUnits;
massUnits: op=('g' | 'kg' | 'lb' | 'oz');
pressure: 'pressure in ' pressureUnits;
pressureUnits: op=('Pa' | 'psi' | 'bar');
distance: 'distance in ' distanceUnits;
distanceUnits: op=('m' | 'cm' | 'mi' | 'in');
angle: 'angle in ' angleUnits;
angleUnits: op=('degree' | 'rad');
volume: 'volume in ' volumeUnits;
volumeUnits: op=('l' | 'gal' | 'ml' | 'floz');
area: 'area in ' areaUnits;
areaUnits: op=('m2' | 'ha' | 'ac' | 'mm');
rain: 'rain in ' rainUnits;
rainUnits: op='mm';
density: 'density in ' densityUnits;
densityUnits: op='kgm3';
latitude: 'latitude in ' latitudeUnits;
latitudeUnits: op='degree';
longitude: 'longitude in ' longitudeUnits;
longitudeUnits: op='degree';
speed: 'speed in ' speedUnits;
speedUnits: op=('ms' | 'cms' | 'kms' | 'mih');
volumeflow: 'volumeflow in ' volumeflowUnits;
volumeflowUnits: op=('m3s'| 'scmm' | 'ls' | 'lm' | 'ft3s' | 'galm');
energy: 'energy in ' energyUnits;
energyUnits: op=('J' | 'nm' | 'Wh' | 'KWh' | 'cal' | 'Kcal');
