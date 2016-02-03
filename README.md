dnsmasq-china-list
==================

Configuration for hot China domains (or CDN domains that have node in China) to accelerate via Dnsmasq (Now also includes bogus-nxdomain lines to stop common DNS servers from hijacking NXDOMAIN results)

更新工具
=======

Golang 版本更新工具，可以分析已开启 `-log-queries` 参数的 dnsmasq 日志中的域名，根据NS记录判定是否中国域名。 使用 `go build` 或 `env GOOS=linux GOARCH=amd64 go build` 编译。

用法：

- `-confdir=`: 遍历目录下的现有 dnsmasq 配置文件，例如 ` -confdir=/etc/dnsmasq.d/`

- `-logfile=`: dsnmasq 日志文件路径，例如 `-logfile=/var/log/dnsmasq.log`

- `-output=` : 将新的中国域名配置写入文件，默认值为当前目录下的  `accelerated-domains2.china.conf`

Content
=======

- `accelerated-domains.china.conf`: Acceleratable Domains.

  The domain should have a better resolving speed or result when using a Chinese DNS server.

  To determine if a domain is eligible, one of the criteria below must be met:

 - The domain's NS server is located in China.

 - The domain will resolve to an IP located in China mainland when using a Chinese DNS server, but _not_ always do when using a foreign DNS server (For example, CDN accelerated sites that have node in China). This however does _not_ include those having node _near_ China mainland, like in Japan, Hong Kong, Taiwan, etc.

  Please don't add subdomains if the top domain is already in the list. This includes all .cn domains which are already matched by the `/cn/` rule.

- `bogus-nxdomain.china.conf`: Known addresses that are hijacking NXDOMAIN results returned by DNS servers.

- `google.china.conf`: Acceleratable Google domains.

  These domains are resolved to Google China servers when using a Chinese DNS. In most conditions this will yield better page load time for sites using Google's web services, e.g. Google Web Fonts and AdSense.

  Bear in mind that they are _not_ considered stable. **Use at your own risk**.

Usage
=====

1. Place accelerated-domains.china.conf, bogus-nxdomain.china.conf (and optionally google.china.conf) under /etc/dnsmasq.d/ (Create the folder if it does not exist).
2. Uncomment "conf-dir=/etc/dnsmasq.d" in /etc/dnsmasq.conf
3. (Optional) Place dnsmasq-update-china-list into /usr/bin/

License
=======

```
Copyright © 2015 Felix Yan <felixonmars@archlinux.org>
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file for more details.
```
