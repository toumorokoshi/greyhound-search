greyhound-search
================

A web application for searching for files and finding text within them.

Requirements
------------

go1.1 or higher is required


Runbook
-------

No executables exist as of yet. If you have golang you can start greyhound
by cloning the repository, navigating the directory root, and running
the following:

    export GOPATH=`cwd`
    go get
    go run main.go

This will build the repository, and run the main executable for GreyhoundSearch

Configuration
-------------
An example configuration is below:

    {"Projects": {
        "code": {
            "Root": "/home/tsutsumi/Workspace/",
            "Exclusions": [".*\\.class", ".*\\.pyc"]
        },
        "statics": {
            "Root": "/home/tsutsumi/Opensource/"}
        }
    }

This configuration Illustrates the following:

* specifying a project, which requires:
  * a project name
  * a 'Root' key, pointing to the rootpath of that GHS should index
  * 'Exclusions', which is a list of regex files to exclude

TODO:
-----

* Image Preview / load
