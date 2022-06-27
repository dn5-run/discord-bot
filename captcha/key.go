package captcha

import "time"

var keyMap = map[string]string{}

func registerKey(ID string, KEY string) string {
	keyMap[ID] = KEY
	go func() {
		time.Sleep(5 * time.Minute)
		unregisterKey(ID)
	}()
	return KEY
}

func getKey(ID string) string {
	session, ok := keyMap[ID]
	if !ok {
		return ""
	}
	return session
}

func unregisterKey(ID string) {
	delete(keyMap, ID)
}
