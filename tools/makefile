
academic:
	rm -rf academia.csv
	head -n 1 proposal.csv > academia.csv
	sed -s '1d' proposal.csv >> academia.csv
	sed -s '1d' dia.csv >> academia.csv
	go run main.go academia.csv
	mv academia.json ../data

life:
	rm -rf life.csv
	head -n 1 general.csv > life.csv 
	sed -s '1d' general.csv >> life.csv
	sed -s '1d' art.csv >> life.csv
	go run main.go life.csv
	mv life.json ../data
