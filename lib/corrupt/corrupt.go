package corrupt

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func Init(minsec int, maxsec int, corruptdirs []string, corruptnbytes int, errLog *Logger) {
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		//We sleep for a range between min and max.
		time.Sleep(time.Duration(r.Intn(maxsec-minsec)+minsec) * time.Second)

		//right here we decide a path from the corruptdir

		ourDir := corruptdirs[r.Intn(len(t))]
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

	//open file
	f, err := os.Open(basedir + target.Name())
	defer f.Close()
	if err != nil {
		return
	}

	//determine needed buffer size and allocate it
	if target.Size() > 2560000 { //This number is magic - fear it (~2.5M)
		return errors.New(fmt.Sprintf("%s file read failed: File too large", target.Name()))
	}

	buffer := make([]byte, target.Size())
	fmt.Println(len(buffer))
	//read full file into buffer
	_, err = f.Read(buffer)
	if err != nil {
		return
	}
	fmt.Println(string(buffer))

	//generate the bytes to write to the file
	btc := make([]byte, bytesToChange)

	for i, _ := range btc {
		//this is a byte, t-trust me
		btc[i] = byte(r.Intn(255))
	}

	//pick bytesToChange places in the file and modify there.
	for i, _ := range btc {
		buffer[r.Intn(len(buffer))] = btc[i]
	}

	fmt.Println(string(buffer))
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
