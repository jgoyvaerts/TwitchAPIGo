# TwitchAPIGo

An API library for the Twitch API v3 written in Go.

Should be mostly complete, let me know if something is missing and then I can add it (or feel free to create a pull request).

## Installation

Simple as:

    $ go get github.com/jgoyvaerts/TwitchAPIGo/twitch

And in your code:

    import "github.com/jgoyvaerts/TwitchAPIGo/twitch"

## Usage

```go
package main

import "github.com/jgoyvaerts/TwitchAPIGo/twitch"
import "fmt"

func main() {
    client := twitch.NewClient("CLIENTID")
    client.ItemsPerPage = 25 //this is default

    channel, err := client.GetChannel("twitch")
    if err != nil  {
      // Handle error
    }
    //do something with channel here

    summary, err := client.GetStreamsSummary("League of Legends")
    if err != nil {
      // Handle error
    }
    fmt.Printf("League of Legends has %v viewers across %v channels\n", summary.Viewers, summary.Channels)

    
}
```

See the `examples` directory for more (coming soon).

## Authentication

If you want to use an OAuth token, you can do so setting the OAuthToken property on the client. Every call from then on will have the OAuthToken added to it.

## Debugging

If you want to enable the logging of requests and responses, you can do so by using SetDebug(true). By default the logging will be sent to os.Stdout but you can modify this by using SetDebugOutput(io.Writer) like this:

```go
package main

import "github.com/jgoyvaerts/TwitchAPIGo/twitch"
import "fmt"
import "os"

func main() {
    client := twitch.NewClient("CLIENTID")
    client.SetDebug(true)
    file, err = os.Create("/path/to/logFile.log")
    if nil != err {
        panic(err.Error())
    }
    client.SetDebugOutput(file)
    //requests and responses will now be logged to your logfile
    
}
```

## License

MIT, see the LICENSE file.

## Help

If you have any questions, send me a tweet [@jgoyvaerts](http://twitter.com/jgoyvaerts)