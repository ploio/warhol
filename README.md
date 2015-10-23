# Warhol

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/portefaix%2Fwarhol.svg)](https://badge.fury.io/gh/nlamirault%2Fwarhol)

Master :
* [![Circle CI](https://circleci.com/gh/portefaix/warhol/tree/master.svg?style=svg)](https://circleci.com/gh/portefaix/warhol/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/portefaix/warhol/tree/develop.svg?style=svg)](https://circleci.com/gh/portefaix/warhol/tree/develop)

Warhol is a Docker image factory system. It receives web hook from Git providers,
build Docker image and push them to a Docker registry.
Providers supported :

* [Gitlab](https://gitlab.com/)
* [Github](https://github.com/)


## Usage

### Docker registry

Configure Native basic auth:

    $ mkdir auth
    $ docker run --rm --entrypoint htpasswd registry:2 -Bbn warhol warhol > auth/htpasswd

Generate TLS Certificates for the Registry

    $ openssl req \
         -newkey rsa:2048 -nodes -keyout certs/spc.key \
         -x509 -days 365 -out certs/spc.crt

Start the registry :

    $ docker run --rm=true -p 5000:5000 \
        -v `pwd`/auth:/auth \
        -e "REGISTRY_AUTH=htpasswd" \
        -e "REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm" \
        -e REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd \
        -v `pwd`/certs:/certs \
        -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/spc.crt \
        -e REGISTRY_HTTP_TLS_KEY=/certs/spc.key \
        --name registry registry:2

Check authentication :

    $ docker login 127.0.0.1:5000

### Messaging

Run Redis :

    $ docker run --rm=true -p 6379:6379 --name redis redis:3


### Factory web service

Launch the web service :

    $ bin/warhol -d=true \
        -registry-username=warhol -registry-password=warhol --registry-url=127.0.0.1:5000 \
        -redis-host=127.0.0.1
	2015/09/09 12:11:09 [INFO] [warhol] Creates the Docker builder
	2015/09/09 12:11:09 [DEBUG] [api] Creates webservice
	2015/09/09 12:11:09 [INFO] [warhol] Warhol is ready on 8080

Setup your webhooks URI (Ex with Gitlab) :

- Tag push events : http://x.x.x.x:8080/api/v1/notification/gitlab/tag
- Push events : http://x.x.x.x:8080/api/v1/notification/gitlab/push

Then when a webhook tag is received, the image is built and pushed :

    2015/09/11 00:06:51 [INFO] [warhol] Warhol is ready on 8080
    2015/09/11 00:07:02 [INFO] [gitlab] receive Tag event notification
    2015/09/11 00:07:02 [DEBUG] [gitlab] Tag webhook: gitlab.TagWebhook{Before:"1ad2471c103fba37d002529b06596f38ba5ab264", After: ....
    2015/09/11 00:07:02 [INFO] [gitlab] Tag for project foo
    2015/09/11 00:07:02 [INFO] [docker] Send project to pipeline
    2015/09/11 00:07:02 POST /api/v1/notification/gitlab/tag 200 2.197243ms 16
    2015/09/11 00:07:02 [INFO] [docker] Start building project
    2015/09/11 00:07:02 [DEBUG] [docker] Building image : 127.0.0.1:5000/warhol/foo
    [...]
    2015/09/11 00:09:08 [INFO] [docker] Build image done : 127.0.0.1:5000/warhol/foo
    2015/09/11 00:09:08 [INFO] [docker] Start pushing project
    2015/09/11 00:09:08 [DEBUG] [docker] Pushing image : 127.0.0.1:5000/warhol/foo
    2015/09/11 00:09:12 [DEBUG] [docker] The push refers to a repository [127.0.0.1:5000/warhol/foo] (len: 1)
    2015/09/11 00:09:12 [DEBUG] [docker] Sending image list
    2015/09/11 00:09:13 [DEBUG] [docker] Pushing repository 127.0.0.1:5000/warhol/foo (1 tags)
    2015/09/11 00:09:13 [DEBUG] [docker] Image 9a61b6b1315e already pushed, skipping
    [...]
    Pushing/11 00:09:26 [DEBUG] [docker] Pushing [==================================================>] 17.86 MB/17.86 MB
    2015/09/11 00:09:26 [DEBUG] [docker] Image successfully pushed
    2015/09/11 00:09:26 [DEBUG] [docker] Pushing tag for rev [7b0fe638246e] on {http://127.0.0.1:5000/v1/repositories/warhol/foo/tags/latest}
    2015/09/11 00:09:28 [INFO] [docker] Push image done : 127.0.0.1:5000/warhol/foo


## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test

## Run on localhost




## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>
