#!/bin/bash
openssl req \
    -newkey rsa:2048 \
    -x509 \
    -nodes \
    -keyout dev-key.pem \
    -new \
    -out dev-cert.pem \
    -subj /CN=localhost \
    -reqexts SAN \
    -extensions SAN \
    -config <(cat /System/Library/OpenSSL/openssl.cnf \
        <(printf '[SAN]\nsubjectAltName=DNS:localhost')) \
    -sha256 \
    -days 3650