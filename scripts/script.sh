echo "GET http://localhost:8080/orders/" | vegeta attack -rate=50 -duration=5s | tee results.bin | vegeta report
vegeta plot -title="Attack Rate" results.bin > plot.html