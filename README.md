# Sabanci University TUI Application

A simple command line application for viewing the one's student information
at Sabanci University. It is written in Go using the Bubbletea library.

Currently, it supports the following features:
- Viewing your schedule for the current day
- Viewing your wallet balance of your student card
- Viewing the cafeteria meal plan

You can login with your Sabanci University credentials. The application
will not store your credentials. It is only used for logging in to the
CAS server.
The application will store your session cookies so the next time you
run the application, you will not need to login manually again.

## Contributing

The application is written in Go. You can use the newest Go toolchain
to build the application.

Contributions are welcome. You can open an issue or a pull request
to contribute to the project.
Some ideas for contributions:
- Adding support for the shuttle bus schedule
- Adding support for the library information

The application is not affiliated with Sabanci University.
The information is not retrieved via an API. Rather, it is scraped
from the Sabanci University website, usually by calling the right ajax
endpoints. You can find examples in the codebase.
Every endpoint needs to have a `retrieve` function and a `parse` function.
Finally, the service needs to be called by a bubbletea component.

I am only studying at Sabanci University for one semester as an exchange
student. If there is interest in this project beyond my stay, I may
transfer the ownership of the project to someone else.
