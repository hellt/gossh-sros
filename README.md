A basic example that demonstrates how to use Go [golang.org/x/crypto/ssh](https://pkg.go.dev/golang.org/x/crypto@v0.0.0-20201112155050-0c6587e931a9/ssh) package to remotely execute commands on Nokia SR OS routers.

Nokia SR OS requires a login shell to be spawned and bytes written to Stdin.

Based on https://stackoverflow.com/questions/48468318/golang-filtering-stdout-of-active-ssh-session

## Example output

```
❯ go run main.go        
Connecting to  10.2.0.11:22

 SR OS Software
 Copyright (c) Nokia 2020.  All Rights Reserved.
 
 Trademarks
 
 Nokia and the Nokia logo are registered trademarks of Nokia. All other
 trademarks are the property of their respective owners.
 
 IMPORTANT: READ CAREFULLY
 
 The SR OS Software (the "Software") is proprietary to Nokia and is subject
 to and governed by the terms and conditions of the End User License
 Agreement accompanying the product, made available at the time of your order,
 or posted on the Nokia website (collectively, the "EULA").  As set forth
 more fully in the EULA, use of the Software is strictly limited to your 
 internal use.  Downloading, installing, or using the Software constitutes
 acceptance of the EULA and you are binding yourself and the business entity
 that you represent to the EULA.  If you do not agree to all of the terms of
 the EULA, then Nokia is unwilling to license the Software to you and (a) you
 may not download, install or use the Software, and (b) you may return the
 Software as more fully set forth in the EULA.
 
 This product contains cryptographic features and is subject to United States 
 and local country laws governing import, export, transfer and use. Delivery 
 of Nokia cryptographic products does not imply third-party authority to 
 import, export, distribute or use encryption.
 
 If you require further assistance please contact us by sending an email
 to support@nokia.com.

A:vSIM# show version 
TiMOS-B-19.10.R3 both/x86_64 Nokia 7750 SR Copyright (c) 2000-2020 Nokia.
All rights reserved. All use subject to applicable license agreements.
Built on Wed Feb 12 19:18:39 PST 2020 by builder in /builds/c/1910B/R3/panos/main
A:vSIM# logout 
```