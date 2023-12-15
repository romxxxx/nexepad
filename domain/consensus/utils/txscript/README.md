txscript
========

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://godoc.org/github.com/romxxxx/nexepad/txscript?status.png)](http://godoc.org/github.com/romxxxx/nexepad/txscript)

Package txscript implements the nexepa transaction script language. There is
a comprehensive test suite.

## nexepa Scripts

nexepa provides a stack-based, FORTH-like language for the scripts in
the nexepa transactions. This language is not turing complete
although it is still fairly powerful. 

## Examples

* [Standard Pay-to-pubkey Script](http://godoc.org/github.com/romxxxx/nexepad/txscript#example-PayToAddrScript)  
  Demonstrates creating a script which pays to a nexepa address. It also
  prints the created script hex and uses the DisasmString function to display
  the disassembled script.

* [Extracting Details from Standard Scripts](http://godoc.org/github.com/romxxxx/nexepad/txscript#example-ExtractPkScriptAddrs)  
  Demonstrates extracting information from a standard public key script.
