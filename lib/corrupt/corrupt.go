package corrupt

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func Init(minsec int, maxsec int, corruptdirs []string, corruptnbytes int, errLog *log.Logger) {
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		time.Sleep(time.Duration(r.Intn(maxsec-minsec)+minsec) * time.Second)

		ourDir := corruptdirs[r.Intn(len(corruptdirs))]
		a, err := ioutil.ReadDir(ourDir)
		if err != nil {
			errLog.Println(err)
		}

		f := a[r.Intn(len(a))]
		err = corrupt(ourDir, f, corruptnbytes)
		if err != nil {
			errLog.Printf("Failed to corrupt file %s in directory %s: %s\n", f.Name(), ourDir, err)
		} else {
			errLog.Printf("Successfully corrupted file %s in directory %s\n", f.Name(), ourDir)
		}
	}

}

func corrupt(basedir string, target os.FileInfo, bytesToChange int) (err error) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	f, err := os.Open(basedir + target.Name())
	defer f.Close()
	if err != nil {
		return
	}

	//determine needed buffer size and allocate it
	//fsck off if you're too large
	if target.Size() > 2560000 { //This number is magic - fear it (~2.5M)
		return errors.New(fmt.Sprintf("%s file read failed: File too large", target.Name()))
	}

	buffer := make([]byte, target.Size())

	//read full file into buffer
	_, err = f.Read(buffer)
	if err != nil {
		return
	}

	//generate the bytes to write to the file
	btc := make([]byte, bytesToChange)

	for i, _ := range btc {
		//this is a byte, t-trust me
		btc[i] = byte(r.Intn(255))
	}

	for i, _ := range btc {
		buffer[r.Intn(len(buffer))] = btc[i]
	}

	f, err = os.OpenFile(basedir+target.Name(), os.O_RDWR, 0660)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write(buffer)
	if err != nil {
		return
	}
	return nil

}
