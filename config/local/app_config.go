package config

// AppConfig ...
var AppConfig = `{
  "env": "local",
  "appID": "bff-web",
  "version": "1.0.0",
  "sitePrefix": "",
  "port": ":1301",
  "log": {
    "apllog": {
      "level": 6,
      "outputType": "os.Stdout",
      "outputFilePath": "log/app.log",
      "outputFileMaxSize": 1,
      "outputFileMaxBackups": 3,
      "outputFileMaxAge": 30,
      "outputFileLocalTime": true,
      "outputFileCompress": false
    },
    "errlog": {
      "level": 6,
      "outputType": "os.Stderr",
      "outputFilePath": "log/err.log",
      "outputFileMaxSize": 1,
      "outputFileMaxBackups": 3,
      "outputFileMaxAge": 30,
      "outputFileLocalTime": true,
      "outputFileCompress": false
    },
    "accesslog": {
      "level": 4,
      "outputType": "os.Stdout",
      "outputFilePath": "log/access.log",
      "outputFileMaxSize": 1,
      "outputFileMaxBackups": 3,
      "outputFileMaxAge": 30,
      "outputFileLocalTime": true,
      "outputFileCompress": false
    },
    "externaljournallog": {
      "level": 4,
      "outputType": "os.Stdout",
      "outputFilePath": "log/externaljournal.log",
      "outputFileMaxSize": 1,
      "outputFileMaxBackups": 3,
      "outputFileMaxAge": 30,
      "outputFileLocalTime": true,
      "outputFileCompress": false
    },
    "internaljournallog": {
      "level": 4,
      "outputType": "os.Stdout",
      "outputFilePath": "log/internaljournal.log",
      "outputFileMaxSize": 1,
      "outputFileMaxBackups": 3,
      "outputFileMaxAge": 30,
      "outputFileLocalTime": true,
      "outputFileCompress": false
    },
    "notificationlog": {
      "level": 6,
      "outputType": "os.Stdout",
      "outputFilePath": "log/notification.log",
      "outputFileMaxSize": 1,
      "outputFileMaxBackups": 3,
      "outputFileMaxAge": 30,
      "outputFileLocalTime": true,
      "outputFileCompress": false
    }
  },
  "db": [
    {
      "dbHost": "",
      "dbName": "",
      "dbUserName": "",
      "dbPassword": "",
      "dbDnsParam": "parseTime=true&loc=Asia%2FTokyo&rejectReadOnly=true",
      "dbDns": "",
      "maxIdleConns": 20,
      "maxOpenConns": 20,
      "connMaxLifetime": 20,
      "logLevel": 4
    },
    {
      "dbHost": "",
      "dbName": "",
      "dbUserName": "",
      "dbPassword": "",
      "dbDnsParam": "parseTime=true&loc=Asia%2FTokyo",
      "dbDns": "",
      "maxIdleConns": 20,
      "maxOpenConns": 20,
      "connMaxLifetime": 20,
      "logLevel": 4
    }
  ],
  "session": [
    {
      "dbHost": "",
      "dbName": "",
      "dbUserName": "",
      "dbPassword": "",
      "dbDnsParam": "parseTime=true&loc=Asia%2FTokyo&rejectReadOnly=true",
      "dbDns": "",
      "maxIdleConns": 20,
      "maxOpenConns": 20,
      "connMaxLifetime": 20,
      "logMode": true,
      "sessionID": "NWSESSIONID",
      "sessionResetKEY": "RESETKEY",
      "table": "nudge-dynamodb-bff-web-manage-sessions",
      "sessionCookiePath": "/",
      "expiresInSeconds": 900,
      "secret": "2716da4c86fa4d76b43743933fd4624805aa094fe3dd4082b365b46fac93714b",
      "options": {
        "path": "/",
        "domain": "",
        "maxAge": 900,
        "secure": false,
        "httpOnly": true,
        "sameSiteMode": "None"
      }
    }
  ],
  "secretsManager": [
    {
      "secretId": "2057e2ace77042deb77d52c1188c9f34"
    }
  ],
  "internalService": {
    "svcRepayment": {
      "basePath": "http:\/\/localhost:1390"
    },
    "svcPayment": {
      "basePath": "http:\/\/localhost:1390"
    }
  },
  "externalService": {
    "xxxXxxxxxx": {
      "basePath": ""
    }
  },
  "app": {
    "cors": {
      "corsSetting": true,
      "headerAccessControlAllowOrigin": [
        "http://localhost:8080",
        "*"
      ],
      "headerAccessControlAllowCredentials": true
    },
    "request": {
      "contentLengthLimit": "100M"
    },
    "sevenBank": {
      "clientSecret": "58b817533ee94204b31a1624a6551392"
    },
    "commonKey": {
      "commonKeyCryptIV": "",
      "commonKeyCryptKey": "210dc29797f84f969a2a93acb542851d5d44c02de68a4a81840e79464d57d507"
    },
    "operator": {
      "pwCryptIV": "70f039795bb82b55283b45ab7d469666",
      "pwCryptKey": "40e24576dc02541d99b40e306414652224aa8dbb84b3fecf5a16cdca1c51fda5",
      "maximumNumberLoginHistory": 20,
      "maximumNumberLoginPasswordChangeHistory": 20,
      "maximumRecordsPerPage01": 100,
      "maximumNumberLoginFailNumForRevoke": 3
    },
    "repaymentErrMessage": {
      "cache": {
        "interval": 60
      }
    }
  },
  "dynamoDB" : {
    "session" : "nudge-dynamodb-bff-web-sessions",
    "reset-password" : "nudge-dynamodb-bff-web-reset-passwords"
  },
  "endpoint" : {
    "svc-operator" : " https://web-api.dev.nudge-platform.com",
    "svc-user" : "http://svc-user.nudge.local",
    "svc-card" : "http://svc-card.nudge.local",
    "svc-payment" : "http://svc-payment.nudge.local",
    "svc-club" : "http://svc-club.nudge.local",
    "svc-notification" : "http://svc-notification.nudge.local",
    "svc-application" : "http://svc-application.nudge.local",
    "frontend" : "http://operator.dev.nudge-platform.com"
  }
}`
