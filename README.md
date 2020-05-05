# spannerDdl2ModelJson

## Installation

(need go 1.12+ to run it)

``` shell
go build
```

## Usage

the program will receive DDL statement from os.Stdin and output a model json file to os.Stdout. so you can run it like this (base on command-line tool jq to format result, remove it if you don't need it)

``` shell
echo ' CREATE TABLE students (
    Id INT64 NOT NULL,
    FirstName STRING(1024) NOT NULL, /* first name of student */
    LastName STRING(1024) NOT NULL, /* last name of student */
) PRIMARY KEY (Id);' | ./spannerDdl2ModelJson | jq
```
output will be like this:

``` json
{
  "type": "object",
  "format": "students",
  "required": true,
  "description": "",
  "properties": {
    "FirstName": {
      "type": "string",
      "format": "string",
      "required": true,
      "description": "first name of student"
    },
    "Id": {
      "type": "number",
      "format": "number",
      "required": true,
      "description": ""
    },
    "LastName": {
      "type": "string",
      "format": "string",
      "required": true,
      "description": "last name of student"
    }
  }
}
```
