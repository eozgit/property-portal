package main

var data = `1,Birmingham,"1,141,816","Metropolitan borough, City (1889)",West Midlands,West Midlands
2,Leeds,"793,139","Metropolitan borough, City (1893)",West Yorkshire,Yorkshire and the Humber
3,Sheffield,"584,853","Metropolitan borough, City (1893)",South Yorkshire,Yorkshire and the Humber
4,Cornwall,"569,578",Unitary authority,Cornwall,South West
5,Manchester,"552,858","Metropolitan borough, City (1853)",Greater Manchester,North West
6,Buckinghamshire,"543,973",Unitary authority,Buckinghamshire,South East
7,Bradford,"539,776","Metropolitan borough, City (1897)",West Yorkshire,Yorkshire and the Humber
8,County Durham,"530,094",Unitary authority,Durham,North East
9,Wiltshire,"500,024",Unitary authority,Wiltshire,South West
10,Liverpool,"498,042","Metropolitan borough, City (1880)",Merseyside,North West
11,Bristol,"463,377","Unitary authority, City (1542)",Bristol,South West
12,Kirklees,"439,787",Metropolitan borough,West Yorkshire,Yorkshire and the Humber
13,Barnet,"395,869",London borough,Greater London,London
14,"Bournemouth, Christchurch and Poole","395,331",Unitary authority,Dorset,South West
15,Croydon,"386,710",London borough,Greater London,London
16,Cheshire East,"384,152",Unitary authority,Cheshire,North West
17,Dorset,"378,508",Unitary authority,Dorset,South West
18,Coventry,"371,521","Metropolitan borough, City (1345)",West Midlands,West Midlands
19,Leicester,"354,224","Unitary authority, City (1919)",Leicestershire,East Midlands
20,Newham,"353,134",London borough,Greater London,London
21,Wakefield,"348,312","Metropolitan borough, City (1888)",West Yorkshire,Yorkshire and the Humber
22,Cheshire West and Chester,"343,071",Unitary authority,Cheshire,North West
23,Ealing,"341,806",London borough,Greater London,London
24,East Riding of Yorkshire,"341,173",Unitary authority,East Riding of Yorkshire,Yorkshire and the Humber
25,Enfield,"333,794",London borough,Greater London,London
26,Nottingham,"332,900","Unitary authority, City (1897)",Nottinghamshire,East Midlands
27,Bromley,"332,336",London borough,Greater London,London
28,Brent,"329,771",London borough,Greater London,London
29,Wandsworth,"329,677",London borough,Greater London,London
30,Wigan,"328,662",Metropolitan borough,Greater Manchester,North West
31,Sandwell,"328,450",Metropolitan borough,West Midlands,West Midlands
32,Lambeth,"326,034",London borough,Greater London,London
33,Tower Hamlets,"324,745",London borough,Greater London,London
34,Wirral,"324,011",Metropolitan borough,Merseyside,North West
35,Shropshire,"323,136",Unitary authority,Shropshire,West Midlands
36,Northumberland,"322,434",Unitary authority,Northumberland,North East
37,Dudley,"321,596",Metropolitan borough,West Midlands,West Midlands
38,Southwark,"318,830",London borough,Greater London,London
39,Doncaster,"311,890",Metropolitan borough,South Yorkshire,Yorkshire and the Humber
40,Hillingdon,"306,870",London borough,Greater London,London
41,Lewisham,"305,842",London borough,Greater London,London
42,Redbridge,"305,222",London borough,Greater London,London
43,Newcastle upon Tyne,"302,820","Metropolitan borough, City (1882)",Tyne and Wear,North East
44,Stockport,"293,423",Metropolitan borough,Greater Manchester,North West
45,Brighton & Hove,"290,885","Unitary authority, City (2000)",East Sussex,South East
46,Central Bedfordshire,"288,648",Unitary authority,Bedfordshire,East of England
47,Greenwich,"287,942","London borough, Royal Borough",Greater London,London
48,Bolton,"287,550",Metropolitan borough,Greater Manchester,North West
49,Walsall,"285,478",Metropolitan borough,West Midlands,West Midlands
50,South Gloucestershire,"285,093",Unitary authority,Gloucestershire,South West
51,Hackney,"281,120",London borough,Greater London,London
52,Medway,"278,556",Unitary authority,Kent,South East
53,Sunderland,"277,705","Metropolitan borough, City (1992)",Tyne and Wear,North East
54,Waltham Forest,"276,983",London borough,Greater London,London
55,Sefton,"276,410",Metropolitan borough,Merseyside,North West
56,Hounslow,"271,523",London borough,Greater London,London
57,Camden,"270,029",London borough,Greater London,London
58,Milton Keynes,"269,457","Unitary authority, Borough",Buckinghamshire,South East
59,Haringey,"268,647",London borough,Greater London,London
60,Rotherham,"265,411",Metropolitan borough,South Yorkshire,Yorkshire and the Humber
61,Wolverhampton,"263,357","Metropolitan borough, City (2000)",West Midlands,West Midlands
62,Plymouth,"262,100","Unitary authority, City (1928)",Devon,South West
63,Westminster,"261,317","London borough, City (1540)",Greater London,London
64,Kingston upon Hull,"259,778","Unitary authority, City (1299)",East Riding of Yorkshire,Yorkshire and the Humber
65,Havering,"259,552",London borough,Greater London,London
66,Salford,"258,834","Metropolitan borough, City (1926)",Greater Manchester,North West
67,Derby,"257,302","Unitary authority, City (1977)",Derbyshire,East Midlands
68,Stoke-on-Trent,"256,375","Unitary authority, City (1925)",Staffordshire,West Midlands
69,Southampton,"252,520","Unitary authority, City (1964)",Hampshire,South East
70,Harrow,"251,160",London borough,Greater London,London
71,East Suffolk,"249,461",Non-metropolitan district,Suffolk,East of England
72,Bexley,"248,287",London borough,Greater London,London
73,Barnsley,"246,866",Metropolitan borough,South Yorkshire,Yorkshire and the Humber
74,Islington,"242,467",London borough,Greater London,London
75,Trafford,"237,354",Metropolitan borough,Greater Manchester,North West
76,Oldham,"237,110",Metropolitan borough,Greater Manchester,North West
77,Tameside,"226,493",Metropolitan borough,Greater Manchester,North West
78,Northampton,"224,610","Non-metropolitan district, Borough",Northamptonshire,East Midlands
79,Rochdale,"222,412",Metropolitan borough,Greater Manchester,North West
80,Swindon,"222,193","Unitary authority, Borough",Wiltshire,South West
81,Solihull,"216,374",Metropolitan borough,West Midlands,West Midlands
82,North Somerset,"215,052",Unitary authority,Somerset,South West
83,Portsmouth,"214,905","Unitary authority, City (1926)",Hampshire,South East
84,Luton,"213,052","Unitary authority, Borough",Bedfordshire,East of England
85,Barking and Dagenham,"212,906",London borough,Greater London,London
86,Calderdale,"211,455",Metropolitan borough,West Yorkshire,Yorkshire and the Humber
87,York,"210,618","Unitary authority, City (TI)",North Yorkshire,Yorkshire and the Humber
88,Warrington,"210,014","Unitary authority, Borough",Cheshire,North West
89,North Tyneside,"207,913",Metropolitan borough,Tyne and Wear,North East
90,Merton,"206,548",London borough,Greater London,London
91,Sutton,"206,349",London borough,Greater London,London
92,Peterborough,"202,259","Unitary authority, City (1541)",Cambridgeshire,East of England
93,Gateshead,"202,055",Metropolitan borough,Tyne and Wear,North East
94,Richmond upon Thames,"198,019",London borough,Greater London,London
95,Stockton-on-Tees,"197,348","Unitary authority, Borough",Durham andNorth Yorkshire,North East
96,Colchester,"194,706","Non-metropolitan district, Borough",Essex,East of England
97,Bath and North East Somerset,"193,282",Unitary authority,Somerset,South West
98,Herefordshire,"192,801",Unitary authority,Herefordshire,West Midlands
99,Bury,"190,990",Metropolitan borough,Greater Manchester,North West
100,Basildon,"187,199","Non-metropolitan district, Borough",Essex,East of England
101,Charnwood,"185,851","Non-metropolitan district, Borough",Leicestershire,East Midlands
102,Hammersmith and Fulham,"185,143",London borough,Greater London,London
103,Southend-on-Sea,"183,125","Unitary authority, Borough",Essex,East of England
104,St Helens,"180,585",Metropolitan borough,Merseyside,North West
105,New Forest,"180,086",Non-metropolitan district,Hampshire,South East
106,Telford and Wrekin,"179,854","Unitary authority, Borough",Shropshire,West Midlands
107,West Suffolk,"179,045",Non-metropolitan district,Suffolk,East of England
108,Chelmsford,"178,388","Non-metropolitan district, City (2012)",Essex,East of England
109,Huntingdonshire,"177,963",Non-metropolitan district,Cambridgeshire,East of England
110,Kingston upon Thames,"177,507","London borough, Royal Borough",Greater London,London
111,Basingstoke and Deane,"176,582","Non-metropolitan district, Borough",Hampshire,South East
112,Thurrock,"174,341","Unitary authority, Borough",Essex,East of England
113,Bedford,"173,292","Unitary authority, Borough",Bedfordshire,East of England
114,North Lincolnshire,"172,292","Unitary authority, Borough",Lincolnshire,Yorkshire and the Humber
115,Maidstone,"171,826","Non-metropolitan district, Borough",Kent,South East
116,Wokingham,"171,119",Unitary authority,Berkshire,South East
117,South Somerset,"168,345",Non-metropolitan district,Somerset,South West
118,Canterbury,"165,394","Non-metropolitan district, City (TI)",Kent,South East
119,Reading,"161,780","Unitary authority, Borough",Berkshire,South East
120,Wealden,"161,475",Non-metropolitan district,East Sussex,South East
121,Harrogate,"160,831","Non-metropolitan district, Borough",North Yorkshire,Yorkshire and the Humber
122,Arun,"160,758",Non-metropolitan district,West Sussex,South East
123,North East Lincolnshire,"159,563","Unitary authority, Borough",Lincolnshire,Yorkshire and the Humber
124,South Cambridgeshire,"159,086",Non-metropolitan district,Cambridgeshire,East of England
125,West Berkshire,"158,450",Unitary authority,Berkshire,South East
126,Kensington and Chelsea,"156,129","London borough, Royal Borough",Greater London,London
127,Somerset West and Taunton,"155,115",Non-metropolitan district,Somerset,South West
128,Dacorum,"154,763","Non-metropolitan district, Borough",Hertfordshire,East of England
129,Braintree,"152,604",Non-metropolitan district,Essex,East of England
130,Oxford,"152,457","Non-metropolitan district, City (1542)",Oxfordshire,South East
131,Windsor and Maidenhead,"151,422","Unitary authority, Royal Borough",Berkshire,South East
132,King's Lynn and West Norfolk,"151,383","Non-metropolitan district, Borough",Norfolk,East of England
133,Mid Sussex,"151,022",Non-metropolitan district,West Sussex,South East
134,South Tyneside,"150,976",Metropolitan borough,Tyne and Wear,North East
135,Knowsley,"150,862",Metropolitan borough,Merseyside,North West
136,Cherwell,"150,503",Non-metropolitan district,Oxfordshire,South East
137,Swale,"150,082","Non-metropolitan district, Borough",Kent,South East
138,East Hertfordshire,"149,748",Non-metropolitan district,Hertfordshire,East of England
139,Blackburn with Darwen,"149,696","Unitary authority, Borough",Lancashire,North West
140,Slough,"149,539","Unitary authority, Borough",Berkshire,South East
141,Guildford,"148,998","Non-metropolitan district, Borough",Surrey,South East
142,Reigate and Banstead,"148,748","Non-metropolitan district, Borough",Surrey,South East
143,St Albans,"148,452","Non-metropolitan district, City (1877)",Hertfordshire,East of England
144,Tendring,"146,561",Non-metropolitan district,Essex,East of England
145,East Devon,"146,284",Non-metropolitan district,Devon,South West
146,Lancaster,"146,038","Non-metropolitan district, City (1937)",Lancashire,North West
147,Horsham,"143,791",Non-metropolitan district,West Sussex,South East
148,Warwick,"143,753",Non-metropolitan district,Warwickshire,West Midlands
149,Preston,"143,135","Non-metropolitan district, City (2002)",Lancashire,North West
150,South Kesteven,"142,424",Non-metropolitan district,Lincolnshire,East Midlands
151,South Oxfordshire,"142,057",Non-metropolitan district,Oxfordshire,South East
152,Thanet,"141,922",Non-metropolitan district,Kent,South East
153,Isle of Wight,"141,771","Unitary authority, County",Isle of Wight,South East
154,East Lindsey,"141,727",Non-metropolitan district,Lincolnshire,East Midlands
155,Middlesbrough,"140,980","Unitary authority, Borough",North Yorkshire,North East
156,South Norfolk,"140,880",Non-metropolitan district,Norfolk,East of England
157,Norwich,"140,573","Non-metropolitan district, City (1195)",Norfolk,East of England
158,Breckland,"139,968",Non-metropolitan district,Norfolk,East of England
159,Blackpool,"139,446","Unitary authority, Borough",Lancashire,North West
160,Stafford,"137,280",Non-metropolitan district,Staffordshire,West Midlands
161,Redcar and Cleveland,"137,150","Unitary authority, Borough",North Yorkshire,North East
162,Ipswich,"136,913","Non-metropolitan district, Borough",Suffolk,East of England
163,Elmbridge,"136,795","Non-metropolitan district, Borough",Surrey,South East
164,Torbay,"136,264","Unitary authority, Borough",Devon,South West
165,Vale of White Horse,"136,007",Non-metropolitan district,Oxfordshire,South East
166,Teignbridge,"134,163",Non-metropolitan district,Devon,South West
167,Eastleigh,"133,584","Non-metropolitan district, Borough",Hampshire,South East
168,North Hertfordshire,"133,570",Non-metropolitan district,Hertfordshire,East of England
169,Tonbridge and Malling,"132,153","Non-metropolitan district, Borough",Kent,South East
170,Epping Forest,"131,689",Non-metropolitan district,Essex,East of England
171,Exeter,"131,405","Non-metropolitan district, City (TI)",Devon,South West
172,Broadland,"130,783",Non-metropolitan district,Norfolk,East of England
173,Stratford-on-Avon,"130,098",Non-metropolitan district,Warwickshire,West Midlands
174,Ashford,"130,032","Non-metropolitan district, Borough",Kent,South East
175,Nuneaton and Bedworth,"129,883","Non-metropolitan district, Borough",Warwickshire,West Midlands
176,Newcastle-under-Lyme,"129,441","Non-metropolitan district, Borough",Staffordshire,West Midlands
177,Wychavon,"129,433",Non-metropolitan district,Worcestershire,West Midlands
178,Halton,"129,410","Unitary authority, Borough",Cheshire,North West
179,Gloucester,"129,128","Non-metropolitan district, City (1541)",Gloucestershire,South West
180,Amber Valley,"128,147","Non-metropolitan district, Borough",Derbyshire,East Midlands
181,Ashfield,"127,918",Non-metropolitan district,Nottinghamshire,East Midlands
182,Waverley,"126,328","Non-metropolitan district, Borough",Surrey,South East
183,Havant,"126,220","Non-metropolitan district, Borough",Hampshire,South East
184,Test Valley,"126,160","Non-metropolitan district, Borough",Hampshire,South East
185,Winchester,"124,859","Non-metropolitan district, City (TI)",Hampshire,South East
186,Cambridge,"124,798","Non-metropolitan district, City (1951)",Cambridgeshire,East of England
187,Sedgemoor,"123,178",Non-metropolitan district,Somerset,South West
188,Welwyn Hatfield,"123,043",Non-metropolitan district,Hertfordshire,East of England
189,Bracknell Forest,"122,549","Unitary authority, Borough",Berkshire,South East
190,Newark and Sherwood,"122,421",Non-metropolitan district,Nottinghamshire,East Midlands
191,East Hampshire,"122,308",Non-metropolitan district,Hampshire,South East
192,Chichester,"121,129",Non-metropolitan district,West Sussex,South East
193,Sevenoaks,"120,750",Non-metropolitan district,Kent,South East
194,Stroud,"119,964",Non-metropolitan district,Gloucestershire,South West
195,East Staffordshire,"119,754","Non-metropolitan district, Borough",Staffordshire,West Midlands
196,Rushcliffe,"119,184","Non-metropolitan district, Borough",Nottinghamshire,East Midlands
197,Tunbridge Wells,"118,724","Non-metropolitan district, Borough",Kent,South East
198,Chorley,"118,216","Non-metropolitan district, Borough",Lancashire,North West
199,Dover,"118,131",Non-metropolitan district,Kent,South East
200,Gedling,"117,896","Non-metropolitan district, Borough",Nottinghamshire,East Midlands
201,Bassetlaw,"117,459",Non-metropolitan district,Nottinghamshire,East Midlands
202,North Kesteven,"116,915",Non-metropolitan district,Lincolnshire,East Midlands
203,Cheltenham,"116,306","Non-metropolitan district, Borough",Gloucestershire,South West
204,Fareham,"116,233","Non-metropolitan district, Borough",Hampshire,South East
205,Mendip,"115,587",Non-metropolitan district,Somerset,South West
206,Erewash,"115,371","Non-metropolitan district, Borough",Derbyshire,East Midlands
207,West Lancashire,"114,306",Non-metropolitan district,Lancashire,North West
208,Broxtowe,"114,033","Non-metropolitan district, Borough",Nottinghamshire,East Midlands
209,Hinckley and Bosworth,"113,136","Non-metropolitan district, Borough",Leicestershire,East Midlands
210,Folkestone and Hythe,"112,996",Non-metropolitan district,Kent,South East
211,Dartford,"112,606","Non-metropolitan district, Borough",Kent,South East
212,South Staffordshire,"112,436",Non-metropolitan district,Staffordshire,West Midlands
213,Crawley,"112,409","Non-metropolitan district, Borough",West Sussex,South East
214,Wyre,"112,091","Non-metropolitan district, Borough",Lancashire,North West
215,South Ribble,"110,788","Non-metropolitan district, Borough",Lancashire,North West
216,West Oxfordshire,"110,643",Non-metropolitan district,Oxfordshire,South East
217,Worthing,"110,570","Non-metropolitan district, Borough",West Sussex,South East
218,Mansfield,"109,313",Non-metropolitan district,Nottinghamshire,East Midlands
219,Rugby,"108,935","Non-metropolitan district, Borough",Warwickshire,West Midlands
220,Scarborough,"108,757","Non-metropolitan district, Borough",North Yorkshire,Yorkshire and the Humber
221,Carlisle,"108,678","Non-metropolitan district, City (TI)",Cumbria,North West
222,South Derbyshire,"107,261",Non-metropolitan district,Derbyshire,East Midlands
223,Gravesham,"106,939","Non-metropolitan district, Borough",Kent,South East
224,Darlington,"106,803","Unitary authority, Borough",Durham,North East
225,South Lakeland,"105,088",Non-metropolitan district,Cumbria,North West
226,Hertsmere,"104,919","Non-metropolitan district, Borough",Hertfordshire,East of England
227,Chesterfield,"104,900","Non-metropolitan district, Borough",Derbyshire,East Midlands
228,North Norfolk,"104,837",Non-metropolitan district,Norfolk,East of England
229,Lichfield,"104,756",Non-metropolitan district,Staffordshire,West Midlands
230,Mid Suffolk,"103,895",Non-metropolitan district,Suffolk,East of England
231,Eastbourne,"103,745","Non-metropolitan district, Borough",East Sussex,South East
232,North West Leicestershire,"103,611",Non-metropolitan district,Leicestershire,East Midlands
233,Lewes,"103,268",Non-metropolitan district,East Sussex,South East
234,Fenland,"101,850",Non-metropolitan district,Cambridgeshire,East of England
235,Kettering,"101,776","Non-metropolitan district, Borough",Northamptonshire,East Midlands
236,Blaby,"101,526",Non-metropolitan district,Leicestershire,East Midlands
237,North East Derbyshire,"101,462",Non-metropolitan district,Derbyshire,East Midlands
238,Wyre Forest,"101,291",Non-metropolitan district,Worcestershire,West Midlands
239,Worcester,"101,222","Non-metropolitan district, City (1189)",Worcestershire,West Midlands
240,Woking,"100,793","Non-metropolitan district, Borough",Surrey,South East
241,Cannock Chase,"100,762",Non-metropolitan district,Staffordshire,West Midlands
242,Bromsgrove,"99,881",Non-metropolitan district,Worcestershire,West Midlands
243,Spelthorne,"99,844","Non-metropolitan district, Borough",Surrey,South East
244,Great Yarmouth,"99,336","Non-metropolitan district, Borough",Norfolk,East of England
245,Lincoln,"99,299","Non-metropolitan district, City (TI)",Lincolnshire,East Midlands
246,Staffordshire Moorlands,"98,435",Non-metropolitan district,Staffordshire,West Midlands
247,Allerdale,"97,761","Non-metropolitan district, Borough",Cumbria,North West
248,Broxbourne,"97,279","Non-metropolitan district, Borough",Hertfordshire,East of England
249,North Devon,"97,145",Non-metropolitan district,Devon,South West
250,Hart,"97,073",Non-metropolitan district,Hampshire,South East
251,Watford,"96,577","Non-metropolitan district, Borough",Hertfordshire,East of England
252,Rother,"96,080",Non-metropolitan district,East Sussex,South East
253,West Lindsey,"95,667",Non-metropolitan district,Lincolnshire,East Midlands
254,South Holland,"95,019",Non-metropolitan district,Lincolnshire,East Midlands
254,Tewkesbury,"95,019","Non-metropolitan district, Borough",Gloucestershire,South West
256,Rushmoor,"94,599","Non-metropolitan district, Borough",Hampshire,South East
257,East Northamptonshire,"94,527",Non-metropolitan district,Northamptonshire,East Midlands
258,South Northamptonshire,"94,490",Non-metropolitan district,Northamptonshire,East Midlands
259,Harborough,"93,807",Non-metropolitan district,Leicestershire,East Midlands
260,Hartlepool,"93,663","Unitary authority, Borough",Durham,North East
261,Three Rivers,"93,323",Non-metropolitan district,Hertfordshire,East of England
262,High Peak,"92,666","Non-metropolitan district, Borough",Derbyshire,East Midlands
263,Hastings,"92,661","Non-metropolitan district, Borough",East Sussex,South East
264,Pendle,"92,112","Non-metropolitan district, Borough",Lancashire,North West
265,Babergh,"92,036",Non-metropolitan district,Suffolk,East of England
266,Hambleton,"91,594",Non-metropolitan district,North Yorkshire,Yorkshire and the Humber
267,Uttlesford,"91,284",Non-metropolitan district,Essex,East of England
268,Selby,"90,620",Non-metropolitan district,North Yorkshire,Yorkshire and the Humber
269,Castle Point,"90,376","Non-metropolitan district, Borough",Essex,East of England
270,Cotswold,"89,862",Non-metropolitan district,Gloucestershire,South West
271,East Cambridgeshire,"89,840",Non-metropolitan district,Cambridgeshire,East of England
272,Runnymede,"89,424","Non-metropolitan district, Borough",Surrey,South East
273,Surrey Heath,"89,305","Non-metropolitan district, Borough",Surrey,South East
274,Burnley,"88,920","Non-metropolitan district, Borough",Lancashire,North West
275,Tandridge,"88,129",Non-metropolitan district,Surrey,South East
276,Stevenage,"87,845","Non-metropolitan district, Borough",Hertfordshire,East of England
277,Rochford,"87,368",Non-metropolitan district,Essex,East of England
278,Mole Valley,"87,245",Non-metropolitan district,Surrey,South East
279,Harlow,"87,067",Non-metropolitan district,Essex,East of England
280,South Hams,"87,004",Non-metropolitan district,Devon,South West
281,Forest of Dean,"86,791",Non-metropolitan district,Gloucestershire,South West
282,Daventry,"85,950",Non-metropolitan district,Northamptonshire,East Midlands
283,Redditch,"85,261","Non-metropolitan district, Borough",Worcestershire,West Midlands
284,Gosport,"84,838","Non-metropolitan district, Borough",Hampshire,South East
285,Mid Devon,"82,311",Non-metropolitan district,Devon,South West
286,Hyndburn,"81,043","Non-metropolitan district, Borough",Lancashire,North West
287,Fylde,"80,780","Non-metropolitan district, Borough",Lancashire,North West
288,Epsom and Ewell,"80,627","Non-metropolitan district, Borough",Surrey,South East
289,Bolsover,"80,562",Non-metropolitan district,Derbyshire,East Midlands
290,Wellingborough,"79,707","Non-metropolitan district, Borough",Northamptonshire,East Midlands
291,Malvern Hills,"78,698",Non-metropolitan district,Worcestershire,West Midlands
292,Brentwood,"77,021","Non-metropolitan district, Borough",Essex,East of England
293,Tamworth,"76,696","Non-metropolitan district, Borough",Staffordshire,West Midlands
294,Derbyshire Dales,"72,325",Non-metropolitan district,Derbyshire,East Midlands
295,Corby,"72,218","Non-metropolitan district, Borough",Northamptonshire,East Midlands
296,Rossendale,"71,482","Non-metropolitan district, Borough",Lancashire,North West
297,Boston,"70,173","Non-metropolitan district, Borough",Lincolnshire,East Midlands
298,Torridge,"68,267",Non-metropolitan district,Devon,South West
299,Copeland,"68,183","Non-metropolitan district, Borough",Cumbria,North West
300,Barrow-in-Furness,"67,049","Non-metropolitan district, Borough",Cumbria,North West
301,North Warwickshire,"65,264","Non-metropolitan district, Borough",Warwickshire,West Midlands
302,Maldon,"64,926",Non-metropolitan district,Essex,East of England
303,Adur,"64,301",Non-metropolitan district,West Sussex,South East
304,Ribble Valley,"60,888","Non-metropolitan district, Borough",Lancashire,North West
305,Craven,"57,142",Non-metropolitan district,North Yorkshire,Yorkshire and the Humber
306,Oadby and Wigston,"57,015","Non-metropolitan district, Borough",Leicestershire,East Midlands
307,West Devon,"55,796","Non-metropolitan district, Borough",Devon,South West
308,Ryedale,"55,380",Non-metropolitan district,North Yorkshire,Yorkshire and the Humber
309,Richmondshire,"53,730",Non-metropolitan district,North Yorkshire,Yorkshire and the Humber
310,Eden,"53,253",Non-metropolitan district,Cumbria,North West
311,Melton,"51,209","Non-metropolitan district, Borough",Leicestershire,East Midlands
312,Rutland,"39,927",Unitary authority,Rutland,East Midlands
313,City of London,"9,721","sui generis, City (TI)",City of London,London
314,Isles of Scilly,"2,224",sui generis,Cornwall,South West`
