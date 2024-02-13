[![Go Report Card](https://goreportcard.com/badge/github.com/jspc/synthdown)](https://goreportcard.com/report/github.com/jspc/synthdown)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=jspc_synthdown&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=jspc_synthdown)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=jspc_synthdown&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=jspc_synthdown)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=jspc_synthdown&metric=bugs)](https://sonarcloud.io/summary/new_code?id=jspc_synthdown)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=jspc_synthdown&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=jspc_synthdown)
[![Coverage Status](https://coveralls.io/repos/github/jspc/synthdown/badge.svg?branch=main)](https://coveralls.io/github/jspc/synthdown?branch=main)

# synthdown

Synthdown is:

1. A notation format for describing modular synthesiser patches
2. A reference implementation of a parser for this notation
3. A CLI that reads valid notations and generates all kinds of diagrams and stuff

## The Synthdown Notation

Synthdown notation is based on linking modules together with patch cables.

Which is fair, because that's what modular synthensisers are based on.

A simple patch might look like:

```
sequencer[bpm: 90](output)
  -> (trigger)envelope[A:0, D:2, S:0, R:2](out)
  -> (control)sine[Level: 10, Tone: 1.5](out);
```

This, in essence, says

1. Set your sequencer to 90bpm and patch from the output jack to the trigger jack of an envelope
2. Set the envelope ADSR values accordingly, then patch from the out jack to the control jack of a sine module
2. Set control to 10 and the tone to 1.5 and patch from the out jack to wherever you're off next

The patching its self is pretty obvious; the symbol `->` represents a patch cable.

A module, such as `(trigger)envelope[A:0, D:2, S:0, R:2](out)` is made up of:

* `(trigger)` - the input jack, which can be either audio or CV (we don't make any distinction)
* `envelope[A:0, D:2, S:0, R:2]` - the module its self (`envelope`) and whichever knobs and twiddler values apply
* `(out)` - the output jack

There is one special case; the first module in a patch. This will error if an input jack is set, since setting an input to the first module is an absurdity.

Finally, a patch ends with a semicolon; this allows us to add many patches to a single input file, should we want to.

### EBNF

The following [EBNF](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form) describes the notation above.

```ebnf
SynthdownFile = (Patch ";")* .
Patch = Module ("-" ">" Module)* .
Module = Jack? <ident> "[" Arg* ("," Arg)* "]" Jack .
Jack = "(" <ident> ")" .
Arg = <ident> ":" Value .
Value = <string> | <float> | <int> .
```


## Licence

BSD 3-Clause License

Copyright (c) 2024, James Condron
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
