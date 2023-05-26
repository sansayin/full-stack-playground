*** Settings ***
Library               RequestsLibrary
Library               JSONLibrary
Library               Collections

#Suite Setup    Create Session  jsonplaceholder  http://127.0.0.1/api
Suite Setup    Create Session  jsonplaceholder  http://172.22.255.1/api
*** Variables ***
*** Test Cases ***

Create User
    &{data}=    Create dictionary  name=iBab892763y
    ${resp}=    POST On Session    jsonplaceholder  /user  json=${data}  expected_status=anything
    Status Should Be                 201  ${resp}
    Dictionary Should Contain Key    ${resp.json()}  id
    ${id}=   Get Value From Json   ${resp.json()}  id
    Set Global Variable    ${USER_ID}    ${id}[0]
    Log To Console    ${USER_ID} 
    Log To Console    ${resp.json()}

    
Get User
    Log To Console    ${USER_ID} 
    ${endpoint}=  Set Variable   /user/${USER_ID}
    ${resp}=  GET On Session    jsonplaceholder    ${endpoint}  
    Should Be Equal As Strings          ${resp.reason}  OK
    Log To Console    ${resp.json()}

Update User
    Log To Console    ${USER_ID} 
    ${endpoint}=  Set Variable   /user
    &{data}=    Create dictionary   id=1 name=New Baby
    ${resp}=    PUT On Session    jsonplaceholder  ${endpoint}  json=${data}  expected_status=anything
    Status Should Be                 200  ${resp}
    Dictionary Should Contain Key    ${resp.json()}  id
    Log To Console    ${resp.json()}

Delete User
    ${endpoint}=  Set Variable   /user/${USER_ID}
    ${resp}=    DELETE On Session   jsonplaceholder  ${endpoint}  
    Status Should Be                 200  ${resp}
