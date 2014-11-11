# toget

a timeout layer on top of `net/http` client basic functionality

	import "github.com/jasonmoo/toget"

	resp, err = toget.Get("https://www.google.com", time.Millisecond)
	if err != TimeoutError {
		log.Println("You got some fast internets!")
	}

### Docs
[http://godoc.org/github.com/jasonmoo/toget](http://godoc.org/github.com/jasonmoo/toget)