package tmz

// CountryZonesMap is a map of countries to the timezones they have.
// generated using: https://gist.github.com/prnvbn/487507ce1caa2f160e54e0d6e0eab717
//
// Additional Timezones added manually
var CountryZonesMap = map[string][]Zone{
	"GMT":                                    {"Europe/Warsaw"},
	"UTC":                                    {"Europe/Warsaw"},
	"BST":                                    {"Europe/London"},
	"IST":                                    {"Asia/Kolkata"},
	"CET":                                    {"Europe/Paris", "Europe/Berlin", "Europe/Rome"},
	"CDT":                                    {"America/Chicago", "America/Mexico_City"},
	"PT":                                     {"America/Los_Angeles", "America/Vancouver"},
	"Aruba":                                  {"America/Aruba"},
	"Afghanistan":                            {"Asia/Kabul"},
	"Angola":                                 {"Africa/Luanda"},
	"Anguilla":                               {"America/Anguilla"},
	"Åland Islands":                          {"Europe/Mariehamn"},
	"Albania":                                {"Europe/Tirane"},
	"Andorra":                                {"Europe/Andorra"},
	"United Arab Emirates":                   {"Asia/Dubai"},
	"Argentina":                              {"America/Argentina/Buenos_Aires", "America/Argentina/Cordoba", "America/Argentina/Salta", "America/Argentina/Jujuy", "America/Argentina/Tucuman", "America/Argentina/Catamarca", "America/Argentina/La_Rioja", "America/Argentina/San_Juan", "America/Argentina/Mendoza", "America/Argentina/San_Luis", "America/Argentina/Rio_Gallegos", "America/Argentina/Ushuaia"},
	"Armenia":                                {"Asia/Yerevan"},
	"American Samoa":                         {"Pacific/Pago_Pago"},
	"Antarctica":                             {"Antarctica/McMurdo", "Antarctica/Casey", "Antarctica/Davis", "Antarctica/DumontDUrville", "Antarctica/Mawson", "Antarctica/Palmer", "Antarctica/Rothera", "Antarctica/Syowa", "Antarctica/Troll", "Antarctica/Vostok"},
	"French Southern Territories":            {"Indian/Kerguelen"},
	"Antigua and Barbuda":                    {"America/Antigua"},
	"Australia":                              {"Australia/Lord_Howe", "Antarctica/Macquarie", "Australia/Hobart", "Australia/Melbourne", "Australia/Sydney", "Australia/Broken_Hill", "Australia/Brisbane", "Australia/Lindeman", "Australia/Adelaide", "Australia/Darwin", "Australia/Perth", "Australia/Eucla"},
	"Austria":                                {"Europe/Vienna"},
	"Azerbaijan":                             {"Asia/Baku"},
	"Burundi":                                {"Africa/Bujumbura"},
	"Belgium":                                {"Europe/Brussels"},
	"Benin":                                  {"Africa/Porto-Novo"},
	"Bonaire, Sint Eustatius and Saba":       {"America/Kralendijk"},
	"Burkina Faso":                           {"Africa/Ouagadougou"},
	"Bangladesh":                             {"Asia/Dhaka"},
	"Bulgaria":                               {"Europe/Sofia"},
	"Bahrain":                                {"Asia/Bahrain"},
	"Bahamas":                                {"America/Nassau"},
	"Bosnia and Herzegovina":                 {"Europe/Sarajevo"},
	"Saint Barthélemy":                       {"America/St_Barthelemy"},
	"Belarus":                                {"Europe/Minsk"},
	"Belize":                                 {"America/Belize"},
	"Bermuda":                                {"Atlantic/Bermuda"},
	"Bolivia, Plurinational State of":        {"America/La_Paz"},
	"Brazil":                                 {"America/Noronha", "America/Belem", "America/Fortaleza", "America/Recife", "America/Araguaina", "America/Maceio", "America/Bahia", "America/Sao_Paulo", "America/Campo_Grande", "America/Cuiaba", "America/Santarem", "America/Porto_Velho", "America/Boa_Vista", "America/Manaus", "America/Eirunepe", "America/Rio_Branco"},
	"Barbados":                               {"America/Barbados"},
	"Brunei Darussalam":                      {"Asia/Brunei"},
	"Bhutan":                                 {"Asia/Thimphu"},
	"Botswana":                               {"Africa/Gaborone"},
	"Central African Republic":               {"Africa/Bangui"},
	"Canada":                                 {"America/St_Johns", "America/Halifax", "America/Glace_Bay", "America/Moncton", "America/Goose_Bay", "America/Blanc-Sablon", "America/Toronto", "America/Nipigon", "America/Thunder_Bay", "America/Iqaluit", "America/Pangnirtung", "America/Atikokan", "America/Winnipeg", "America/Rainy_River", "America/Resolute", "America/Rankin_Inlet", "America/Regina", "America/Swift_Current", "America/Edmonton", "America/Cambridge_Bay", "America/Yellowknife", "America/Inuvik", "America/Creston", "America/Dawson_Creek", "America/Fort_Nelson", "America/Whitehorse", "America/Dawson", "America/Vancouver"},
	"Cocos (Keeling) Islands":                {"Indian/Cocos"},
	"Switzerland":                            {"Europe/Zurich"},
	"Chile":                                  {"America/Santiago", "America/Punta_Arenas", "Pacific/Easter"},
	"China":                                  {"Asia/Shanghai", "Asia/Urumqi"},
	"Côte d'Ivoire":                          {"Africa/Abidjan"},
	"Cameroon":                               {"Africa/Douala"},
	"Congo, The Democratic Republic of the":  {"Africa/Kinshasa", "Africa/Lubumbashi"},
	"Congo":                                  {"Africa/Brazzaville"},
	"Cook Islands":                           {"Pacific/Rarotonga"},
	"Colombia":                               {"America/Bogota"},
	"Comoros":                                {"Indian/Comoro"},
	"Cabo Verde":                             {"Atlantic/Cape_Verde"},
	"Costa Rica":                             {"America/Costa_Rica"},
	"Cuba":                                   {"America/Havana"},
	"Curaçao":                                {"America/Curacao"},
	"Christmas Island":                       {"Indian/Christmas"},
	"Cayman Islands":                         {"America/Cayman"},
	"Cyprus":                                 {"Asia/Nicosia", "Asia/Famagusta"},
	"Czechia":                                {"Europe/Prague"},
	"Germany":                                {"Europe/Berlin", "Europe/Busingen"},
	"Djibouti":                               {"Africa/Djibouti"},
	"Dominica":                               {"America/Dominica"},
	"Denmark":                                {"Europe/Copenhagen"},
	"Dominican Republic":                     {"America/Santo_Domingo"},
	"Algeria":                                {"Africa/Algiers"},
	"Ecuador":                                {"America/Guayaquil", "Pacific/Galapagos"},
	"Egypt":                                  {"Africa/Cairo"},
	"Eritrea":                                {"Africa/Asmara"},
	"Western Sahara":                         {"Africa/El_Aaiun"},
	"Spain":                                  {"Europe/Madrid", "Africa/Ceuta", "Atlantic/Canary"},
	"Estonia":                                {"Europe/Tallinn"},
	"Ethiopia":                               {"Africa/Addis_Ababa"},
	"Finland":                                {"Europe/Helsinki"},
	"Fiji":                                   {"Pacific/Fiji"},
	"Falkland Islands (Malvinas)":            {"Atlantic/Stanley"},
	"France":                                 {"Europe/Paris"},
	"Faroe Islands":                          {"Atlantic/Faroe"},
	"Micronesia, Federated States of":        {"Pacific/Chuuk", "Pacific/Pohnpei", "Pacific/Kosrae"},
	"Gabon":                                  {"Africa/Libreville"},
	"United Kingdom":                         {"Europe/London"},
	"Georgia":                                {"Asia/Tbilisi"},
	"Guernsey":                               {"Europe/Guernsey"},
	"Ghana":                                  {"Africa/Accra"},
	"Gibraltar":                              {"Europe/Gibraltar"},
	"Guinea":                                 {"Africa/Conakry"},
	"Guadeloupe":                             {"America/Guadeloupe"},
	"Gambia":                                 {"Africa/Banjul"},
	"Guinea-Bissau":                          {"Africa/Bissau"},
	"Equatorial Guinea":                      {"Africa/Malabo"},
	"Greece":                                 {"Europe/Athens"},
	"Grenada":                                {"America/Grenada"},
	"Greenland":                              {"America/Nuuk", "America/Danmarkshavn", "America/Scoresbysund", "America/Thule"},
	"Guatemala":                              {"America/Guatemala"},
	"French Guiana":                          {"America/Cayenne"},
	"Guam":                                   {"Pacific/Guam"},
	"Guyana":                                 {"America/Guyana"},
	"Hong Kong":                              {"Asia/Hong_Kong"},
	"Honduras":                               {"America/Tegucigalpa"},
	"Croatia":                                {"Europe/Zagreb"},
	"Haiti":                                  {"America/Port-au-Prince"},
	"Hungary":                                {"Europe/Budapest"},
	"Indonesia":                              {"Asia/Jakarta", "Asia/Pontianak", "Asia/Makassar", "Asia/Jayapura"},
	"Isle of Man":                            {"Europe/Isle_of_Man"},
	"India":                                  {"Asia/Kolkata"},
	"British Indian Ocean Territory":         {"Indian/Chagos"},
	"Ireland":                                {"Europe/Dublin"},
	"Iran, Islamic Republic of":              {"Asia/Tehran"},
	"Iraq":                                   {"Asia/Baghdad"},
	"Iceland":                                {"Atlantic/Reykjavik"},
	"Israel":                                 {"Asia/Jerusalem"},
	"Italy":                                  {"Europe/Rome"},
	"Jamaica":                                {"America/Jamaica"},
	"Jersey":                                 {"Europe/Jersey"},
	"Jordan":                                 {"Asia/Amman"},
	"Japan":                                  {"Asia/Tokyo"},
	"Kazakhstan":                             {"Asia/Almaty", "Asia/Qyzylorda", "Asia/Qostanay", "Asia/Aqtobe", "Asia/Aqtau", "Asia/Atyrau", "Asia/Oral"},
	"Kenya":                                  {"Africa/Nairobi"},
	"Kyrgyzstan":                             {"Asia/Bishkek"},
	"Cambodia":                               {"Asia/Phnom_Penh"},
	"Kiribati":                               {"Pacific/Tarawa", "Pacific/Enderbury", "Pacific/Kiritimati"},
	"Saint Kitts and Nevis":                  {"America/St_Kitts"},
	"Korea, Republic of":                     {"Asia/Seoul"},
	"Kuwait":                                 {"Asia/Kuwait"},
	"Lao People's Democratic Republic":       {"Asia/Vientiane"},
	"Lebanon":                                {"Asia/Beirut"},
	"Liberia":                                {"Africa/Monrovia"},
	"Libya":                                  {"Africa/Tripoli"},
	"Saint Lucia":                            {"America/St_Lucia"},
	"Liechtenstein":                          {"Europe/Vaduz"},
	"Sri Lanka":                              {"Asia/Colombo"},
	"Lesotho":                                {"Africa/Maseru"},
	"Lithuania":                              {"Europe/Vilnius"},
	"Luxembourg":                             {"Europe/Luxembourg"},
	"Latvia":                                 {"Europe/Riga"},
	"Macao":                                  {"Asia/Macau"},
	"Saint Martin (French part)":             {"America/Marigot"},
	"Morocco":                                {"Africa/Casablanca"},
	"Monaco":                                 {"Europe/Monaco"},
	"Moldova, Republic of":                   {"Europe/Chisinau"},
	"Madagascar":                             {"Indian/Antananarivo"},
	"Maldives":                               {"Indian/Maldives"},
	"Mexico":                                 {"America/Mexico_City", "America/Cancun", "America/Merida", "America/Monterrey", "America/Matamoros", "America/Mazatlan", "America/Chihuahua", "America/Ojinaga", "America/Hermosillo", "America/Tijuana", "America/Bahia_Banderas"},
	"Marshall Islands":                       {"Pacific/Majuro", "Pacific/Kwajalein"},
	"North Macedonia":                        {"Europe/Skopje"},
	"Mali":                                   {"Africa/Bamako"},
	"Malta":                                  {"Europe/Malta"},
	"Myanmar":                                {"Asia/Yangon"},
	"Montenegro":                             {"Europe/Podgorica"},
	"Mongolia":                               {"Asia/Ulaanbaatar", "Asia/Hovd", "Asia/Choibalsan"},
	"Northern Mariana Islands":               {"Pacific/Saipan"},
	"Mozambique":                             {"Africa/Maputo"},
	"Mauritania":                             {"Africa/Nouakchott"},
	"Montserrat":                             {"America/Montserrat"},
	"Martinique":                             {"America/Martinique"},
	"Mauritius":                              {"Indian/Mauritius"},
	"Malawi":                                 {"Africa/Blantyre"},
	"Malaysia":                               {"Asia/Kuala_Lumpur", "Asia/Kuching"},
	"Mayotte":                                {"Indian/Mayotte"},
	"Namibia":                                {"Africa/Windhoek"},
	"New Caledonia":                          {"Pacific/Noumea"},
	"Niger":                                  {"Africa/Niamey"},
	"Norfolk Island":                         {"Pacific/Norfolk"},
	"Nigeria":                                {"Africa/Lagos"},
	"Nicaragua":                              {"America/Managua"},
	"Niue":                                   {"Pacific/Niue"},
	"Netherlands":                            {"Europe/Amsterdam"},
	"Norway":                                 {"Europe/Oslo"},
	"Nepal":                                  {"Asia/Kathmandu"},
	"Nauru":                                  {"Pacific/Nauru"},
	"New Zealand":                            {"Pacific/Auckland", "Pacific/Chatham"},
	"Oman":                                   {"Asia/Muscat"},
	"Pakistan":                               {"Asia/Karachi"},
	"Panama":                                 {"America/Panama"},
	"Pitcairn":                               {"Pacific/Pitcairn"},
	"Peru":                                   {"America/Lima"},
	"Philippines":                            {"Asia/Manila"},
	"Palau":                                  {"Pacific/Palau"},
	"Papua New Guinea":                       {"Pacific/Port_Moresby", "Pacific/Bougainville"},
	"Poland":                                 {"Europe/Warsaw"},
	"Puerto Rico":                            {"America/Puerto_Rico"},
	"Korea, Democratic People's Republic of": {"Asia/Pyongyang"},
	"Portugal":                               {"Europe/Lisbon", "Atlantic/Madeira", "Atlantic/Azores"},
	"Paraguay":                               {"America/Asuncion"},
	"Palestine, State of":                    {"Asia/Gaza", "Asia/Hebron"},
	"French Polynesia":                       {"Pacific/Tahiti", "Pacific/Marquesas", "Pacific/Gambier"},
	"Qatar":                                  {"Asia/Qatar"},
	"Réunion":                                {"Indian/Reunion"},
	"Romania":                                {"Europe/Bucharest"},
	"Russian Federation":                     {"Europe/Kaliningrad", "Europe/Moscow", "Europe/Kirov", "Europe/Volgograd", "Europe/Astrakhan", "Europe/Saratov", "Europe/Ulyanovsk", "Europe/Samara", "Asia/Yekaterinburg", "Asia/Omsk", "Asia/Novosibirsk", "Asia/Barnaul", "Asia/Tomsk", "Asia/Novokuznetsk", "Asia/Krasnoyarsk", "Asia/Irkutsk", "Asia/Chita", "Asia/Yakutsk", "Asia/Khandyga", "Asia/Vladivostok", "Asia/Ust-Nera", "Asia/Magadan", "Asia/Sakhalin", "Asia/Srednekolymsk", "Asia/Kamchatka", "Asia/Anadyr"},
	"Rwanda":                                 {"Africa/Kigali"},
	"Saudi Arabia":                           {"Asia/Riyadh"},
	"Sudan":                                  {"Africa/Khartoum"},
	"Senegal":                                {"Africa/Dakar"},
	"Singapore":                              {"Asia/Singapore"},
	"South Georgia and the South Sandwich Islands": {"Atlantic/South_Georgia"},
	"Saint Helena, Ascension and Tristan da Cunha": {"Atlantic/St_Helena"},
	"Svalbard and Jan Mayen":                       {"Arctic/Longyearbyen"},
	"Solomon Islands":                              {"Pacific/Guadalcanal"},
	"Sierra Leone":                                 {"Africa/Freetown"},
	"El Salvador":                                  {"America/El_Salvador"},
	"San Marino":                                   {"Europe/San_Marino"},
	"Somalia":                                      {"Africa/Mogadishu"},
	"Saint Pierre and Miquelon":                    {"America/Miquelon"},
	"Serbia":                                       {"Europe/Belgrade"},
	"South Sudan":                                  {"Africa/Juba"},
	"Sao Tome and Principe":                        {"Africa/Sao_Tome"},
	"Suriname":                                     {"America/Paramaribo"},
	"Slovakia":                                     {"Europe/Bratislava"},
	"Slovenia":                                     {"Europe/Ljubljana"},
	"Sweden":                                       {"Europe/Stockholm"},
	"Eswatini":                                     {"Africa/Mbabane"},
	"Sint Maarten (Dutch part)":                    {"America/Lower_Princes"},
	"Seychelles":                                   {"Indian/Mahe"},
	"Syrian Arab Republic":                         {"Asia/Damascus"},
	"Turks and Caicos Islands":                     {"America/Grand_Turk"},
	"Chad":                                         {"Africa/Ndjamena"},
	"Togo":                                         {"Africa/Lome"},
	"Thailand":                                     {"Asia/Bangkok"},
	"Tajikistan":                                   {"Asia/Dushanbe"},
	"Tokelau":                                      {"Pacific/Fakaofo"},
	"Turkmenistan":                                 {"Asia/Ashgabat"},
	"Timor-Leste":                                  {"Asia/Dili"},
	"Tonga":                                        {"Pacific/Tongatapu"},
	"Trinidad and Tobago":                          {"America/Port_of_Spain"},
	"Tunisia":                                      {"Africa/Tunis"},
	"Türkiye":                                      {"Europe/Istanbul"},
	"Tuvalu":                                       {"Pacific/Funafuti"},
	"Taiwan, Province of China":                    {"Asia/Taipei"},
	"Tanzania, United Republic of":                 {"Africa/Dar_es_Salaam"},
	"Uganda":                                       {"Africa/Kampala"},
	"Ukraine":                                      {"Europe/Simferopol", "Europe/Kiev", "Europe/Uzhgorod", "Europe/Zaporozhye"},
	"United States Minor Outlying Islands":         {"Pacific/Midway", "Pacific/Wake"},
	"Uruguay":                                      {"America/Montevideo"},
	"United States":                                {"America/New_York", "America/Detroit", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Indiana/Indianapolis", "America/Indiana/Vincennes", "America/Indiana/Winamac", "America/Indiana/Marengo", "America/Indiana/Petersburg", "America/Indiana/Vevay", "America/Chicago", "America/Indiana/Tell_City", "America/Indiana/Knox", "America/Menominee", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/North_Dakota/Beulah", "America/Denver", "America/Boise", "America/Phoenix", "America/Los_Angeles", "America/Anchorage", "America/Juneau", "America/Sitka", "America/Metlakatla", "America/Yakutat", "America/Nome", "America/Adak", "Pacific/Honolulu"},
	"Uzbekistan":                                   {"Asia/Samarkand", "Asia/Tashkent"},
	"Holy See (Vatican City State)":                {"Europe/Vatican"},
	"Saint Vincent and the Grenadines":             {"America/St_Vincent"},
	"Venezuela, Bolivarian Republic of":            {"America/Caracas"},
	"Virgin Islands, British":                      {"America/Tortola"},
	"Virgin Islands, U.S.":                         {"America/St_Thomas"},
	"Viet Nam":                                     {"Asia/Ho_Chi_Minh"},
	"Vanuatu":                                      {"Pacific/Efate"},
	"Wallis and Futuna":                            {"Pacific/Wallis"},
	"Samoa":                                        {"Pacific/Apia"},
	"Yemen":                                        {"Asia/Aden"},
	"South Africa":                                 {"Africa/Johannesburg"},
	"Zambia":                                       {"Africa/Lusaka"},
	"Zimbabwe":                                     {"Africa/Harare"},
}
