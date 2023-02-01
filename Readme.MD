# Итоговая аттестация по программе профессиональной переподготовки «Go-разработчик»

## Конфигурация
настройки хроняться в файле my_conf.yml
```yaml 
apport: 8080    #Порт
logfile: log.txt #Путь до файла логов
smsdata: ..\skillbox-diploma\sms.data #Путь до файла с данными симулятора о СМС
voicedata: ..\skillbox-diploma\voice.data #Путь до файла с данными симулятора со звонками
incidentdata: http://127.0.0.1:8383/accendent #Ссылка на API с тикетами
mmsdata: http://127.0.0.1:8383/mms #Ссылка на API с ММС
supportdata: http://127.0.0.1:8383/support #Ссылка на API с тикетами
emaildata: ..\skillbox-diploma\email.data #Путь до файла с данными симулятора о почте
billing: ..\skillbox-diploma\billing.data #Путь до файла с данными симулятора о билленге 
```

## Запуск
В соотвествии с условиями ТЗ запускаем симулятор, предоварительно указав файлы с данными в my_conf.yml
```bash
Diploma_Skillbox> go run cmd\main.go
```
Стартует сетевая служба на порту указанного в файле конфигурации
Создается запись в лог файле о старте сервиса
```json
{"level":"info","msg":"Server Started","time":"2023-02-01T20:49:44+03:00"}
```

## Работа с сервисом

В соотвествии с заданием  сетевая служба получает GET запрос и возвращает набор данных в формате json в соотвествии с ТЗ
```curl
GET http://localhost:8080/
```

