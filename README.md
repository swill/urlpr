urlpr
-----

A very simple webserver that enables files referenced via a URL to be printed on a local printer.  It is currently only compiled for a Mac, but it can be compiled for Linux as well.  It depends on the `lpr` unix style command, so Windows support is not currently built.

**Get the binary**
``` bash
$ wget -O /usr/local/bin/urlpr https://github.com/swill/urlpr/blob/master/bin/urlpr_darwin_amd64
```

**Run it manually**
``` bash
$ /usr/local/bin/urlpr &
```

**Or run it on boot**
``` bash
$ cd ~/Library/LaunchAgents
$ touch urlpr.plist
$ open -e urlpr.plist
```

Paste in the following.  Make sure the `Program` string matches where you installed the binary.

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>Label</key>
        <string>urlpr</string>
        <key>Program</key>
        <string>/usr/local/bin/urlpr</string>
        <key>RunAtLoad</key>
        <true/>
    </dict>
</plist>
```


