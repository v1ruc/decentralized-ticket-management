# Decentralized event ticket management

Digital identity and taking ownership of your data we believe are key things towards mass adoption of de-centralization. While participating at the conference digital identity was raised multiple times. What if a digital identity could not only grant you access to digital world but unify access to physical world as well. This is how we came up with a ticket management dApp idea for hackathons like [EthBerlinZwei](https://twitter.com/ETHBerlin).

## Actors of the dApp

- _Event Participant_. Person who is willing to participate in the event.
  - Creates or provides his digital identity wallet where event related information will be stored.
  - Submits his will to participate in the event.
- _Event Organizer_.
  - Creates a digital identity for the event.
  - Review a list of participants entering the event.
- _Event Volunteers_, Security-check/Canteen worker etc.
  - Scan participants' QR code upon request to validate the ticket

## Workflows

*Event Participant*

1. Each _Event Participant_ during signup creates a [digital identity](https://github.com/monetha/js-verifiable-data#Deploying-digital-identity) (aka Passport) which is later to be used to collect a ticket for the event.
2. Upon submition of data of participation a sensitive data fact is being written into the digital identity of _Event Organizer_

*Event Organizer*

1. Monitors a list of newly registered _Event Participant_ and creates tickets by a click of a button.
2. Creating a ticket writes a sensitive fact to the digital identity of the _Event Participant_. A sensitive fact contains an encrypted _event_id + participant's digital identity address_

## Tools to be used for final product

- Monetha platform SDK. Maintain digital identity and data (aka information) ownership. A concept of interaction is shown via command line utility prepared during hackathon.
- OpenZeppelin GSN "meta transactions" to make a seamless onboarding for users who sign-up for event.
- React for App interface
