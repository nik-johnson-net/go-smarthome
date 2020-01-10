# TP-Link Smart Home (Kasa) Golang Client

This is a golang TP-Link Smart Home (Kasa) client for communicating with Kasa devices directly. It's built and tested against the HS110 and HS300 smart plug devices. Other models have not been tested. It communicates directly with devices via TP-Link's custom protocol, **not** the Kasa cloud API.

<https://enphase.com/en-us/support/what-envoy>

## Example

```go
import "github.com/nik-johnson-net/go-smarthome"

// Create the client
client := smarthome.NewClient("192.168.0.201")

// Get the list of children
info, err := client.SysInfo()

// For each child, get the current power usage
for _, child := range info.Children {
    meter, err := client.EMeter(child.ID)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s: %fW\n", child.Alias, meter.PowerW)
}
```

## Thanks

Shoutout to Lubomir Stroetmann and Tobias Esser, who wrote the fantastic blog post [Reverse Engineering the TP-Link HS110](https://www.softscheck.com/en/reverse-engineering-tp-link-hs110/). Their work made this library possible.

## License

This library is provided under the [MIT License](LICENSE.md)
