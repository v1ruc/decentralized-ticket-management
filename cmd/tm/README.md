# Ticket management (tm) tool

Event organizer DID address: [0xEDf8e2Bb4871f2ff76D3aa8B9Dc3252dA06c4a52](https://scanner.monetha.io/identity/0xedf8e2bb4871f2ff76d3aa8b9dc3252da06c4a52?network=ropsten)
Participant 1 DID address: [0x778Eb97D8F3938a13aBB9Ba26176752511d1eD70](https://scanner.monetha.io/identity/0x778eb97d8f3938a13abb9ba26176752511d1ed70?network=ropsten)
Participant 2 DID address: [0xf39c56A379f6aD182b359C0296ddf3ac30A3C871](https://scanner.monetha.io/identity/0xf39c56a379f6ad182b359c0296ddf3ac30a3c871?network=ropsten)

## Participant 1 sign up

```bash
./bin/tm signup \
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

## Participant 2 sign up

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

## Event organizer gets the list of registered participants

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

## Event organizer creates ticket for participant 1 (Dima)

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

## Participant 1 (Dima) reads the ticket

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

## Volunteer validates the ticket of participant 1 (Dima)

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