
LOVE=love

test:
	go run server.go --test_conn_and_quit &
	# hacky way of waiting for the server to be up and listening before connecting the client
	sleep 1.0
	$(LOVE) client/ --test_conn_and_quit

