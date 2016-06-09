# Kafkoop

## Setup

### Fetch dependencies

`bin/bootstrap`

### Start hadoop container

In a terminal run:

* `bin/start-hadoop`

An interactive shell should start connected to the container. Run the following
commands:

* `groupadd supergroup`
* `adduser hdfs -g supergroup`

### Start kafka container

In another terminal, start kafka container:

* `bin/start-kafka`

## Inspecting HDFS

* cd /usr/local/hadoop

List files:

* `bin/hdfs dfs -ls /`

Show contents of file:

* `bin/hdfs dfs -cat /tmp/my-file`

Remove file(s):

* `bin/hdfs dfs -rm -r /tmp`


