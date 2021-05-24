package config

// MessagesConfig ...
var MessagesConfig = `{
  "010000500900101": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900101",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000000900102": {
    "messageLevel": "status-judgment",
    "statusCode": 0,
    "messageType": "InternalServerError",
    "messageCode": "250000000900102",
    "internalMessage": "共通エラーハンドラー個別ハンドル対象外メッセージ。共通エラーハンドラーで受け取ったHTTPステータスを応答する。",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000400900103": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400900103",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（400）",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "010000401900104": {
    "messageLevel": "info",
    "statusCode": 401,
    "messageType": "Unauthorized",
    "messageCode": "250000401900104",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（401）",
    "message": {
      "ja": "Unauthorized",
      "en": "Unauthorized"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is denied for some reason.",
      "en": "Accessing the page or resource you were trying to reach is denied for some reason."
    },
    "action": "NoAutoRetry"
  },
  "010000403900105": {
    "messageLevel": "info",
    "statusCode": 403,
    "messageType": "Forbidden",
    "messageCode": "250000403900105",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（403）",
    "message": {
      "ja": "Forbidden",
      "en": "Forbidden"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is forbidden for some reason.",
      "en": "Accessing the page or resource you were trying to reach is forbidden for some reason."
    },
    "action": "NoAutoRetry"
  },
  "010000404900106": {
    "messageLevel": "info",
    "statusCode": 404,
    "messageType": "NotFound",
    "messageCode": "250000404900106",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（404）",
    "message": {
      "ja": "Not Found",
      "en": "Not Found"
    },
    "messageDetail": {
      "ja": "The URI requested is invalid or the resource requested does not exists.",
      "en": "The URI requested is invalid or the resource requested does not exists."
    },
    "action": "NoAutoRetry"
  },
  "010000500900107": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900107",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（500）",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000503900108": {
    "messageLevel": "info",
    "statusCode": 503,
    "messageType": "ServiceUnavailable",
    "messageCode": "250000503900108",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（503）",
    "message": {
      "ja": "Service Unavailable",
      "en": "Service Unavailable"
    },
    "messageDetail": {
      "ja": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate.",
      "en": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate."
    },
    "action": "NoAutoRetry"
  },
  "010000405900109": {
    "messageLevel": "info",
    "statusCode": 405,
    "messageType": "MethodNotAllowed",
    "messageCode": "250000405900109",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（405）",
    "message": {
      "ja": "Method Not Allowed",
      "en": "Method Not Allowed"
    },
    "messageDetail": {
      "ja": "The method received in the request-line is known by the origin server but not supported by the target resource.",
      "en": "The method received in the request-line is known by the origin server but not supported by the target resource."
    },
    "action": "NoAutoRetry"
  },
  "010000429900110": {
    "messageLevel": "warn",
    "statusCode": 429,
    "messageType": "TooManyRequests",
    "messageCode": "250000429900110",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（429）",
    "message": {
      "ja": "Too Many Requests",
      "en": "Too Many Requests"
    },
    "messageDetail": {
      "ja": "The user has sent too many requests in a given amount of time. (\"rate limiting\")",
      "en": "The user has sent too many requests in a given amount of time. (\"rate limiting\")"
    },
    "action": "NoAutoRetry"
  },
  "010000413900111": {
    "messageLevel": "info",
    "statusCode": 413,
    "messageType": "PayloadTooLarge",
    "messageCode": "250000413900111",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（413）",
    "message": {
      "ja": "Payload Too Large",
      "en": "Payload Too Large"
    },
    "messageDetail": {
      "ja": "The request you sent is too large.",
      "en": "The request you sent is too large."
    },
    "action": "NoAutoRetry"
  },
  "010000500900112": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900112",
    "internalMessage": "環境設定失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900201": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900201",
    "internalMessage": "トランザクション開始失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000500900202": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900202",
    "internalMessage": "トランザクションコミット失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000500900203": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900203",
    "internalMessage": "トランザクションロールバック失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000500900204": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900204",
    "internalMessage": "DB検索失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000500900205": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900205",
    "internalMessage": "DB登録失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000500900206": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900206",
    "internalMessage": "DB更新失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000500900207": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900207",
    "internalMessage": "リードオンリーレプリカDB登録",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900208": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900208",
    "internalMessage": "リードオンリーレプリカDB更新",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900210": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900210",
    "internalMessage": "データ変換処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900212": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900212",
    "internalMessage": "DB接続失敗（プライマリ）",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000500900213": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900213",
    "internalMessage": "DB接続失敗（レプリカ）",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000500900215": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900215",
    "internalMessage": "DB削除失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000500900216": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900216",
    "internalMessage": "リードオンリーレプリカDB削除",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900217": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900217",
    "internalMessage": "デッドロック",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000500900218": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900218",
    "internalMessage": "データ処理中",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000409900219": {
    "messageLevel": "info",
    "statusCode": 409,
    "messageType": "Conflict",
    "messageCode": "250000409900219",
    "internalMessage": "ロックウェイトタイムアウト",
    "message": {
      "ja": "Conflict",
      "en": "Conflict"
    },
    "messageDetail": {
      "ja": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state.",
      "en": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state."
    },
    "action": "Retryable"
  },
  "010000500900220": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900220",
    "internalMessage": "一意制約違反",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900221": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900221",
    "internalMessage": "一意制約違反",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "010000500900301": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900301",
    "internalMessage": "暗号化／復号失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900302": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900302",
    "internalMessage": "日付処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900303": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900303",
    "internalMessage": "数値処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900304": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900304",
    "internalMessage": "文字列処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900305": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900305",
    "internalMessage": "データ変換処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900401": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900401",
    "internalMessage": "設定値取得失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900402": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900402",
    "internalMessage": "SecretsManager処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900403": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900403",
    "internalMessage": "データ変換処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000503900501": {
    "messageLevel": "info",
    "statusCode": 503,
    "messageType": "ServiceUnavailable",
    "messageCode": "250000503900501",
    "internalMessage": "RDS作業サービス閉塞",
    "message": {
      "ja": "Service Unavailable",
      "en": "Service Unavailable"
    },
    "messageDetail": {
      "ja": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate.",
      "en": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate."
    },
    "action": "NoAutoRetry"
  },
  "010000500900701": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900701",
    "internalMessage": "HTTPクライアント処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900702": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900702",
    "internalMessage": "HTTPクライアントAPI呼び出し処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000500900703": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900703",
    "internalMessage": "データ変換処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "010000422900704": {
    "messageLevel": "info",
    "statusCode": 422,
    "messageType": "UnprocessableEntity",
    "messageCode": "250000422900704",
    "internalMessage": "フォーマット不正",
    "message": {
      "ja": "Unprocessable Entity",
      "en": "Unprocessable Entity"
    },
    "messageDetail": {
      "ja": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation.",
      "en": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation."
    },
    "action": "NoAutoRetry"
  },
  "010000500900705": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "010000500900705",
    "internalMessage": "エラーマッピング失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900101": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900101",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000000900102": {
    "messageLevel": "status-judgment",
    "statusCode": 0,
    "messageType": "InternalServerError",
    "messageCode": "250000000900102",
    "internalMessage": "共通エラーハンドラー個別ハンドル対象外メッセージ。共通エラーハンドラーで受け取ったHTTPステータスを応答する。",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000400900103": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400900103",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（400）",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000401900104": {
    "messageLevel": "info",
    "statusCode": 401,
    "messageType": "Unauthorized",
    "messageCode": "250000401900104",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（401）",
    "message": {
      "ja": "Unauthorized",
      "en": "Unauthorized"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is denied for some reason.",
      "en": "Accessing the page or resource you were trying to reach is denied for some reason."
    },
    "action": "NoAutoRetry"
  },
  "250000403900105": {
    "messageLevel": "info",
    "statusCode": 403,
    "messageType": "Forbidden",
    "messageCode": "250000403900105",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（403）",
    "message": {
      "ja": "Forbidden",
      "en": "Forbidden"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is forbidden for some reason.",
      "en": "Accessing the page or resource you were trying to reach is forbidden for some reason."
    },
    "action": "NoAutoRetry"
  },
  "250000404900106": {
    "messageLevel": "info",
    "statusCode": 404,
    "messageType": "NotFound",
    "messageCode": "250000404900106",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（404）",
    "message": {
      "ja": "Not Found",
      "en": "Not Found"
    },
    "messageDetail": {
      "ja": "The URI requested is invalid or the resource requested does not exists.",
      "en": "The URI requested is invalid or the resource requested does not exists."
    },
    "action": "NoAutoRetry"
  },
  "250000500900107": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900107",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（500）",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000503900108": {
    "messageLevel": "info",
    "statusCode": 503,
    "messageType": "ServiceUnavailable",
    "messageCode": "250000503900108",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（503）",
    "message": {
      "ja": "Service Unavailable",
      "en": "Service Unavailable"
    },
    "messageDetail": {
      "ja": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate.",
      "en": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate."
    },
    "action": "NoAutoRetry"
  },
  "250000405900109": {
    "messageLevel": "info",
    "statusCode": 405,
    "messageType": "MethodNotAllowed",
    "messageCode": "250000405900109",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（405）",
    "message": {
      "ja": "Method Not Allowed",
      "en": "Method Not Allowed"
    },
    "messageDetail": {
      "ja": "The method received in the request-line is known by the origin server but not supported by the target resource.",
      "en": "The method received in the request-line is known by the origin server but not supported by the target resource."
    },
    "action": "NoAutoRetry"
  },
  "250000429900110": {
    "messageLevel": "warn",
    "statusCode": 429,
    "messageType": "TooManyRequests",
    "messageCode": "250000429900110",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（429）",
    "message": {
      "ja": "Too Many Requests",
      "en": "Too Many Requests"
    },
    "messageDetail": {
      "ja": "The user has sent too many requests in a given amount of time. (\"rate limiting\")",
      "en": "The user has sent too many requests in a given amount of time. (\"rate limiting\")"
    },
    "action": "NoAutoRetry"
  },
  "250000413900111": {
    "messageLevel": "info",
    "statusCode": 413,
    "messageType": "PayloadTooLarge",
    "messageCode": "250000413900111",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（413）",
    "message": {
      "ja": "Payload Too Large",
      "en": "Payload Too Large"
    },
    "messageDetail": {
      "ja": "The request you sent is too large.",
      "en": "The request you sent is too large."
    },
    "action": "NoAutoRetry"
  },
  "250000500900112": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900112",
    "internalMessage": "環境設定失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900201": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900201",
    "internalMessage": "トランザクション開始失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000500900202": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900202",
    "internalMessage": "トランザクションコミット失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000500900203": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900203",
    "internalMessage": "トランザクションロールバック失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000500900204": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900204",
    "internalMessage": "DB検索失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000500900205": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900205",
    "internalMessage": "DB登録失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000500900206": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900206",
    "internalMessage": "DB更新失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000500900207": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900207",
    "internalMessage": "リードオンリーレプリカDB登録",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900208": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900208",
    "internalMessage": "リードオンリーレプリカDB更新",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900210": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900210",
    "internalMessage": "データ変換処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900212": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900212",
    "internalMessage": "DB接続失敗（プライマリ）",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000500900213": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900213",
    "internalMessage": "DB接続失敗（レプリカ）",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000500900215": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900215",
    "internalMessage": "DB削除失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000500900216": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900216",
    "internalMessage": "リードオンリーレプリカDB削除",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900217": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900217",
    "internalMessage": "デッドロック",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000500900218": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900218",
    "internalMessage": "データ処理中",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000409900219": {
    "messageLevel": "info",
    "statusCode": 409,
    "messageType": "Conflict",
    "messageCode": "250000409900219",
    "internalMessage": "ロックウェイトタイムアウト",
    "message": {
      "ja": "Conflict",
      "en": "Conflict"
    },
    "messageDetail": {
      "ja": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state.",
      "en": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state."
    },
    "action": "Retryable"
  },
  "250000500900220": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900220",
    "internalMessage": "一意制約違反",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900221": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900221",
    "internalMessage": "一意制約違反",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "Retryable"
  },
  "250000500900301": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900301",
    "internalMessage": "暗号化／復号失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900302": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900302",
    "internalMessage": "日付処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900303": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900303",
    "internalMessage": "数値処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900304": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900304",
    "internalMessage": "文字列処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900305": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900305",
    "internalMessage": "データ変換処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900401": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900401",
    "internalMessage": "設定値取得失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900402": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900402",
    "internalMessage": "SecretsManager処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900403": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900403",
    "internalMessage": "データ変換処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000503900501": {
    "messageLevel": "info",
    "statusCode": 503,
    "messageType": "ServiceUnavailable",
    "messageCode": "250000503900501",
    "internalMessage": "RDS作業サービス閉塞",
    "message": {
      "ja": "Service Unavailable",
      "en": "Service Unavailable"
    },
    "messageDetail": {
      "ja": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate.",
      "en": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate."
    },
    "action": "NoAutoRetry"
  },
  "250000400900601": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400900601",
    "internalMessage": "リクエストが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000400900602": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400900602",
    "internalMessage": "Content-Typeが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000400900603": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400900603",
    "internalMessage": "User-Agentが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000400900604": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400900604",
    "internalMessage": "X-Nudge-Interaction-Idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000400900605": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400900605",
    "internalMessage": "X-Nudge-User-Idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000400900606": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400900606",
    "internalMessage": "X-Nudge-Security-Requestが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000401900607": {
    "messageLevel": "info",
    "statusCode": 401,
    "messageType": "Unauthorized",
    "messageCode": "250000401900607",
    "internalMessage": "Authorizationが間違っている",
    "message": {
      "ja": "Unauthorized",
      "en": "Unauthorized"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is denied for some reason.",
      "en": "Accessing the page or resource you were trying to reach is denied for some reason."
    },
    "action": "Reauth"
  },
  "250000400900608": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400900608",
    "internalMessage": "Hostが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000400900609": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400900609",
    "internalMessage": "リクエストボディが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000500900701": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900701",
    "internalMessage": "HTTPクライアント処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900702": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900702",
    "internalMessage": "HTTPクライアントAPI呼び出し処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500900703": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900703",
    "internalMessage": "データ変換処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000422900704": {
    "messageLevel": "info",
    "statusCode": 422,
    "messageType": "UnprocessableEntity",
    "messageCode": "250000422900704",
    "internalMessage": "フォーマット不正",
    "message": {
      "ja": "Unprocessable Entity",
      "en": "Unprocessable Entity"
    },
    "messageDetail": {
      "ja": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation.",
      "en": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation."
    },
    "action": "NoAutoRetry"
  },
  "250000500900705": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500900705",
    "internalMessage": "エラーマッピング失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000400010101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400010101",
    "internalMessage": "リクエストが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000401010102": {
    "messageLevel": "info",
    "statusCode": 401,
    "messageType": "Unauthorized",
    "messageCode": "250000401010102",
    "internalMessage": "署名が間違っている",
    "message": {
      "ja": "Unauthorized",
      "en": "Unauthorized"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is denied for some reason.",
      "en": "Accessing the page or resource you were trying to reach is denied for some reason."
    },
    "action": "NoAutoRetry"
  },
  "250000400010103": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400010103",
    "internalMessage": "JWSヘッダが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000400010104": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400010104",
    "internalMessage": "JWSペイロードが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000400010105": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400010105",
    "internalMessage": "JWSペイロード内セキュア情報が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000400010106": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400010106",
    "internalMessage": "JWSペイロード内暗号化情報が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250000422010107": {
    "messageLevel": "info",
    "statusCode": 422,
    "messageType": "UnprocessableEntity",
    "messageCode": "250000422010107",
    "internalMessage": "ATM取引特定情報なし",
    "message": {
      "ja": "Unprocessable Entity",
      "en": "Unprocessable Entity"
    },
    "messageDetail": {
      "ja": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation.",
      "en": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation."
    },
    "action": "NoAutoRetry"
  },
  "250000422010108": {
    "messageLevel": "info",
    "statusCode": 422,
    "messageType": "UnprocessableEntity",
    "messageCode": "250000422010108",
    "internalMessage": "共通鍵情報なし",
    "message": {
      "ja": "Unprocessable Entity",
      "en": "Unprocessable Entity"
    },
    "messageDetail": {
      "ja": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation.",
      "en": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation."
    },
    "action": "NoAutoRetry"
  },
  "250000500010109": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500010109",
    "internalMessage": "ATMメッセージマス情報なし",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500010110": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500010110",
    "internalMessage": "応答データマスタ情報なし",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500010111": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500010111",
    "internalMessage": "返済エラーマスタ情報なし",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500010112": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500010112",
    "internalMessage": "要求データマスタ情報なし",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000500010113": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250000500010113",
    "internalMessage": "JWSペイロード内セキュア情報が間違っている",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250000400800101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250000400800101",
    "internalMessage": "x-fapi-interaction-idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400900601": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400900601",
    "internalMessage": "リクエストが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400900602": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400900602",
    "internalMessage": "Content-Typeが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400900603": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400900603",
    "internalMessage": "User-Agentが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400900604": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400900604",
    "internalMessage": "X-Nudge-Interaction-Idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400900605": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400900605",
    "internalMessage": "X-Nudge-User-Idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400900606": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400900606",
    "internalMessage": "X-Nudge-Security-Requestが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504401900607": {
    "messageLevel": "info",
    "statusCode": 401,
    "messageType": "Unauthorized",
    "messageCode": "250504401900607",
    "internalMessage": "Authorizationが間違っている",
    "message": {
      "ja": "Unauthorized",
      "en": "Unauthorized"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is denied for some reason.",
      "en": "Accessing the page or resource you were trying to reach is denied for some reason."
    },
    "action": "Reauth"
  },
  "250504400900608": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400900608",
    "internalMessage": "Hostが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400900609": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400900609",
    "internalMessage": "リクエストボディが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010101",
    "internalMessage": "リクエストボディ - JWS - JWSヘッダ - JWTのコンテンツのタイプ が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010102": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010102",
    "internalMessage": "リクエストボディ - JWS - JWSヘッダ - アルゴリズム が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010103": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010103",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - アクセストークン が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010104": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010104",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - イシュア が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010105": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010105",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - サブジェクト が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010106": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010106",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - オーディエンス が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010107": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010107",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - JWT識別子 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010108": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010108",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - JWT発行時刻 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010109": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010109",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - JWT有効期限 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010110": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010110",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 電文種別 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010111": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010111",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 取引指定区分 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010112": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010112",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 取引金額 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010113": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010113",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 処理通番 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010114": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010114",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 電文枝番 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010115": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010115",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - トランザクション日時 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010116": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010116",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 要求機関運用センタコード が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010117": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010117",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - ＡＴＭ番号 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010118": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010118",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 宛先機関識別コード が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010119": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010119",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 受信機関識別コード が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010120": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010120",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 受信機関サブセンタコード が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010121": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010121",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - セキュア情報(バイナリ) が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010122": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010122",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 暗号鍵情報(バイナリ) が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010123": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010123",
    "internalMessage": "リクエストボディ - JWS - JWS署名 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010124": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010124",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - セキュア情報(バイナリ) - QRコード情報 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010125": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010125",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 暗号鍵情報(バイナリ) - IntTranKeyタイプ が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010126": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010126",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 暗号鍵情報(バイナリ) - IntTranKeyバージョン が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400010127": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400010127",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 暗号鍵情報(バイナリ) - IntTranKeyレングス が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504400800101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250504400800101",
    "internalMessage": "x-fapi-interaction-idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250504404010201": {
    "messageLevel": "info",
    "statusCode": 404,
    "messageType": "NotFound",
    "messageCode": "250504404010201",
    "internalMessage": "QRコード情報なし（取引予約なし）",
    "message": {
      "ja": "Not Found",
      "en": "Not Found"
    },
    "messageDetail": {
      "ja": "The URI requested is invalid or the resource requested does not exists.",
      "en": "The URI requested is invalid or the resource requested does not exists."
    },
    "action": "NoAutoRetry"
  },
  "250505400900601": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400900601",
    "internalMessage": "リクエストが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400900602": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400900602",
    "internalMessage": "Content-Typeが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400900603": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400900603",
    "internalMessage": "User-Agentが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400900604": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400900604",
    "internalMessage": "X-Nudge-Interaction-Idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400900605": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400900605",
    "internalMessage": "X-Nudge-User-Idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400900606": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400900606",
    "internalMessage": "X-Nudge-Security-Requestが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505401900607": {
    "messageLevel": "info",
    "statusCode": 401,
    "messageType": "Unauthorized",
    "messageCode": "250505401900607",
    "internalMessage": "Authorizationが間違っている",
    "message": {
      "ja": "Unauthorized",
      "en": "Unauthorized"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is denied for some reason.",
      "en": "Accessing the page or resource you were trying to reach is denied for some reason."
    },
    "action": "Reauth"
  },
  "250505400900608": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400900608",
    "internalMessage": "Hostが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400900609": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400900609",
    "internalMessage": "リクエストボディが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010101",
    "internalMessage": "リクエストボディ - JWS - JWSヘッダ - JWTのコンテンツのタイプ が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010102": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010102",
    "internalMessage": "リクエストボディ - JWS - JWSヘッダ - アルゴリズム が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010103": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010103",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - アクセストークン が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010104": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010104",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - イシュア が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010105": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010105",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - サブジェクト が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010106": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010106",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - オーディエンス が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010107": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010107",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - JWT識別子 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010108": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010108",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - JWT発行時刻 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010109": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010109",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - JWT有効期限 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010110": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010110",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 電文種別 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010111": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010111",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 取引指定区分 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010112": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010112",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 取引金額 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010113": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010113",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 処理通番 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010114": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010114",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 電文枝番 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010115": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010115",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - トランザクション日時 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010116": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010116",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 要求機関運用センタコード が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010117": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010117",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - ＡＴＭ番号 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010118": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010118",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 宛先機関識別コード が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010119": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010119",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 受信機関識別コード が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010120": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010120",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 受信機関サブセンタコード が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010121": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010121",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - セキュア情報(バイナリ) が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010122": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010122",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 暗号鍵情報(バイナリ) が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010123": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010123",
    "internalMessage": "リクエストボディ - JWS - JWS署名 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010124": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010124",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - セキュア情報(バイナリ) - お客様番号 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010125": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010125",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 暗号鍵情報(バイナリ) - IntTranKeyタイプ が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010126": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010126",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 暗号鍵情報(バイナリ) - IntTranKeyバージョン が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010127": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010127",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - 暗号鍵情報(バイナリ) - IntTranKeyレングス が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400010128": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400010128",
    "internalMessage": "リクエストボディ - JWS - JWSペイロード - セキュア情報(バイナリ) - 入力暗証番号／入力認証番号 が間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505400800101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400800101",
    "internalMessage": "x-fapi-interaction-idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505404010201": {
    "messageLevel": "info",
    "statusCode": 404,
    "messageType": "NotFound",
    "messageCode": "250505404010201",
    "internalMessage": "取引情報なし",
    "message": {
      "ja": "Not Found",
      "en": "Not Found"
    },
    "messageDetail": {
      "ja": "The URI requested is invalid or the resource requested does not exists.",
      "en": "The URI requested is invalid or the resource requested does not exists."
    },
    "action": "NoAutoRetry"
  },
  "250505409010202": {
    "messageLevel": "info",
    "statusCode": 409,
    "messageType": "Conflict",
    "messageCode": "250505409010202",
    "internalMessage": "取引情報 - 取引完了済",
    "message": {
      "ja": "Conflict",
      "en": "Conflict"
    },
    "messageDetail": {
      "ja": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state.",
      "en": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state."
    },
    "action": "NoAutoRetry"
  },
  "250505422010203": {
    "messageLevel": "info",
    "statusCode": 422,
    "messageType": "UnprocessableEntity",
    "messageCode": "250505422010203",
    "internalMessage": "返済明細情報なし",
    "message": {
      "ja": "Unprocessable Entity",
      "en": "Unprocessable Entity"
    },
    "messageDetail": {
      "ja": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation.",
      "en": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation."
    },
    "action": "NoAutoRetry"
  },
  "250505400140001": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250505400140001",
    "internalMessage": "400エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250505401140101": {
    "messageLevel": "info",
    "statusCode": 401,
    "messageType": "Unauthorized",
    "messageCode": "250505401140101",
    "internalMessage": "401エラー",
    "message": {
      "ja": "Unauthorized",
      "en": "Unauthorized"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is denied for some reason.",
      "en": "Accessing the page or resource you were trying to reach is denied for some reason."
    },
    "action": "NoAutoRetry"
  },
  "250505403140301": {
    "messageLevel": "info",
    "statusCode": 403,
    "messageType": "Forbidden",
    "messageCode": "250505403140301",
    "internalMessage": "403エラー",
    "message": {
      "ja": "Forbidden",
      "en": "Forbidden"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is forbidden for some reason.",
      "en": "Accessing the page or resource you were trying to reach is forbidden for some reason."
    },
    "action": "NoAutoRetry"
  },
  "250505404140401": {
    "messageLevel": "info",
    "statusCode": 404,
    "messageType": "NotFound",
    "messageCode": "250505404140401",
    "internalMessage": "404エラー",
    "message": {
      "ja": "Not Found",
      "en": "Not Found"
    },
    "messageDetail": {
      "ja": "The URI requested is invalid or the resource requested does not exists.",
      "en": "The URI requested is invalid or the resource requested does not exists."
    },
    "action": "NoAutoRetry"
  },
  "250505405140501": {
    "messageLevel": "info",
    "statusCode": 405,
    "messageType": "MethodNotAllowed",
    "messageCode": "250505405140501",
    "internalMessage": "405エラー",
    "message": {
      "ja": "Method Not Allowed",
      "en": "Method Not Allowed"
    },
    "messageDetail": {
      "ja": "The method received in the request-line is known by the origin server but not supported by the target resource.",
      "en": "The method received in the request-line is known by the origin server but not supported by the target resource."
    },
    "action": "NoAutoRetry"
  },
  "250505409140901": {
    "messageLevel": "info",
    "statusCode": 409,
    "messageType": "Conflict",
    "messageCode": "250505409140901",
    "internalMessage": "409エラー",
    "message": {
      "ja": "Conflict",
      "en": "Conflict"
    },
    "messageDetail": {
      "ja": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state.",
      "en": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state."
    },
    "action": "NoAutoRetry"
  },
  "250505410141001": {
    "messageLevel": "info",
    "statusCode": 410,
    "messageType": "Gone",
    "messageCode": "250505410141001",
    "internalMessage": "409エラー",
    "message": {
      "ja": "Gone",
      "en": "Gone"
    },
    "messageDetail": {
      "ja": "The URI requested is invalid or the resource requested does not exists.",
      "en": "The URI requested is invalid or the resource requested does not exists."
    },
    "action": "NoAutoRetry"
  },
  "250505413141301": {
    "messageLevel": "info",
    "statusCode": 413,
    "messageType": "PayloadTooLarge",
    "messageCode": "250505413141301",
    "internalMessage": "413エラー",
    "message": {
      "ja": "Payload Too Large",
      "en": "Payload Too Large"
    },
    "messageDetail": {
      "ja": "The request you sent is too large.",
      "en": "The request you sent is too large."
    },
    "action": "NoAutoRetry"
  },
  "250505422142201": {
    "messageLevel": "info",
    "statusCode": 422,
    "messageType": "UnprocessableEntity",
    "messageCode": "250505422142201",
    "internalMessage": "422エラー",
    "message": {
      "ja": "Unprocessable Entity",
      "en": "Unprocessable Entity"
    },
    "messageDetail": {
      "ja": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation.",
      "en": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation."
    },
    "action": "NoAutoRetry"
  },
  "250505429142901": {
    "messageLevel": "warn",
    "statusCode": 429,
    "messageType": "TooManyRequests",
    "messageCode": "250505429142901",
    "internalMessage": "429エラー",
    "message": {
      "ja": "Too Many Requests",
      "en": "Too Many Requests"
    },
    "messageDetail": {
      "ja": "The user has sent too many requests in a given amount of time. (\"rate limiting\")",
      "en": "The user has sent too many requests in a given amount of time. (\"rate limiting\")"
    },
    "action": "NoAutoRetry"
  },
  "250505500150001": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250505500150001",
    "internalMessage": "500エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250505503150301": {
    "messageLevel": "info",
    "statusCode": 503,
    "messageType": "ServiceUnavailable",
    "messageCode": "250505503150301",
    "internalMessage": "503エラー",
    "message": {
      "ja": "Service Unavailable",
      "en": "Service Unavailable"
    },
    "messageDetail": {
      "ja": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate.",
      "en": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate."
    },
    "action": "NoAutoRetry"
  },
  "250101400900601": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400900601",
    "internalMessage": "リクエストが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250101400900602": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400900602",
    "internalMessage": "Content-Typeが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250101400900603": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400900603",
    "internalMessage": "User-Agentが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250101400900604": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400900604",
    "internalMessage": "X-Nudge-Interaction-Idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250101400900605": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400900605",
    "internalMessage": "X-Nudge-User-Idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250101400900606": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400900606",
    "internalMessage": "X-Nudge-Security-Requestが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250101401900607": {
    "messageLevel": "info",
    "statusCode": 401,
    "messageType": "Unauthorized",
    "messageCode": "250101401900607",
    "internalMessage": "Authorizationが間違っている",
    "message": {
      "ja": "Unauthorized",
      "en": "Unauthorized"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is denied for some reason.",
      "en": "Accessing the page or resource you were trying to reach is denied for some reason."
    },
    "action": "Reauth"
  },
  "250101400900608": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400900608",
    "internalMessage": "Hostが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250101400900609": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400900609",
    "internalMessage": "リクエストボディが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },


  "250101400010101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400010101",
    "internalMessage": "repaymentIdが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250101400010102": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400010102",
    "internalMessage": "userIdが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250101400010103": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400010103",
    "internalMessage": "repaymentAmountが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250101500010201": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250101500010201",
    "internalMessage": "返済必要額情報なし",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250101500010202": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250101500010202",
    "internalMessage": "ユーザ利息金額合計.利息金額合計 < 0",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250101500010203": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250101500010203",
    "internalMessage": "返済必要額.利息計算対象金額 < 0",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250101500010204": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250101500010204",
    "internalMessage": "月別返済必要額データ不整合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250101409010205": {
    "messageLevel": "info",
    "statusCode": 409,
    "messageType": "Conflict",
    "messageCode": "250101409010205",
    "internalMessage": "返済明細登録一意制約違反",
    "message": {
      "ja": "Conflict",
      "en": "Conflict"
    },
    "messageDetail": {
      "ja": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state.",
      "en": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state."
    },
    "action": "NoAutoRetry"
  },
  "250101422010206": {
    "messageLevel": "info",
    "statusCode": 422,
    "messageType": "UnprocessableEntity",
    "messageCode": "250101422010206",
    "internalMessage": "返済必要額.返済必要額 <= 0",
    "message": {
      "ja": "Unprocessable Entity",
      "en": "Unprocessable Entity"
    },
    "messageDetail": {
      "ja": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation.",
      "en": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation."
    },
    "action": "NoAutoRetry"
  },
  "250101400140001": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "250101400140001",
    "internalMessage": "400エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "250101401140101": {
    "messageLevel": "info",
    "statusCode": 401,
    "messageType": "Unauthorized",
    "messageCode": "250101401140101",
    "internalMessage": "401エラー",
    "message": {
      "ja": "Unauthorized",
      "en": "Unauthorized"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is denied for some reason.",
      "en": "Accessing the page or resource you were trying to reach is denied for some reason."
    },
    "action": "NoAutoRetry"
  },
  "250101403140301": {
    "messageLevel": "info",
    "statusCode": 403,
    "messageType": "Forbidden",
    "messageCode": "250101403140301",
    "internalMessage": "403エラー",
    "message": {
      "ja": "Forbidden",
      "en": "Forbidden"
    },
    "messageDetail": {
      "ja": "Accessing the page or resource you were trying to reach is forbidden for some reason.",
      "en": "Accessing the page or resource you were trying to reach is forbidden for some reason."
    },
    "action": "NoAutoRetry"
  },
  "250101404140401": {
    "messageLevel": "info",
    "statusCode": 404,
    "messageType": "NotFound",
    "messageCode": "250101404140401",
    "internalMessage": "404エラー",
    "message": {
      "ja": "Not Found",
      "en": "Not Found"
    },
    "messageDetail": {
      "ja": "The URI requested is invalid or the resource requested does not exists.",
      "en": "The URI requested is invalid or the resource requested does not exists."
    },
    "action": "NoAutoRetry"
  },
  "250101405140501": {
    "messageLevel": "info",
    "statusCode": 405,
    "messageType": "MethodNotAllowed",
    "messageCode": "250101405140501",
    "internalMessage": "405エラー",
    "message": {
      "ja": "Method Not Allowed",
      "en": "Method Not Allowed"
    },
    "messageDetail": {
      "ja": "The method received in the request-line is known by the origin server but not supported by the target resource.",
      "en": "The method received in the request-line is known by the origin server but not supported by the target resource."
    },
    "action": "NoAutoRetry"
  },
  "250101409140901": {
    "messageLevel": "info",
    "statusCode": 409,
    "messageType": "Conflict",
    "messageCode": "250101409140901",
    "internalMessage": "409エラー",
    "message": {
      "ja": "Conflict",
      "en": "Conflict"
    },
    "messageDetail": {
      "ja": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state.",
      "en": "A conflict exists with the current state of the resource. The requested action cannot be performed on the resource in its current state."
    },
    "action": "NoAutoRetry"
  },
  "250101410141001": {
    "messageLevel": "info",
    "statusCode": 410,
    "messageType": "Gone",
    "messageCode": "250101410141001",
    "internalMessage": "409エラー",
    "message": {
      "ja": "Gone",
      "en": "Gone"
    },
    "messageDetail": {
      "ja": "The URI requested is invalid or the resource requested does not exists.",
      "en": "The URI requested is invalid or the resource requested does not exists."
    },
    "action": "NoAutoRetry"
  },
  "250101413141301": {
    "messageLevel": "info",
    "statusCode": 413,
    "messageType": "PayloadTooLarge",
    "messageCode": "250101413141301",
    "internalMessage": "413エラー",
    "message": {
      "ja": "Payload Too Large",
      "en": "Payload Too Large"
    },
    "messageDetail": {
      "ja": "The request you sent is too large.",
      "en": "The request you sent is too large."
    },
    "action": "NoAutoRetry"
  },
  "250101422142201": {
    "messageLevel": "info",
    "statusCode": 422,
    "messageType": "UnprocessableEntity",
    "messageCode": "250101422142201",
    "internalMessage": "422エラー",
    "message": {
      "ja": "Unprocessable Entity",
      "en": "Unprocessable Entity"
    },
    "messageDetail": {
      "ja": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation.",
      "en": "The designated transaction could not be completed. Sorry to trouble you, but please retry the operation."
    },
    "action": "NoAutoRetry"
  },
  "250101429142901": {
    "messageLevel": "warn",
    "statusCode": 429,
    "messageType": "TooManyRequests",
    "messageCode": "250101429142901",
    "internalMessage": "429エラー",
    "message": {
      "ja": "Too Many Requests",
      "en": "Too Many Requests"
    },
    "messageDetail": {
      "ja": "The user has sent too many requests in a given amount of time. (\"rate limiting\")",
      "en": "The user has sent too many requests in a given amount of time. (\"rate limiting\")"
    },
    "action": "NoAutoRetry"
  },
  "250101500150001": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "250101500150001",
    "internalMessage": "500エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "250101503150301": {
    "messageLevel": "info",
    "statusCode": 503,
    "messageType": "ServiceUnavailable",
    "messageCode": "250101503150301",
    "internalMessage": "503エラー",
    "message": {
      "ja": "Service Unavailable",
      "en": "Service Unavailable"
    },
    "messageDetail": {
      "ja": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate.",
      "en": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate."
    },
    "action": "NoAutoRetry"
  },
  "0100005000107": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000107",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ（500）",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000403": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000403",
    "internalMessage": "数値処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000206": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000206",
    "internalMessage": "DB更新失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000207": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000207",
    "internalMessage": "リードオンリーレプリカDB登録",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000208": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000208",
    "internalMessage": "リードオンリーレプリカDB更新",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000210": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000210",
    "internalMessage": "データ変換処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000212": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000212",
    "internalMessage": "DB接続失敗（プライマリ）",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000213": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000213",
    "internalMessage": "DB接続失敗（レプリカ）",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000215": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000215",
    "internalMessage": "DB削除失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000216": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000216",
    "internalMessage": "リードオンリーレプリカDB削除",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000217": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000217",
    "internalMessage": "デッドロック",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000218": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000218",
    "internalMessage": "S3処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005000219": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000219",
    "internalMessage": "S3処理失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100004220711": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "0100004220711",
    "internalMessage": "ユーザーがロックされています",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "0100005000401": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005000401",
    "internalMessage": "Fail to descrypt Password",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005001001": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005001001",
    "internalMessage": "セッション取得失敗",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01010101",
    "internalMessage": "idの値が不正の場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010102": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01010102",
    "internalMessage": "passwordが間違っている（入力規則）",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010103": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01010103",
    "internalMessage": "sms_codeの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010106": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01010106",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010107": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01010107",
    "internalMessage": "SESSION.auth_statusが2以外の場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010201": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01010201",
    "internalMessage": "mail_addrが設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010202": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01010202",
    "internalMessage": "mail_addrの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  
  "BFWE01010301": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01010301",
    "internalMessage": "keyの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010302": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01010302",
    "internalMessage": "keyの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010303": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01010303",
    "internalMessage": "passwordの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010304": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01010304",
    "internalMessage": "passwordが間違っている（入力規則）",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010401": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01010401",
    "internalMessage": "passwordが設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010402": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01010402",
    "internalMessage": "passwordが間違っている（入力規則）",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010405": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01010405",
    "internalMessage": "共通エラーハンドラーデフォルトメッセージ",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01010406": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01010406",
    "internalMessage": "SESSION.auth_statusが3以外の場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020101",
    "internalMessage": "search_textの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020102": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020102",
    "internalMessage": "sort_keyの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020103": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020103",
    "internalMessage": "sort_orderの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020104": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020104",
    "internalMessage": "page_numの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020105": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020105",
    "internalMessage": "page_sizeの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020106": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020106",
    "internalMessage": "member_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020109": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020109",
    "internalMessage": "member_idに値が設定されているとき、他のリクエスト項目が設定されている場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020110": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020110",
    "internalMessage": "member_idに値が設定されていないとき、sort_keyが設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020111": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020111",
    "internalMessage": "sort_keyが設定されているかつ06以外のとき、search_text,member_id以外の値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020112": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020112",
    "internalMessage": "sort_keyが設定されているかつ06のとき、search_text,sort_orderが設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020113": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020113",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020114": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020114",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020115": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020115",
    "internalMessage": "account_statusが01でない場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030101",
    "internalMessage": "club_owner_nameの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030102": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030102",
    "internalMessage": "club_owner_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030103": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030103",
    "internalMessage": "mail_addressの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030104": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030104",
    "internalMessage": "mail_addressの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030106": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030106",
    "internalMessage": "phone_numberの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030108": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030108",
    "internalMessage": "company_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030110": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030110",
    "internalMessage": "in_charge_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030113": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01030113",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030114": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01030114",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030115": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01030115",
    "internalMessage": "account_statusが01でない場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030201": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030201",
    "internalMessage": "club_owner_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030202": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030202",
    "internalMessage": "club_owner_idの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030203": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030203",
    "internalMessage": "image_urlの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030204": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030204",
    "internalMessage": "mail_addressの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030205": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030205",
    "internalMessage": "phone_numberの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030206": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030206",
    "internalMessage": "company_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030207": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030207",
    "internalMessage": "in_charge_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030208": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030208",
    "internalMessage": "bank_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030209": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030209",
    "internalMessage": "branch_codeの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030210": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030210",
    "internalMessage": "branch_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030211": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030211",
    "internalMessage": "account_typeの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030212": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030212",
    "internalMessage": "account_noの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030213": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030213",
    "internalMessage": "account_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030214": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030214",
    "internalMessage": "account_name_kanaの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030215": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030215",
    "internalMessage": "release_statusの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030216": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030216",
    "internalMessage": "release_start_dateの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030217": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030217",
    "internalMessage": "release_end_dateの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030220": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01030220",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030221": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01030221",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030222": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01030222",
    "internalMessage": "account_statusが01でない場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040201": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040201",
    "internalMessage": "club_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040202": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040202",
    "internalMessage": "club_idの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040203": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040203",
    "internalMessage": "content_typeの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040204": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040204",
    "internalMessage": "club_image_urlの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040205": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040205",
    "internalMessage": "cover_image_urlの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040206": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040206",
    "internalMessage": "club_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040207": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040207",
    "internalMessage": "club_owner_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040208": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040208",
    "internalMessage": "descriptionの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040209": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040209",
    "internalMessage": "theme_colorの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040210": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040210",
    "internalMessage": "category_codeの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040211": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040211",
    "internalMessage": "release_statusの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040212": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040212",
    "internalMessage": "release_start_dateの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040213": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040213",
    "internalMessage": "release_end_dateの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040214": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040214",
    "internalMessage": "max_membersの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040215": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040215",
    "internalMessage": "club_participantsの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040216": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040216",
    "internalMessage": "max_membersip_noの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040219": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040219",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040220": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040220",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040221": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040221",
    "internalMessage": "account_statusが01でない場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040601": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040601",
    "internalMessage": "club_idの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040602": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040602",
    "internalMessage": "club_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040603": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040603",
    "internalMessage": "reward_nameの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040604": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040604",
    "internalMessage": "reward_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040605": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040605",
    "internalMessage": "sectionの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040606": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040606",
    "internalMessage": "sectionの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040607": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040607",
    "internalMessage": "achievement_amountの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040608": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040608",
    "internalMessage": "achievement_amountの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040609": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040609",
    "internalMessage": "createdの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040610": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040610",
    "internalMessage": "createdの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040611": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040611",
    "internalMessage": "shipment_sourceの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040612": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040612",
    "internalMessage": "shipment_sourceの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040613": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040613",
    "internalMessage": "distribution_periodの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040614": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040614",
    "internalMessage": "distribution_statusの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040615": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040615",
    "internalMessage": "total_order_quantityの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040616": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040616",
    "internalMessage": "reward_contentの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040617": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040617",
    "internalMessage": "reward_contentの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040618": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040618",
    "internalMessage": "reward_imageの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040619": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040619",
    "internalMessage": "reward_thumbnailの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040622": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040622",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040623": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040623",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040624": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040624",
    "internalMessage": "account_statusが01でない場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040701": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040701",
    "internalMessage": "reward_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040702": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040702",
    "internalMessage": "reward_idの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040703": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040703",
    "internalMessage": "content_typeの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040704": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040704",
    "internalMessage": "club_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040705": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040705",
    "internalMessage": "reward_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040706": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040706",
    "internalMessage": "distribution_periodの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040707": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040707",
    "internalMessage": "distribution_statusの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040708": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040708",
    "internalMessage": "total_order_quantityの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040709": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040709",
    "internalMessage": "downloadURLの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040710": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040710",
    "internalMessage": "reward_contentの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040711": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040711",
    "internalMessage": "reward_imageの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040712": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040712",
    "internalMessage": "reward_thumbnailの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040715": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040715",
    "internalMessage": "content_typeが設定されていないとき、reward_id, content_type以外の値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040716": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040716",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040717": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040717",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040718": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040718",
    "internalMessage": "account_statusが01でない場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040801": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040801",
    "internalMessage": "club_idの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040802": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01040802",
    "internalMessage": "club_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040805": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040805",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040806": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040806",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01040807": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01040807",
    "internalMessage": "account_statusが01でない場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020301": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020301",
    "internalMessage": "member_idの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020302": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020302",
    "internalMessage": "member_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020303": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020303",
    "internalMessage": "year_monthの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020304": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020304",
    "internalMessage": "year_monthの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020305": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020305",
    "internalMessage": "tran_monthの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020306": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020306",
    "internalMessage": "tran_monthの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020307": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020307",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020308": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020308",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020401": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020401",
    "internalMessage": "user_idの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020402": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020402",
    "internalMessage": "user_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020403": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020403",
    "internalMessage": "repay_monthの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020404": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020404",
    "internalMessage": "repay_monthの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020407": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020407",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020408": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020408",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020501": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020501",
    "internalMessage": "member_idの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020502": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020502",
    "internalMessage": "member_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020505": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020505",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020506": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020506",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020507": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020507",
    "internalMessage": "account_statusが01でない場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020601": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020601",
    "internalMessage": "member_idの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020602": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020602",
    "internalMessage": "member_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020603": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020603",
    "internalMessage": "sort_keyの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020604": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020604",
    "internalMessage": "sort_keyの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020605": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020605",
    "internalMessage": "sort_orderの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020606": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020606",
    "internalMessage": "sort_orderの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020607": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020607",
    "internalMessage": "page_numの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020608": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020608",
    "internalMessage": "page_numの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020609": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020609",
    "internalMessage": "page_sizeの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020610": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020610",
    "internalMessage": "page_sizeの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020611": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020611",
    "internalMessage": "deleted_dataの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020612": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01020612",
    "internalMessage": "deleted_dataの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020615": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020615",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020616": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020616",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01020617": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01020617",
    "internalMessage": "account_statusが01でない場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030301": {
    "messageLevel": "error",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030301",
    "internalMessage": "typeの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030302": {
    "messageLevel": "error",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030302",
    "internalMessage": "sort_keyの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030303": {
    "messageLevel": "error",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030303",
    "internalMessage": "sort_orderの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030304": {
    "messageLevel": "error",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030304",
    "internalMessage": "page_numの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030305": {
    "messageLevel": "error",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030305",
    "internalMessage": "page_sizeの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030306": {
    "messageLevel": "error",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030306",
    "internalMessage": "search_textの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030307": {
    "messageLevel": "error",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030307",
    "internalMessage": "club_owner_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030308": {
    "messageLevel": "info",
    "statusCode": 503,
    "messageType": "ServiceUnavailable",
    "messageCode": "BFWE01030308",
    "internalMessage": "サービス利用不可（システムビジー）の場合",
    "message": {
      "ja": "Service Unavailable",
      "en": "Service Unavailable"
    },
    "messageDetail": {
      "ja": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate.",
      "en": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030309": {
    "messageLevel": "info",
    "statusCode": 503,
    "messageType": "ServiceUnavailable",
    "messageCode": "BFWE01030309",
    "internalMessage": "サービス利用不可（サービス時間外）の場合",
    "message": {
      "ja": "Service Unavailable",
      "en": "Service Unavailable"
    },
    "messageDetail": {
      "ja": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate.",
      "en": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030310": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030310",
    "internalMessage": "club_owner_idが設定されているとき、他のパラメータが設定されている場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030311": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030311",
    "internalMessage": "club_owner_idが設定されていないとき、club_owner_idとsearch_text以外が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030312": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01030312",
    "internalMessage": "typeに0または1以外が設定されているとき、search_textが空の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030313": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01030313",
    "internalMessage": "SESSION.auth_statusが未設定の場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030314": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01030314",
    "internalMessage": "SESSION.auth_statusが1以外の場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01030315": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01030315",
    "internalMessage": "account_statusが01でない場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050101",
    "internalMessage": "user_nameの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050102": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050102",
    "internalMessage": "user_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050103": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050103",
    "internalMessage": "first_nameの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050104": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050104",
    "internalMessage": "first_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050105": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050105",
    "internalMessage": "belongの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050106": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050106",
    "internalMessage": "belongの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050107": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050107",
    "internalMessage": "mailの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050108": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050108",
    "internalMessage": "mailの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050109": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050109",
    "internalMessage": "mailの値がメール形式(xxx@yyy.xxx)でない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050110": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050110",
    "internalMessage": "telの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050111": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050111",
    "internalMessage": "telの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050112": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050112",
    "internalMessage": "passwordの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050113": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050113",
    "internalMessage": "passwordの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050114": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050114",
    "internalMessage": "passwordの値が英小文字大文字記号をすべて含まない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050115": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050115",
    "internalMessage": "roleの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050116": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050116",
    "internalMessage": "roleの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050117": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050117",
    "internalMessage": "first_login_flgの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050118": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050118",
    "internalMessage": "first_login_flgの値がtrue以外の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050121": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01050121",
    "internalMessage": "SESSION.auth_statusが未設定の場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050122": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01050122",
    "internalMessage": "SESSION.auth_statusが1以外の場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050123": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050123",
    "internalMessage": "telの桁数が10 or 11以外の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050124": {
    "messageLevel": "info",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01050124",
    "internalMessage": "account_statusが01でない場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050124": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050124",
    "internalMessage": "first_nameの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050125": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050125",
    "internalMessage": "first_nameの値がtrue以外の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050126": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050126",
    "internalMessage": "last_nameの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050127": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050127",
    "internalMessage": "last_nameの値がtrue以外の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050201": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050201",
    "internalMessage": "operator_idが設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050202": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050202",
    "internalMessage": "operator_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050203": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050203",
    "internalMessage": "last_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050204": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050204",
    "internalMessage": "first_nameの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050205": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050205",
    "internalMessage": "belongの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050206": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050206",
    "internalMessage": "mailの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050207": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050207",
    "internalMessage": "mailの値がメール形式(xxx@yyy.xxx)でない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050208": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050208",
    "internalMessage": "telの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050209": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050209",
    "internalMessage": "passwordの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050210": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050210",
    "internalMessage": "passwordの値が英小文字大文字記号をすべて含まない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050211": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050211",
    "internalMessage": "roleの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050212": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050212",
    "internalMessage": "first_login_flgの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050213": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050213",
    "internalMessage": "passwordに値が設定されている場合、first_login_flgの値がtrue以外の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050214": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050214",
    "internalMessage": "リクエスト項目のうち、operator_id以外の項目が存在しない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050217": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01050217",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050218": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01050218",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050219": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01050219",
    "internalMessage": "account_statusが01でない場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050220": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050220",
    "internalMessage": "リクエスト項目のうち、last_nameが設定されているとき、first_name項目が存在しない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050221": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050221",
    "internalMessage": "リクエスト項目のうち、first_nameが設定されているとき、last_name項目が存在しない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050301": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050301",
    "internalMessage": "operator_idの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050302": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050302",
    "internalMessage": "operator_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050303": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050303",
    "internalMessage": "operator_statusの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050304": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050304",
    "internalMessage": "operator_statusの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050307": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01050307",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050308": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01050308",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050401": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050401",
    "internalMessage": "search_textの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050402": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050402",
    "internalMessage": "sort_keyの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050403": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050403",
    "internalMessage": "sort_keyの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050404": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050404",
    "internalMessage": "sort_orderの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050405": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050405",
    "internalMessage": "sort_orderの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050406": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050406",
    "internalMessage": "page_numの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050407": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050407",
    "internalMessage": "page_numの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050408": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050408",
    "internalMessage": "page_sizeの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050409": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050409",
    "internalMessage": "page_sizeの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050410": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01050410",
    "internalMessage": "operator_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050413": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01050413",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01050414": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01050414",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060101": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060101",
    "internalMessage": "message_sectionの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060102": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060102",
    "internalMessage": "message_sectionの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060103": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060103",
    "internalMessage": "member_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060104": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060104",
    "internalMessage": "club_idの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060105": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060105",
    "internalMessage": "titleの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060106": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060106",
    "internalMessage": "titleの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060107": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060107",
    "internalMessage": "contentの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060108": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060108",
    "internalMessage": "contentの値が不正の場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060111": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060111",
    "internalMessage": "message_sectionが02の時、member_idが設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060112": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060112",
    "internalMessage": "message_sectionが03の時、club_idが設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060113": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01060113",
    "internalMessage": "SESSION.auth_statusが未設定の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060114": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01060114",
    "internalMessage": "SESSION.auth_statusが1以外の場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060115": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01060115",
    "internalMessage": "account_statusが01でない場合",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060201": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060201",
    "internalMessage": "message_sectionの値が設定されていない場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060202": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060202",
    "internalMessage": "message_sectionの値が不正の場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060203": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060203",
    "internalMessage": "sort_keyの値が設定されていない場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060204": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060204",
    "internalMessage": "sort_keyの値が不正の場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060205": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060205",
    "internalMessage": "sort_orderの値が設定されていない場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060206": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060206",
    "internalMessage": "sort_orderの値が不正の場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060207": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060207",
    "internalMessage": "page_numの値が設定されていない場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060208": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060208",
    "internalMessage": "page_numの値が不正の場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060209": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060209",
    "internalMessage": "page_sizeの値が設定されていない場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060210": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060210",
    "internalMessage": "page_sizeの値が不正の場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060211": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060211",
    "internalMessage": "member_idの値が不正の場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060212": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060212",
    "internalMessage": "club_nameの値が不正の場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060213": {
    "messageLevel": "info",
    "statusCode": 503,
    "messageType": "ServiceUnavailable",
    "messageCode": "BFWE01060213",
    "internalMessage": "サービス利用不可（システムビジー）の場合、エラー",
    "message": {
      "ja": "Service Unavailable",
      "en": "Service Unavailable"
    },
    "messageDetail": {
      "ja": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate.",
      "en": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060214": {
    "messageLevel": "info",
    "statusCode": 503,
    "messageType": "ServiceUnavailable",
    "messageCode": "BFWE01060214",
    "internalMessage": "サービス利用不可（システムビジー）の場合、エラー",
    "message": {
      "ja": "Service Unavailable",
      "en": "Service Unavailable"
    },
    "messageDetail": {
      "ja": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate.",
      "en": "API service is temporarily unavailable or maintenance. If the problem persists after maintenance period, contact service provider and we will investigate."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060215": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060215",
    "internalMessage": "message_sectionが02の時、member_idが設定されていない場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060216": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE01060216",
    "internalMessage": "message_sectionが03の時、club_nameが設定されていない場合、エラー",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060217": {
    "messageLevel": "info",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01060217",
    "internalMessage": "SESSION.auth_statusが未設定の場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060218": {
    "messageLevel": "info",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01060218",
    "internalMessage": "SESSION.auth_statusが1以外の場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "BFWE01060219": {
    "messageLevel": "info",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "BFWE01060219",
    "internalMessage": "account_statusが01でない場合エラー",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005001001": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005001001",
    "internalMessage": "ユーザーのステータスが無効",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005001002": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005001002",
    "internalMessage": "Fail to create Session",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005001003": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005001003",
    "internalMessage": "Fail to marshal Record",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005001004": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005001004",
    "internalMessage": "Fail to Insert record",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005001005": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005001005",
    "internalMessage": "Fail to Get record",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005001006": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005001006",
    "internalMessage": "Fail to Unmarshal record",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "0100005001007": {
    "messageLevel": "error",
    "statusCode": 500,
    "messageType": "InternalServerError",
    "messageCode": "0100005001007",
    "internalMessage": "Fail to Delete record",
    "message": {
      "ja": "Internal Server Error",
      "en": "Internal Server Error"
    },
    "messageDetail": {
      "ja": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it.",
      "en": "An unknown internal error occurred. If the error persists, please contact service provider to figure out what has happened and how to fix it."
    },
    "action": "NoAutoRetry"
  },
  "01010301": {
    "messageLevel": "info",
    "statusCode": 404,
    "messageType": "NotFound",
    "messageCode": "01010301",
    "internalMessage": "Key has expired",
    "message": {
      "ja": "Not Found",
      "en": "Not Found"
    },
    "messageDetail": {
      "ja": "The URI requested is invalid or the resource requested does not exists.",
      "en": "The URI requested is invalid or the resource requested does not exists."
    },
    "action": "NoAutoRetry"
  },
  "01010302": {
    "messageLevel": "info",
    "statusCode": 404,
    "messageType": "NotFound",
    "messageCode": "01010302",
    "internalMessage": "Key not found",
    "message": {
      "ja": "Not Found",
      "en": "Not Found"
    },
    "messageDetail": {
      "ja": "The URI requested is invalid or the resource requested does not exists.",
      "en": "The URI requested is invalid or the resource requested does not exists."
    },
    "action": "NoAutoRetry"
  },
  "BFWE00000001": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE00000001",
    "internalMessage": "リクエストが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE00000002": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE00000002",
    "internalMessage": "Content-Typeの値が設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE00000003": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE00000003",
    "internalMessage": "Content-Typeが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE00000004": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE00000004",
    "internalMessage": "User-Agentが設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE00000005": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE00000005",
    "internalMessage": "User-Agentが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE00000006": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE00000006",
    "internalMessage": "X-Requested-Withが設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE00000007": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE00000007",
    "internalMessage": "X-Requested-Withが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE00000008": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE00000008",
    "internalMessage": "X-Nudge-Interaction-Idが設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE00000009": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE00000009",
    "internalMessage": "X-Nudge-Interaction-Idが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE00000010": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE00000010",
    "internalMessage": "X-API-KEYが設定されていない場合",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  },
  "BFWE00000011": {
    "messageLevel": "info",
    "statusCode": 400,
    "messageType": "BadRequest",
    "messageCode": "BFWE00000011",
    "internalMessage": "X-API-Keyが間違っている",
    "message": {
      "ja": "Bad Request",
      "en": "Bad Request"
    },
    "messageDetail": {
      "ja": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send.",
      "en": "Any case where a parameter is invalid, or a required parameter is missing. Please check the data you are trying to send."
    },
    "action": "NoAutoRetry"
  }
}`
