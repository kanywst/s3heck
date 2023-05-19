Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            9b:c6:eb:61:e1:1c:c3:08:10:fd:cb:86:7d:9d:21:f5
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = US, O = Google Trust Services LLC, CN = GTS CA 1C3
        Validity
            Not Before: Apr 17 08:16:32 2023 GMT
            Not After : Jul 10 08:16:31 2023 GMT
        Subject: CN = *.google.com
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:55:44:28:5d:3e:34:42:3f:9f:a0:13:a3:36:d6:
                    f2:ee:39:32:5e:0d:dd:b6:1c:56:17:bd:d7:f7:71:
                    48:ad:1f:e6:28:c8:8d:df:31:a3:91:67:de:64:a8:
                    51:bb:c7:23:fd:9b:51:86:e4:ac:c8:c1:f8:85:80:
                    bb:77:8d:c3:15
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature
            X509v3 Extended Key Usage: 
                TLS Web Server Authentication
            X509v3 Basic Constraints: critical
                CA:FALSE
            X509v3 Subject Key Identifier: 
                3F:BB:C1:4E:C7:0F:64:F0:75:01:D2:F5:F6:32:5D:8D:D0:00:48:A7
            X509v3 Authority Key Identifier: 
                8A:74:7F:AF:85:CD:EE:95:CD:3D:9C:D0:E2:46:14:F3:71:35:1D:27
            Authority Information Access: 
                OCSP - URI:http://ocsp.pki.goog/gts1c3
                CA Issuers - URI:http://pki.goog/repo/certs/gts1c3.der
            X509v3 Subject Alternative Name: 
                DNS:*.google.com, DNS:*.appengine.google.com, DNS:*.bdn.dev, DNS:*.origin-test.bdn.dev, DNS:*.cloud.google.com, DNS:*.crowdsource.google.com, DNS:*.datacompute.google.com, DNS:*.google.ca, DNS:*.google.cl, DNS:*.google.co.in, DNS:*.google.co.jp, DNS:*.google.co.uk, DNS:*.google.com.ar, DNS:*.google.com.au, DNS:*.google.com.br, DNS:*.google.com.co, DNS:*.google.com.mx, DNS:*.google.com.tr, DNS:*.google.com.vn, DNS:*.google.de, DNS:*.google.es, DNS:*.google.fr, DNS:*.google.hu, DNS:*.google.it, DNS:*.google.nl, DNS:*.google.pl, DNS:*.google.pt, DNS:*.googleadapis.com, DNS:*.googleapis.cn, DNS:*.googlevideo.com, DNS:*.gstatic.cn, DNS:*.gstatic-cn.com, DNS:googlecnapps.cn, DNS:*.googlecnapps.cn, DNS:googleapps-cn.com, DNS:*.googleapps-cn.com, DNS:gkecnapps.cn, DNS:*.gkecnapps.cn, DNS:googledownloads.cn, DNS:*.googledownloads.cn, DNS:recaptcha.net.cn, DNS:*.recaptcha.net.cn, DNS:recaptcha-cn.net, DNS:*.recaptcha-cn.net, DNS:widevine.cn, DNS:*.widevine.cn, DNS:ampproject.org.cn, DNS:*.ampproject.org.cn, DNS:ampproject.net.cn, DNS:*.ampproject.net.cn, DNS:google-analytics-cn.com, DNS:*.google-analytics-cn.com, DNS:googleadservices-cn.com, DNS:*.googleadservices-cn.com, DNS:googlevads-cn.com, DNS:*.googlevads-cn.com, DNS:googleapis-cn.com, DNS:*.googleapis-cn.com, DNS:googleoptimize-cn.com, DNS:*.googleoptimize-cn.com, DNS:doubleclick-cn.net, DNS:*.doubleclick-cn.net, DNS:*.fls.doubleclick-cn.net, DNS:*.g.doubleclick-cn.net, DNS:doubleclick.cn, DNS:*.doubleclick.cn, DNS:*.fls.doubleclick.cn, DNS:*.g.doubleclick.cn, DNS:dartsearch-cn.net, DNS:*.dartsearch-cn.net, DNS:googletraveladservices-cn.com, DNS:*.googletraveladservices-cn.com, DNS:googletagservices-cn.com, DNS:*.googletagservices-cn.com, DNS:googletagmanager-cn.com, DNS:*.googletagmanager-cn.com, DNS:googlesyndication-cn.com, DNS:*.googlesyndication-cn.com, DNS:*.safeframe.googlesyndication-cn.com, DNS:app-measurement-cn.com, DNS:*.app-measurement-cn.com, DNS:gvt1-cn.com, DNS:*.gvt1-cn.com, DNS:gvt2-cn.com, DNS:*.gvt2-cn.com, DNS:2mdn-cn.net, DNS:*.2mdn-cn.net, DNS:googleflights-cn.net, DNS:*.googleflights-cn.net, DNS:admob-cn.com, DNS:*.admob-cn.com, DNS:googlesandbox-cn.com, DNS:*.googlesandbox-cn.com, DNS:*.safenup.googlesandbox-cn.com, DNS:*.gstatic.com, DNS:*.metric.gstatic.com, DNS:*.gvt1.com, DNS:*.gcpcdn.gvt1.com, DNS:*.gvt2.com, DNS:*.gcp.gvt2.com, DNS:*.url.google.com, DNS:*.youtube-nocookie.com, DNS:*.ytimg.com, DNS:android.com, DNS:*.android.com, DNS:*.flash.android.com, DNS:g.cn, DNS:*.g.cn, DNS:g.co, DNS:*.g.co, DNS:goo.gl, DNS:www.goo.gl, DNS:google-analytics.com, DNS:*.google-analytics.com, DNS:google.com, DNS:googlecommerce.com, DNS:*.googlecommerce.com, DNS:ggpht.cn, DNS:*.ggpht.cn, DNS:urchin.com, DNS:*.urchin.com, DNS:youtu.be, DNS:youtube.com, DNS:*.youtube.com, DNS:youtubeeducation.com, DNS:*.youtubeeducation.com, DNS:youtubekids.com, DNS:*.youtubekids.com, DNS:yt.be, DNS:*.yt.be, DNS:android.clients.google.com, DNS:developer.android.google.cn, DNS:developers.android.google.cn, DNS:source.android.google.cn
            X509v3 Certificate Policies: 
                Policy: 2.23.140.1.2.1
                Policy: 1.3.6.1.4.1.11129.2.5.3
            X509v3 CRL Distribution Points: 
                Full Name:
                  URI:http://crls.pki.goog/gts1c3/QOvJ0N1sT2A.crl
            CT Precertificate SCTs: 
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : E8:3E:D0:DA:3E:F5:06:35:32:E7:57:28:BC:89:6B:C9:
                                03:D3:CB:D1:11:6B:EC:EB:69:E1:77:7D:6D:06:BD:6E
                    Timestamp : Apr 17 09:16:37.609 2023 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:45:02:21:00:CF:20:2B:33:C9:2F:FE:6A:50:68:1E:
                                0F:6D:20:C2:42:E1:9D:83:10:30:2A:AF:07:A2:C5:58:
                                E7:82:8B:18:1B:02:20:25:15:50:BF:8D:EA:08:C3:72:
                                1C:99:17:3F:DA:CD:6D:7F:78:9F:5B:67:1C:E5:83:48:
                                43:8C:F9:CA:97:A4:E3
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : B3:73:77:07:E1:84:50:F8:63:86:D6:05:A9:DC:11:09:
                                4A:79:2D:B1:67:0C:0B:87:DC:F0:03:0E:79:36:A5:9A
                    Timestamp : Apr 17 09:16:37.629 2023 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:45:02:20:75:A7:DE:F0:DE:75:64:10:29:13:84:4D:
                                63:B9:93:21:19:1B:A3:EB:23:1C:1C:81:22:AD:CB:11:
                                56:C7:3B:E8:02:21:00:D9:40:A5:CC:81:8C:FD:28:40:
                                A1:3F:D6:4E:DA:0A:E0:0A:6B:4A:61:E8:DA:17:42:61:
                                EE:AA:A7:4B:63:85:2C
    Signature Algorithm: sha256WithRSAEncryption
    Signature Value:
        b2:ae:18:12:a0:66:91:ac:77:7b:21:aa:88:7d:76:15:d7:aa:
        c2:1e:f5:cd:05:20:a2:61:a3:da:7e:e5:c8:00:20:9c:4e:55:
        71:a2:bb:4b:70:05:5c:78:2f:ca:2b:43:15:13:49:1a:57:e4:
        0e:2c:11:a9:11:4c:20:d2:53:40:8d:2e:4d:00:28:a2:00:c5:
        74:a9:1d:f5:19:95:f0:4e:a3:76:68:12:fd:3b:c0:7f:c9:fe:
        ff:f4:6d:b9:d5:8e:58:38:fd:cf:61:87:5a:44:ed:97:5f:6a:
        d1:22:25:5b:4e:ac:e7:88:a7:47:71:9f:2c:b3:2c:f3:df:a5:
        43:f4:64:7c:1b:31:91:2c:03:70:d5:a7:39:b9:65:67:c9:fc:
        cb:dc:12:62:98:68:eb:d4:c8:d1:fa:42:78:ba:2a:85:9c:6e:
        7a:1d:74:bc:33:aa:50:8d:51:6c:ce:17:3c:11:00:31:11:21:
        06:a8:09:c1:41:ec:a4:cf:8f:26:b3:36:0e:af:c6:78:54:db:
        d6:c8:c1:a9:c7:40:ba:58:76:2c:64:f7:12:cd:46:da:20:d3:
        63:25:d2:fd:70:c5:94:13:0f:2f:73:4e:86:0c:a6:53:41:a1:
        e8:53:d6:36:c9:ab:0a:b4:c0:d7:4c:3a:fc:8b:e7:3e:01:ae:
        df:e8:bd:f9
