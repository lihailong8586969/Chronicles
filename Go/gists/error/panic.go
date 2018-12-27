func postToSlack(panic interface{}) {
}

func reportPanics() {
  if panic := recover(); panic != nil {
    postToSlack(panic)
  }
}

func runFunctionWithPanicReporting() {
  go func() {
    defer reportPanics()
    doSomething()
  }()
}
