## Migration to DEON Controller Package

- ```pkg/msp/caclient.go```
    - invoke registration request to indy for signing DID and verkey
    - create ad hoc client controller instance from DEON ```off-grid-block/controller``` package
    - instruct client agent to create signing DID and verkey pair for client application
    - store DID/key on VON network ledger

- ```pkg/fab/signingmgr```
    - replace original SDK signing function with call to client agent to sign message

- ```pkg/fab/events/deliverclient/connection.go```
    - events which are signed by indy client agent if DID is present