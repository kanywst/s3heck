Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            6f:71:75:21:63:a4:7f:9d:09:3c:ec:98:7b:5c:be:12
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = US, O = Google Trust Services LLC, CN = GTS CA 1C3
        Validity
            Not Before: Apr 24 11:56:06 2023 GMT
            Not After : Jul 17 11:56:05 2023 GMT
        Subject: CN = *.google.com
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:2c:86:44:ff:e4:af:fb:85:fd:36:f9:89:8a:bb:
                    b4:e6:c3:97:f2:c4:8d:a2:44:3d:1b:89:ac:de:d3:
                    34:aa:fb:42:26:7d:bf:80:73:0b:1f:69:2b:f9:77:
                    7a:4e:2c:8f:30:5a:7d:1b:d2:32:e1:8d:c2:8c:4f:
                    dd:37:a5:c7:ae
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
                9C:9D:2C:57:BE:FE:AD:8E:94:44:F1:38:67:C5:7D:D0:5E:B6:AC:98
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
                  URI:http://crls.pki.goog/gts1c3/fVJxbV-Ktmk.crl
            CT Precertificate SCTs: 
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : AD:F7:BE:FA:7C:FF:10:C8:8B:9D:3D:9C:1E:3E:18:6A:
                                B4:67:29:5D:CF:B1:0C:24:CA:85:86:34:EB:DC:82:8A
                    Timestamp : Apr 24 12:56:11.001 2023 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:45:02:21:00:E3:4A:6E:9D:9A:18:88:F5:86:E9:58:
                                C4:9A:9F:0C:93:D1:0E:0A:B3:F1:70:A9:4B:B7:5B:7D:
                                DA:69:00:0A:24:02:20:09:35:40:1E:82:87:60:A2:89:
                                3D:0D:32:80:B9:8A:A3:52:F2:17:D7:07:8F:9B:4A:41:
                                EC:51:C0:14:B3:24:6D
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : 7A:32:8C:54:D8:B7:2D:B6:20:EA:38:E0:52:1E:E9:84:
                                16:70:32:13:85:4D:3B:D2:2B:C1:3A:57:A3:52:EB:52
                    Timestamp : Apr 24 12:56:11.006 2023 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:45:02:20:32:A7:6B:0C:BE:4E:A9:AC:F3:A2:F7:A2:
                                27:76:D0:B3:07:25:A3:F1:E3:EE:1F:92:9B:E4:B3:B7:
                                90:C1:11:62:02:21:00:E4:7C:89:2D:D0:CE:B4:F7:A5:
                                73:6F:EC:07:61:CE:83:5F:0F:69:F1:3F:6A:67:C8:92:
                                59:7B:17:10:5C:76:06
    Signature Algorithm: sha256WithRSAEncryption
    Signature Value:
        7a:a3:66:62:44:5e:2b:69:06:4f:c4:86:a4:e7:66:62:54:1b:
        59:2f:bf:2a:d0:4c:be:62:e2:74:71:a7:66:00:a4:e5:51:3f:
        e2:f1:bf:5e:77:03:e2:39:49:b7:28:fb:ce:cb:81:57:33:b4:
        3d:ee:ac:f8:31:e0:9f:3d:36:f9:f2:10:b1:a7:1f:7f:32:d0:
        84:2c:58:08:2b:c0:d4:41:31:ca:2a:86:dd:ed:96:9d:34:85:
        fb:a1:f9:f5:60:f4:a4:0d:fd:01:85:72:b9:9e:6e:77:47:63:
        70:5f:6c:2c:6a:cc:d1:b8:c5:85:66:ab:1b:f0:67:50:38:57:
        ba:18:6c:da:21:20:00:04:6c:7c:31:39:26:f3:9e:d1:4d:0f:
        4d:e3:60:5d:0f:28:a3:b7:e5:57:b5:e8:13:26:91:62:c5:70:
        6c:42:25:2e:0e:aa:f2:41:ec:16:ad:f4:f6:98:33:84:0e:96:
        e3:80:a5:35:26:8b:e9:bc:b1:20:af:3a:81:fb:f0:97:f0:35:
        87:8a:87:27:2f:4f:24:ff:f3:22:da:45:e9:4d:9a:4c:6e:b5:
        eb:5f:da:93:48:92:d2:e0:e2:3e:dc:a9:e2:ba:e4:c2:29:c9:
        f6:f7:01:7d:60:2f:4e:84:de:be:35:ff:8a:93:85:17:4e:09:
        45:d7:62:d0
