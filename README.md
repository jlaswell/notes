# notes

notes is a simple note taking utility that was developed for simplicity and ease. The following examples are some common ways to use notes.

### Getting Started

```sh
go get -u github.com/jlaswell/notes
cd $GOPATH/github.com/jlaswell/notes
go install
cd -
```

### Examples

To create or edit a note

```sh
notes @ meetings/sprint15/retro-minutes
```

To list notes

```sh
notes ls meetings/sprint15

sprint15
└── grooming
└── retro-minutes
```

To delete a note

```sh
notes del meetings/sprint15/retro-minutes
```
