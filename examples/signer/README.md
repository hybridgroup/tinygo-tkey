# signer

This TKey application is intended to provide the same functionality as the https://github.com/tillitis/tkey-device-signer application, except written entirely using TinyGo.

It is an ed25519 signing tool that runs on the hardware device. It can sign messages up to 4 kByte in length. Just like the `tkey-device-signer` application it can be used by the [`tkey-ssh-agent`](https://github.com/tillitis/tkey-ssh-agent) application for SSH authentication, and by the [`tkey-sign`](https://github.com/tillitis/tkey-sign-cli) application for providing digital signatures of files.

## Building/flashing the `signer` application

Flashing the `signer` application:

```shell
tinygo flash -size full -target=tkey ./examples/signer/app/
```

Using it with the `tkey-sign` CLI tool to obtain key:

```shell
tkey-sign -G -p tkey.pub
```

Using it with the `tkey-sign` CLI tool to sign a document, in this example the "README.md" file:

```shell
tkey-sign -S -m README.md -p tkey.pub
```
