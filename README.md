# govern

from Ancient Greek κυβερνάω (kubernáō, “I steer, drive, govern”)

This is a tool that figures out what I intend to do when working with a file in
vim. For example in ruby, I want to run rspec but in clojure test files I want
to evaluate the current file and additionally run tests if clojure.test is
required.

### Usage

`govern -file ~/.vimrc`
`govern -file ~/ruby-project/spec/some_spec.rb`
`govern -file ~/clojure-project/clj/core.clj`
`govern -file ~/some.md`

When a file is given to govern, it will determine based on it's extension what
acction should be performed and provide a response to this activity to stdout.
This can be used from within vim using the wrapper function:

`right here`

### Roadmap

markdown should be converted to html in tmp and opened in a browser

# notes

task: read tests and evaluate when govern can do
task: create a roadmap of the next 3 things you want to accomplish.
task: do them.
