[![wercker status](https://app.wercker.com/status/28206e111a8436a88af297f075ed8005/m/ "wercker status")](https://app.wercker.com/project/bykey/28206e111a8436a88af297f075ed8005)

[![GoDoc](https://godoc.org/github.com/aquilax/tripcode?status.png)](https://godoc.org/github.com/aquilax/tripcode)

Tripcode generation library for golang.

Based on code from [avimedia](http://avimedia.livejournal.com/1583.html), [KenanY](https://github.com/KenanY/tripcode) and [SaveTheInternet](https://github.com/savetheinternet/Tinyboard/blob/master/inc/functions.php#L1969-L2003).

## Usage
    tripcode.Tripcode("password")
and

    tripcode.SecureTripcode("password", "secure salt")
