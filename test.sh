#!/bin/bash

(cd appengine && go test -v) || exit 10
(cd fasthttp && go test -v) || exit 10
(cd standard && go test -v) || exit 10
