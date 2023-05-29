# s3heck

- [s3heck](#s3heck)
  - [Installation](#installation)
    - [Recommendation](#recommendation)
  - [Usage](#usage)
    - [Subcommands](#subcommands)
      - [Example](#example)
        - [short](#short)
        - [Other](#other)
  - [Contributing](#contributing)

s3heck is a command line interface that can perform checks on the certificate chain.

## Installation

```bash
git clone https://github.com/kanywst/s3heck.git
cd s3heck
go install
```

### Recommendation

```bash
echo 'alias s3=s3heck' >> ~/.zshrc
source ~/.zshrc
```

## Usage

```bash
$ s3heck -h
s3heck is a command line interface that can perform checks on the certificate chain.

Usage:
  s3heck [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  x509        Output necessary information about x509 certificates

Flags:
  -h, --help   help for s3heck

Use "s3heck [command] --help" for more information about a command.
```

```bash
s3heck x509 -h
Output necessary information about x509 certificates

Usage:
  s3heck x509 [flags]

Flags:
  -c, --cert string   specify server certificate
  -d, --dns           Output subject alternative name
  -h, --help          help for x509
  -i, --issuer        Output issuer
  -t, --short         Output minimum information
  -s, --subject       Output subject
  -v, --validity      Output validity
```

### Subcommands

Use the `x509` to have outputs the desired information about the x509 certificate

#### Example

##### short

```bash
$ s3heck x509 --cert ./sample/google.com.cer --short
Certificate: SerialNumber is 6f71752163a47f9d093cec987b5cbe12
  Issuer: GTS CA 1C3
  Subject: *.google.com *.google.com, *.appengine.google.com, *.bdn.dev, *.origin-test.bdn.dev, *.cloud.google.com, *.crowdsource.google.com, *.datacompute.google.com, *.google.ca, *.google.cl, *.google.co.in, *.google.co.jp, *.google.co.uk, *.google.com.ar, *.google.com.au, *.google.com.br, *.google.com.co, *.google.com.mx, *.google.com.tr, *.google.com.vn, *.google.de, *.google.es, *.google.fr, *.google.hu, *.google.it, *.google.nl, *.google.pl, *.google.pt, *.googleadapis.com, *.googleapis.cn, *.googlevideo.com, *.gstatic.cn, *.gstatic-cn.com, googlecnapps.cn, *.googlecnapps.cn, googleapps-cn.com, *.googleapps-cn.com, gkecnapps.cn, *.gkecnapps.cn, googledownloads.cn, *.googledownloads.cn, recaptcha.net.cn, *.recaptcha.net.cn, recaptcha-cn.net, *.recaptcha-cn.net, widevine.cn, *.widevine.cn, ampproject.org.cn, *.ampproject.org.cn, ampproject.net.cn, *.ampproject.net.cn, google-analytics-cn.com, *.google-analytics-cn.com, googleadservices-cn.com, *.googleadservices-cn.com, googlevads-cn.com, *.googlevads-cn.com, googleapis-cn.com, *.googleapis-cn.com, googleoptimize-cn.com, *.googleoptimize-cn.com, doubleclick-cn.net, *.doubleclick-cn.net, *.fls.doubleclick-cn.net, *.g.doubleclick-cn.net, doubleclick.cn, *.doubleclick.cn, *.fls.doubleclick.cn, *.g.doubleclick.cn, dartsearch-cn.net, *.dartsearch-cn.net, googletraveladservices-cn.com, *.googletraveladservices-cn.com, googletagservices-cn.com, *.googletagservices-cn.com, googletagmanager-cn.com, *.googletagmanager-cn.com, googlesyndication-cn.com, *.googlesyndication-cn.com, *.safeframe.googlesyndication-cn.com, app-measurement-cn.com, *.app-measurement-cn.com, gvt1-cn.com, *.gvt1-cn.com, gvt2-cn.com, *.gvt2-cn.com, 2mdn-cn.net, *.2mdn-cn.net, googleflights-cn.net, *.googleflights-cn.net, admob-cn.com, *.admob-cn.com, googlesandbox-cn.com, *.googlesandbox-cn.com, *.safenup.googlesandbox-cn.com, *.gstatic.com, *.metric.gstatic.com, *.gvt1.com, *.gcpcdn.gvt1.com, *.gvt2.com, *.gcp.gvt2.com, *.url.google.com, *.youtube-nocookie.com, *.ytimg.com, android.com, *.android.com, *.flash.android.com, g.cn, *.g.cn, g.co, *.g.co, goo.gl, www.goo.gl, google-analytics.com, *.google-analytics.com, google.com, googlecommerce.com, *.googlecommerce.com, ggpht.cn, *.ggpht.cn, urchin.com, *.urchin.com, youtu.be, youtube.com, *.youtube.com, youtubeeducation.com, *.youtubeeducation.com, youtubekids.com, *.youtubekids.com, yt.be, *.yt.be, android.clients.google.com, developer.android.google.cn, developers.android.google.cn, source.android.google.cn
  Validity
    from : 2023-04-24 11:56:06 +0000 UTC
    to   : 2023-07-17 11:56:05 +0000 UTC
-------------------------------------------------------------
Certificate: SerialNumber is 203bc53596b34c718f5015066
  Issuer: GTS Root R1
  Subject: GTS CA 1C3
  Validity
    from : 2020-08-13 00:00:42 +0000 UTC
    to   : 2027-09-30 00:00:42 +0000 UTC
------------------------------------------------------
Certificate: SerialNumber is 203e5936f31b01349886ba217
  Issuer: GTS Root R1
  Subject: GTS Root R1
  Validity
    from : 2016-06-22 00:00:00 +0000 UTC
    to   : 2036-06-22 00:00:00 +0000 UTC
------------------------------------------------------
```

##### Other

```bash
$ s3heck x509 --cert ./sample/google.com.cer --issuer --subject --validity --dns
Certificate: SerialNumber is 6f71752163a47f9d093cec987b5cbe12
  Issuer: GTS CA 1C3
  Subject: *.google.com
  Validity
    from : 2023-04-24 11:56:06 +0000 UTC
    to   : 2023-07-17 11:56:05 +0000 UTC
  X509v3 extensions
    X509v3 Subject Alternative Name: *.google.com, *.appengine.google.com, *.bdn.dev, *.origin-test.bdn.dev, *.cloud.google.com, *.crowdsource.google.com, *.datacompute.google.com, *.google.ca, *.google.cl, *.google.co.in, *.google.co.jp, *.google.co.uk, *.google.com.ar, *.google.com.au, *.google.com.br, *.google.com.co, *.google.com.mx, *.google.com.tr, *.google.com.vn, *.google.de, *.google.es, *.google.fr, *.google.hu, *.google.it, *.google.nl, *.google.pl, *.google.pt, *.googleadapis.com, *.googleapis.cn, *.googlevideo.com, *.gstatic.cn, *.gstatic-cn.com, googlecnapps.cn, *.googlecnapps.cn, googleapps-cn.com, *.googleapps-cn.com, gkecnapps.cn, *.gkecnapps.cn, googledownloads.cn, *.googledownloads.cn, recaptcha.net.cn, *.recaptcha.net.cn, recaptcha-cn.net, *.recaptcha-cn.net, widevine.cn, *.widevine.cn, ampproject.org.cn, *.ampproject.org.cn, ampproject.net.cn, *.ampproject.net.cn, google-analytics-cn.com, *.google-analytics-cn.com, googleadservices-cn.com, *.googleadservices-cn.com, googlevads-cn.com, *.googlevads-cn.com, googleapis-cn.com, *.googleapis-cn.com, googleoptimize-cn.com, *.googleoptimize-cn.com, doubleclick-cn.net, *.doubleclick-cn.net, *.fls.doubleclick-cn.net, *.g.doubleclick-cn.net, doubleclick.cn, *.doubleclick.cn, *.fls.doubleclick.cn, *.g.doubleclick.cn, dartsearch-cn.net, *.dartsearch-cn.net, googletraveladservices-cn.com, *.googletraveladservices-cn.com, googletagservices-cn.com, *.googletagservices-cn.com, googletagmanager-cn.com, *.googletagmanager-cn.com, googlesyndication-cn.com, *.googlesyndication-cn.com, *.safeframe.googlesyndication-cn.com, app-measurement-cn.com, *.app-measurement-cn.com, gvt1-cn.com, *.gvt1-cn.com, gvt2-cn.com, *.gvt2-cn.com, 2mdn-cn.net, *.2mdn-cn.net, googleflights-cn.net, *.googleflights-cn.net, admob-cn.com, *.admob-cn.com, googlesandbox-cn.com, *.googlesandbox-cn.com, *.safenup.googlesandbox-cn.com, *.gstatic.com, *.metric.gstatic.com, *.gvt1.com, *.gcpcdn.gvt1.com, *.gvt2.com, *.gcp.gvt2.com, *.url.google.com, *.youtube-nocookie.com, *.ytimg.com, android.com, *.android.com, *.flash.android.com, g.cn, *.g.cn, g.co, *.g.co, goo.gl, www.goo.gl, google-analytics.com, *.google-analytics.com, google.com, googlecommerce.com, *.googlecommerce.com, ggpht.cn, *.ggpht.cn, urchin.com, *.urchin.com, youtu.be, youtube.com, *.youtube.com, youtubeeducation.com, *.youtubeeducation.com, youtubekids.com, *.youtubekids.com, yt.be, *.yt.be, android.clients.google.com, developer.android.google.cn, developers.android.google.cn, source.android.google.cn
-------------------------------------------------------------
Certificate: SerialNumber is 203bc53596b34c718f5015066
  Issuer: GTS Root R1
  Subject: GTS CA 1C3
  Validity
    from : 2020-08-13 00:00:42 +0000 UTC
    to   : 2027-09-30 00:00:42 +0000 UTC
  X509v3 extensions
    X509v3 Subject Alternative Name:
------------------------------------------------------
Certificate: SerialNumber is 203e5936f31b01349886ba217
  Issuer: GTS Root R1
  Subject: GTS Root R1
  Validity
    from : 2016-06-22 00:00:00 +0000 UTC
    to   : 2036-06-22 00:00:00 +0000 UTC
  X509v3 extensions
    X509v3 Subject Alternative Name:
------------------------------------------------------
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am "Add some feature"`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