```json
http://localhost:8080

HTTP/1.1 200 OK
Date: Wed, 01 Feb 2023 17:52:52 GMT
Content-Type: text/plain; charset=utf-8
Transfer-Encoding: chunked

{
  "status": true,
  "data": {
   "sms": [
    [
     {
      "country": "New Zealand",
      "bandwidth": "59",
      "response_time": "1660",
      "provider": "Kildy"
     },
     {
      "country": "Saint Barthélemy",
      "bandwidth": "45",
      "response_time": "1153",
      "provider": "Kildy"
     },
     {
      "country": "Bulgaria",
      "bandwidth": "26",
      "response_time": "188",
      "provider": "Rond"
     },
     {
      "country": "United States of America",
      "bandwidth": "46",
      "response_time": "1757",
      "provider": "Rond"
     },
     {
      "country": "Türkiye",
      "bandwidth": "18",
      "response_time": "134",
      "provider": "Rond"
     },
     {
      "country": "Canada",
      "bandwidth": "19",
      "response_time": "89",
      "provider": "Rond"
     },
     {
      "country": "France",
      "bandwidth": "55",
      "response_time": "1105",
      "provider": "Topolo"
     },
     {
      "country": "Denmark",
      "bandwidth": "9",
      "response_time": "521",
      "provider": "Topolo"
     },
     {
      "country": "Austria",
      "bandwidth": "63",
      "response_time": "1922",
      "provider": "Topolo"
     },
     {
      "country": "Spain",
      "bandwidth": "70",
      "response_time": "582",
      "provider": "Topolo"
     },
     {
      "country": "Switzerland",
      "bandwidth": "65",
      "response_time": "1635",
      "provider": "Topolo"
     },
     {
      "country": "Russian Federation",
      "bandwidth": "81",
      "response_time": "799",
      "provider": "Topolo"
     },
     {
      "country": "Peru",
      "bandwidth": "64",
      "response_time": "1556",
      "provider": "Topolo"
     },
     {
      "country": "United Kingdom of Great Britain and Northern Ireland",
      "bandwidth": "12",
      "response_time": "1639",
      "provider": "Topolo"
     }
    ],
    [
     {
      "country": "Austria",
      "bandwidth": "63",
      "response_time": "1922",
      "provider": "Topolo"
     },
     {
      "country": "Bulgaria",
      "bandwidth": "26",
      "response_time": "188",
      "provider": "Rond"
     },
     {
      "country": "Canada",
      "bandwidth": "19",
      "response_time": "89",
      "provider": "Rond"
     },
     {
      "country": "Denmark",
      "bandwidth": "9",
      "response_time": "521",
      "provider": "Topolo"
     },
     {
      "country": "France",
      "bandwidth": "55",
      "response_time": "1105",
      "provider": "Topolo"
     },
     {
      "country": "New Zealand",
      "bandwidth": "59",
      "response_time": "1660",
      "provider": "Kildy"
     },
     {
      "country": "Peru",
      "bandwidth": "64",
      "response_time": "1556",
      "provider": "Topolo"
     },
     {
      "country": "Russian Federation",
      "bandwidth": "81",
      "response_time": "799",
      "provider": "Topolo"
     },
     {
      "country": "Saint Barthélemy",
      "bandwidth": "45",
      "response_time": "1153",
      "provider": "Kildy"
     },
     {
      "country": "Spain",
      "bandwidth": "70",
      "response_time": "582",
      "provider": "Topolo"
     },
     {
      "country": "Switzerland",
      "bandwidth": "65",
      "response_time": "1635",
      "provider": "Topolo"
     },
     {
      "country": "Türkiye",
      "bandwidth": "18",
      "response_time": "134",
      "provider": "Rond"
     },
     {
      "country": "United Kingdom of Great Britain and Northern Ireland",
      "bandwidth": "12",
      "response_time": "1639",
      "provider": "Topolo"
     },
     {
      "country": "United States of America",
      "bandwidth": "46",
      "response_time": "1757",
      "provider": "Rond"
     }
    ]
   ],
   "mms": [
    [
     {
      "country": "Monaco",
      "provider": "Kildy",
      "bandwidth": "29",
      "response_time": "392"
     },
     {
      "country": "New Zealand",
      "provider": "Kildy",
      "bandwidth": "6",
      "response_time": "1229"
     },
     {
      "country": "Saint Barthélemy",
      "provider": "Kildy",
      "bandwidth": "45",
      "response_time": "83"
     },
     {
      "country": "Canada",
      "provider": "Rond",
      "bandwidth": "19",
      "response_time": "1916"
     },
     {
      "country": "United States of America",
      "provider": "Rond",
      "bandwidth": "12",
      "response_time": "1032"
     },
     {
      "country": "Türkiye",
      "provider": "Rond",
      "bandwidth": "74",
      "response_time": "1896"
     },
     {
      "country": "Bulgaria",
      "provider": "Rond",
      "bandwidth": "9",
      "response_time": "236"
     },
     {
      "country": "France",
      "provider": "Topolo",
      "bandwidth": "74",
      "response_time": "772"
     },
     {
      "country": "Denmark",
      "provider": "Topolo",
      "bandwidth": "59",
      "response_time": "1202"
     },
     {
      "country": "Spain",
      "provider": "Topolo",
      "bandwidth": "67",
      "response_time": "56"
     },
     {
      "country": "Switzerland",
      "provider": "Topolo",
      "bandwidth": "97",
      "response_time": "1622"
     },
     {
      "country": "Austria",
      "provider": "Topolo",
      "bandwidth": "93",
      "response_time": "1635"
     },
     {
      "country": "Peru",
      "provider": "Topolo",
      "bandwidth": "63",
      "response_time": "1275"
     },
     {
      "country": "Russian Federation",
      "provider": "Topolo",
      "bandwidth": "12",
      "response_time": "1708"
     },
     {
      "country": "United Kingdom of Great Britain and Northern Ireland",
      "provider": "Topolo",
      "bandwidth": "94",
      "response_time": "927"
     }
    ],
    [
     {
      "country": "Austria",
      "provider": "Topolo",
      "bandwidth": "93",
      "response_time": "1635"
     },
     {
      "country": "Bulgaria",
      "provider": "Rond",
      "bandwidth": "9",
      "response_time": "236"
     },
     {
      "country": "Canada",
      "provider": "Rond",
      "bandwidth": "19",
      "response_time": "1916"
     },
     {
      "country": "Denmark",
      "provider": "Topolo",
      "bandwidth": "59",
      "response_time": "1202"
     },
     {
      "country": "France",
      "provider": "Topolo",
      "bandwidth": "74",
      "response_time": "772"
     },
     {
      "country": "Monaco",
      "provider": "Kildy",
      "bandwidth": "29",
      "response_time": "392"
     },
     {
      "country": "New Zealand",
      "provider": "Kildy",
      "bandwidth": "6",
      "response_time": "1229"
     },
     {
      "country": "Peru",
      "provider": "Topolo",
      "bandwidth": "63",
      "response_time": "1275"
     },
     {
      "country": "Russian Federation",
      "provider": "Topolo",
      "bandwidth": "12",
      "response_time": "1708"
     },
     {
      "country": "Saint Barthélemy",
      "provider": "Kildy",
      "bandwidth": "45",
      "response_time": "83"
     },
     {
      "country": "Spain",
      "provider": "Topolo",
      "bandwidth": "67",
      "response_time": "56"
     },
     {
      "country": "Switzerland",
      "provider": "Topolo",
      "bandwidth": "97",
      "response_time": "1622"
     },
     {
      "country": "Türkiye",
      "provider": "Rond",
      "bandwidth": "74",
      "response_time": "1896"
     },
     {
      "country": "United Kingdom of Great Britain and Northern Ireland",
      "provider": "Topolo",
      "bandwidth": "94",
      "response_time": "927"
     },
     {
      "country": "United States of America",
      "provider": "Rond",
      "bandwidth": "12",
      "response_time": "1032"
     }
    ]
   ],
   "voice_call": [
    {
     "country": "RU",
     "bandwidth": "77",
     "response_time": "508",
     "provider": "TransparentCalls",
     "connection_stability": 0.7,
     "ttfb": "25",
     "voice_purity": "22",
     "median_of_calls_time": "28"
    },
    {
     "country": "US",
     "bandwidth": "10",
     "response_time": "575",
     "provider": "E-Voice",
     "connection_stability": 0.8,
     "ttfb": "317",
     "voice_purity": "87",
     "median_of_calls_time": "4"
    },
    {
     "country": "GB",
     "bandwidth": "47",
     "response_time": "347",
     "provider": "TransparentCalls",
     "connection_stability": 0.98,
     "ttfb": "369",
     "voice_purity": "32",
     "median_of_calls_time": "51"
    },
    {
     "country": "FR",
     "bandwidth": "74",
     "response_time": "1462",
     "provider": "TransparentCalls",
     "connection_stability": 0.64,
     "ttfb": "464",
     "voice_purity": "22",
     "median_of_calls_time": "52"
    },
    {
     "country": "BL",
     "bandwidth": "38",
     "response_time": "960",
     "provider": "E-Voice",
     "connection_stability": 0.74,
     "ttfb": "549",
     "voice_purity": "22",
     "median_of_calls_time": "47"
    },
    {
     "country": "AT",
     "bandwidth": "51",
     "response_time": "1137",
     "provider": "TransparentCalls",
     "connection_stability": 0.77,
     "ttfb": "377",
     "voice_purity": "33",
     "median_of_calls_time": "50"
    },
    {
     "country": "DK",
     "bandwidth": "99",
     "response_time": "371",
     "provider": "JustPhone",
     "connection_stability": 0.79,
     "ttfb": "356",
     "voice_purity": "36",
     "median_of_calls_time": "43"
    },
    {
     "country": "CA",
     "bandwidth": "56",
     "response_time": "975",
     "provider": "JustPhone",
     "connection_stability": 0.92,
     "ttfb": "808",
     "voice_purity": "3",
     "median_of_calls_time": "9"
    },
    {
     "country": "ES",
     "bandwidth": "89",
     "response_time": "1986",
     "provider": "E-Voice",
     "connection_stability": 0.99,
     "ttfb": "746",
     "voice_purity": "32",
     "median_of_calls_time": "33"
    },
    {
     "country": "CH",
     "bandwidth": "82",
     "response_time": "1275",
     "provider": "JustPhone",
     "connection_stability": 0.97,
     "ttfb": "722",
     "voice_purity": "10",
     "median_of_calls_time": "8"
    },
    {
     "country": "TR",
     "bandwidth": "15",
     "response_time": "535",
     "provider": "TransparentCalls",
     "connection_stability": 0.76,
     "ttfb": "775",
     "voice_purity": "41",
     "median_of_calls_time": "22"
    },
    {
     "country": "PE",
     "bandwidth": "22",
     "response_time": "37",
     "provider": "JustPhone",
     "connection_stability": 0.8,
     "ttfb": "314",
     "voice_purity": "20",
     "median_of_calls_time": "8"
    },
    {
     "country": "NZ",
     "bandwidth": "81",
     "response_time": "271",
     "provider": "JustPhone",
     "connection_stability": 0.95,
     "ttfb": "303",
     "voice_purity": "12",
     "median_of_calls_time": "26"
    },
    {
     "country": "MC",
     "bandwidth": "41",
     "response_time": "1220",
     "provider": "E-Voice",
     "connection_stability": 0.67,
     "ttfb": "940",
     "voice_purity": "75",
     "median_of_calls_time": "36"
    }
   ],
   "email": {
    "AT": [
     [
      {
       "country": "AT",
       "provider": "Orange",
       "delivery_time": 81
      },
      {
       "country": "AT",
       "provider": "AOL",
       "delivery_time": 141
      },
      {
       "country": "AT",
       "provider": "GMX",
       "delivery_time": 218
      }
     ],
     [
      {
       "country": "AT",
       "provider": "Yahoo",
       "delivery_time": 538
      },
      {
       "country": "AT",
       "provider": "MSN",
       "delivery_time": 528
      },
      {
       "country": "AT",
       "provider": "Live",
       "delivery_time": 475
      }
     ]
    ],
    "BG": [
     [
      {
       "country": "BG",
       "provider": "Hotmail",
       "delivery_time": 10
      },
      {
       "country": "BG",
       "provider": "MSN",
       "delivery_time": 32
      },
      {
       "country": "BG",
       "provider": "Comcast",
       "delivery_time": 52
      }
     ],
     [
      {
       "country": "BG",
       "provider": "GMX",
       "delivery_time": 492
      },
      {
       "country": "BG",
       "provider": "Orange",
       "delivery_time": 481
      },
      {
       "country": "BG",
       "provider": "RediffMail",
       "delivery_time": 412
      }
     ]
    ],
    "BL": [
     [
      {
       "country": "BL",
       "provider": "Comcast",
       "delivery_time": 77
      },
      {
       "country": "BL",
       "provider": "Yahoo",
       "delivery_time": 80
      },
      {
       "country": "BL",
       "provider": "Mail.ru",
       "delivery_time": 116
      }
     ],
     [
      {
       "country": "BL",
       "provider": "GMX",
       "delivery_time": 573
      },
      {
       "country": "BL",
       "provider": "Yandex",
       "delivery_time": 491
      },
      {
       "country": "BL",
       "provider": "Gmail",
       "delivery_time": 457
      }
     ]
    ],
    "CA": [
     [
      {
       "country": "CA",
       "provider": "Comcast",
       "delivery_time": 134
      },
      {
       "country": "CA",
       "provider": "AOL",
       "delivery_time": 150
      },
      {
       "country": "CA",
       "provider": "RediffMail",
       "delivery_time": 218
      }
     ],
     [
      {
       "country": "CA",
       "provider": "MSN",
       "delivery_time": 564
      },
      {
       "country": "CA",
       "provider": "Orange",
       "delivery_time": 493
      },
      {
       "country": "CA",
       "provider": "Yahoo",
       "delivery_time": 458
      }
     ]
    ],
    "CH": [
     [
      {
       "country": "CH",
       "provider": "RediffMail",
       "delivery_time": 6
      },
      {
       "country": "CH",
       "provider": "GMX",
       "delivery_time": 160
      },
      {
       "country": "CH",
       "provider": "AOL",
       "delivery_time": 277
      }
     ],
     [
      {
       "country": "CH",
       "provider": "Gmail",
       "delivery_time": 582
      },
      {
       "country": "CH",
       "provider": "Yandex",
       "delivery_time": 581
      },
      {
       "country": "CH",
       "provider": "Live",
       "delivery_time": 572
      }
     ]
    ],
    "DK": [
     [
      {
       "country": "DK",
       "provider": "RediffMail",
       "delivery_time": 4
      },
      {
       "country": "DK",
       "provider": "Orange",
       "delivery_time": 10
      },
      {
       "country": "DK",
       "provider": "Yahoo",
       "delivery_time": 24
      }
     ],
     [
      {
       "country": "DK",
       "provider": "Hotmail",
       "delivery_time": 587
      },
      {
       "country": "DK",
       "provider": "GMX",
       "delivery_time": 261
      },
      {
       "country": "DK",
       "provider": "Yandex",
       "delivery_time": 251
      }
     ]
    ],
    "ES": [
     [
      {
       "country": "ES",
       "provider": "Yandex",
       "delivery_time": 32
      },
      {
       "country": "ES",
       "provider": "MSN",
       "delivery_time": 48
      },
      {
       "country": "ES",
       "provider": "GMX",
       "delivery_time": 71
      }
     ],
     [
      {
       "country": "ES",
       "provider": "Gmail",
       "delivery_time": 555
      },
      {
       "country": "ES",
       "provider": "AOL",
       "delivery_time": 521
      },
      {
       "country": "ES",
       "provider": "Orange",
       "delivery_time": 420
      }
     ]
    ],
    "FR": [
     [
      {
       "country": "FR",
       "provider": "Orange",
       "delivery_time": 3
      },
      {
       "country": "FR",
       "provider": "RediffMail",
       "delivery_time": 32
      },
      {
       "country": "FR",
       "provider": "Yahoo",
       "delivery_time": 32
      }
     ],
     [
      {
       "country": "FR",
       "provider": "Gmail",
       "delivery_time": 581
      },
      {
       "country": "FR",
       "provider": "Yandex",
       "delivery_time": 514
      },
      {
       "country": "FR",
       "provider": "MSN",
       "delivery_time": 442
      }
     ]
    ],
    "GB": [
     [
      {
       "country": "GB",
       "provider": "Live",
       "delivery_time": 37
      },
      {
       "country": "GB",
       "provider": "MSN",
       "delivery_time": 41
      },
      {
       "country": "GB",
       "provider": "AOL",
       "delivery_time": 75
      }
     ],
     [
      {
       "country": "GB",
       "provider": "Gmail",
       "delivery_time": 581
      },
      {
       "country": "GB",
       "provider": "Comcast",
       "delivery_time": 517
      },
      {
       "country": "GB",
       "provider": "Yandex",
       "delivery_time": 489
      }
     ]
    ],
    "MC": [
     [
      {
       "country": "MC",
       "provider": "Hotmail",
       "delivery_time": 36
      },
      {
       "country": "MC",
       "provider": "Yandex",
       "delivery_time": 66
      },
      {
       "country": "MC",
       "provider": "Orange",
       "delivery_time": 142
      }
     ],
     [
      {
       "country": "MC",
       "provider": "MSN",
       "delivery_time": 531
      },
      {
       "country": "MC",
       "provider": "Live",
       "delivery_time": 437
      },
      {
       "country": "MC",
       "provider": "Comcast",
       "delivery_time": 429
      }
     ]
    ],
    "NZ": [
     [
      {
       "country": "NZ",
       "provider": "Mail.ru",
       "delivery_time": 22
      },
      {
       "country": "NZ",
       "provider": "Orange",
       "delivery_time": 131
      },
      {
       "country": "NZ",
       "provider": "Hotmail",
       "delivery_time": 164
      }
     ],
     [
      {
       "country": "NZ",
       "provider": "Comcast",
       "delivery_time": 538
      },
      {
       "country": "NZ",
       "provider": "Live",
       "delivery_time": 520
      },
      {
       "country": "NZ",
       "provider": "Yahoo",
       "delivery_time": 485
      }
     ]
    ],
    "PE": [
     [
      {
       "country": "PE",
       "provider": "GMX",
       "delivery_time": 48
      },
      {
       "country": "PE",
       "provider": "MSN",
       "delivery_time": 84
      },
      {
       "country": "PE",
       "provider": "Mail.ru",
       "delivery_time": 171
      }
     ],
     [
      {
       "country": "PE",
       "provider": "Yandex",
       "delivery_time": 580
      },
      {
       "country": "PE",
       "provider": "RediffMail",
       "delivery_time": 513
      },
      {
       "country": "PE",
       "provider": "Live",
       "delivery_time": 489
      }
     ]
    ],
    "RU": [
     [
      {
       "country": "RU",
       "provider": "Live",
       "delivery_time": 51
      },
      {
       "country": "RU",
       "provider": "Gmail",
       "delivery_time": 124
      },
      {
       "country": "RU",
       "provider": "AOL",
       "delivery_time": 131
      }
     ],
     [
      {
       "country": "RU",
       "provider": "Orange",
       "delivery_time": 585
      },
      {
       "country": "RU",
       "provider": "Comcast",
       "delivery_time": 484
      },
      {
       "country": "RU",
       "provider": "Mail.ru",
       "delivery_time": 389
      }
     ]
    ],
    "TR": [
     [
      {
       "country": "TR",
       "provider": "MSN",
       "delivery_time": 17
      },
      {
       "country": "TR",
       "provider": "Live",
       "delivery_time": 149
      },
      {
       "country": "TR",
       "provider": "Comcast",
       "delivery_time": 179
      }
     ],
     [
      {
       "country": "TR",
       "provider": "Yahoo",
       "delivery_time": 569
      },
      {
       "country": "TR",
       "provider": "Gmail",
       "delivery_time": 516
      },
      {
       "country": "TR",
       "provider": "AOL",
       "delivery_time": 511
      }
     ]
    ],
    "US": [
     [
      {
       "country": "US",
       "provider": "Orange",
       "delivery_time": 14
      },
      {
       "country": "US",
       "provider": "MSN",
       "delivery_time": 37
      },
      {
       "country": "US",
       "provider": "Live",
       "delivery_time": 66
      }
     ],
     [
      {
       "country": "US",
       "provider": "Hotmail",
       "delivery_time": 572
      },
      {
       "country": "US",
       "provider": "GMX",
       "delivery_time": 486
      },
      {
       "country": "US",
       "provider": "Gmail",
       "delivery_time": 462
      }
     ]
    ]
   },
   "billing": {
    "create_customer": true,
    "purchase": true,
    "payout": false,
    "recurring": true,
    "fraud_control": true,
    "checkout_page": true
   },
   "support": [
    3,
    87
   ],
   "incident": [
    {
     "topic": "SMS delivery in EU",
     "status": "closed"
    },
    {
     "topic": "MMS connection stability",
     "status": "closed"
    },
    {
     "topic": "Voice call connection purity",
     "status": "closed"
    },
    {
     "topic": "Checkout page is down",
     "status": "closed"
    },
    {
     "topic": "Support overload",
     "status": "closed"
    },
    {
     "topic": "Buy phone number not working in US",
     "status": "closed"
    },
    {
     "topic": "API Slow latency",
     "status": "closed"
    }
   ]
  },
  "error": ""
 }

Response code: 200 (OK); Time: 47ms (47 ms); Content length: 21220 bytes (21,22 kB)

```
Факт запроса записывается в лог файл, в запись попадает IP и UserAgent клиента
```json
{"level":"info","msg":"userIP: 127.0.0.1:61415, User-Agent: Apache-HttpClient/4.5.13 (Java/17.0.5)","time":"2023-02-01T20:52:52+03:00"}
```



