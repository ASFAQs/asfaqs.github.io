datadir?=AS-AIGFAQ/examples

clone:
	rm -rf AS-AIGFAQ
	git clone git@github.com:AS-AIGC/AS-AIGFAQ.git

academicfr:
	head -n 1 ${datadir}/AS-proposal-QA_fr.csv > academia.fr.csv
	sed -s '1d' ${datadir}/AS-proposal-QA_fr.csv >> academia.fr.csv
	sed -s '1d' ${datadir}/AS-dia-QA_fr.csv >> academia.fr.csv
	go run main.go academia.fr.csv
	mv academia.fr.json ../data/
	rm -rf academia.fr.csv

academicja:
	head -n 1 ${datadir}/AS-proposal-QA_ja.csv > academia.ja.csv
	sed -s '1d' ${datadir}/AS-proposal-QA_ja.csv >> academia.ja.csv
	sed -s '1d' ${datadir}/AS-dia-QA_ja.csv >> academia.ja.csv
	go run main.go academia.ja.csv
	mv academia.ja.json ../data/
	rm -rf academia.ja.csv

academicen:
	head -n 1 ${datadir}/AS-proposal-QA_en.csv > academia.en.csv
	sed -s '1d' ${datadir}/AS-proposal-QA_en.csv >> academia.en.csv
	echo "\n" >> academia.en.csv
	sed -s '1d' ${datadir}/AS-dia-QA_en.csv >> academia.en.csv
	go run main.go academia.en.csv
	mv academia.en.json ../data/
	rm -rf academia.en.csv

academic:
	head -n 1 ${datadir}/AS-proposal-QA.csv > academia.csv
	sed -s '1d' ${datadir}/AS-proposal-QA.csv >> academia.csv
	echo "\n" >> academia.csv
	sed -s '1d' ${datadir}/AS-dia-QA.csv >> academia.csv
	go run main.go academia.csv
	mv academia.json ../data/
	rm -rf academia.csv

lifefr:
	head -n 1 ${datadir}/AS-general-QA_fr.csv > life.fr.csv 
	sed -s '1d' ${datadir}/AS-general-QA_fr.csv >> life.fr.csv
	sed -s '1d' ${datadir}/AS-art-QA_fr.csv >> life.fr.csv
	go run main.go life.fr.csv
	mv life.fr.json ../data
	rm -rf life.fr.csv


lifeja:
	head -n 1 ${datadir}/AS-general-QA_ja.csv > life.ja.csv 
	sed -s '1d' ${datadir}/AS-general-QA_ja.csv >> life.ja.csv
	sed -s '1d' ${datadir}/AS-art-QA_ja.csv >> life.ja.csv
	go run main.go life.ja.csv
	mv life.ja.json ../data
	rm -rf life.ja.csv

lifeen:
	head -n 1 ${datadir}/AS-general-QA_en.csv > life.en.csv 
	sed -s '1d' ${datadir}/AS-general-QA_en.csv >> life.en.csv
	sed -s '1d' ${datadir}/AS-art-QA_en.csv >> life.en.csv
	go run main.go life.en.csv
	mv life.en.json ../data
	rm -rf life.en.csv

lifetw:
	head -n 1 ${datadir}/AS-general-QA.csv > life.csv 
	sed -s '1d' ${datadir}/AS-general-QA.csv >> life.csv
	echo "\n" >> life.csv
	sed -s '1d' ${datadir}/AS-art-QA.csv >> life.csv
	go run main.go life.csv
	mv life.json ../data
	rm -rf life.csv

itstw:
	head -n 1 ${datadir}/AS-ITS-QA.csv > its.csv 
	sed -s '1d' ${datadir}/AS-ITS-QA.csv >> its.csv
	go run main.go its.csv
	mv its.json ../data
	rm -rf its.csv
itsen:
	head -n 1 ${datadir}/AS-ITS-QA_en.csv > its.en.csv 
	sed -s '1d' ${datadir}/AS-ITS-QA_en.csv >> its.en.csv
	go run main.go its.en.csv
	mv its.en.json ../data
	rm -rf its.en.csv
itsja:
	head -n 1 ${datadir}/AS-ITS-QA_ja.csv > its.ja.csv 
	sed -s '1d' ${datadir}/AS-ITS-QA_ja.csv >> its.ja.csv
	go run main.go its.ja.csv
	mv its.ja.json ../data
	rm -rf its.ja.csv

itsfr:
	head -n 1 ${datadir}/AS-ITS-QA_fr.csv > its.fr.csv 
	sed -s '1d' ${datadir}/AS-ITS-QA_fr.csv >> its.fr.csv
	go run main.go its.fr.csv
	mv its.fr.json ../data
	rm -rf its.fr.csv

budgettw:
	head -n 1 ${datadir}/AS-dgbas-QA.csv > budget.csv 
	sed -s '1d' ${datadir}/AS-dgbas-QA.csv >> budget.csv
	go run main.go budget.csv
	mv budget.json ../data
	rm -rf budget.csv
budgeten:
	head -n 1 ${datadir}/AS-dgbas-QA_en.csv > budget.en.csv 
	sed -s '1d' ${datadir}/AS-dgbas-QA_en.csv >> budget.en.csv
	go run main.go budget.en.csv
	mv budget.en.json ../data
	rm -rf budget.en.csv
budgetja:
	head -n 1 ${datadir}/AS-dgbas-QA_ja.csv > budget.ja.csv 
	sed -s '1d' ${datadir}/AS-dgbas-QA_ja.csv >> budget.ja.csv
	go run main.go budget.ja.csv
	mv budget.ja.json ../data
	rm -rf budget.ja.csv

budgetfr:
	head -n 1 ${datadir}/AS-dgbas-QA_fr.csv > budget.fr.csv 
	sed -s '1d' ${datadir}/AS-dgbas-QA_fr.csv >> budget.fr.csv
	go run main.go budget.fr.csv
	mv budget.fr.json ../data
	rm -rf budget.fr.csv

hrtw:
	head -n 1 ${datadir}/AS-hro-QA.csv > hr.csv 
	sed -s '1d' ${datadir}/AS-hro-QA.csv >> hr.csv
	go run main.go hr.csv
	mv hr.json ../data
	rm -rf hr.csv

hren:
	head -n 1 ${datadir}/AS-hro-QA_en.csv > hr.en.csv 
	sed -s '1d' ${datadir}/AS-hro-QA_en.csv >> hr.en.csv
	go run main.go hr.en.csv
	mv hr.en.json ../data
	rm -rf hr.en.csv

hrja:
	head -n 1 ${datadir}/AS-hro-QA_ja.csv > hr.ja.csv 
	sed -s '1d' ${datadir}/AS-hro-QA_ja.csv >> hr.ja.csv
	go run main.go hr.ja.csv
	mv hr.ja.json ../data
	rm -rf hr.ja.csv
hrfr:
	head -n 1 ${datadir}/AS-hro-QA_fr.csv > hr.fr.csv 
	sed -s '1d' ${datadir}/AS-hro-QA_fr.csv >> hr.fr.csv
	go run main.go hr.fr.csv
	mv hr.fr.json ../data
	rm -rf hr.fr.csv


academia: academic academicen academicja academicfr
#academia: academic academicja academicfr

life:lifetw lifeen lifeja lifefr

its: itstw itsja itsen itsfr

hr: hrtw hrja hren hrfr

budget: budgettw budgeten budgetja budgetfr

#run: life its budget hr academia
run: hr

re: clone run
