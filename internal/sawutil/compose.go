// Copyright 2020 Paul Sitoh
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sawutil

const composeContent = `version: "2.1"

services:

  settings-tp:
    image: hyperledger/sawtooth-settings-tp:chime
    container_name: sawtooth-settings-tp-default
    entrypoint: settings-tp -vv -C tcp://validator:4004
    depends_on:
      - validator

  intkey-tp-python:
    image: hyperledger/sawtooth-intkey-tp-python:chime
    container_name: sawtooth-intkey-tp-python-default
    entrypoint: intkey-tp-python -vv -C tcp://validator:4004
    depends_on:
      - validator

  xo-tp-python:
    image: hyperledger/sawtooth-xo-tp-python:chime
    container_name: sawtooth-xo-tp-python-default
    entrypoint: xo-tp-python -vv -C tcp://validator:4004
    depends_on:
      - validator

  validator:
    image: hyperledger/sawtooth-validator:chime
    container_name: sawtooth-validator-default
    expose:
      - 4004
    ports:
      - "4004:4004"
    # start the validator with an empty genesis batch
    entrypoint: "bash -c \"\
        sawadm keygen && \
        sawtooth keygen my_key && \
        sawset genesis -k /root/.sawtooth/keys/my_key.priv && \
        sawset proposal create \
          -k /root/.sawtooth/keys/my_key.priv \
          sawtooth.consensus.algorithm.name=Devmode \
          sawtooth.consensus.algorithm.version=0.1 \
          -o config.batch && \
        sawadm genesis config-genesis.batch config.batch && \
        sawtooth-validator -vv \
          --endpoint tcp://validator:8800 \
          --bind component:tcp://eth0:4004 \
          --bind network:tcp://eth0:8800 \
          --bind consensus:tcp://eth0:5050 \
        \""

  devmode-engine:
    image: hyperledger/sawtooth-devmode-engine-rust:chime
    container_name: sawtooth-devmode-engine-rust-default
    entrypoint: devmode-engine-rust -C tcp://validator:5050
    depends_on:
      - validator

  rest-api:
    image: hyperledger/sawtooth-rest-api:chime
    container_name: sawtooth-rest-api-default
    entrypoint: sawtooth-rest-api -C tcp://validator:4004 --bind rest-api:8008
    ports:
      - "8008:8008"
    depends_on:
      - validator`

func GenerateCompose() []byte {
	return []byte(composeContent)
}
