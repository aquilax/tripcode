[![GoDoc](https://godoc.org/github.com/ComSecNinja/tripcode?status.png)](https://godoc.org/github.com/ComSecNinja/tripcode)

Tripcode generation library for golang.

Based on code from [avimedia](http://avimedia.livejournal.com/1583.html), [KenanY](https://github.com/KenanY/tripcode) and [SaveTheInternet](https://github.com/savetheinternet/Tinyboard/blob/master/inc/functions.php#L1969-L2003).

## Usage
    tripcode.Tripcode("password")
and
    tripcode.SecureTripcode("password", "secure salt")
