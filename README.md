# gobf

Another brainfuck interpreter in go

## Usage

Its very simple:

`./gobf main.bf`

## Compiling

### Requirements

- upx
- go

_This is for windows, if you use linux you can probably figure this out yourself_

First do:
`go build -ldflags "-s -w" -o dist/gobf.exe`
to create a executable

Then to pack it do: `upx -9 dist/gobf.exe`

Then you have your executable in dist/gobf.exe.
