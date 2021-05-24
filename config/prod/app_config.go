package config

// AppConfig ...
var AppConfig = `{
  "env": "prod",
  "appID": "",
  "version": "1.0.0",
  "sitePrefix": "",
  "port": ":80",
  "log": {
    "apllog": {
      "level": 2,
      "outputType": "os.Stdout",
      "outputFilePath": "log/app.log",
      "outputFileMaxSize": 1,
      "outputFileMaxBackups": 3,
      "outputFileMaxAge": 30,
      "outputFileLocalTime": true,
      "outputFileCompress": false
    },
    "errlog": {
      "level": 3,
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
      "logLevel": 1
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
      "logLevel": 1
    }
  ],
  "secretsManager": [
    {
      "secretId": "dummy"
    }
  ],
  "internalService": {
    "svcRepayment": {
      "basePath": ""
    },
    "svcPayment": {
      "basePath": ""
    }
  },
  "externalService": {
    "xxxXxxxxxx": {
      "basePath": ""
    }
  },
  "app": {
    "request": {
      "contentLengthLimit": "100M"
    },
    "sevenBank": {
      "clientSecret": ""
    },
    "commonKey": {
      "commonKeyCryptIV": "",
      "commonKeyCryptKey": ""
    },
    "repaymentErrMessage": {
      "cache": {
        "interval": 60
      }
    }
  }
}`
