# Decentralized event ticket management <!-- omit in toc -->

Presenting the idea - [PDF](Decentralized-ticket-management.pdf)

Digital identity and taking back ownership of your data. These two things are the key towards mass adoption of de-centralization. While participating at the Dappcon Berlin 2019 digital identity was raised multiple times. However what if a digital identity could not only grant you access to digital world but unify access to physical world as well.

This project is was prepared during [EthBerlinZwei](https://twitter.com/ETHBerlin) Hackathon 2019.

- [Actors of the dApp](#actors-of-the-dapp)
- [Workflows](#workflows)
  - [Event Participant](#event-participant)
  - [Event Organizer](#event-organizer)
  - [Participation validators](#participation-validators)
- [Implementation](#implementation)
- [TODO](#todo)

## Actors of the dApp

- _Event Participant_. Person who is willing to participate in the event.
  - Creates or provides his digital identity wallet where event related information will be stored.
  - Submits his will to participate in the event.
- _Event Organizer_.
  - Creates a digital identity for the event.
  - Review a list of participants entering the event.
- _Event Volunteers, Security-check/Canteen worker etc._
  - Scan participants' QR code upon request to validate the ticket

## Workflows

### Event Participant

1. Each _Event Participant_ during signup creates a [digital identity](https://github.com/monetha/js-verifiable-data#Deploying-digital-identity) (aka Passport) which is later to be used to collect a ticket for the event.
2. Upon submition of data of participation a sensitive data fact is being written into the digital identity of _Event Organizer_

### Event Organizer

1. Monitors a list of newly registered _Event Participant_ and creates tickets by a click of a button.
2. Creating a ticket writes a sensitive fact to the digital identity of the _Event Participant_. A sensitive fact contains a signed _participant's information_ (provided during signup) + _address of event digital identity_.

### Participation validators

1. Asks to show his ticket via Event dApp. A QR generated from _Event Participant_ digital identity `ticket` data point provided by _Event Organizer_.
2. Ticket scanner validates QR and makes sure that it was written by the _Event Organizer_ for this Particular digital identity.
3. _Optional:_ write a fact with an _Event Organizer_ private key to mark that _Participants_ key was used.

## Implementation

![UI meme](assets/ui-meme.jpg)

We do understand the necessity of the good UI/UX. However all we did come up during hackathon is the command line utility that resembles the flow that would be a part of the dApp.

## TODO

Because not all awesome things can be finished in 36 hours.

- [OpenZeppelin GSN](https://docs.openzeppelin.com/openzeppelin/gsn/getting-started.html) "meta transactions" to make a seamless onboarding for users who sign-up for event. Because that's the only right way for user onboarding.
- React for the dApp interface. Because we still need to make functional testing of the "backend".
- Prepare an initial setup for the _Event Orgnizator_ who would setup a IPFS local node using [#dappNode](https://dappnode.io) for storing tickets only during event and turn of the node afterwards.
