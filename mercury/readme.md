# GoMercury - Mercury Web Parser wrapper

This library is the wrapper around the [Mercury Web Parser](https://mercury.postlight.com/web-parser/) REST API by Postlight Labs.

## Example:
    c := &MercuryConfig{
		ApiKey: os.Getenv("MERCURY_APIKEY"),
	}

	client := New(c)
	doc, err := client.Parse("https://trackchanges.postlight.com/building-awesome-cms-f034344d8ed")

