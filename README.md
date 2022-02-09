# Bank

An online, full fledged bank system - ATM, online bank, transactions, bank cards

- Online Bank: **The online banking system can be used to send transactions, check balances, get your account info, bank addresses, and your bank cards. This also simulates a ATM machine, for withdrawing and depositing money. For security, the page has only one endpoint, /. It uses tokens that expire after 30 minutes of being created, so even if someone forgot to close the tab, it would become useless after 30 minutes of logging in.**
- Bank Cards: **Bank cards are generated based on the bank account, and can be linked to a certain address. To check with the ATM for the correct details, the app encrypts the card details in SHA256 and compares them with the ATMs data.**

The internal banking system - API, card generation, addresses and customer data, is written in Go-Lang. The API uses the gin framework.

The online banking system - ATM, transactions, is written in Node.JS with Pug for the front-end.
