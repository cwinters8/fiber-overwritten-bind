# overwritten bind demo

simple demonstration of bind variables being overwritten by defaults set in fiber middleware

## running locally

first clone this repo and `cd` to it

```sh
go build -o tmp/demo && ./tmp/demo
```

then navigate to `localhost:3000` in your browser

you will see that "default" is being rendered in the h1 tag, instead of the expected "home" value
