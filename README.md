# Online voting using Blind Signatures

The model of online voting we have implemented relies on the use of blind signatures to dissociate voter preference from vote identity.



The voting is done through a client server system built using go. The server can implement mulitple ballot boxes corresponding to different elections at the same time. Inorder to facilitate this, the server has been built to allow safe multi-threaded execution. All the voter needs to do is input his/her credential and candidate preference in the client program and it will take care of the authentication of the user and subsequent casting of the vote in the chosen ballot box. 

---

## Backend

The backend has been built in golang using the crypto libraries go has.

---

## Frontend

The development of frontend is still in progress. (Using ReasonML)
