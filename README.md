# ec2-tag-spots

Create tags for AWC EC2 spot fleets.

## Install

```
$ go get github.com/travisjeffery/ec2-tag-spots
```

## Example

```
$ ec2-tag-spots --spot-fleet-request-id=some-request-id --tags=Name:Adele --tags=Song:Hello
success creating tags
```

## Author

[Travis Jeffery](http://twitter.com/travisjeffery)

## License

MIT
