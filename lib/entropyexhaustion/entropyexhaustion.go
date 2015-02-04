package entropyexhaustion

import (
	"io/ioutil"
	"log"
)

/*
 * NOTE: running this on a machine with haveged
 * is useless, as it adds a static amount of
 * good (non-static) entropy.
 *
 * Most machines don't have that installed though.
 */

func Init(errLog *log.Logger) {
	errLog.Println("Starting entropy exhaustion. ")
	for {
		/* ReadFile reads until EOF. If that were to happen, we just do it again. */
		/* We really don't care about the result, it's just random things anyways. */
		_, err := ioutil.ReadFile("/dev/random")
		/* EOF isn't an error. */
		if err != nil {
			errLog.Println("entropyexhaustion failed: ", err)
		}

	}
}
