package entropyexhaustion

import "ioutil"

func Init(errLog *Logger) {
	errLog.Println("Starting entropy exhaustion. ")
	for {
		/* ReadFile reads until EOF. If that were to happen, we just do it again. */
		/* We really don't care about the result, it's just random things anyways. */
		_, err := ioutil.ReadFile("/dev/random")
		/* EOF isn't an error. */
		if err != nil {
			log.Println("entropyexhaustion failed: ", err)
		}

	}
}
