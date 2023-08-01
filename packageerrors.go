//pkg/errors example
package main

import(
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

//Config holds configuration
type Config struct {
	//confuguration fields go here (redacted)
}

func readConfig(path string)(*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		//errors.Wrap is coming from pkg errors : we are wrapping the current error that came back from OS open with our own message
		return nil, errors.Wrap(err, "can't open configuration file")
	}
	defer file.Close()

	cfg := &Config{}
	//Parse file here (redacted)
	return cfg, nil
}

func setUpLogging(){
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY,064)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main(){
	setUpLogging()
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		fmt.Fprint(os.Stderr, "error: %s\n", err)
		log.Printf("error: +%v", err)
		os.Exit(1)

		//Normal operation (redacted
		fmt.Println(cfg)

	}
}