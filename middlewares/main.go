package middlewares

import "time"

const APPLICATION_NAME = "yourAppName"
const JWT_SIGNATURE_KEY = "abc123456789"

var JWT_EXPIRES_AT = time.Now().Add(time.Hour * 1).Unix()
