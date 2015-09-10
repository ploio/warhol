# Warhol

[![Travis](https://img.shields.io/travis/portefaix/warhol.svg)]()

Warhol is a Docker image factory system. It receives web hook from Git providers,
build Docker image and push them to a Docker registry.
Providers supported :

* [x] Gitlab
* [ ] Github

## Usage

Launch the web service :

	$ warhol -d
	2015/09/09 12:11:09 [INFO] [warhol] Creates the Docker builder
	2015/09/09 12:11:09 [DEBUG] [api] Creates webservice
	2015/09/09 12:11:09 [INFO] [warhol] Warhol is ready on 8080

Setup your webhooks URI (Ex with Gitlab) :

- Tag push events : http://x.x.x.x:8080/api/v1/notification/gitlab/tag
- Push events : http://x.x.x.x:8080/api/v1/notification/gitlab/push


## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>
