package utils

import socketio "github.com/googollee/go-socket.io"

const TWILIO_ACCOUNT_SID string ="AC0915dbf93f898f54027ef27788e0c282"

const TWILIO_VERIFY_SERVICE_SID string ="VAf1e9a94afb5b63d5f051b59deab51077"

var SocketServerInstance =socketio.NewServer(nil)