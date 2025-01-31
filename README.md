<a id="markdown-component-accesslogs---settings-component-for-logging-http-accesslogs" name="markdown-component-accesslogs---settings-component-for-logging-http-accesslogs" ></a>
# component-accesslog - Settings component for logging HTTP accesslogs
[![GoDoc](https://godoc.org/github.com/asecurityteam/component-accesslog?status.svg)](https://godoc.org/github.com/asecurityteam/component-accesslog)

[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-accesslog&metric=bugs)](https://sonarcloud.io/dashboard?id=asecurityteam_component-accesslog)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-accesslog&metric=code_smells)](https://sonarcloud.io/dashboard?id=asecurityteam_component-accesslog)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-accesslog&metric=coverage)](https://sonarcloud.io/dashboard?id=asecurityteam_component-accesslog)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-accesslog&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=asecurityteam_component-accesslog)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-accesslog&metric=ncloc)](https://sonarcloud.io/dashboard?id=asecurityteam_component-accesslog)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-accesslog&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-accesslog)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-accesslog&metric=alert_status)](https://sonarcloud.io/dashboard?id=asecurityteam_component-accesslog)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-accesslog&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-accesslog)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-accesslog&metric=security_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-accesslog)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-accesslog&metric=sqale_index)](https://sonarcloud.io/dashboard?id=asecurityteam_component-accesslog)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-accesslog&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=asecurityteam_component-accesslog)

<!-- TOC -->

- [component-accesslog - Settings component for logging HTTP accesslogs](#component-accesslog---settings-component-for-logging-http-accesslogs)
    - [Overview](#overview)
    - [Quick Start](#quick-start)
    - [Status](#status)
    - [Contributing](#contributing)
        - [Building And Testing](#building-and-testing)
        - [License](#license)
        - [Contributing Agreement](#contributing-agreement)

<!-- /TOC -->

<a id="markdown-overview" name="overview"></a>
## Overview

This is a [`settings`](https://github.com/asecurityteam/settings) that enables
constructing a `RoundTripper` that logs meta data on `RoundTrip`, such as the HTTP Method, time it takes to complete the HTTP request, path and query parameters, etc.

<a id="markdown-quick-start" name="quick-start"></a>
## Quick Start

```golang
package main

import (
    "context"
    "net/http"

    accesslog "github.com/asecurityteam/component-accesslog"
)

func main() {
    ctx := context.Background()
	accessLogComponent := AccessLogComponent{}
    accessLogConfig := accessLogComponent.Settings()
    wrapper, _ := accessLogComponent.New(ctx, accessLogConfig)
    transport := wrapper(http.DefaultTransport)
    client := &http.Client{Transport: transport}
    req, _ := http.NewRequest(http.MethodGet, "www.google.com", http.NoBody)

    // should see accesslogs
    _, _ := j.HTTPClient.Do(req)

}
```

<a id="markdown-status" name="status"></a>
## Status

This project is in incubation which means we are not yet operating this tool in
production and the interfaces are subject to change.

<a id="markdown-contributing" name="contributing"></a>
## Contributing

<a id="markdown-building-and-testing" name="building-and-testing"></a>
### Building And Testing

We publish a docker image called [SDCLI](https://github.com/asecurityteam/sdcli) that
bundles all of our build dependencies. It is used by the included Makefile to help
make building and testing a bit easier. The following actions are available through
the Makefile:

-   make dep

    Install the project dependencies into a vendor directory

-   make lint

    Run our static analysis suite

-   make test

    Run unit tests and generate a coverage artifact

-   make integration

    Run integration tests and generate a coverage artifact

-   make coverage

    Report the combined coverage for unit and integration tests

<a id="markdown-license" name="license"></a>
### License

This project is licensed under Apache 2.0. See LICENSE.txt for details.

<a id="markdown-contributing-agreement" name="contributing-agreement"></a>
### Contributing Agreement

Atlassian requires signing a contributor's agreement before we can accept a patch. If
you are an individual you can fill out the [individual
CLA](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=3f94fbdc-2fbe-46ac-b14c-5d152700ae5d).
If you are contributing on behalf of your company then please fill out the [corporate
CLA](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=e1c17c66-ca4d-4aab-a953-2c231af4a20b).


