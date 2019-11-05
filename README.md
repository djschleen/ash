[![DepShield Badge](https://depshield.sonatype.org/badges/djschleen/ash/depshield.svg)](https://depshield.github.io)

# Application Security Health Score (ASH)

*** NOTE:  This is a working POC but hasn't been tested at a massive scale  ***

I've been looking for a replacement for Security Defect Density that can provide a more precise measurement of the security health of an application. Enter the Application Security Health Score (ASH) - a calculation providing a single number similar to a credit score to describe application security risk.

If you were a bank and your application was applying for a loan, would you give it one with a credit score of 70, or a score of 810?

## Contributing

_Contributions are definitely encouraged!_. The scoring calculation has many TODO's that would be great to implement. Create a pull request and let's get more accurate in scoring 

## Build from Source

This application is built using go version 1.13. Ensure you are using a version of go that supports modules. 

View module help with the following command:

```bash
go help modules
```

Additional help on using go modules can be found in a blog entry by Niraj Foneska on [Medium](https://medium.com/@fonseka.live/getting-started-with-go-modules-b3dac652066d).

Once the repository is cloned, run the following:

```bash
go build
```

## Playing Around

The application can be run via Visual Studio Code with a few preconfigured launch configurations, or try the following:

``` bash
ash calculate --identifiers CVE-2010-3333,CVE-2018-11776,CVE-2017-9791,CVE-2018-5407
```
This will give a sample health score for many high severity vulnerabilities (including Apache Struts) and one low vulnerability.

The following is a sample low vulnerability calculation:

``` bash
ash calculate --identifiers CVE-2018-5407
```

The following is an example of various levels of low severity vulnerabilities:

``` bash
ash calculate --identifiers CVE-2019-1563,CVE-2019-1549,CVE-2019-1547
```



