// 代码来自 ： https://gist.github.com/up1/3ef77e23a1db736bb3781872a8529ea0

import log "github.com/Sirupsen/logrus"

func main() {
  log.SetFormatter(&log.JSONFormatter{})

  log.WithFields(log.Fields{
    "id": 1,
    "server": "www.somkiat.cc",
  }).Info("Redirecting user")

}