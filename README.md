# Buf Language Server

`bufls` is a prototype for the beginnings of a Protobuf language server compatible with
[Buf](https://github.com/bufbuild/buf) modules and workspaces. This currently
only supports go-to-definition.

**This is a proof-of-concept** that we wanted to share with
the community. We do not actively maintain this, and there are no guarantees
in terms of stability, but we want to hear your feedback!

For details on where we could go with this, please refer to [future work](#future-work).

## Usage

Regardless of the LSP-compatible editor you use, you'll need to install
`bufls` so that it's available on your `$PATH`.

```bash
go install github.com/bufbuild/buf-language-server/cmd/bufls
```

### Vim

With [vim-lsp], the only configuration you need is the following:

```vim
Plug 'prabirshrestha/vim-lsp'

augroup LspBuf
  au!
  autocmd User lsp_setup call lsp#register_server({
      \ 'name': 'bufls',
      \ 'cmd': {server_info->['bufls', 'serve']},
      \ 'whitelist': ['proto'],
      \ })
  autocmd FileType proto nmap <buffer> gd <plug>(lsp-definition)
augroup END
```

  [vim-lsp]: https://github.com/prabirshrestha/vim-lsp

## Supported features

Buf's language server behaves similarly to the rest of the `buf` CLI. If
the user has a `buf.work.yaml` defined, the modules defined in the workspace
will take precedence over the modules specified in the `buf.lock` (i.e. the
modules found in the module cache). The language server requires that inputs
are of the [protofile] type.

  [protofile]: https://docs.buf.build/reference/inputs#protofile

### Go to definition

Go to definition resolves the definition location of a symbol at a
given text document position (i.e. [textDocument/definition]).

This feature is currently only implemented on the `textDocument/definition`
endpoint. It may make sense to move this to `textDocument/typeDefinition`
and/or `textDocument/typeImplementation`. The Protobuf grammar is far more
limited than a programming language grammar, so not all of the semantics
for each LSP endpoint apply here.

Today, this feature is only supported for messages and enums. The well-known
types (WKT), and [custom] options are not yet supported.

  [textDocument/definition]: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_definition

## Implementation

Protobuf compilation is _fast_, so the implementation is currently naive. Every
editor command will compile the input file (e.g. `file://proto/pet/v1/pet.proto`)
from scratch (there isn't any caching). Simple caching is fairly straightforward,
but the cache would need to be cleared whenever a file is edited during the same
language server session, which would require a file watcher. For now, performance
is fine as-is (even for workspaces and large modules), but we might need to revisit
this later as build graphs continue to grow.

## Future work

### More LSP features

This is just the tip of the iceberg - there's way more that a fully-featured Protobuf
language server can do for the Protobuf community. For starters, the following set of
endpoints are next in line:

 - [textDocument/completion]
 - [textDocument/codeLens]
 - [textDocument/foldingRange]
 - [textDocument/formatting]
 - [textDocument/hover]

  [textDocument/completion]: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion
  [textDocument/codeLens]: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_codeLens
  [textDocument/foldingRange]: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_foldingRange
  [textDocument/formatting]: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_formatting
  [textDocument/hover]: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_hover

### Go to definition

A couple features remain for full go to definition support:

 - Add go to definition support for [custom] options.
 - Add go to definition support for the well-known types (i.e. synthesize the WKT in the module cache).
