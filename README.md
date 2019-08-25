# Decentralized event ticket management <!-- omit in toc --> [![GoDoc][1]][2] [![Build Status][3]][4] [![Go Report Card][5]][6]
                                                             
[1]: https://godoc.org/github.com/v1ruc/decentralized-ticket-management?status.svg
[2]: https://godoc.org/github.com/v1ruc/decentralized-ticket-management
[3]: https://travis-ci.org/v1ruc/decentralized-ticket-management.svg?branch=master
[4]: https://travis-ci.org/v1ruc/decentralized-ticket-management
[5]: https://goreportcard.com/badge/github.com/v1ruc/decentralized-ticket-management
[6]: https://goreportcard.com/report/github.com/v1ruc/decentralized-ticket-management

[(Decentralized ticket management intro.pdf](Decentralized-ticket-management.pdf) presenting the idea.

There is a big focus on Wallet implementations, but what if we would extend it to Digital Identity which might drive mass adoption of de-centralization.
We suggest that community starts using blockchain in their life e.g. event organization (bit bot limited to),

This project was prepared during [EthBerlinZwei](https://twitter.com/ETHBerlin) Hackathon 2019.

- [Actors of the dApp](#actors-of-the-dapp)
- [Workflows](#workflows)
  - [Event Participant](#event-participant)
  - [Event Organizer](#event-organizer)
  - [Participation validators (aka Volunteers)](#participation-validators-aka-volunteers)
- [Implementation](#implementation)
  - [Setting up simulation](#setting-up-simulation)
    - [Creating DID of Event Organizer](#creating-did-of-event-organizer)
    - [Creating DID of Participant 1](#creating-did-of-participant-1)
    - [Creating DID of Participant 2](#creating-did-of-participant-2)
    - [Digital identity addresses](#digital-identity-addresses)
  - [Participation in the event](#participation-in-the-event)
    - [Participant 1 sign up](#participant-1-sign-up)
      - [Participant 2 sign up](#participant-2-sign-up)
    - [Event organizer gets the list of registered participants](#event-organizer-gets-the-list-of-registered-participants)
      - [Event organizer creates ticket for participant 1 (Dima)](#event-organizer-creates-ticket-for-participant-1-dima)
    - [Participant 1 (Dima) reads the ticket](#participant-1-dima-reads-the-ticket)
      - [Volunteer validates the ticket of participant 1 (Dima)](#volunteer-validates-the-ticket-of-participant-1-dima)
- [TODO](#todo)

## Actors of the dApp

- _Event Participant_. Person who is willing to participate in the event (hackathon or conference).
  - Manages his digital identity wallet where event related information will be stored.
  - Signs up for the event.
- _Event Organizer_.
  - Manages digital identity of the event.
  - Manages event participants.
- _Event Volunteers, Security-check/Canteen worker etc._
  - Scan participants' QR code upon request to validate the ticket

## Workflows

### Event Participant

1. Each _Event Participant_ during (or before) signup creates a [digital identity](https://github.com/monetha/js-verifiable-data#Deploying-digital-identity) which is later to be used to store a ticket for the event.
2. Upon submition of data of participation a ticket information is being stored into the digital identity of _Event Organizer_

### Event Organizer

1. Monitors a list of newly registered _Event Participant_ and issues tickets for those who performed a payment.
2. During ticket issuing information is stored to the digital identity of the _Event Participant_. Stored data contains a signed _participant's information_ (provided during signup) + _address of event organizer's digital identity_.

### Participation validators (aka Volunteers)

1. _Participant_ shows his ticket as QR code. A QR code contains a signature of _Event Organizer_ and signed ticket information which is Organizer's DID address and Participant's DID address.
2. Ticket scanner validates the signature and verifies that it was submited by the _Event Organizer_ for a specified Participant's DID address.

## Implementation

We do understand the necessity of the good UI/UX. However all we did come up during Hackathon is the command line utility that resembles the [workflows](#workflows) that would be a part of the dApp.

Command-line utility is using **Ropsten** testnet and [scanner.monetha.io](https://scanner.monetha.io) can be used in order to explore digital identities and facts that are written there (aka data points).

![UI meme](assets/ui-meme.jpg)

As any systems there are some prerequisites for further simulation of behavior. For simplicity we've created 3 digital identities (aka DIDs): _Event Organizer_, _Participant 1_, _Participant 2_.

### Setting up simulation
#### Creating DID of Event Organizer

In order to create DID you need to use `deploy-passport` tool which is a part of Monetha's Decentralized Reputation Framework (Go SDK). Follow [README](https://github.com/monetha/go-verifiable-data) instructions to build the tool.

After that, you can use the tool to create a digital identity. Note: private keys of digital identity owner should be stored in file in plain-hex format.

```sh
$ ./deploy-passport \
  -ownerkey ./event-organizer.key \
  -factoryaddr 0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2 \
  -backendurl https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041
```

```sh
WARN [08-24|20:33:16.360] Loaded configuration                     owner_address=0xC501F94F846ff03291f760f44e438DeAb1Eb3dCB backend_url=https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041 factory=0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2
WARN [08-24|20:33:17.162] Getting balance                          address=0xC501F94F846ff03291f760f44e438DeAb1Eb3dCB
WARN [08-24|20:33:17.304] Initializing PassportFactory contract    factory=0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2
WARN [08-24|20:33:17.305] Deploying Passport contract              owner_address=0xC501F94F846ff03291f760f44e438DeAb1Eb3dCB
WARN [08-24|20:33:17.927] Waiting for transaction                  hash=0xbbda36637e92b559d6fa45a025f3c81bbfe18dc6dc19fab96d9795d90d3455e0
WARN [08-24|20:33:34.550] Transaction successfully mined           tx_hash=0xbbda36637e92b559d6fa45a025f3c81bbfe18dc6dc19fab96d9795d90d3455e0 gas_used=445061
WARN [08-24|20:33:34.550] New passport deployed                    gas_used=445061
WARN [08-24|20:33:34.555] Passport deployed                        contract_address=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52
WARN [08-24|20:33:34.555] Initializing Passport contract           passport=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52
WARN [08-24|20:33:34.555] Claiming ownership                       owner_address=0xC501F94F846ff03291f760f44e438DeAb1Eb3dCB
WARN [08-24|20:33:35.101] Waiting for transaction                  hash=0x469c63e5d29fc1a13d8a771fac62a6255131a18b45510654ecaa82a9a8dbb475
WARN [08-24|20:33:43.364] Transaction successfully mined           tx_hash=0x469c63e5d29fc1a13d8a771fac62a6255131a18b45510654ecaa82a9a8dbb475 gas_used=19397
WARN [08-24|20:33:43.364] Ownership claimed successfully           gas_used=19397
WARN [08-24|20:33:43.364] Done.
```

#### Creating DID of Participant 1

```sh
$ ./deploy-passport \
  -ownerkey ./event-participant1.key \
  -factoryaddr 0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2 \
  -backendurl https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041
```

```sh
WARN [08-24|20:37:24.013] Loaded configuration                     owner_address=0x1faF3952f34A936E042D33d58A9BBF0930886D2C backend_url=https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041 factory=0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2
WARN [08-24|20:37:25.746] Getting balance                          address=0x1faF3952f34A936E042D33d58A9BBF0930886D2C
WARN [08-24|20:37:25.881] Initializing PassportFactory contract    factory=0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2
WARN [08-24|20:37:25.882] Deploying Passport contract              owner_address=0x1faF3952f34A936E042D33d58A9BBF0930886D2C
WARN [08-24|20:37:26.409] Waiting for transaction                  hash=0x8a188bae73e3ee9ac53e875d8880573f4d50a3cfc769d93ed673dc4f4a78f1d0
WARN [08-24|20:38:21.733] Transaction successfully mined           tx_hash=0x8a188bae73e3ee9ac53e875d8880573f4d50a3cfc769d93ed673dc4f4a78f1d0 gas_used=445061
WARN [08-24|20:38:21.733] New passport deployed                    gas_used=445061
WARN [08-24|20:38:21.747] Passport deployed                        contract_address=0x778Eb97D8F3938a13aBB9Ba26176752511d1eD70
WARN [08-24|20:38:21.747] Initializing Passport contract           passport=0x778Eb97D8F3938a13aBB9Ba26176752511d1eD70
WARN [08-24|20:38:21.748] Claiming ownership                       owner_address=0x1faF3952f34A936E042D33d58A9BBF0930886D2C
WARN [08-24|20:38:22.301] Waiting for transaction                  hash=0xe268c5ade69406a3f7e75387fbd8fd969ef8708fa2374eb0cef47aaad315cd82
WARN [08-24|20:39:28.777] Transaction successfully mined           tx_hash=0xe268c5ade69406a3f7e75387fbd8fd969ef8708fa2374eb0cef47aaad315cd82 gas_used=19397
WARN [08-24|20:39:28.777] Ownership claimed successfully           gas_used=19397
WARN [08-24|20:39:28.778] Done.
```

#### Creating DID of Participant 2

```sh
$ ./deploy-passport \
  -ownerkey ./event-participant2.key \
  -factoryaddr 0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2 \
  -backendurl https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041
```

```sh
WARN [08-24|20:41:36.259] Loaded configuration                     owner_address=0x88BDb5deeaA40E7c980b1f52ecD2113bBF424b5C backend_url=https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041 factory=0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2
WARN [08-24|20:41:38.200] Getting balance                          address=0x88BDb5deeaA40E7c980b1f52ecD2113bBF424b5C
WARN [08-24|20:41:38.353] Initializing PassportFactory contract    factory=0x35Cb95Db8E6d56D1CF8D5877EB13e9EE74e457F2
WARN [08-24|20:41:38.355] Deploying Passport contract              owner_address=0x88BDb5deeaA40E7c980b1f52ecD2113bBF424b5C
WARN [08-24|20:41:39.067] Waiting for transaction                  hash=0x7ca14782021390e7a03653c8461fa07e985e2d7a6b74b0df3385a236985ebc0a
WARN [08-24|20:42:30.674] Transaction successfully mined           tx_hash=0x7ca14782021390e7a03653c8461fa07e985e2d7a6b74b0df3385a236985ebc0a gas_used=445061
WARN [08-24|20:42:30.674] New passport deployed                    gas_used=445061
WARN [08-24|20:42:30.680] Passport deployed                        contract_address=0xf39c56A379f6aD182b359C0296ddf3ac30A3C871
WARN [08-24|20:42:30.680] Initializing Passport contract           passport=0xf39c56A379f6aD182b359C0296ddf3ac30A3C871
WARN [08-24|20:42:30.681] Claiming ownership                       owner_address=0x88BDb5deeaA40E7c980b1f52ecD2113bBF424b5C
WARN [08-24|20:42:31.214] Waiting for transaction                  hash=0x9756b2a2a2ebef405999e5c4f2daaf4844b4b15512493d964e1e60127d02e30b
WARN [08-24|20:42:43.850] Transaction successfully mined           tx_hash=0x9756b2a2a2ebef405999e5c4f2daaf4844b4b15512493d964e1e60127d02e30b gas_used=19397
WARN [08-24|20:42:43.850] Ownership claimed successfully           gas_used=19397
WARN [08-24|20:42:43.850] Done.
```

#### Digital identity addresses

- Event organizer DID address: [0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52](https://scanner.monetha.io/identity/0xedf8e2bb4871f2ff76d3aa8b9dc3252da06c4a52?network=ropsten)
- Participant 1 DID address: [0x778Eb97D8F3938a13aBB9Ba26176752511d1eD70](https://scanner.monetha.io/identity/0x778eb97d8f3938a13abb9ba26176752511d1ed70?network=ropsten)
- Participant 2 DID address: [0xf39c56A379f6aD182b359C0296ddf3ac30A3C871](https://scanner.monetha.io/identity/0xf39c56a379f6ad182b359c0296ddf3ac30a3c871?network=ropsten)

### Participation in the event

#### Participant 1 sign up

```bash
$ ./bin/tm signup \
  --eventaddr=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52 \
  --participantaddr=0x778Eb97D8F3938a13aBB9Ba26176752511d1eD70 \
  --participantname=Dima \
  --privatekey=./event-participant1.key \
  --backendurl=https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041 \
  --ipfsurl=http://127.0.0.1:5001/
```

Output
```
INFO [08-24|21:15:40.453] Filtering OwnershipTransferred           newOwner=0xC501F94F846ff03291f760f44e438DeAb1Eb3dCB
INFO [08-24|21:15:40.616] Getting transaction by hash              tx_hash=0x469c63e5d29fc1a13d8a771fac62a6255131a18b45510654ecaa82a9a8dbb475
INFO [08-24|21:15:40.740] Writing ephemeral public key to IPFS...
INFO [08-24|21:15:40.754] Ephemeral public key added to IPFS       hash=QmSi62mJQ4MWdAgHq1m6TBBJSPqfAQTLt2NkiHP4H3NUNL size=73
INFO [08-24|21:15:40.754] Writing encrypted message to IPFS...
INFO [08-24|21:15:40.758] Encrypted message added to IPFS          hash=QmeABowQ8hiEkZLvPVCSwwFozyV7J7gp4ciAHgaTrD9isa size=127
INFO [08-24|21:15:40.758] Writing message HMAC to IPFS...
INFO [08-24|21:15:40.762] Message HMAC added to IPFS               hash=QmaBBmXNC23sVwyb2N3sc2EYNCJ5TdXYjoHWsZh52na66p size=40
INFO [08-24|21:15:40.762] Creating directory in IPFS...
INFO [08-24|21:15:40.770] Directory created in IPFS                hash=QmbyZgioijHBTJX88ChbGRkHUPMhf11fMZ3CBXEkJYRz37
INFO [08-24|21:15:40.770] Writing private data hashes to Ethereum  passport=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52 fact_key="[115 105 103 110 117 112 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]" ipfs_hash=QmbyZgioijHBTJX88ChbGRkHUPMhf11fMZ3CBXEkJYRz37 data_key_hash=0x9e16a2bc26c410a38d49b39830703e96c7f91b325d96f32c982e4b5034dfa089
INFO [08-24|21:15:40.897] Writing IPFS private data hashes to passport fact_provider=0x1faF3952f34A936E042D33d58A9BBF0930886D2C key="[115 105 103 110 117 112 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"
INFO [08-24|21:15:41.881] Waiting for transaction                  hash=0x62b6433f5816f48237fc39cdbf07b42728b97be713cf8ca4353f5c5ac7e26912
INFO [08-24|21:15:50.172] Transaction successfully mined           tx_hash=0x62b6433f5816f48237fc39cdbf07b42728b97be713cf8ca4353f5c5ac7e26912 gas_used=60519
INFO [08-24|21:15:50.172] Signup info                              tx_hash=0x62b6433f5816f48237fc39cdbf07b42728b97be713cf8ca4353f5c5ac7e26912 ipfs_hash=QmbyZgioijHBTJX88ChbGRkHUPMhf11fMZ3CBXEkJYRz37```
```

##### Participant 2 sign up

```
./bin/tm signup \
  --eventaddr=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52 \
  --participantaddr=0xf39c56A379f6aD182b359C0296ddf3ac30A3C871 \
  --participantname=Slava \
  --privatekey=./event-participant2.key \
  --backendurl=https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041 \
  --ipfsurl=http://127.0.0.1:5001/
```

Output

```
INFO [08-24|21:18:13.029] Filtering OwnershipTransferred           newOwner=0xC501F94F846ff03291f760f44e438DeAb1Eb3dCB
INFO [08-24|21:18:13.168] Getting transaction by hash              tx_hash=0x469c63e5d29fc1a13d8a771fac62a6255131a18b45510654ecaa82a9a8dbb475
INFO [08-24|21:18:13.298] Writing ephemeral public key to IPFS...
INFO [08-24|21:18:13.309] Ephemeral public key added to IPFS       hash=QmTfpZByJwpBwPKtN9yKHkVHpTL9TwYeu9gV14aXTMNVFR size=73
INFO [08-24|21:18:13.309] Writing encrypted message to IPFS...
INFO [08-24|21:18:13.314] Encrypted message added to IPFS          hash=QmQjofP7CjNhFk3ovTaHBp8YtUtjsBUFAHykZqSgabGYaB size=128
INFO [08-24|21:18:13.314] Writing message HMAC to IPFS...
INFO [08-24|21:18:13.319] Message HMAC added to IPFS               hash=QmVuuBfWVQ11sy2nyRhmywSqvkc34v2V2PRroDCtFrKvNY size=40
INFO [08-24|21:18:13.319] Creating directory in IPFS...
INFO [08-24|21:18:13.322] Directory created in IPFS                hash=QmeF8fF5kYjvMXt7voa3SxSt14VRTMrtZ9KtEJKhk1aBkP
INFO [08-24|21:18:13.322] Writing private data hashes to Ethereum  passport=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52 fact_key="[115 105 103 110 117 112 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]" ipfs_hash=QmeF8fF5kYjvMXt7voa3SxSt14VRTMrtZ9KtEJKhk1aBkP data_key_hash=0xcaf670898b6ff3045a8c868477cc4bf0996d13aae742f76e6fa4165750d799df
INFO [08-24|21:18:13.445] Writing IPFS private data hashes to passport fact_provider=0x88BDb5deeaA40E7c980b1f52ecD2113bBF424b5C key="[115 105 103 110 117 112 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"
INFO [08-24|21:18:14.161] Waiting for transaction                  hash=0x409b9827096106ad8a0c2b6aad5616283a3465195a063e4825d5795b7c32110f
INFO [08-24|21:18:30.748] Transaction successfully mined           tx_hash=0x409b9827096106ad8a0c2b6aad5616283a3465195a063e4825d5795b7c32110f gas_used=135519
INFO [08-24|21:18:30.748] Signup info                              tx_hash=0x409b9827096106ad8a0c2b6aad5616283a3465195a063e4825d5795b7c32110f ipfs_hash=QmeF8fF5kYjvMXt7voa3SxSt14VRTMrtZ9KtEJKhk1aBkP
```

#### Event organizer gets the list of registered participants

```bash
./bin/tm signup-list --eventaddr=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52 \
  --privatekey=./event-organizer.key \
  --backendurl=https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041 \
  --ipfsurl=http://127.0.0.1:5001/
```

```
INFO [08-24|21:19:08.781] Reading data hashes from Ethereum transaction passport=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52 tx_hash=62b643…e26912
INFO [08-24|21:19:08.781] Getting transaction by hash              tx_hash=0x62b6433f5816f48237fc39cdbf07b42728b97be713cf8ca4353f5c5ac7e26912
INFO [08-24|21:19:08.919] Reading ephemeral public key from IPFS   hash=QmbyZgioijHBTJX88ChbGRkHUPMhf11fMZ3CBXEkJYRz37 filename=public_key
INFO [08-24|21:19:08.922] Reading encrypted message from IPFS      hash=QmbyZgioijHBTJX88ChbGRkHUPMhf11fMZ3CBXEkJYRz37 filename=encrypted_message
INFO [08-24|21:19:08.923] Reading message HMAC from IPFS           hash=QmbyZgioijHBTJX88ChbGRkHUPMhf11fMZ3CBXEkJYRz37 filename=hmac
INFO [08-24|21:19:08.924] Reading data hashes from Ethereum transaction passport=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52 tx_hash=409b98…32110f
INFO [08-24|21:19:08.925] Getting transaction by hash              tx_hash=0x409b9827096106ad8a0c2b6aad5616283a3465195a063e4825d5795b7c32110f
INFO [08-24|21:19:09.050] Reading ephemeral public key from IPFS   hash=QmeF8fF5kYjvMXt7voa3SxSt14VRTMrtZ9KtEJKhk1aBkP filename=public_key
INFO [08-24|21:19:09.054] Reading encrypted message from IPFS      hash=QmeF8fF5kYjvMXt7voa3SxSt14VRTMrtZ9KtEJKhk1aBkP filename=encrypted_message
INFO [08-24|21:19:09.055] Reading message HMAC from IPFS           hash=QmeF8fF5kYjvMXt7voa3SxSt14VRTMrtZ9KtEJKhk1aBkP filename=hmac
+--------------------------------------------+-----------+
|              PARTICIPANT DID               | FULL NAME |
+--------------------------------------------+-----------+
| 0x778Eb97D8F3938a13aBB9Ba26176752511d1eD70 | Dima      |
| 0xf39c56A379f6aD182b359C0296ddf3ac30A3C871 | Slava     |
+--------------------------------------------+-----------+
```

##### Event organizer creates ticket for participant 1 (Dima)

```bash
./bin/tm create-ticket --eventaddr=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52 \
  --participantaddr=0x778Eb97D8F3938a13aBB9Ba26176752511d1eD70 \
  --privatekey=./event-organizer.key \
  --backendurl=https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041 \
  --ipfsurl=http://127.0.0.1:5001/
```

Output

```
INFO [08-24|21:21:13.302] Filtering OwnershipTransferred           newOwner=0x1faF3952f34A936E042D33d58A9BBF0930886D2C
INFO [08-24|21:21:13.431] Getting transaction by hash              tx_hash=0xe268c5ade69406a3f7e75387fbd8fd969ef8708fa2374eb0cef47aaad315cd82
INFO [08-24|21:21:13.553] Writing ephemeral public key to IPFS...
INFO [08-24|21:21:13.561] Ephemeral public key added to IPFS       hash=QmWdG7DftBUEm1FxmYuscKQ5WMWaDFZZacJYWiao63fPJk size=73
INFO [08-24|21:21:13.561] Writing encrypted message to IPFS...
INFO [08-24|21:21:13.566] Encrypted message added to IPFS          hash=QmW5xqpNqM4aJD9yLwxeDhPdXH1ebaFmxAov5LLCGyCL92 size=653
INFO [08-24|21:21:13.567] Writing message HMAC to IPFS...
INFO [08-24|21:21:13.573] Message HMAC added to IPFS               hash=QmP3QhACNH1gXecicYAQx7iydjaVZLYHMJyVZ33QoScCVz size=40
INFO [08-24|21:21:13.573] Creating directory in IPFS...
INFO [08-24|21:21:13.579] Directory created in IPFS                hash=QmentPHPSk9aj4Xb8m3oHdTDcQCVyq5bp7j9X9WZ5b8vqM
INFO [08-24|21:21:13.579] Writing private data hashes to Ethereum  passport=0x778Eb97D8F3938a13aBB9Ba26176752511d1eD70 fact_key="[237 248 226 187 72 113 242 255 118 211 170 139 157 195 37 45 160 108 74 82 0 0 0 0 0 0 0 0 0 0 0 0]" ipfs_hash=QmentPHPSk9aj4Xb8m3oHdTDcQCVyq5bp7j9X9WZ5b8vqM data_key_hash=0x014406cd45ae83fa817e763b9f6c1ca79c3e817f4f08fae49459316119e742af
INFO [08-24|21:21:13.707] Writing IPFS private data hashes to passport fact_provider=0xC501F94F846ff03291f760f44e438DeAb1Eb3dCB key="[237 248 226 187 72 113 242 255 118 211 170 139 157 195 37 45 160 108 74 82 0 0 0 0 0 0 0 0 0 0 0 0]"
INFO [08-24|21:21:14.294] Waiting for transaction                  hash=0xb6b15d23e004eac5baf0ff6b09289467eeb345ae4514dc8eecfe3fa5243b5101
INFO [08-24|21:21:22.558] Transaction successfully mined           tx_hash=0xb6b15d23e004eac5baf0ff6b09289467eeb345ae4514dc8eecfe3fa5243b5101 gas_used=136415
INFO [08-24|21:21:22.558] Ticket info                              tx_hash=0xb6b15d23e004eac5baf0ff6b09289467eeb345ae4514dc8eecfe3fa5243b5101 ipfs_hash=QmentPHPSk9aj4Xb8m3oHdTDcQCVyq5bp7j9X9WZ5b8vqM
```

#### Participant 1 (Dima) reads the ticket

```bash
./bin/tm read-ticket \
  --eventaddr=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52 \
  --participantaddr=0x778Eb97D8F3938a13aBB9Ba26176752511d1eD70 \
  --privatekey=./event-participant1.key \
  --backendurl=https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041 \
  --ipfsurl=http://127.0.0.1:5001/
```

Output:

```
INFO [08-24|21:24:33.573] Reading private data hashes from Ethereum passport=0x778Eb97D8F3938a13aBB9Ba26176752511d1eD70 fact_provider=0xC501F94F846ff03291f760f44e438DeAb1Eb3dCB fact_key="[237 248 226 187 72 113 242 255 118 211 170 139 157 195 37 45 160 108 74 82 0 0 0 0 0 0 0 0 0 0 0 0]"
INFO [08-24|21:24:33.574] Getting IPFS private data hashes         fact_provider=0xC501F94F846ff03291f760f44e438DeAb1Eb3dCB key="[237 248 226 187 72 113 242 255 118 211 170 139 157 195 37 45 160 108 74 82 0 0 0 0 0 0 0 0 0 0 0 0]"
INFO [08-24|21:24:33.703] Reading ephemeral public key from IPFS   hash=QmentPHPSk9aj4Xb8m3oHdTDcQCVyq5bp7j9X9WZ5b8vqM filename=public_key
INFO [08-24|21:24:33.711] Reading encrypted message from IPFS      hash=QmentPHPSk9aj4Xb8m3oHdTDcQCVyq5bp7j9X9WZ5b8vqM filename=encrypted_message
INFO [08-24|21:24:33.712] Reading message HMAC from IPFS           hash=QmentPHPSk9aj4Xb8m3oHdTDcQCVyq5bp7j9X9WZ5b8vqM filename=hmac
Ticket QR: 0x7b227469636b65745f64617461223a2265794a6c646d56756446396b615752665957526b636d567a63794936496a42345a57526d4f475579596d49304f4463785a6a4a6d5a6a63325a444e68595468694f57526a4d7a49314d6d52684d445a6a4e4745314d694973496e4268636e527059326c7759573530583252705a4639685a4752795a584e7a496a6f694d4867334e7a686c596a6b335a44686d4d7a6b7a4f4745784d324669596a6c69595449324d5463324e7a55794e5445785a44466c5a446377496e303d222c227469636b65745f7369676e6174757265223a227238783246317354374678685743665050636b6c484e62646b47575973755950506762356b684f64662b354b62466357474837483777736242637173355956527a3753314c68544239516834445478767a506a6a2f77413d227d
```

##### Volunteer validates the ticket of participant 1 (Dima)

```bash
./bin/tm validate-ticket \
  --eventaddr=0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52 \
  --ticket=0x7b227469636b65745f64617461223a2265794a6c646d56756446396b615752665957526b636d567a63794936496a42345a57526d4f475579596d49304f4463785a6a4a6d5a6a63325a444e68595468694f57526a4d7a49314d6d52684d445a6a4e4745314d694973496e4268636e527059326c7759573530583252705a4639685a4752795a584e7a496a6f694d4867334e7a686c596a6b335a44686d4d7a6b7a4f4745784d324669596a6c69595449324d5463324e7a55794e5445785a44466c5a446377496e303d222c227469636b65745f7369676e6174757265223a227238783246317354374678685743665050636b6c484e62646b47575973755950506762356b684f64662b354b62466357474837483777736242637173355956527a3753314c68544239516834445478767a506a6a2f77413d227d \
  --backendurl=https://ropsten.infura.io/v3/e82ae97170b0425d99684d6767e20041
```

Output:

```
Valid ticket for participant DID: 0x778Eb97D8F3938a13aBB9Ba26176752511d1eD70
```


## TODO

Because not all awesome things can be finished in 36 hours.

- [OpenZeppelin GSN](https://docs.openzeppelin.com/openzeppelin/gsn/getting-started.html) "meta transactions" to make a seamless onboarding for users who sign-up for event. Because that's the only right way for user onboarding.
- React for the dApp interface. Because we still need to make functional testing of the "backend".
- _Event Organizer_ must instantiate a local IPFS node during event using [#dappNode](https://dappnode.io) for storing tickets only during event and turn of the node afterwards.
- A Participation validator should be writing a fact to _Participant's_ DID contract that a ticket was already used/validated.
